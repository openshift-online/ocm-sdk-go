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

// PermissionServer represents the interface the manages the 'permission' resource.
type PermissionServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the permission.
	Get(ctx context.Context, request *PermissionGetServerRequest, response *PermissionGetServerResponse) error

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the permission.
	Delete(ctx context.Context, request *PermissionDeleteServerRequest, response *PermissionDeleteServerResponse) error
}

// PermissionGetServerRequest is the request for the 'get' method.
type PermissionGetServerRequest struct {
	path  string
	query url.Values
}

// PermissionGetServerResponse is the response for the 'get' method.
type PermissionGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Permission
}

// Body sets the value of the 'body' parameter.
//
//
func (r *PermissionGetServerResponse) Body(value *Permission) *PermissionGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *PermissionGetServerResponse) SetStatusCode(status int) *PermissionGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *PermissionGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// PermissionDeleteServerRequest is the request for the 'delete' method.
type PermissionDeleteServerRequest struct {
	path  string
	query url.Values
}

// PermissionDeleteServerResponse is the response for the 'delete' method.
type PermissionDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *PermissionDeleteServerResponse) SetStatusCode(status int) *PermissionDeleteServerResponse {
	r.status = status
	return r
}

// PermissionServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type PermissionServerAdapter struct {
	server PermissionServer
	router *mux.Router
}

func NewPermissionServerAdapter(server PermissionServer, router *mux.Router) *PermissionServerAdapter {
	adapter := new(PermissionServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	adapter.router.Methods("DELETE").Path("").HandlerFunc(adapter.deleteHandler)
	return adapter
}
func (a *PermissionServerAdapter) readPermissionGetServerRequest(r *http.Request) (*PermissionGetServerRequest, error) {
	result := new(PermissionGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *PermissionServerAdapter) writePermissionGetServerResponse(w http.ResponseWriter, r *PermissionGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *PermissionServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readPermissionGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(PermissionGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writePermissionGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *PermissionServerAdapter) readPermissionDeleteServerRequest(r *http.Request) (*PermissionDeleteServerRequest, error) {
	result := new(PermissionDeleteServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *PermissionServerAdapter) writePermissionDeleteServerResponse(w http.ResponseWriter, r *PermissionDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *PermissionServerAdapter) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readPermissionDeleteServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(PermissionDeleteServerResponse)
	err = a.server.Delete(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Delete: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writePermissionDeleteServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *PermissionServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
