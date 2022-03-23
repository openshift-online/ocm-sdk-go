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

package leadership

import (
	"database/sql"
	"time"

	. "github.com/onsi/ginkgo/v2/dsl/core"                // nolint
	. "github.com/onsi/gomega"                            // nolint
	. "github.com/openshift-online/ocm-sdk-go/v2/testing" // nolint
)

var _ = Describe("Flag behaviour", func() {
	var dbObject *Database
	var dbHandle *sql.DB

	var CreateTable = func() {
		_, err := dbHandle.Exec(`
			create table leadership_flags (
				name text not null primary key,
				holder text not null,
				version bigint not null,
				timestamp timestamp with time zone not null
			)
		`)
		Expect(err).ToNot(HaveOccurred())
	}

	BeforeEach(func() {
		// Create a database:
		dbObject = dbServer.MakeDatabase()
		dbHandle = dbObject.MakeHandle()
	})

	AfterEach(func() {
		dbObject.Close()
	})

	It("Can't be created without a logger", func() {
		_, err := NewFlag().
			Handle(dbHandle).
			Name("my_flag").
			Process("my_process").
			Build()
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("logger"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Can't be created without a database handle", func() {
		_, err := NewFlag().
			Logger(logger).
			Name("my_flag").
			Process("my_process").
			Build()
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("database"))
		Expect(message).To(ContainSubstring("handle"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Can't be created without a name", func() {
		_, err := NewFlag().
			Logger(logger).
			Handle(dbHandle).
			Process("my_process").
			Build()
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("name"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Can't be created without a process name", func() {
		_, err := NewFlag().
			Logger(logger).
			Handle(dbHandle).
			Name("my_flag").
			Build()
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("process"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Creates the table if it doesn't exist", func() {
		// Create the flag:
		flag, err := NewFlag().
			Logger(logger).
			Handle(dbHandle).
			Name("my_flag").
			Process("my_process").
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = flag.Close()
			Expect(err).ToNot(HaveOccurred())
		}()

		// Check that the table exists:
		rows, err := dbHandle.Query(`
			select
				name,
				holder,
				version,
				timestamp
			from
				leadership_flags
		`)
		Expect(err).ToNot(HaveOccurred())
		err = rows.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can be created if the table already exists", func() {
		// Create the database table:
		CreateTable()

		// Create the flag object:
		flag, err := NewFlag().
			Logger(logger).
			Handle(dbHandle).
			Name("my_flag").
			Process("my_process").
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = flag.Close()
			Expect(err).ToNot(HaveOccurred())
		}()
	})

	When("Doesn't exist", func() {
		var flag *Flag

		BeforeEach(func() {
			var err error

			// Create the flag object:
			flag, err = NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("my_process").
				Interval(200 * time.Millisecond).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			err := flag.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		It("It is quickly raised ", func() {
			time.Sleep(40 * time.Millisecond)
			Expect(flag.Raised()).To(BeTrue())
		})
	})

	When("Is held by another process and not expired", func() {
		var flag *Flag

		BeforeEach(func() {
			var err error

			// Create the database table:
			CreateTable()

			// Create the database row so that the flag is already held by another
			// process that will eventually fail to update it:
			_, err = dbHandle.Exec(`
				insert into leadership_flags (
					name,
					holder,
					version,
					timestamp
				) values (
					'my_flag',
					'your_process',
					123,
					now()
				)
			`)
			Expect(err).ToNot(HaveOccurred())

			// Create the object:
			flag, err = NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("my_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			err := flag.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		It("It isn't quickly raised", func() {
			time.Sleep(40 * time.Millisecond)
			Expect(flag.Raised()).To(BeFalse())
		})

		It("It isn't raised while the previous holder can still renew", func() {
			time.Sleep(100 * time.Millisecond)
			Expect(flag.Raised()).To(BeFalse())
		})

		It("It is raised when the previous holder fails to renew", func() {
			time.Sleep(400 * time.Millisecond)
			Expect(flag.Raised()).To(BeTrue())
		})
	})

	When("Is held by another process but already expired", func() {
		var flag *Flag

		BeforeEach(func() {
			var err error

			// Create the database table:
			CreateTable()

			// Create the database row so that the flag is already held by another
			// process that already failed to renew it:
			_, err = dbHandle.Exec(`
				insert into leadership_flags (
					name,
					holder,
					version,
					timestamp
				) values (
					'my_flag',
					'your_process',
					123,
					now() - interval '1 second'
				)
			`)
			Expect(err).ToNot(HaveOccurred())

			// Create the object:
			flag, err = NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("my_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			err := flag.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		It("It is quickly raised", func() {
			time.Sleep(40 * time.Millisecond)
			Expect(flag.Raised()).To(BeTrue())
		})
	})

	When("Is held by this process and not expired", func() {
		var flag *Flag

		BeforeEach(func() {
			var err error

			// Create the database table:
			CreateTable()

			// Create the database row so that the flag is already held by this process:
			_, err = dbHandle.Exec(`
				insert into leadership_flags (
					name,
					holder,
					version,
					timestamp
				) values (
					'my_flag',
					'my_process',
					123,
					now()
				)
			`)
			Expect(err).ToNot(HaveOccurred())

			// Create the object:
			flag, err = NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("my_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			err := flag.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		It("It is quickly raised", func() {
			time.Sleep(40 * time.Millisecond)
			Expect(flag.Raised()).To(BeTrue())
		})
	})

	When("Is held by this process and already expired", func() {
		var flag *Flag

		BeforeEach(func() {
			var err error

			// Create the database table:
			CreateTable()

			// Create the database row so that the flag is held by this process but
			// expired:
			_, err = dbHandle.Exec(`
				insert into leadership_flags (
					name,
					holder,
					version,
					timestamp
				) values (
					'my_flag',
					'my_process',
					123,
					now() - interval '1 second'
				)
			`)
			Expect(err).ToNot(HaveOccurred())

			// Create the object:
			flag, err = NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("my_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			err := flag.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		It("It isn't quickly raised", func() {
			time.Sleep(40 * time.Millisecond)
			Expect(flag.Raised()).To(BeTrue())
		})
	})

	When("Current holder closes the flag", func() {
		It("Is raised by another process", func() {
			var err error

			// Create the first process:
			first, err := NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("first_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Give the first process some time to get hold of the flag and then check
			// that it did:
			time.Sleep(200 * time.Millisecond)
			Expect(first.Raised()).To(BeTrue())

			// Create the second process:
			second, err := NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("second_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = second.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Give the second process some time to try to get hold of the flag and
			// check that it didn't:
			time.Sleep(200 * time.Millisecond)
			Expect(second.Raised()).To(BeFalse())

			// Close the first process so that it will fail to renew the flag:
			err = first.Close()
			Expect(err).ToNot(HaveOccurred())

			// Allow time for the second process to get hold of the flag and check that
			// it did:
			time.Sleep(400 * time.Millisecond)
			Expect(second.Raised()).To(BeTrue())
		})
	})

	When("Current holder loses database connection", func() {
		It("Is raised by another process", func() {
			var err error

			// Create the first process, but using a separate database handle, so that
			// we can close it without affecting the second process:
			altHandle := dbObject.MakeHandle()
			first, err := NewFlag().
				Logger(logger).
				Handle(altHandle).
				Name("my_flag").
				Process("first_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = first.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Give the first process some time to get hold of the flag and then check
			// that it did:
			time.Sleep(200 * time.Millisecond)
			Expect(first.Raised()).To(BeTrue())

			// Create the second process:
			second, err := NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("second_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = second.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Give the second process some time to try to get hold of the flag and then
			// check that it didn't:
			time.Sleep(200 * time.Millisecond)
			Expect(second.Raised()).To(BeFalse())

			// Close the database connection of the first process:
			err = altHandle.Close()
			Expect(err).ToNot(HaveOccurred())

			// Allow time for the second process to get hold of the flag and then check
			// that it did:
			time.Sleep(400 * time.Millisecond)
			Expect(second.Raised()).To(BeTrue())
		})
	})

	When("Stolen", func() {
		It("Is lowers", func() {
			var err error

			// Create the flag:
			flag, err := NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("my_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = flag.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Give the process some time to get hold of the flag and check that it did:
			time.Sleep(200 * time.Millisecond)
			Expect(flag.Raised()).To(BeTrue())

			// Steal the flag updating the database directly:
			_, err = dbHandle.Exec(`
				update
					leadership_flags
				set
					holder = 'your_process',
					version = version + 1,
					timestamp = now()
				where
					name = 'my_flag'
			`)
			Expect(err).ToNot(HaveOccurred())

			// Give the process some time to detect the situation and then check that it
			// lowers the flag:
			time.Sleep(200 * time.Millisecond)
			Expect(flag.Raised()).To(BeFalse())
		})
	})

	When("Stolen and then expired", func() {
		It("Is recovers and raises", func() {
			var err error

			// Create the flag:
			flag, err := NewFlag().
				Logger(logger).
				Handle(dbHandle).
				Name("my_flag").
				Process("my_process").
				Interval(200 * time.Millisecond).
				Jitter(0).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = flag.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Give the process some time to get hold of the flag and then check that
			// it did:
			time.Sleep(200 * time.Millisecond)
			Expect(flag.Raised()).To(BeTrue())

			// Force another holder updating the database directly:
			_, err = dbHandle.Exec(`
				update
					leadership_flags
				set
					holder = 'your_process',
					version = version + 1,
					timestamp = now()
				where
					name = 'my_flag'
			`)
			Expect(err).ToNot(HaveOccurred())

			// Give the process some time to detect the situation and check that it
			// lowers the flag:
			time.Sleep(200 * time.Millisecond)
			Expect(flag.Raised()).To(BeFalse())

			// Give it more time, so that the forced holder fails to renew and then check
			// that it recovers and raises it:
			time.Sleep(200 * time.Millisecond)
			Expect(flag.Raised()).To(BeTrue())
		})
	})
})

var _ = Describe("Flag metrics enabled", func() {
	var dbObject *Database
	var dbHandle *sql.DB
	var metricsServer *MetricsServer

	BeforeEach(func() {
		// Create a database:
		dbObject = dbServer.MakeDatabase()
		dbHandle = dbObject.MakeHandle()

		// Create the metrics server:
		metricsServer = NewMetricsServer()
	})

	AfterEach(func() {
		// Delete the database:
		dbObject.Close()

		// Stop the metrics server:
		metricsServer.Close()
	})

	It("Generates state metrics", func() {
		// Create the first process:
		first, err := NewFlag().
			Logger(logger).
			Handle(dbHandle).
			Name("my_flag").
			Process("first_process").
			Interval(200 * time.Millisecond).
			Jitter(0).
			MetricsSubsystem("my").
			MetricsRegisterer(metricsServer.Registry()).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = first.Close()
			Expect(err).ToNot(HaveOccurred())
		}()

		// Give it time to raise:
		time.Sleep(200 * time.Millisecond)
		Expect(first.Raised()).To(BeTrue())

		// Create the second process:
		second, err := NewFlag().
			Logger(logger).
			Handle(dbHandle).
			Name("my_flag").
			Process("second_process").
			Interval(200 * time.Millisecond).
			Jitter(0).
			MetricsSubsystem("my").
			MetricsRegisterer(metricsServer.Registry()).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = second.Close()
			Expect(err).ToNot(HaveOccurred())
		}()

		// Git it time to lower:
		time.Sleep(200 * time.Millisecond)
		Expect(second.Raised()).To(BeFalse())

		// Verify the metrics:
		metrics := metricsServer.Metrics()
		Expect(metrics).To(MatchLine(`^my_leadership_flag_state\{name="my_flag",process="first_process"\} 1$`))
		Expect(metrics).To(MatchLine(`^my_leadership_flag_state\{name="my_flag",process="second_process"\} 0$`))
	})
})
