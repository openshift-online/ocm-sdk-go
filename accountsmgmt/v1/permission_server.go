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

// PermissionServer represents the interface the manages the 'permission' resource.
type PermissionServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the permission.
	Delete(ctx context.Context, request *PermissionDeleteServerRequest, response *PermissionDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the permission.
	Get(ctx context.Context, request *PermissionGetServerRequest, response *PermissionGetServerResponse) error
}

// PermissionDeleteServerRequest is the request for the 'delete' method.
type PermissionDeleteServerRequest struct {
}

// PermissionDeleteServerResponse is the response for the 'delete' method.
type PermissionDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *PermissionDeleteServerResponse) Status(value int) *PermissionDeleteServerResponse {
	r.status = value
	return r
}

// PermissionGetServerRequest is the request for the 'get' method.
type PermissionGetServerRequest struct {
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

// Status sets the status code.
func (r *PermissionGetServerResponse) Status(value int) *PermissionGetServerResponse {
	r.status = value
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

// PermissionAdapter represents the structs that adapts Requests and Response to internal
// structs.
type PermissionAdapter struct {
	server PermissionServer
	router *mux.Router
}

func NewPermissionAdapter(server PermissionServer, router *mux.Router) *PermissionAdapter {
	adapter := new(PermissionAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodDelete).Path("").HandlerFunc(adapter.handlerDelete)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *PermissionAdapter) readDeleteRequest(r *http.Request) (*PermissionDeleteServerRequest, error) {
	var err error
	result := new(PermissionDeleteServerRequest)
	return result, err
}
func (a *PermissionAdapter) writeDeleteResponse(w http.ResponseWriter, r *PermissionDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *PermissionAdapter) handlerDelete(w http.ResponseWriter, r *http.Request) {
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
	response := new(PermissionDeleteServerResponse)
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
func (a *PermissionAdapter) readGetRequest(r *http.Request) (*PermissionGetServerRequest, error) {
	var err error
	result := new(PermissionGetServerRequest)
	return result, err
}
func (a *PermissionAdapter) writeGetResponse(w http.ResponseWriter, r *PermissionGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *PermissionAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(PermissionGetServerResponse)
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
func (a *PermissionAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
