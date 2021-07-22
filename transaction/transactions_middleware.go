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

// This file contains an HTTP middleware that associates a transaction to each request.
package transaction

import (
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/errors"
)

// TransactionResponseWriter is a wrapper for http.ResponseWriter,
// it attempts to resolve (commit or rollback) the transaction on the context of the request
// if the transaction failed for some reason it catches the error and returns a 500 response
// with a clear error indicating the failure.
type TransactionResponseWriter struct {
	http.ResponseWriter
	request        *http.Request
	transactionErr error
	headerWritten  bool
}

// Header returns the https response header.
func (t *TransactionResponseWriter) Header() http.Header {
	return t.ResponseWriter.Header()
}

// WriteHeader writes the status code to the response.
// In this implementation we DO NOT write the actual status code to the
// underlying header to avoid multiple response.WriteHeader calls error
// but rather we save the status code and write it back only if the transaction
// resolved successfully.
func (t *TransactionResponseWriter) WriteHeader(statusCode int) {
	// TODO: Note that this approach of resolving the transaction when the application has
	// started writing the response isn't generally correct. It works because all our handlers
	// currently complete the database activities before starting to write the response. But
	// that may change and then we would potentially be committing unfinished transactions.
	t.headerWritten = true
	u, err := FromContext(t.request.Context())
	if err != nil {
		t.transactionErr = err
	}
	if err = u.Resolve(); err != nil {
		t.transactionErr = err
	}
	if t.transactionErr != nil {
		t.ResponseWriter.Header().Set("Content-Type", "application/json")
		t.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	t.ResponseWriter.WriteHeader(statusCode)
}

func (t *TransactionResponseWriter) Write(payload []byte) (int, error) {
	// Some applications may start writing the response body before calling the `WriteHeader`
	// method. That is completely OK according to the documentation. But if we don't call it
	// explicitly the transaction will never be resolved and the connection will be leaked.
	if !t.headerWritten {
		t.WriteHeader(http.StatusOK)
	}
	if t.transactionErr != nil {
		errors.SendInternalServerError(t.ResponseWriter, t.request)
		return
	}
	return t.ResponseWriter.Write(payload)
}

func (t *TransactionResponseWriter) WriteErrIfNothingWritten(r *http.Request) {
	if !t.headerWritten {
		errors.SendInternalServerError(t.ResponseWriter, r)
	}
}

// TransactionMiddleware creates a new HTTP middleware that begins a database transaction and stores
// it in the request context.
func TransactionMiddleware(next http.Handler, manager *Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new Context with the transaction stored in it.
		u, err := manager.Begin()
		if err != nil {
			errors.SendInternalServerError(w, r)
		}
		ctx, err := IntoContext(u, r.Context())
		if err != nil {
			errors.SendInternalServerError(w, r)
			return
		}

		// Make a copy of the request context with the new context key,value stored in it.
		reqWithTx := r.WithContext(ctx)

		txWriter := TransactionResponseWriter{
			ResponseWriter: w,
			request:        reqWithTx,
			headerWritten:  false,
		}

		defer func() {
			if f := recover(); f != nil {
				manager.RecoverPanic(reqWithTx.Context(), f)
				txWriter.WriteErrIfNothingWritten(r)
			}
		}()

		// Continue handling requests.
		next.ServeHTTP(&txWriter, reqWithTx)
	})
}
