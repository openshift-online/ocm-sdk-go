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

package transaction

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	// nolint
	. "github.com/onsi/ginkgo"
	// nolint
	. "github.com/onsi/gomega"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

// Sends a request through TransactionMiddleware -> nextHandler, returns response info.
func invokeMiddleware(nextHandler http.Handler) (httpCode int, responseBody string) {
	handlerToTest := TransactionMiddleware(nextHandler)
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()
	handlerToTest.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	Expect(err).ToNot(HaveOccurred())
	return response.StatusCode, string(body)
}

var _ = Describe("Transaction middleware", func() {
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var err error
		var db *sql.DB
		db, mock, err = sqlmock.New()
		Expect(err).ToNot(HaveOccurred())
		err = transaction.SetDB(db)
		Expect(err).ToNot(HaveOccurred())
	})

	It("Commits when all is good", func() {
		mock.ExpectBegin()
		mock.ExpectCommit()

		code, responseBody := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			Expect(w.Write([]byte("{}"))).To(Equal(2))
		}))
		Expect(code).To(Equal(http.StatusOK))
		Expect(responseBody).To(Equal("{}"))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	// TODO: https://issues.redhat.com/browse/SDA-1925
	// The middleware should not assume WriteHeader() is called,
	// the ResponseWriter contract allows skipping it and directly calling Write().
	XIt("Commits when WriteHeader() was skipped", func() {
		mock.ExpectBegin()
		mock.ExpectCommit()

		code, responseBody := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(w.Write([]byte("{}"))).To(Equal(2))
		}))
		Expect(code).To(Equal(http.StatusOK))
		Expect(responseBody).To(Equal("{}"))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("Sends body correctly when Write() called more then once", func() {
		mock.ExpectBegin()
		mock.ExpectCommit()

		code, responseBody := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			Expect(w.Write([]byte("{"))).To(Equal(1))
			Expect(w.Write([]byte("..."))).To(Equal(3))
			Expect(w.Write([]byte("}"))).To(Equal(1))
		}))
		Expect(code).To(Equal(http.StatusOK))
		Expect(responseBody).To(Equal("{...}"))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("Rolls back when marked for rollback", func() {
		mock.ExpectBegin()
		mock.ExpectRollback()

		code, responseBody := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			MarkForRollback(r.Context())
			w.WriteHeader(http.StatusTeapot)
			Expect(w.Write([]byte("I'm a teapot"))).To(Equal(12))
		}))
		Expect(code).To(Equal(http.StatusTeapot))
		Expect(responseBody).To(Equal("I'm a teapot"))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("Rolls back and returns 500 if no header written and a panic occurred", func() {
		mock.ExpectBegin()
		mock.ExpectRollback()

		code, _ := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("Something went wrong ....")
		}))

		Expect(code).To(Equal(http.StatusInternalServerError))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("Returns original header written if a header was written before a panic occurred", func() {
		mock.ExpectBegin()
		mock.ExpectCommit()

		code, _ := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			panic("Something went wrong ....")
		}))

		Expect(code).To(Equal(http.StatusOK))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("Returns 500 when DB is unreachable", func() {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("[simulated] connect: timeout"))

		code, responseBody := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Fail("next handler should not run")
		}))

		Expect(code).To(Equal(http.StatusInternalServerError))
		Expect(responseBody).To(ContainSubstring("CLUSTERS-MGMT-500"))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("Returns 500 when commit fails", func() {
		mock.ExpectBegin()
		mock.ExpectCommit().WillReturnError(fmt.Errorf("[simulated] uniqueness constraint violation"))

		code, responseBody := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			Expect(w.Write([]byte("{}"))).To(BeNumerically(">", 0))
		}))

		Expect(code).To(Equal(http.StatusInternalServerError))
		Expect(responseBody).To(ContainSubstring("CLUSTERS-MGMT-500"))
		Expect(responseBody).To(ContainSubstring("uniqueness constraint violation"))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("Returns 500 when rollback fails", func() {
		mock.ExpectBegin()
		mock.ExpectRollback().WillReturnError(fmt.Errorf("[simulated] connect: timeout"))

		code, responseBody := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			transaction.MarkForRollback(r.Context())
			w.WriteHeader(http.StatusTeapot)
			Expect(w.Write([]byte("I'm a teapot"))).To(BeNumerically(">", 0))
		}))

		Expect(code).To(Equal(http.StatusInternalServerError))
		Expect(responseBody).To(ContainSubstring("CLUSTERS-MGMT-500"))
		Expect(responseBody).To(ContainSubstring("connect: timeout"))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	// TODO: https://issues.redhat.com/browse/SDA-1925
	// The middleware should not assume Write() is called exactly once,
	// the ResponseWriter contract allows calling it many times.
	XIt("Sends single error JSON when commit fails and Write() called more than once", func() {
		mock.ExpectBegin()
		mock.ExpectCommit().WillReturnError(fmt.Errorf("[simulated] uniqueness constraint violation"))

		code, responseBody := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			Expect(w.Write([]byte("{"))).To(BeNumerically(">", 0))
			Expect(w.Write([]byte("..."))).To(BeNumerically(">", 0))
			Expect(w.Write([]byte("}"))).To(BeNumerically(">", 0))
		}))

		Expect(code).To(Equal(http.StatusInternalServerError))
		Expect(strings.Count(responseBody, "CLUSTERS-MGMT-500")).To(Equal(1), "count of error JSONs")
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("Commits transaction even if handler doesn't call `WriteHeader`", func() {
		mock.ExpectBegin()
		mock.ExpectCommit()
		code, body := invokeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Note that this writes some body without explicitly writing the headers.
			_, err := w.Write([]byte("{}"))
			Expect(err).ToNot(HaveOccurred())
		}))
		Expect(code).To(Equal(http.StatusOK))
		Expect(body).To(Equal("{}"))
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})
})
