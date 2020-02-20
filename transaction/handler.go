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

package transaction

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/openshift-online/ocm-sdk-go"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// HandlerBuilder contains the data and logic needed to create a new transaction handler. Don't
// create objects of this type directly, use the NewHandler function instead.
type HandlerBuilder struct {
	logger  sdk.Logger
	manager *Manager
	next    http.Handler
}

// Handler is an HTTP handler that creates a new transaction for each inbound request and adds it
// to the context.
type Handler struct {
	logger  sdk.Logger
	manager *Manager
	next    http.Handler
}

// NewHandler creates a builder that can then be configured and used to create transaction handlers.
func NewHandler() *HandlerBuilder {
	return &HandlerBuilder{}
}

// Logger sets the logger that the handler will use to send messages to the log. This is mandatory.
func (b *HandlerBuilder) Logger(value sdk.Logger) *HandlerBuilder {
	b.logger = value
	return b
}

// Manager sets the transaction manager that the handler will use to create and complete
// transactions.
func (b *HandlerBuilder) Manager(value *Manager) *HandlerBuilder {
	b.manager = value
	return b
}

// Next sets the HTTP handler that will be called after the transaction has been created and added
// to the context.
func (b *HandlerBuilder) Next(value http.Handler) *HandlerBuilder {
	b.next = value
	return b
}

// Build uses the data stored in the builder to create a new authentication handler.
func (b *HandlerBuilder) Build() (handler *Handler, err error) {
	// Check parameters:
	if b.logger == nil {
		err = fmt.Errorf("logger is mandatory")
		return
	}
	if b.manager == nil {
		err = fmt.Errorf("transaction manager is mandatory")
		return
	}
	if b.next == nil {
		err = fmt.Errorf("next handler is mandatory")
		return
	}

	// Create and populate the object:
	handler = &Handler{
		logger:  b.logger,
		manager: b.manager,
		next:    b.next,
	}

	return
}

// ServeHTTP is the implementation of the HTTP handler interface.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the context:
	ctx := r.Context()

	// Begin the transaction and remember to complete it regardless of the result of the
	// processing of the HTTP request. This completion should be done even if sending the body
	// is aborted due to a communications error, for example.
	object, err := h.manager.Begin(ctx)
	if err != nil {
		h.logger.Error(ctx, "Can't create transaction: %v", err)
		h.sendError(w, r)
	}
	defer func() {
		err := h.manager.Complete(ctx, object)
		if err != nil {
			h.logger.Error(ctx, "Can't complete transaction: %v", err)
		}
	}()

	// Replace the HTTP request with a new one that contains the transaction:
	r = r.WithContext(ToContext(ctx, object))

	// Call the next handler:
	h.next.ServeHTTP(w, r)
}

// sendError sends an error response to the client.
func (h *Handler) sendError(w http.ResponseWriter, r *http.Request) {
	// Get the context:
	ctx := r.Context()

	// Prepare the body:
	segments := strings.Split(r.URL.Path, "/")
	builder := errors.NewError()
	builder.ID(fmt.Sprintf("%d", http.StatusInternalServerError))
	if len(segments) >= 4 {
		prefix := segments[1]
		service := segments[2]
		version := segments[3]
		builder.HREF(fmt.Sprintf(
			"/%s/%s/%s/errors/%d",
			prefix, service, version, http.StatusInternalServerError,
		))
		builder.Code(fmt.Sprintf(
			"%s-%d",
			strings.ToUpper(strings.ReplaceAll(service, "_", "-")),
			http.StatusUnauthorized,
		))
	}
	builder.Reason("Transaction error")
	body, err := builder.Build()
	if err != nil {
		h.logger.Error(ctx, "Can't build error response: %v", err)
		errors.SendPanic(w, r)
	}

	// Send the response:
	errors.SendError(w, r, body)
}
