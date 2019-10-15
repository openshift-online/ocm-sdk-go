/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// SubscriptionServer represents the interface the manages the 'subscription' resource.
type SubscriptionServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the subscription.
	Delete(ctx context.Context, request *SubscriptionDeleteServerRequest, response *SubscriptionDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the subscription.
	Get(ctx context.Context, request *SubscriptionGetServerRequest, response *SubscriptionGetServerResponse) error
}

// SubscriptionDeleteServerRequest is the request for the 'delete' method.
type SubscriptionDeleteServerRequest struct {
}

// SubscriptionDeleteServerResponse is the response for the 'delete' method.
type SubscriptionDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *SubscriptionDeleteServerResponse) Status(value int) *SubscriptionDeleteServerResponse {
	r.status = value
	return r
}

// SubscriptionGetServerRequest is the request for the 'get' method.
type SubscriptionGetServerRequest struct {
}

// SubscriptionGetServerResponse is the response for the 'get' method.
type SubscriptionGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Subscription
}

// Body sets the value of the 'body' parameter.
//
//
func (r *SubscriptionGetServerResponse) Body(value *Subscription) *SubscriptionGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *SubscriptionGetServerResponse) Status(value int) *SubscriptionGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *SubscriptionGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// SubscriptionAdapter represents the structs that adapts Requests and Response to internal
// structs.
type SubscriptionAdapter struct {
	server SubscriptionServer
	router *mux.Router
}

func NewSubscriptionAdapter(server SubscriptionServer, router *mux.Router) *SubscriptionAdapter {
	adapter := new(SubscriptionAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodDelete).Path("").HandlerFunc(adapter.handlerDelete)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *SubscriptionAdapter) readDeleteRequest(r *http.Request) (*SubscriptionDeleteServerRequest, error) {
	var err error
	result := new(SubscriptionDeleteServerRequest)
	return result, err
}
func (a *SubscriptionAdapter) writeDeleteResponse(w http.ResponseWriter, r *SubscriptionDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *SubscriptionAdapter) handlerDelete(w http.ResponseWriter, r *http.Request) {
	request, err := a.readDeleteRequest(r)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to read request from client: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
		return
	}
	response := new(SubscriptionDeleteServerResponse)
	response.status = http.StatusOK
	err = a.server.Delete(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method Delete: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeDeleteResponse(w, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to write response for client: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
}
func (a *SubscriptionAdapter) readGetRequest(r *http.Request) (*SubscriptionGetServerRequest, error) {
	var err error
	result := new(SubscriptionGetServerRequest)
	return result, err
}
func (a *SubscriptionAdapter) writeGetResponse(w http.ResponseWriter, r *SubscriptionGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *SubscriptionAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
	request, err := a.readGetRequest(r)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to read request from client: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
		return
	}
	response := new(SubscriptionGetServerResponse)
	response.status = http.StatusOK
	err = a.server.Get(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method Get: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeGetResponse(w, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to write response for client: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
}
func (a *SubscriptionAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
