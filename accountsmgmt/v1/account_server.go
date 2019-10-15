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

// Status sets the status code.
func (r *AccountGetServerResponse) Status(value int) *AccountGetServerResponse {
	r.status = value
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
	body *Account
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
}

// Status sets the status code.
func (r *AccountUpdateServerResponse) Status(value int) *AccountUpdateServerResponse {
	r.status = value
	return r
}

// AccountAdapter represents the structs that adapts Requests and Response to internal
// structs.
type AccountAdapter struct {
	server AccountServer
	router *mux.Router
}

func NewAccountAdapter(server AccountServer, router *mux.Router) *AccountAdapter {
	adapter := new(AccountAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	adapter.router.Methods(http.MethodPatch).Path("").HandlerFunc(adapter.handlerUpdate)
	return adapter
}
func (a *AccountAdapter) readGetRequest(r *http.Request) (*AccountGetServerRequest, error) {
	var err error
	result := new(AccountGetServerRequest)
	return result, err
}
func (a *AccountAdapter) writeGetResponse(w http.ResponseWriter, r *AccountGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *AccountAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(AccountGetServerResponse)
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
func (a *AccountAdapter) readUpdateRequest(r *http.Request) (*AccountUpdateServerRequest, error) {
	var err error
	result := new(AccountUpdateServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *AccountAdapter) writeUpdateResponse(w http.ResponseWriter, r *AccountUpdateServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *AccountAdapter) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	request, err := a.readUpdateRequest(r)
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
	response := new(AccountUpdateServerResponse)
	response.status = http.StatusOK
	err = a.server.Update(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method Update: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeUpdateResponse(w, response)
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
func (a *AccountAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
