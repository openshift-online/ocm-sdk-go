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
	"context"
	"database/sql"
	"time"

	. "github.com/onsi/ginkgo"                         // nolint
	. "github.com/onsi/gomega"                         // nolint
	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

var _ = Describe("Change log behaviour", func() {
	var ctx context.Context
	var dbObject *Database
	var dbURL string
	var dbHandle *sql.DB

	BeforeEach(func() {
		// Create a context:
		ctx = context.Background()

		// Create a database:
		dbObject = dbServer.MakeDatabase()
		dbURL = dbObject.MakeURL()
		dbHandle = dbObject.MakeHandle()

		// Create the data table:
		_, err := dbHandle.Exec(`
			create table my_data (
				id integer not null primary key,
				name text not null
			)
		`)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		// Close the database handle:
		err := dbHandle.Close()
		Expect(err).ToNot(HaveOccurred())

		// Close the database server:
		dbObject.Close()
	})

	// NopCallback is a callback that doesn nothing with the change.
	var NopCallback = func(ctx context.Context, change *ChangeQueueItem) {
		// Do nothing, as the name says.
	}

	It("Can't be created without a logger", func() {
		_, err := NewChangeQueue().
			URL(dbURL).
			Install(true).
			Callback(NopCallback).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("logger"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Can't be created without a database URL", func() {
		_, err := NewChangeQueue().
			Logger(logger).
			Install(true).
			Callback(NopCallback).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("database"))
		Expect(message).To(ContainSubstring("URL"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Can't be created with zero interval", func() {
		_, err := NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Install(true).
			Callback(NopCallback).
			Interval(0).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("interval"))
		Expect(message).To(ContainSubstring("greater or equal than zero"))
	})

	It("Can't be created with zero timeout", func() {
		_, err := NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Install(true).
			Callback(NopCallback).
			Timeout(0).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("timeout"))
		Expect(message).To(ContainSubstring("greater or equal than zero"))
	})

	It("Creates the changes table if it doesn't exist", func() {
		// Create the queue:
		queue, err := NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Name("my_changes").
			Install(true).
			Callback(NopCallback).
			Interval(100 * time.Millisecond).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = queue.Close()
			Expect(err).ToNot(HaveOccurred())
		}()

		// Check that the table exists:
		rows, err := dbHandle.Query(`
			select
				serial,
				timestamp,
				source,
				operation,
				old,
				new
			from
				my_changes
		`)
		Expect(err).ToNot(HaveOccurred())
		err = rows.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can be created if the changes table already exists", func() {
		// Create the changes table:
		_, err := dbHandle.Exec(`
			create table my_changes (
				serial serial primary key,
				timestamp timestamp with time zone not null default now(),
				source text,
				operation text,
				old jsonb,
				new jsonb
			)
		`)
		Expect(err).ToNot(HaveOccurred())

		// Create the queue:
		queue, err := NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Name("my_changes").
			Install(true).
			Callback(NopCallback).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = queue.Close()
			Expect(err).ToNot(HaveOccurred())
		}()
	})

	It("Processes insert, update and delete", func() {
		// Create a callback function that stores the changes in an array:
		var changes []*ChangeQueueItem
		callback := func(ctx context.Context, item *ChangeQueueItem) {
			changes = append(changes, item)
		}

		// Create the queue:
		queue, err := NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Name("my_changes").
			Table("my_data").
			Install(true).
			Callback(callback).
			Interval(100 * time.Millisecond).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = queue.Close()
			Expect(err).ToNot(HaveOccurred())
		}()

		// Insert:
		_, err = dbHandle.Exec(`insert into my_data (id, name) values (123, 'my_name')`)
		Expect(err).ToNot(HaveOccurred())

		// Update:
		_, err = dbHandle.Exec(`update my_data set name = 'your_name' where id = 123`)
		Expect(err).ToNot(HaveOccurred())

		// Delete:
		_, err = dbHandle.Exec(`delete from my_data where id = 123`)
		Expect(err).ToNot(HaveOccurred())

		// Wait a bit so that the changes can be processed:
		time.Sleep(150 * time.Millisecond)

		// Check the number of changes:
		Expect(changes).To(HaveLen(3))

		// Check the insert:
		insertChange := changes[0]
		Expect(insertChange.Serial).To(Equal(1))
		Expect(insertChange.Source).To(Equal("my_data"))
		Expect(insertChange.Operation).To(Equal("insert"))
		Expect(insertChange.Old).To(BeEmpty())
		Expect(insertChange.New).To(MatchJSON(`{
			"id": 123,
			"name": "my_name"
		}`))

		// Check the update:
		updateChange := changes[1]
		Expect(updateChange.Serial).To(Equal(2))
		Expect(updateChange.Source).To(Equal("my_data"))
		Expect(updateChange.Operation).To(Equal("update"))
		Expect(updateChange.Old).To(MatchJSON(`{
			"id": 123,
			"name": "my_name"
		}`))
		Expect(updateChange.New).To(MatchJSON(`{
			"id": 123,
			"name": "your_name"
		}`))

		// Check the delete:
		deleteChange := changes[2]
		Expect(deleteChange.Serial).To(Equal(3))
		Expect(deleteChange.Source).To(Equal("my_data"))
		Expect(deleteChange.Operation).To(Equal("delete"))
		Expect(deleteChange.Old).To(MatchJSON(`{
			"id": 123,
			"name": "your_name"
		}`))
		Expect(deleteChange.New).To(BeEmpty())
	})

	It("Processes changes made before creation", func() {
		// Create the queue and close it inmediately so that it will only install the
		// database objects but not process any changes:
		queue, err := NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Name("my_changes").
			Table("my_data").
			Install(true).
			Callback(NopCallback).
			Interval(100 * time.Millisecond).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		err = queue.Close()
		Expect(err).ToNot(HaveOccurred())

		// Make a change:
		_, err = dbHandle.Exec(`insert into my_data (id, name) values (123, 'my_name')`)
		Expect(err).ToNot(HaveOccurred())

		// Create the queue again, this time letting it run till the end of the test:
		var change *ChangeQueueItem
		callback := func(ctx context.Context, item *ChangeQueueItem) {
			change = item
		}
		queue, err = NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Name("my_changes").
			Table("my_data").
			Install(true).
			Callback(callback).
			Interval(100 * time.Millisecond).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = queue.Close()
			Expect(err).ToNot(HaveOccurred())
		}()

		// Wait a bit so that the changes can be processed:
		time.Sleep(150 * time.Millisecond)

		// Check the change:
		Expect(change).ToNot(BeNil())
		Expect(change.Serial).To(Equal(1))
		Expect(change.Source).To(Equal("my_data"))
		Expect(change.Operation).To(Equal("insert"))
		Expect(change.Old).To(BeEmpty())
		Expect(change.New).To(MatchJSON(`{
			"id": 123,
			"name": "my_name"
		}`))
	})

	It("Process changes made after creation", func() {
		// Create the queue:
		var change *ChangeQueueItem
		callback := func(ctx context.Context, item *ChangeQueueItem) {
			change = item
		}
		queue, err := NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Name("my_changes").
			Table("my_data").
			Install(true).
			Callback(callback).
			Interval(100 * time.Millisecond).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = queue.Close()
			Expect(err).ToNot(HaveOccurred())
		}()

		// Do a change:
		_, err = dbHandle.Exec(`insert into my_data (id, name) values (123, 'my_name')`)
		Expect(err).ToNot(HaveOccurred())

		// Wait a bit so that the changes can be processed:
		time.Sleep(150 * time.Millisecond)

		// Check the change:
		Expect(change).ToNot(BeNil())
		Expect(change.Serial).To(Equal(1))
		Expect(change.Source).To(Equal("my_data"))
		Expect(change.Operation).To(Equal("insert"))
		Expect(change.Old).To(BeEmpty())
		Expect(change.New).To(MatchJSON(`{
			"id": 123,
			"name": "my_name"
		}`))
	})

	It("Passes custom context to callback", func() {
		// Create the queue:
		queue, err := NewChangeQueue().
			Logger(logger).
			URL(dbURL).
			Name("my_changes").
			Table("my_data").
			Install(true).
			Context(func() context.Context {
				// nolint
				return context.WithValue(context.Background(), "my_key", "my_value")
			}).
			Callback(func(ctx context.Context, item *ChangeQueueItem) {
				defer GinkgoRecover()
				key := ctx.Value("my_key")
				Expect(key).To(Equal("my_value"))
			}).
			Interval(100 * time.Millisecond).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = queue.Close()
			Expect(err).ToNot(HaveOccurred())
		}()

		// Do a change:
		_, err = dbHandle.Exec(`insert into my_data (id, name) values (123, 'my_name')`)
		Expect(err).ToNot(HaveOccurred())

		// Wait a bit so that the changes can be processed:
		time.Sleep(150 * time.Millisecond)
	})
})
