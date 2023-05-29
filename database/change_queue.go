/*
Copyright (c) 2021 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package database

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"text/template"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	"github.com/openshift-online/ocm-sdk-go/logging"
)

// ChangeQueueBuilder contains the data and logic needed to build a database change queue.
type ChangeQueueBuilder struct {
	logger   logging.Logger
	url      string
	name     string
	tables   []string
	install  bool
	context  func() context.Context
	callback func(context.Context, *ChangeQueueItem)
	interval time.Duration
	timeout  time.Duration
}

// ChangeQueue is a mechanism to asynchronously process changes made to database tables. It uses
// triggers that write to a changes table the details of the changes made to other tables. It then
// waits for data written to that changes table and processes it using the callbaks given in the
// configuration.
type ChangeQueue struct {
	// Basic fields:
	logger        logging.Logger
	url           string
	name          string
	context       func() context.Context
	callback      func(context.Context, *ChangeQueueItem)
	waitIterval   time.Duration
	retryInterval time.Duration
	timeout       time.Duration

	// We use the flag to ask the loop to stop as soon as possible, and then the loop uses the
	// channel to tell the close method that it actually finished.
	closeFlag int32
	closeChan chan struct{}

	// Database connection used to receive notifications.
	waitConn *pgx.Conn

	// Function to cancel the context used to run the current wait.
	waitCancel func()

	// Database connection used to fetch changes.
	fetchConn *pgx.Conn

	// Precalculated SQL for frequently used statements:
	fetchSQL  string
	listenSQL string
}

// ChangeQueueItem contains the data of one single change.
type ChangeQueueItem struct {
	Serial    int
	Timestamp time.Time
	Source    string
	Operation string
	Old       []byte
	New       []byte
}

// changeQueueRow is used to read rows from the database.
type changeQueueRow struct {
	Serial    int
	Timestamp time.Time
	Source    string
	Operation string
	Old       []byte
	New       []byte
}

// NewChangeQueue creates a builder that can then be used to configure and create a change queue.
// Note that the required database objects (table, functions and triggers) aren't automatically
// created by default. See the documentation of the Install method of the builder for details.
func NewChangeQueue() *ChangeQueueBuilder {
	return &ChangeQueueBuilder{
		name:     defaultChangeQueueName,
		interval: defaultChangeQueueInterval,
		timeout:  defaultChangeQueueTimeout,
		install:  false,
		context:  context.Background,
	}
}

// Logger sets the logger that the queue will use to write to the log. This is mandatory.
func (b *ChangeQueueBuilder) Logger(value logging.Logger) *ChangeQueueBuilder {
	b.logger = value
	return b
}

// URL sets the database URL that the queue will use connect to the database. This is mandatory.
func (b *ChangeQueueBuilder) URL(value string) *ChangeQueueBuilder {
	b.url = value
	return b
}

// Name sets the name of the queue. This can be used to have different queues for different uses, or
// for different sets of tables, or for different environments that happen to share the database.
// The changes table, function and triggers created will have this name. The default is `changes`.
func (b *ChangeQueueBuilder) Name(value string) *ChangeQueueBuilder {
	b.name = value
	return b
}

// Install enables or disables the creation of the database objects (tables, functions and triggers)
// needed by the queue. If set to true the objects will be automatically created if they don't
// exist.
//
// If set to false then you will need to manually create the following objects:
//
// 1. The table where the changes are stored. For example, if the name of the queue is `my_queue`:
//
//	create table if not exists my_queue (
//		serial serial not null primary key,
//		timestamp timestamp with time zone not null default now(),
//		source text,
//		operation text,
//		old jsonb,
//		new jsonb
//	);
//
// 2. The trigger function that will be called by the triggers to write changes to the table:
//
//	create or replace function my_queue_save() returns trigger as $$
//	begin
//		insert into my_queue (
//			source,
//			operation,
//			old,
//			new
//		) values (
//			tg_table_name,
//			lower(tg_op),
//			row_to_json(old.*),
//			row_to_json(new.*)
//		);
//		return null;
//	end;
//	$$ language plpgsql;
//
// 3. The trigger function that will be called by triggers to send notifications:
//
//	create or replace function my_queue_notify() returns trigger as $$
//	begin
//		notify my_queue;
//		return null;
//	end;
//	$$ language plpgsql;
//
// 4. For each data table the trigger that calls the save function. For example, if the name of the
// data table is `my_data`:
//
//	drop trigger if exists my_data_my_queue_save on my_data;
//	create trigger my_data_my_queue_save
//		after insert or update or delete on my_data
//		for each row execute function my_queue_save();
//
// 5. For each data table the trigger that calls the notify function:
//
//	create trigger my_data_my_queue_notify
//		after insert or update or delete on my_data
//		execute function my_queue_notify();
//
// The default value is false.
func (b *ChangeQueueBuilder) Install(value bool) *ChangeQueueBuilder {
	b.install = value
	return b
}

// Table adds a table that will be configured so that changes are written to the change queue.
func (b *ChangeQueueBuilder) Table(value string) *ChangeQueueBuilder {
	b.tables = append(b.tables, value)
	return b
}

// Table adds a collection of tables that will be configured so that changes are written to the
// changes queue.
func (b *ChangeQueueBuilder) Tables(values ...string) *ChangeQueueBuilder {
	b.tables = append(b.tables, values...)
	return b
}

// Context sets a function that the queue will use to create contexts. The default is to create
// contexts using the context.Background function.
func (b *ChangeQueueBuilder) Context(value func() context.Context) *ChangeQueueBuilder {
	b.context = value
	return b
}

// Callback sets the function that will be called to process changes. This function will be called
// in the same goroutine that reads the item from the database, so processing of later changes will
// not proceed till this returns.
func (b *ChangeQueueBuilder) Callback(
	value func(context.Context, *ChangeQueueItem)) *ChangeQueueBuilder {
	b.callback = value
	return b
}

// Interval sets the iterval for periodic checks. Usually changes from the queue will be processed
// inmediately because the database sends notifications when new changes are available. If that
// fails, for whatever the raeson, changes will be processed after this interval. Default value is
// 30 seconds.
func (b *ChangeQueueBuilder) Interval(value time.Duration) *ChangeQueueBuilder {
	b.interval = value
	return b
}

// Timeout sets the timeout for database operations. The default is one second.
func (b *ChangeQueueBuilder) Timeout(value time.Duration) *ChangeQueueBuilder {
	b.timeout = value
	return b
}

// Build uses the data stored in the builder to configure and create a new change queue.
func (b *ChangeQueueBuilder) Build(ctx context.Context) (result *ChangeQueue, err error) {
	// Check parameters:
	if b.logger == nil {
		err = errors.New("logger is mandatory")
		return
	}
	if b.url == "" {
		err = errors.New("database URL is mandatory")
		return
	}
	if b.name == "" {
		err = errors.New("name is mandatory")
		return
	}
	if b.context == nil {
		err = errors.New("context function is mandatory")
		return
	}
	if b.callback == nil {
		err = errors.New("callback function is mandatory")
		return
	}
	if b.interval <= 0 {
		err = fmt.Errorf(
			"check interval %s isn't valid, should be greater or equal than zero",
			b.interval,
		)
		return
	}
	if b.timeout <= 0 {
		err = fmt.Errorf(
			"timeout %s isn't valid, should be greater or equal than zero",
			b.timeout,
		)
		return
	}

	// Create the database objects if needed. Note that if the install flag is false this will
	// only write the SQL code to the log and will not try to create the objects.
	err = b.createObjects(ctx)
	if err != nil {
		return
	}

	// Calculate specific intervals from the general interval given in the configuration:
	waitInterval := b.interval
	retryInterval := b.interval / 10

	// Calculate the SQL for frequently used statements:
	fetchSQL, err := evaluateTemplate(
		changeQueueFetchTemplate,
		"Name", b.name,
	)
	if err != nil {
		return
	}
	listenSQL, err := evaluateTemplate(
		changeQueueListenTemplate,
		"Name", b.name,
	)
	if err != nil {
		return
	}

	// Create and populate the object:
	result = &ChangeQueue{
		logger:        b.logger,
		url:           b.url,
		name:          b.name,
		context:       b.context,
		callback:      b.callback,
		waitIterval:   waitInterval,
		retryInterval: retryInterval,
		timeout:       b.timeout,
		fetchSQL:      fetchSQL,
		listenSQL:     listenSQL,
		closeFlag:     0,
		closeChan:     make(chan struct{}),
	}

	// Start the loop:
	go result.loop()

	return
}

// createObjects creates the database objects needed by the queue: tables, triggers and functions.
func (b *ChangeQueueBuilder) createObjects(ctx context.Context) error {
	var err error

	// Create the database connection to install database objects:
	var conn *pgx.Conn
	if b.install {
		conn, err = pgx.Connect(ctx, b.url)
		if err != nil {
			return err
		}
		defer func() {
			err := conn.Close(ctx)
			if err != nil {
				b.logger.Error(ctx, "Can't close connection: %v", err)
			}
		}()
	}

	// Create the tables:
	err = b.createTables(ctx, conn)
	if err != nil {
		return err
	}

	// Create the functions:
	err = b.createFunctions(ctx, conn)
	if err != nil {
		return err
	}

	// Configure the tables:
	for _, table := range b.tables {
		err = b.configureTable(ctx, conn, table)
		if err != nil {
			return err
		}
	}
	return nil
}

// createTables creates the changes table if it doesn't already exist.
func (b *ChangeQueueBuilder) createTables(ctx context.Context, conn *pgx.Conn) error {
	var err error
	createTableSQL, err := evaluateTemplate(
		changeQueueTableTemplate,
		"Name", b.name,
	)
	if err != nil {
		return err
	}
	if b.install {
		_, err = conn.Exec(ctx, createTableSQL)
		if err != nil {
			return err
		}
	} else {
		b.logger.Info(
			ctx,
			"To create the changes table for queue '%s' run the "+
				"following SQL: %s",
			b.name, createTableSQL,
		)
	}
	return nil
}

// createFunctions creates the trigger functions if they don't already exist.
func (b *ChangeQueueBuilder) createFunctions(ctx context.Context, conn *pgx.Conn) error {
	var err error
	// Create the save function:
	createSaveFunctionSQL, err := evaluateTemplate(
		changeQueueSaveFunctionTemplate,
		"Name", b.name,
	)
	if err != nil {
		return err
	}
	if b.install {
		_, err = conn.Exec(ctx, createSaveFunctionSQL)
		if err != nil {
			return err
		}
	} else {
		b.logger.Info(
			ctx,
			"To create the save function for queue '%s' run the "+
				"following SQL: %s",
			b.name, createSaveFunctionSQL,
		)
	}

	// Create the notify function:
	createNotifyFunctionSQL, err := evaluateTemplate(
		changeQueueNotifyFunctionTemplate,
		"Name", b.name,
	)
	if err != nil {
		return err
	}
	if b.install {
		_, err = conn.Exec(ctx, createNotifyFunctionSQL)
		if err != nil {
			return err
		}
	} else {
		b.logger.Info(
			ctx,
			"To create the notify function for queue '%s' run the "+
				"following SQL: %s",
			b.name, createNotifyFunctionSQL,
		)
	}

	return nil
}

// configureTable configures the given table with triggers that call the functions that save the
// changes to the changes table and send the notifications.
func (b *ChangeQueueBuilder) configureTable(ctx context.Context, conn *pgx.Conn,
	table string) error {
	// Create the save trigger:
	createSaveTriggerSQL, err := evaluateTemplate(
		changeQueueSaveTriggerTemplate,
		"Name", b.name,
		"Table", table,
	)
	if err != nil {
		return err
	}
	if b.install {
		_, err = conn.Exec(ctx, createSaveTriggerSQL)
		if err != nil {
			return err
		}
	} else {
		b.logger.Info(
			ctx,
			"To create the save trigger for queue '%s' and table '%s' run the "+
				"following SQL: %s",
			b.name, table, createSaveTriggerSQL,
		)
	}

	// Create the notify trigger:
	createNotifyTriggerSQL, err := evaluateTemplate(
		changeQueueNotifyTriggerTemplate,
		"Name", b.name,
		"Table", table,
	)
	if err != nil {
		return err
	}
	if b.install {
		_, err = conn.Exec(ctx, createNotifyTriggerSQL)
		if err != nil {
			return err
		}
	} else {
		b.logger.Info(
			ctx,
			"To create the notify trigger for queue '%s' and table '%s' run the "+
				"following SQL: %s",
			b.name, table, createNotifyTriggerSQL,
		)
	}

	return nil
}

// loop runs the loop that waits for notifications from the database and processes the pending
// changes.
func (q *ChangeQueue) loop() {
	// Create a context:
	ctx := q.context()

	for !q.closing() {
		// Check for pending changes:
		q.check(ctx)

		// Wait for notifications. It is normal if this finishes with a timeout. Any other
		// error isn't normal and we should wait a bit before trying again to avoid too many
		// attempts when the error isn't resolved quickly.
		err := q.wait(ctx)
		if err != nil {
			q.logger.Info(
				ctx,
				"Wait failed, will wait a bit before trying again: %v",
				err,
			)
			time.Sleep(q.retryInterval)
		}
	}

	// Let the close method know that we finished:
	close(q.closeChan)
}

// wait waits for the next notification from the database is received or the operation times out.
func (q *ChangeQueue) wait(ctx context.Context) error {
	var err error

	// Start listening if needed:
	if q.waitConn == nil {
		q.waitConn, err = pgx.Connect(ctx, q.url)
		if err != nil {
			return err
		}
		_, err = q.waitConn.Exec(ctx, q.listenSQL)
		if err != nil {
			return err
		}
	}

	// Wait for a new notification. We set a timeout so that we will process pending changes
	// periodically even if the notification mechanism fails. If the wait results in an error
	// other than a timeout then we will close and discard the connection, so that we recover
	// from database restarts and other errors that may make the connection unusable.
	var waitCtx context.Context
	waitCtx, q.waitCancel = context.WithTimeout(ctx, q.waitIterval)
	defer q.waitCancel()
	_, err = q.waitConn.WaitForNotification(waitCtx)
	if pgconn.Timeout(err) {
		err = nil
	}
	if err != nil {
		q.logger.Debug(ctx, "Wait failed, will close the connection: %v", err)
		closeErr := q.waitConn.Close(ctx)
		if closeErr != nil {
			q.logger.Info(ctx, "Can't close connection: %v", closeErr)
		}
		q.waitConn = nil
		return err
	}

	return nil
}

// check checks the contents of the changes table and process them.
func (q *ChangeQueue) check(ctx context.Context) {
	// Fetch and process all the available changes.
	for !q.closing() {
		// Fetch the next available row:
		found, row, err := q.fetch(ctx)
		if err != nil {
			q.logger.Error(ctx, "Can't fetch change: %v", err)
			return
		}
		if !found {
			break
		}

		// Process the change:
		q.logger.Debug(ctx, "Processing change %d", row.Serial)
		change := &ChangeQueueItem{
			Serial:    row.Serial,
			Timestamp: row.Timestamp,
			Source:    row.Source,
			Operation: row.Operation,
			Old:       row.Old,
			New:       row.New,
		}
		q.callback(ctx, change)
	}
}

// fetch tries to read the next row from the changes table. It returns a boolean flag indicating if
// there was such a row and the row itself.
func (q *ChangeQueue) fetch(ctx context.Context) (found bool, result *changeQueueRow,
	err error) {
	// Create the connection if needed:
	if q.fetchConn == nil {
		q.fetchConn, err = pgx.Connect(ctx, q.url)
		if err != nil {
			return
		}
	}

	// Run the query:
	queryCtx, queryCancel := context.WithTimeout(ctx, q.timeout)
	defer queryCancel()
	row := q.fetchConn.QueryRow(queryCtx, q.fetchSQL)
	var tmp changeQueueRow
	err = row.Scan(
		&tmp.Serial,
		&tmp.Timestamp,
		&tmp.Source,
		&tmp.Operation,
		&tmp.Old,
		&tmp.New,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		err = nil
		return
	}
	if pgconn.Timeout(err) {
		err = nil
		return
	}
	if err != nil {
		q.logger.Debug(ctx, "Fetch failed, will close the connection: %v", err)
		closeErr := q.fetchConn.Close(ctx)
		if closeErr != nil {
			q.logger.Info(ctx, "Can't close connection: %v", closeErr)
		}
		q.fetchConn = nil
		err = nil
	}
	found = true
	result = &tmp
	return
}

func evaluateTemplate(source string, args ...interface{}) (result string,
	err error) {
	// Check that there is an even number of args, and that the first of each pair is a string:
	count := len(args)
	if count%2 != 0 {
		err = fmt.Errorf(
			"template '%s' should have an even number of arguments, but it has %d",
			source, count,
		)
		return
	}
	for i := 0; i < count; i = i + 2 {
		name := args[i]
		_, ok := name.(string)
		if !ok {
			err = fmt.Errorf(
				"argument %d of template '%s' is a key, so it should be a string, "+
					"but its type is %T",
				i, source, name,
			)
			return
		}
	}

	// Put the variables in the map that will be passed as the data object for the execution of
	// the template:
	data := make(map[string]interface{})
	for i := 0; i < count; i = i + 2 {
		name := args[i].(string)
		value := args[i+1]
		data[name] = value
	}

	// Parse the template:
	tmpl, err := template.New("").Parse(source)
	if err != nil {
		err = fmt.Errorf("can't parse template '%s': %v", source, err)
		return
	}

	// Execute the template:
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		err = fmt.Errorf("can't execute template '%s': %v", source, err)
		return
	}

	// Return the result:
	result = buffer.String()
	return
}

// Close releases all the resources used by the queue. Note that the changes will continue to be
// written to the changes table even when the queue is closed, and that may continue to be processed
// by other queue objects running in a this or other processes.
func (q *ChangeQueue) Close() error {
	// Create a context:
	ctx := q.context()

	// Raise the closing flag so that the loop will finish as soon as it checks it:
	q.close()

	// Cancel waiting and close the connection:
	if q.waitCancel != nil {
		q.waitCancel()
	}

	// Wait for the loop to finish:
	<-q.closeChan

	// Close the connection used for fetching changes:
	if q.fetchConn != nil {
		err := q.fetchConn.Close(ctx)
		if err != nil {
			q.logger.Info(ctx, "Can't close connection: %v", err)
		}
	}

	// Close the connection used for listening for notifications:
	if q.waitConn != nil {
		err := q.waitConn.Close(ctx)
		if err != nil {
			q.logger.Info(ctx, "Can't close connection: %v", err)
		}
	}

	return nil
}

// close asks the loop to stop as soon as possible.
func (q *ChangeQueue) close() {
	atomic.StoreInt32(&q.closeFlag, 1)
}

// closing returns true if we are in the process of closing.
func (q *ChangeQueue) closing() bool {
	return atomic.LoadInt32(&q.closeFlag) == 1
}

// template used to create the table.
const changeQueueTableTemplate = `
create table if not exists {{ .Name }} (
	serial serial not null primary key,
	timestamp timestamp with time zone not null default now(),
	source text,
	operation text,
	old jsonb,
	new jsonb
);
`

// template used to create the save function.
const changeQueueSaveFunctionTemplate = `
create or replace function {{ .Name }}_save() returns trigger as $$
begin
	insert into {{ .Name }} (
		source,
		operation,
		old,
		new
	) values (
		tg_table_name,
		lower(tg_op),
		row_to_json(old.*),
		row_to_json(new.*)
	);
	return null;
end;
$$ language plpgsql;
`

// template used to create the notify function.
const changeQueueNotifyFunctionTemplate = `
create or replace function {{ .Name }}_notify() returns trigger as $$
begin
	notify {{ .Name }};
	return null;
end;
$$ language plpgsql;
`

// template used to create the triggers that save changes to the changes table.
const changeQueueSaveTriggerTemplate = `
drop trigger if exists {{ .Table }}_{{ .Name}}_save on {{ .Table }};
create trigger {{ .Table }}_{{ .Name }}_save
	after insert or update or delete on {{ .Table }}
	for each row execute function {{ .Name }}_save();
`

// template used to create the triggers that send notifications.
const changeQueueNotifyTriggerTemplate = `
drop trigger if exists {{ .Table }}_{{ .Name }}_notify on {{ .Table }};
create trigger {{ .Table }}_{{ .Name }}_notify
	after insert or update or delete on {{ .Table }}
	execute function {{ .Name }}_notify();
`

// template used to fetch the next change.
const changeQueueFetchTemplate = `
delete from {{ .Name }} where serial = (
	select serial
	from {{ .Name }}
	order by serial
	for update
	skip locked
	limit 1
)
returning
	serial,
	timestamp,
	source,
	operation,
	old,
	new
;
`

// template used to listen for notifications.
const changeQueueListenTemplate = `
listen {{ .Name }};
`

// Defaults for configuration settings:
const (
	defaultChangeQueueName     = "changes"
	defaultChangeQueueInterval = 30 * time.Second
	defaultChangeQueueTimeout  = 1 * time.Second
)
