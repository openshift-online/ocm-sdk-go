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

// RoleBindingServer represents the interface the manages the 'role_binding' resource.
type RoleBindingServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the role binding.
	Get(ctx context.Context, request *RoleBindingGetServerRequest, response *RoleBindingGetServerResponse) error

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the role binding.
	Delete(ctx context.Context, request *RoleBindingDeleteServerRequest, response *RoleBindingDeleteServerResponse) error
}

// RoleBindingGetServerRequest is the request for the 'get' method.
type RoleBindingGetServerRequest struct {
	path  string
	query url.Values
}

// RoleBindingGetServerResponse is the response for the 'get' method.
type RoleBindingGetServerResponse struct {
	status int
	err    *errors.Error
	body   *RoleBinding
}

// Body sets the value of the 'body' parameter.
//
//
func (r *RoleBindingGetServerResponse) Body(value *RoleBinding) *RoleBindingGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RoleBindingGetServerResponse) SetStatusCode(status int) *RoleBindingGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *RoleBindingGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RoleBindingDeleteServerRequest is the request for the 'delete' method.
type RoleBindingDeleteServerRequest struct {
	path  string
	query url.Values
}

// RoleBindingDeleteServerResponse is the response for the 'delete' method.
type RoleBindingDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RoleBindingDeleteServerResponse) SetStatusCode(status int) *RoleBindingDeleteServerResponse {
	r.status = status
	return r
}

// RoleBindingServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RoleBindingServerAdapter struct {
	server RoleBindingServer
	router *mux.Router
}

func NewRoleBindingServerAdapter(server RoleBindingServer, router *mux.Router) *RoleBindingServerAdapter {
	adapter := new(RoleBindingServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	adapter.router.Methods("DELETE").Path("").HandlerFunc(adapter.deleteHandler)
	return adapter
}
func (a *RoleBindingServerAdapter) readRoleBindingGetServerRequest(r *http.Request) (*RoleBindingGetServerRequest, error) {
	var err error
	result := new(RoleBindingGetServerRequest)
	result.query = r.URL.Query()
	result.path = r.URL.Path
	return result, err
}
func (a *RoleBindingServerAdapter) writeRoleBindingGetServerResponse(w http.ResponseWriter, r *RoleBindingGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RoleBindingServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRoleBindingGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RoleBindingGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRoleBindingGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RoleBindingServerAdapter) readRoleBindingDeleteServerRequest(r *http.Request) (*RoleBindingDeleteServerRequest, error) {
	var err error
	result := new(RoleBindingDeleteServerRequest)
	result.query = r.URL.Query()
	result.path = r.URL.Path
	return result, err
}
func (a *RoleBindingServerAdapter) writeRoleBindingDeleteServerResponse(w http.ResponseWriter, r *RoleBindingDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *RoleBindingServerAdapter) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRoleBindingDeleteServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RoleBindingDeleteServerResponse)
	err = a.server.Delete(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Delete: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRoleBindingDeleteServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RoleBindingServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
