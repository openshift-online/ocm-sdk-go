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

package v1 // github.com/openshift-online/uhc-sdk-go/clustersmgmt/v1

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

// UserServer represents the interface the manages the 'user' resource.
type UserServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the user.
	Get(ctx context.Context, request *UserGetServerRequest, response *UserGetServerResponse) error

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the user.
	Delete(ctx context.Context, request *UserDeleteServerRequest, response *UserDeleteServerResponse) error
}

// UserGetServerRequest is the request for the 'get' method.
type UserGetServerRequest struct {
	path  string
	query url.Values
}

// UserGetServerResponse is the response for the 'get' method.
type UserGetServerResponse struct {
	status int
	err    *errors.Error
	body   *User
}

// Body sets the value of the 'body' parameter.
//
//
func (r *UserGetServerResponse) Body(value *User) *UserGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *UserGetServerResponse) SetStatusCode(status int) *UserGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *UserGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// UserDeleteServerRequest is the request for the 'delete' method.
type UserDeleteServerRequest struct {
	path  string
	query url.Values
}

// UserDeleteServerResponse is the response for the 'delete' method.
type UserDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *UserDeleteServerResponse) SetStatusCode(status int) *UserDeleteServerResponse {
	r.status = status
	return r
}

// UserServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type UserServerAdapter struct {
	server UserServer
	router *mux.Router
}

func NewUserServerAdapter(server UserServer, router *mux.Router) *UserServerAdapter {
	adapter := new(UserServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.HandleFunc("/", adapter.getHandler).Methods("GET")
	adapter.router.HandleFunc("/", adapter.deleteHandler).Methods("DELETE")
	return adapter
}
func (a *UserServerAdapter) readUserGetServerRequest(r *http.Request) (*UserGetServerRequest, error) {
	result := new(UserGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *UserServerAdapter) writeUserGetServerResponse(w http.ResponseWriter, r *UserGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *UserServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readUserGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(UserGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeUserGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *UserServerAdapter) readUserDeleteServerRequest(r *http.Request) (*UserDeleteServerRequest, error) {
	result := new(UserDeleteServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *UserServerAdapter) writeUserDeleteServerResponse(w http.ResponseWriter, r *UserDeleteServerResponse) error {
	w.WriteHeader(r.status)
	return nil
}
func (a *UserServerAdapter) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readUserDeleteServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(UserDeleteServerResponse)
	err = a.server.Delete(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Delete: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeUserDeleteServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *UserServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
