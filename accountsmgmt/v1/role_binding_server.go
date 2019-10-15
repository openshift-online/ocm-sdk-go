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

// RoleBindingServer represents the interface the manages the 'role_binding' resource.
type RoleBindingServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the role binding.
	Delete(ctx context.Context, request *RoleBindingDeleteServerRequest, response *RoleBindingDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the role binding.
	Get(ctx context.Context, request *RoleBindingGetServerRequest, response *RoleBindingGetServerResponse) error
}

// RoleBindingDeleteServerRequest is the request for the 'delete' method.
type RoleBindingDeleteServerRequest struct {
}

// RoleBindingDeleteServerResponse is the response for the 'delete' method.
type RoleBindingDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *RoleBindingDeleteServerResponse) Status(value int) *RoleBindingDeleteServerResponse {
	r.status = value
	return r
}

// RoleBindingGetServerRequest is the request for the 'get' method.
type RoleBindingGetServerRequest struct {
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

// Status sets the status code.
func (r *RoleBindingGetServerResponse) Status(value int) *RoleBindingGetServerResponse {
	r.status = value
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

// RoleBindingAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RoleBindingAdapter struct {
	server RoleBindingServer
	router *mux.Router
}

func NewRoleBindingAdapter(server RoleBindingServer, router *mux.Router) *RoleBindingAdapter {
	adapter := new(RoleBindingAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodDelete).Path("").HandlerFunc(adapter.handlerDelete)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *RoleBindingAdapter) readDeleteRequest(r *http.Request) (*RoleBindingDeleteServerRequest, error) {
	var err error
	result := new(RoleBindingDeleteServerRequest)
	return result, err
}
func (a *RoleBindingAdapter) writeDeleteResponse(w http.ResponseWriter, r *RoleBindingDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *RoleBindingAdapter) handlerDelete(w http.ResponseWriter, r *http.Request) {
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
	response := new(RoleBindingDeleteServerResponse)
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
func (a *RoleBindingAdapter) readGetRequest(r *http.Request) (*RoleBindingGetServerRequest, error) {
	var err error
	result := new(RoleBindingGetServerRequest)
	return result, err
}
func (a *RoleBindingAdapter) writeGetResponse(w http.ResponseWriter, r *RoleBindingGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RoleBindingAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(RoleBindingGetServerResponse)
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
func (a *RoleBindingAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
