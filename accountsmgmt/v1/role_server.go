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

// RoleServer represents the interface the manages the 'role' resource.
type RoleServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the role.
	Delete(ctx context.Context, request *RoleDeleteServerRequest, response *RoleDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the role.
	Get(ctx context.Context, request *RoleGetServerRequest, response *RoleGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	// Updates the role.
	Update(ctx context.Context, request *RoleUpdateServerRequest, response *RoleUpdateServerResponse) error
}

// RoleDeleteServerRequest is the request for the 'delete' method.
type RoleDeleteServerRequest struct {
}

// RoleDeleteServerResponse is the response for the 'delete' method.
type RoleDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *RoleDeleteServerResponse) Status(value int) *RoleDeleteServerResponse {
	r.status = value
	return r
}

// RoleGetServerRequest is the request for the 'get' method.
type RoleGetServerRequest struct {
}

// RoleGetServerResponse is the response for the 'get' method.
type RoleGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Role
}

// Body sets the value of the 'body' parameter.
//
//
func (r *RoleGetServerResponse) Body(value *Role) *RoleGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *RoleGetServerResponse) Status(value int) *RoleGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *RoleGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RoleUpdateServerRequest is the request for the 'update' method.
type RoleUpdateServerRequest struct {
	body *Role
}

// Body returns the value of the 'body' parameter.
//
//
func (r *RoleUpdateServerRequest) Body() *Role {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *RoleUpdateServerRequest) GetBody() (value *Role, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'update' method.
func (r *RoleUpdateServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(roleData)
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

// RoleUpdateServerResponse is the response for the 'update' method.
type RoleUpdateServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *RoleUpdateServerResponse) Status(value int) *RoleUpdateServerResponse {
	r.status = value
	return r
}

// RoleAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RoleAdapter struct {
	server RoleServer
	router *mux.Router
}

func NewRoleAdapter(server RoleServer, router *mux.Router) *RoleAdapter {
	adapter := new(RoleAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodDelete).Path("").HandlerFunc(adapter.handlerDelete)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	adapter.router.Methods(http.MethodPatch).Path("").HandlerFunc(adapter.handlerUpdate)
	return adapter
}
func (a *RoleAdapter) readDeleteRequest(r *http.Request) (*RoleDeleteServerRequest, error) {
	var err error
	result := new(RoleDeleteServerRequest)
	return result, err
}
func (a *RoleAdapter) writeDeleteResponse(w http.ResponseWriter, r *RoleDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *RoleAdapter) handlerDelete(w http.ResponseWriter, r *http.Request) {
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
	response := new(RoleDeleteServerResponse)
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
func (a *RoleAdapter) readGetRequest(r *http.Request) (*RoleGetServerRequest, error) {
	var err error
	result := new(RoleGetServerRequest)
	return result, err
}
func (a *RoleAdapter) writeGetResponse(w http.ResponseWriter, r *RoleGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RoleAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(RoleGetServerResponse)
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
func (a *RoleAdapter) readUpdateRequest(r *http.Request) (*RoleUpdateServerRequest, error) {
	var err error
	result := new(RoleUpdateServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *RoleAdapter) writeUpdateResponse(w http.ResponseWriter, r *RoleUpdateServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *RoleAdapter) handlerUpdate(w http.ResponseWriter, r *http.Request) {
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
	response := new(RoleUpdateServerResponse)
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
func (a *RoleAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
