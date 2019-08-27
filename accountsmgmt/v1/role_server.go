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
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// RoleServer represents the interface the manages the 'role' resource.
type RoleServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the role.
	Get(ctx context.Context, request *RoleGetServerRequest, response *RoleGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	// Updates the role.
	Update(ctx context.Context, request *RoleUpdateServerRequest, response *RoleUpdateServerResponse) error

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the role.
	Delete(ctx context.Context, request *RoleDeleteServerRequest, response *RoleDeleteServerResponse) error
}

// RoleGetServerRequest is the request for the 'get' method.
type RoleGetServerRequest struct {
	path  string
	query url.Values
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

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RoleGetServerResponse) SetStatusCode(status int) *RoleGetServerResponse {
	r.status = status
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
	path  string
	query url.Values
	body  *Role
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
	body   *Role
}

// Body sets the value of the 'body' parameter.
//
//
func (r *RoleUpdateServerResponse) Body(value *Role) *RoleUpdateServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RoleUpdateServerResponse) SetStatusCode(status int) *RoleUpdateServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'update' method.
func (r *RoleUpdateServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RoleDeleteServerRequest is the request for the 'delete' method.
type RoleDeleteServerRequest struct {
	path  string
	query url.Values
}

// RoleDeleteServerResponse is the response for the 'delete' method.
type RoleDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RoleDeleteServerResponse) SetStatusCode(status int) *RoleDeleteServerResponse {
	r.status = status
	return r
}

// RoleServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RoleServerAdapter struct {
	server RoleServer
	router *mux.Router
}

func NewRoleServerAdapter(server RoleServer, router *mux.Router) *RoleServerAdapter {
	adapter := new(RoleServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	adapter.router.Methods("PATCH").Path("").HandlerFunc(adapter.updateHandler)
	adapter.router.Methods("DELETE").Path("").HandlerFunc(adapter.deleteHandler)
	return adapter
}
func (a *RoleServerAdapter) readRoleGetServerRequest(r *http.Request) (*RoleGetServerRequest, error) {
	result := new(RoleGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *RoleServerAdapter) writeRoleGetServerResponse(w http.ResponseWriter, r *RoleGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RoleServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRoleGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RoleGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRoleGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RoleServerAdapter) readRoleUpdateServerRequest(r *http.Request) (*RoleUpdateServerRequest, error) {
	result := new(RoleUpdateServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *RoleServerAdapter) writeRoleUpdateServerResponse(w http.ResponseWriter, r *RoleUpdateServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RoleServerAdapter) updateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRoleUpdateServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RoleUpdateServerResponse)
	err = a.server.Update(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Update: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRoleUpdateServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RoleServerAdapter) readRoleDeleteServerRequest(r *http.Request) (*RoleDeleteServerRequest, error) {
	result := new(RoleDeleteServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *RoleServerAdapter) writeRoleDeleteServerResponse(w http.ResponseWriter, r *RoleDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *RoleServerAdapter) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRoleDeleteServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RoleDeleteServerResponse)
	err = a.server.Delete(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Delete: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRoleDeleteServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RoleServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
