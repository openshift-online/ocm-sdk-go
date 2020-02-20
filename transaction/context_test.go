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

// This file contains tests for the transaction context.

package transaction

import (
	"context"
	"database/sql"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Context", func() {
	// This will contain the database handle that will be used by all the tests.
	var db *sql.DB

	// This will contain the transaction manager that will be used to create the transactions
	// in all the tests:
	var manager *Manager

	BeforeEach(func() {
		var err error

		// Create the database in a temporary file that will be deleted when the database
		// handle is closed:
		db, err = sql.Open("sqlite3", "")
		Expect(err).ToNot(HaveOccurred())

		// Create the manager:
		manager, err = NewManager().
			Logger(logger).
			DB(db).
			Build()
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

	It("Can add and then get a transaction to a context", func() {
		// Create the context:
		ctx := context.Background()

		// Create the transaction:
		object, err := manager.Begin(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Add the transaction to the context:
		ctx = ToContext(ctx, object)

		// Get the transaction from the context:
		object, err = FromContext(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(object).ToNot(BeNil())
	})

	It("Fails if there is not transaction in the context", func() {
		// Create the context:
		ctx := context.Background()

		// Get the transaction from the context:
		object, err := FromContext(ctx)
		Expect(err).To(HaveOccurred())
		Expect(object).To(BeNil())
	})
})
