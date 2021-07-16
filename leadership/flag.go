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

// This file contains functions useful for printing lists in a format that is easy to read for
// humans in log files.

package leadership

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/openshift-online/ocm-sdk-go/database"
	"github.com/openshift-online/ocm-sdk-go/logging"
	"github.com/prometheus/client_golang/prometheus"
)

// FlagBuilder contains the data and logic needed to build leadership flags.
type FlagBuilder struct {
	// Basic fields:
	logger   logging.Logger
	handle   *sql.DB
	name     string
	process  string
	interval time.Duration
	timeout  time.Duration
	jitter   float64

	// Fields used for metrics:
	metricsSubsystem  string
	metricsRegisterer prometheus.Registerer
}

// Flag is a distributed flag intended to manage leadership in a group of processes. Only one of the
// processes using it will see it raised at any point in time.
type Flag struct {
	// Basic fields:
	logger        logging.Logger
	handle        *sql.DB
	name          string
	process       string
	renewInterval time.Duration
	checkInterval time.Duration
	retryInterval time.Duration
	timeout       time.Duration
	jitter        float64
	value         int32
	timer         *time.Timer
	stop          chan struct{}

	// Fields used for metrics:
	stateMetric *prometheus.GaugeVec
}

// NewFlag creates a builder that can then be used to configure and create a leadership flag.
func NewFlag() *FlagBuilder {
	return &FlagBuilder{
		interval:          defaultFlagInterval,
		timeout:           defaultFlagTimeout,
		jitter:            defaultFlagJitter,
		metricsRegisterer: prometheus.DefaultRegisterer,
	}
}

// Logger sets the logger that the flag will use to write to the log. This is mandatory.
func (b *FlagBuilder) Logger(value logging.Logger) *FlagBuilder {
	b.logger = value
	return b
}

// Handle sets the database handle that the flag will use to store its state. This is mandatory.
func (b *FlagBuilder) Handle(value *sql.DB) *FlagBuilder {
	b.handle = value
	return b
}

// Name of the flag. This can be used to have different flags for different uses, or for different
// environments that happen to share the database. This is mandatory.
func (b *FlagBuilder) Name(value string) *FlagBuilder {
	b.name = value
	return b
}

// Process sets the name of the process. This should be unique amonts the set of processes using the
// same flag name. A typical name would be the name of a Kubernetes pod, or the combination of a
// Kubernets cluser name and pod name, to make it unique across different clusters. This is
// mandatory.
func (b *FlagBuilder) Process(value string) *FlagBuilder {
	b.process = value
	return b
}

// Interval sets the interval for renewing the ownership of the flag. The default value is thirty
// seconds.
func (b *FlagBuilder) Interval(value time.Duration) *FlagBuilder {
	b.interval = value
	return b
}

// Timeout sets the timeout for database operations. The default is on second.
func (b *FlagBuilder) Timeout(value time.Duration) *FlagBuilder {
	b.timeout = value
	return b
}

// Jitter sets a factor that will be used to randomize the intervals. For example, if this is set to
// 0.1 then a random adjustment of +10% or -10% will be done to the intervals each time they are
// used. This is intended to reduce simultaneous database accesses by processes that have been
// started simultaneously. The default value is 0.2.
func (b *FlagBuilder) Jitter(value float64) *FlagBuilder {
	b.jitter = value
	return b
}

// MetricsSubsystem sets the name of the subsystem that will be used by the flag to register metrics
// with Prometheus. If this isn't explicitly specified, or if it is an empty string, then no metrics
// will be registered. For example, if the value is `background_tasks` then the following metrics
// will be registered:
//
//	tasks_leadership_flag_state - State of the flag.
//
// The `...leadership_flag_state` metric will have the following labels:
//
//	name - Name of the flag.
//	process - Name of the process.
//
// The value of the `...leaderhsip_flag_state` metric will be one if this process is currently the
// holder of the flag or zero if it isn't.
//
// Note that setting this attribute is not enough to have metrics published, you also need to
// create and start a metrics server, as described in the documentation of the Prometheus library.
func (b *FlagBuilder) MetricsSubsystem(value string) *FlagBuilder {
	b.metricsSubsystem = value
	return b
}

// MetricsRegisterer sets the Prometheus registerer that will be used to register the metrics. The
// default is to use the default Prometheus registerer and there is usually no need to change that.
// This is intended for unit tests, where it is convenient to have a registerer that doesn't
// interfere with the rest of the system.
func (b *FlagBuilder) MetricsRegisterer(value prometheus.Registerer) *FlagBuilder {
	if value == nil {
		value = prometheus.DefaultRegisterer
	}
	b.metricsRegisterer = value
	return b
}

// Build uses the data stored in the builder to configure and create a new leadership flag.
func (b *FlagBuilder) Build(ctx context.Context) (result *Flag, err error) {
	// Check parameters:
	if b.logger == nil {
		err = errors.New("logger is mandatory")
		return
	}
	if b.handle == nil {
		err = errors.New("database handle is mandatory")
		return
	}
	if b.name == "" {
		err = errors.New("name is mandatory")
		return
	}
	if b.process == "" {
		err = errors.New("process is mandatory")
		return
	}
	if b.interval <= 0 {
		err = errors.New("interval should be greater than zero")
		return
	}
	if b.timeout <= 0 {
		err = errors.New("timeout should be greater than zero")
		return
	}
	if b.jitter < 0 || b.jitter > 1 {
		err = errors.New("jitter should be between zero and one")
		return
	}

	// Calculate specific intervals from the general interval given in the configuration:
	renewInterval := b.interval
	checkInterval := b.interval / 2
	retryInterval := b.interval / 10

	// Make sure that the table exists, creating it if needed:
	err = b.ensureTable(ctx)
	if err != nil {
		return
	}

	// Create a timer that will fire inmediatelly, so that the first check will be performed
	// also inmediately after starting the loop:
	timer := time.NewTimer(0)

	// Crete the channel that will be used to stop the loop:
	stop := make(chan struct{})

	// Register the metrics:
	var stateMetric *prometheus.GaugeVec
	if b.metricsSubsystem != "" && b.metricsRegisterer != nil {
		stateMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: b.metricsSubsystem,
				Name:      "leadership_flag_state",
				Help: "State of the leadership flag; one if raised, zero " +
					"if lowered.",
			},
			flagMetricsLabels,
		)
		err = b.metricsRegisterer.Register(stateMetric)
		if err != nil {
			registered, ok := err.(prometheus.AlreadyRegisteredError)
			if ok {
				stateMetric = registered.ExistingCollector.(*prometheus.GaugeVec)
				err = nil
			} else {
				return
			}
		}
	}

	// Create and populate the flag:
	result = &Flag{
		logger:        b.logger,
		handle:        b.handle,
		name:          b.name,
		process:       b.process,
		timeout:       b.timeout,
		renewInterval: renewInterval,
		checkInterval: checkInterval,
		retryInterval: retryInterval,
		jitter:        b.jitter,
		timer:         timer,
		stop:          stop,
		stateMetric:   stateMetric,
	}

	// Run the loop:
	go result.run()

	return
}

// ensureTable creates the table if it doesn't already exist.
func (b *FlagBuilder) ensureTable(ctx context.Context) error {
	var err error
	_, err = b.handle.ExecContext(
		ctx,
		`
		create table if not exists leadership_flags (
			name text not null primary key,
			holder text not null,
			version bigint not null,
			timestamp timestamp with time zone not null
		)
		`,
	)
	return err
}

// Raised returns true if the flag is raised. At any point in time only one of the identities will
// see the flag raised.
func (f *Flag) Raised() bool {
	return atomic.LoadInt32(&f.value) == 1
}

// Close releases all the resources used by the flag.
func (f *Flag) Close() error {
	close(f.stop)
	return nil
}

// run runs the loop that checks the contents of the table and updates it and the state of the flag
// accordinly.
func (f *Flag) run() {
loop:
	for {
		select {
		case <-f.timer.C:
			f.check()
		case <-f.stop:
			break loop
		}
	}
}

// check checks the contents of the table and updates it and the state of the flag accordingly.
func (f *Flag) check() {
	var err error

	// Create a context:
	ctx := context.Background()

	// Get the global time from the database, so that we don't depend on synchronization of the
	// machines that compete for the flag.
	var now time.Time
	now, err = f.now(ctx)
	if err != nil {
		f.logger.Error(
			ctx,
			"Process '%s' can't get current time for flag '%s': %v",
			f.process, f.name, err,
		)
		f.lower(ctx)
		f.schedule(ctx, f.retryInterval)
		return
	}

	// Try to load the state:
	found, holder, version, timestamp, err := f.loadState(ctx)
	if err != nil {
		f.logger.Error(
			ctx,
			"Process '%s' can't load state for flag '%s': %v",
			f.process, f.name, err,
		)
		f.lower(ctx)
		f.schedule(ctx, f.retryInterval)
		return
	}

	// If the state doesn't exist yet then try to create it:
	if !found {
		var created bool
		created, err = f.createState(ctx, now)
		if err != nil {
			f.logger.Error(
				ctx,
				"Process '%s' can't create initial state for flag: %v",
				f.process, f.name, err,
			)
			f.lower(ctx)
			f.schedule(ctx, f.retryInterval)
			return
		}
		if !created {
			f.logger.Debug(
				ctx,
				"Process '%s' found a conflict when trying to create the initial "+
					"state for flag '%s'",
				f.process, f.name,
			)
			f.lower(ctx)
			f.schedule(ctx, f.checkInterval)
			return
		}
		f.logger.Info(
			ctx,
			"Process '%s' successfully created initial state for flag '%s'",
			f.process, f.name,
		)
		f.raise(ctx)
		f.schedule(ctx, f.checkInterval)
		return
	}

	// If we are here then the state already existed and we were able to load it. If we are
	// the current holder then we should extend the renew time and make sure that the flag is
	// raised.
	if holder == f.process {
		var updated bool
		updated, err = f.updateTimestamp(ctx, version, now)
		if err != nil {
			f.logger.Error(
				ctx,
				"Process '%s' can't update the timestamp for flag '%s': %v",
				f.process, f.name, err,
			)
			f.lower(ctx)
			f.schedule(ctx, f.retryInterval)
			return
		}
		if !updated {
			f.logger.Info(
				ctx,
				"Process '%s' found a conflict when trying to update the "+
					"timestamp for flag '%s'",
				f.process, f.name,
			)
			f.lower(ctx)
			f.schedule(ctx, f.checkInterval)
			return
		}
		f.logger.Debug(
			ctx,
			"Process '%s' successfully updated the timestamp for flag '%s'",
			f.process, f.name,
		)
		f.raise(ctx)
		f.schedule(ctx, f.checkInterval)
		return
	}

	// If we aren't the holder then we should check the timestamp and try to become the leader
	// if it hasn't been updated recently enough:
	excess := now.Sub(timestamp) - f.renewInterval
	if excess > 0 {
		f.logger.Info(
			ctx,
			"Process '%s' detected that flag '%s' is currently held by process '%s' "+
				"but it should have been renewed %s ago, will try to get hold "+
				"of it",
			f.process, f.name, holder, excess,
		)
		var updated bool
		updated, err = f.updateHolder(ctx, version, now)
		if err != nil {
			f.logger.Error(
				ctx,
				"Process '%s' can't update holder for flag '%s': %v",
				f.process, f.name,
			)
			f.lower(ctx)
			f.schedule(ctx, f.retryInterval)
			return
		}
		if !updated {
			f.logger.Info(
				ctx,
				"Process '%s' found a conflict when trying to update the holder "+
					"for flag '%s'",
				f.process, f.name,
			)
			f.lower(ctx)
			f.schedule(ctx, f.checkInterval)
			return
		}
		f.logger.Debug(
			ctx,
			"Process '%s' successfully updated holder for flag '%s'",
			f.process, f.name,
		)
		f.raise(ctx)
		f.schedule(ctx, f.checkInterval)
		return
	}

	// If we are here we aren't the holder, and the renew time isn't expired, so all we should
	// do is check again later:
	f.logger.Debug(
		ctx,
		"Process '%s' found that flag '%s' is currently held by process '%s' and it "+
			"should be renewed in %s",
		f.process, f.name, holder, -excess,
	)
	f.lower(ctx)
	f.schedule(ctx, f.checkInterval)
}

// schedule programs the timer so that it fires in the given time from now.
func (f *Flag) schedule(ctx context.Context, d time.Duration) {
	// Adjust the given duration adding or subtracting a random amount. For example, if the
	// random factor given in the configuration is 0.1 will add or sustract up to a 10% of the
	// duration. This is convenient to avoid having all the process doing their checks
	// simultaneously when they have been started simultaneously.
	factor := f.jitter * (1 - 2*rand.Float64())
	delta := time.Duration(float64(d) * factor)
	d += delta

	// Reset the timer:
	f.logger.Debug(ctx, "Process '%s' will check flag '%s' in %s", f.process, f.name, d)
	f.timer.Reset(d)
}

// now returns get the current time from the database, so that there is no need to synchornize the
// clocks of the machines that compete for the flag.
func (f *Flag) now(ctx context.Context) (result time.Time, err error) {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()
	row := f.handle.QueryRowContext(ctx, `select now()`)
	var tmp time.Time
	err = row.Scan(&tmp)
	if err != nil {
		return
	}
	result = tmp
	return
}

// loadState tries to load the database state corresponding to this flag. It returns a flag
// indicating if the state was found and the values.
func (f *Flag) loadState(ctx context.Context) (found bool, holder string, version int64,
	timestamp time.Time, err error) {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()
	row := f.handle.QueryRowContext(
		ctx,
		`
		select
			holder,
			version,
			timestamp
		from
			leadership_flags
		where
			name = $1
		`,
		f.name,
	)
	if err != nil {
		return
	}
	var tmpHolder string
	var tmpVersion int64
	var tmpTimestamp time.Time
	err = row.Scan(
		&tmpHolder,
		&tmpVersion,
		&tmpTimestamp,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return
	}
	found = true
	holder = tmpHolder
	version = tmpVersion
	timestamp = tmpTimestamp
	return
}

// createState tries to save the initial state of the flag. It returns a boolean indicating if the
// state was actually created.
func (f *Flag) createState(ctx context.Context, timestamp time.Time) (created bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()
	_, err = f.handle.ExecContext(
		ctx,
		`
		insert into leadership_flags (
			name,
			holder,
			version,
			timestamp
		) values (
			$1,
			$2,
			0,
			$3
		)
		`,
		f.name,
		f.process,
		timestamp,
	)
	if err != nil {
		// 23505 is the code corresponding to `unique_violation` condition.
		if database.ErrorCode(err) == "23505" {
			err = nil
		}
		return
	}
	created = true
	return
}

// updateTimestamp tries to update the timestamp.
func (f *Flag) updateTimestamp(ctx context.Context, version int64, timestamp time.Time) (updated bool,
	err error) {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()
	result, err := f.handle.ExecContext(
		ctx,
		`
		update
			leadership_flags
		set
			version = $1,
			timestamp = $2
		where
			name = $3 and
			holder = $4 and
			version = $5
		`,
		version+1,
		timestamp,
		f.name,
		f.process,
		version,
	)
	if err != nil {
		return
	}
	count, err := result.RowsAffected()
	if err != nil {
		return
	}
	updated = count == 1
	return
}

// updateHolder tries to update the holder.
func (f *Flag) updateHolder(ctx context.Context, version int64, timestamp time.Time) (updated bool,
	err error) {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()
	result, err := f.handle.ExecContext(
		ctx,
		`
		update
			leadership_flags
		set
			version = $1,
			holder = $2,
			timestamp = $3
		where
			name = $4 and
			version = $5
		`,
		version+1,
		f.process,
		timestamp,
		f.name,
		version,
	)
	if err != nil {
		return
	}
	count, err := result.RowsAffected()
	if err != nil {
		return
	}
	updated = count == 1
	return
}

// raise raises the flag locally, without touching the database.
func (f *Flag) raise(ctx context.Context) {
	old := atomic.SwapInt32(&f.value, 1)
	if old == 0 {
		f.logger.Debug(
			ctx,
			"Process '%s' is now holding flag '%s'",
			f.process, f.name,
		)
	}
	if f.stateMetric != nil {
		f.stateMetric.WithLabelValues(f.name, f.process).Set(1)
	}
}

// lower lowers the flag locally, without touching the database.
func (f *Flag) lower(ctx context.Context) {
	old := atomic.SwapInt32(&f.value, 0)
	if old == 1 {
		f.logger.Debug(
			ctx,
			"Process '%s' is no longer holding flag '%s'",
			f.process, f.name,
		)
	}
	if f.stateMetric != nil {
		f.stateMetric.WithLabelValues(f.name, f.process).Set(0)
	}
}

// Defaults for configuration settings:
const (
	defaultFlagInterval = 30 * time.Second
	defaultFlagTimeout  = 1 * time.Second
	defaultFlagJitter   = 0.2
)

// Names of the labels added to the metrics:
const (
	flagMetricsNameLabel    = "name"
	flagMetricsProcessLabel = "process"
)

// Array of labels added to metrics:
var flagMetricsLabels = []string{
	flagMetricsNameLabel,
	flagMetricsProcessLabel,
}
