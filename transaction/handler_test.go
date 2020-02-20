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
	"database/sql"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
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

		// Create the tables:
		_, err = db.Exec(`CREATE TABLE mytable (mycolumn TEXT)`)
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

	It("Can't be built without a logger", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Manager(manager).
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("logger"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Can't be built without a manager", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("manager"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Can be built with all the required arguments", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Manager(manager).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can't be built without a manager", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("manager"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Can't be built without a next handler", func() {
		_, err := NewHandler().
			Logger(logger).
			Manager(manager).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("next"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Adds transaction to the request context", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			object, err := FromContext(r.Context())
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Manager(manager).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/whatever", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Commits transaction", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the transaction:
			object, err := FromContext(r.Context())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Write to the database:
			_, err = object.TX().Exec(`INSERT INTO mytable (mycolumn) VALUES ('myvalue')`)
			Expect(err).ToNot(HaveOccurred())

			// Send the response:
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Manager(manager).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/whatever", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))

		// Check that the write succeeded:
		row := db.QueryRow(`SELECT mycolumn FROM mytable`)
		var value string
		err = row.Scan(&value)
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal("myvalue"))
	})

	It("Rollbacks transaction", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the transaction:
			object, err := FromContext(r.Context())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Write to the database:
			_, err = object.TX().Exec(`INSERT INTO mytable (mycolumn) VALUES ('myvalue')`)
			Expect(err).ToNot(HaveOccurred())

			// Mark the transaction for rollback:
			object.MarkForRollback()

			// Send the response:
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Manager(manager).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/whatever", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))

		// Check that the write was rolled back:
		row := db.QueryRow(`SELECT COUNT(*) FROM mytable`)
		var count int
		err = row.Scan(&count)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(BeZero())
	})
})
