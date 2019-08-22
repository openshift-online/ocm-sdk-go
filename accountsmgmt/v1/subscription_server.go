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

package v1 // github.com/openshift-online/uhc-sdk-go/accountsmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/uhc-sdk-go/errors"
)

// SubscriptionServer represents the interface the manages the 'subscription' resource.
type SubscriptionServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the subscription.
	Get(ctx context.Context, request *SubscriptionGetServerRequest, response *SubscriptionGetServerResponse) error

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the subscription.
	Delete(ctx context.Context, request *SubscriptionDeleteServerRequest, response *SubscriptionDeleteServerResponse) error
}

// SubscriptionGetServerRequest is the request for the 'get' method.
type SubscriptionGetServerRequest struct {
	path  string
	query url.Values
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

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *SubscriptionGetServerResponse) SetStatusCode(status int) *SubscriptionGetServerResponse {
	r.status = status
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

// SubscriptionDeleteServerRequest is the request for the 'delete' method.
type SubscriptionDeleteServerRequest struct {
	path  string
	query url.Values
}

// SubscriptionDeleteServerResponse is the response for the 'delete' method.
type SubscriptionDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *SubscriptionDeleteServerResponse) SetStatusCode(status int) *SubscriptionDeleteServerResponse {
	r.status = status
	return r
}

// SubscriptionServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type SubscriptionServerAdapter struct {
	server SubscriptionServer
	router *mux.Router
}

func NewSubscriptionServerAdapter(server SubscriptionServer, router *mux.Router) *SubscriptionServerAdapter {
	adapter := new(SubscriptionServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.HandleFunc("/", adapter.getHandler).Methods("GET")
	adapter.router.HandleFunc("/", adapter.deleteHandler).Methods("DELETE")
	return adapter
}
func (a *SubscriptionServerAdapter) readSubscriptionGetServerRequest(r *http.Request) (*SubscriptionGetServerRequest, error) {
	result := new(SubscriptionGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *SubscriptionServerAdapter) writeSubscriptionGetServerResponse(w http.ResponseWriter, r *SubscriptionGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *SubscriptionServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readSubscriptionGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(SubscriptionGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeSubscriptionGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *SubscriptionServerAdapter) readSubscriptionDeleteServerRequest(r *http.Request) (*SubscriptionDeleteServerRequest, error) {
	result := new(SubscriptionDeleteServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *SubscriptionServerAdapter) writeSubscriptionDeleteServerResponse(w http.ResponseWriter, r *SubscriptionDeleteServerResponse) error {
	w.WriteHeader(r.status)
	return nil
}
func (a *SubscriptionServerAdapter) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readSubscriptionDeleteServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(SubscriptionDeleteServerResponse)
	err = a.server.Delete(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Delete: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeSubscriptionDeleteServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *SubscriptionServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
