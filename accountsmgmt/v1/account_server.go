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

// AccountServer represents the interface the manages the 'account' resource.
type AccountServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the account.
	Get(ctx context.Context, request *AccountGetServerRequest, response *AccountGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	// Updates the account.
	Update(ctx context.Context, request *AccountUpdateServerRequest, response *AccountUpdateServerResponse) error
}

// AccountGetServerRequest is the request for the 'get' method.
type AccountGetServerRequest struct {
	path  string
	query url.Values
}

// AccountGetServerResponse is the response for the 'get' method.
type AccountGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Account
}

// Body sets the value of the 'body' parameter.
//
//
func (r *AccountGetServerResponse) Body(value *Account) *AccountGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *AccountGetServerResponse) SetStatusCode(status int) *AccountGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *AccountGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// AccountUpdateServerRequest is the request for the 'update' method.
type AccountUpdateServerRequest struct {
	path  string
	query url.Values
	body  *Account
}

// Body returns the value of the 'body' parameter.
//
//
func (r *AccountUpdateServerRequest) Body() *Account {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *AccountUpdateServerRequest) GetBody() (value *Account, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'update' method.
func (r *AccountUpdateServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(accountData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}

// AccountUpdateServerResponse is the response for the 'update' method.
type AccountUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *Account
}

// Body sets the value of the 'body' parameter.
//
//
func (r *AccountUpdateServerResponse) Body(value *Account) *AccountUpdateServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *AccountUpdateServerResponse) SetStatusCode(status int) *AccountUpdateServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'update' method.
func (r *AccountUpdateServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// AccountServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type AccountServerAdapter struct {
	server AccountServer
	router *mux.Router
}

func NewAccountServerAdapter(server AccountServer, router *mux.Router) *AccountServerAdapter {
	adapter := new(AccountServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").HandlerFunc(adapter.getHandler)
	adapter.router.Methods("PATCH").HandlerFunc(adapter.updateHandler)
	return adapter
}
func (a *AccountServerAdapter) readAccountGetServerRequest(r *http.Request) (*AccountGetServerRequest, error) {
	result := new(AccountGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *AccountServerAdapter) writeAccountGetServerResponse(w http.ResponseWriter, r *AccountGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *AccountServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readAccountGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(AccountGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeAccountGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *AccountServerAdapter) readAccountUpdateServerRequest(r *http.Request) (*AccountUpdateServerRequest, error) {
	result := new(AccountUpdateServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *AccountServerAdapter) writeAccountUpdateServerResponse(w http.ResponseWriter, r *AccountUpdateServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *AccountServerAdapter) updateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readAccountUpdateServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(AccountUpdateServerResponse)
	err = a.server.Update(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Update: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeAccountUpdateServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *AccountServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
