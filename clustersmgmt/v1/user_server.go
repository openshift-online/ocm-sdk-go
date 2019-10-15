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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// UserServer represents the interface the manages the 'user' resource.
type UserServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the user.
	Delete(ctx context.Context, request *UserDeleteServerRequest, response *UserDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the user.
	Get(ctx context.Context, request *UserGetServerRequest, response *UserGetServerResponse) error
}

// UserDeleteServerRequest is the request for the 'delete' method.
type UserDeleteServerRequest struct {
}

// UserDeleteServerResponse is the response for the 'delete' method.
type UserDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *UserDeleteServerResponse) Status(value int) *UserDeleteServerResponse {
	r.status = value
	return r
}

// UserGetServerRequest is the request for the 'get' method.
type UserGetServerRequest struct {
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

// Status sets the status code.
func (r *UserGetServerResponse) Status(value int) *UserGetServerResponse {
	r.status = value
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

// UserAdapter represents the structs that adapts Requests and Response to internal
// structs.
type UserAdapter struct {
	server UserServer
	router *mux.Router
}

func NewUserAdapter(server UserServer, router *mux.Router) *UserAdapter {
	adapter := new(UserAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodDelete).Path("").HandlerFunc(adapter.handlerDelete)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *UserAdapter) readDeleteRequest(r *http.Request) (*UserDeleteServerRequest, error) {
	var err error
	result := new(UserDeleteServerRequest)
	return result, err
}
func (a *UserAdapter) writeDeleteResponse(w http.ResponseWriter, r *UserDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *UserAdapter) handlerDelete(w http.ResponseWriter, r *http.Request) {
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
	response := new(UserDeleteServerResponse)
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
func (a *UserAdapter) readGetRequest(r *http.Request) (*UserGetServerRequest, error) {
	var err error
	result := new(UserGetServerRequest)
	return result, err
}
func (a *UserAdapter) writeGetResponse(w http.ResponseWriter, r *UserGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *UserAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(UserGetServerResponse)
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
func (a *UserAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
