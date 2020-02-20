/*
Copyright (c) 2020 Red Hat, Inc.

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

// This file contains tests for the transaction manager.

package transaction

import (
	"context"
	"database/sql"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	// This will contain the database handle that will be used by all the tests.
	var db *sql.DB

	BeforeEach(func() {
		var err error

		// Create the database in a temporary file that will be deleted when the database
		// handle is closed:
		db, err = sql.Open("sqlite3", "")
		Expect(err).ToNot(HaveOccurred())

		// Create the tables:
		_, err = db.Exec(`CREATE TABLE mytable (mycolumn TEXT)`)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		var err error

		// Close the database. This also deletes the temporary file where it resides.
		if db != nil {
			err = db.Close()
			Expect(err).ToNot(HaveOccurred())
		}
	})

	It("Can't be built without a logger", func() {
		_, err := NewManager().
			DB(db).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("logger"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Can't be built without a database handle", func() {
		_, err := NewManager().
			Logger(logger).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("database"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Can't be built with a logger and a database handle", func() {
		_, err := NewManager().
			Logger(logger).
			DB(db).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can begin and complete an empty transaction", func() {
		// Create a context:
		ctx := context.Background()

		// Create the manager:
		manager, err := NewManager().
			Logger(logger).
			DB(db).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Begin the transaction:
		object, err := manager.Begin(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(object).ToNot(BeNil())
		defer func() {
			err := manager.Complete(ctx, object)
			Expect(err).ToNot(HaveOccurred())
		}()
	})

	It("Can commit a transaction", func() {
		// Create a context:
		ctx := context.Background()

		// Create the manager:
		manager, err := NewManager().
			Logger(logger).
			DB(db).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Begin the transaction:
		object, err := manager.Begin(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Write to the database:
		_, err = object.TX().Exec(`INSERT INTO mytable (mycolumn) VALUES ('myvalue')`)
		Expect(err).ToNot(HaveOccurred())

		// Complete the transaction:
		err = manager.Complete(ctx, object)
		Expect(err).ToNot(HaveOccurred())

		// Check that the write succeeded:
		row := db.QueryRow(`SELECT mycolumn FROM mytable`)
		var value string
		err = row.Scan(&value)
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal("myvalue"))
	})

	It("Can rollback a transaction", func() {
		// Create a context:
		ctx := context.Background()

		// Create the manager:
		manager, err := NewManager().
			Logger(logger).
			DB(db).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Begin the transaction:
		object, err := manager.Begin(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Write to the database:
		_, err = object.TX().Exec(`INSERT INTO mytable (mycolumn) VALUES ('myvalue')`)
		Expect(err).ToNot(HaveOccurred())

		// Mark the transaction for rollback:
		object.MarkForRollback()

		// Complete the transaction:
		err = manager.Complete(ctx, object)
		Expect(err).ToNot(HaveOccurred())

		// Check that the write was rolled back:
		row := db.QueryRow(`SELECT COUNT(*) FROM mytable`)
		var count int
		err = row.Scan(&count)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(BeZero())
	})

	It("Executes post commit callback", func() {
		// Create a context:
		ctx := context.Background()

		// Create the manager:
		manager, err := NewManager().
			Logger(logger).
			DB(db).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Begin the transaction:
		object, err := manager.Begin(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Add a post commit callback that should be executed ant that will notify us when
		// it finishes by closing a channel:
		done := make(chan struct{})
		object.AddPostCommitCallback(func() {
			defer GinkgoRecover()
			close(done)
		})

		// Add a post rollback callback that should not be executed:
		object.AddPostRollbackCallback(func() {
			defer GinkgoRecover()
			Expect(true).To(BeFalse())
		})

		// Complete the transaction:
		err = manager.Complete(ctx, object)
		Expect(err).ToNot(HaveOccurred())

		// Give the callback some time to finish:
		timer := time.NewTimer(1 * time.Millisecond)
		select {
		case <-done:
		case <-timer.C:
			Expect(false).To(BeTrue())
		}
	})

	It("Executes post rollback callback", func() {
		// Create a context:
		ctx := context.Background()

		// Create the manager:
		manager, err := NewManager().
			Logger(logger).
			DB(db).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Begin the transaction:
		object, err := manager.Begin(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Mark the transaction for callback:
		object.MarkForRollback()

		// Add a post rollback callback that should be executed ant that will notify us when
		// it finishes by closing a channel:
		done := make(chan struct{})
		object.AddPostRollbackCallback(func() {
			defer GinkgoRecover()
			close(done)
		})

		// Add a post commit callback that should not be executed:
		object.AddPostCommitCallback(func() {
			defer GinkgoRecover()
			Expect(true).To(BeFalse())
		})

		// Complete the transaction:
		err = manager.Complete(ctx, object)
		Expect(err).ToNot(HaveOccurred())

		// Give the callback some time to finish:
		timer := time.NewTimer(1 * time.Millisecond)
		select {
		case <-done:
		case <-timer.C:
			Expect(false).To(BeTrue())
		}
	})
})
