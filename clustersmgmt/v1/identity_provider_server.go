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

// IdentityProviderServer represents the interface the manages the 'identity_provider' resource.
type IdentityProviderServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the identity provider.
	Delete(ctx context.Context, request *IdentityProviderDeleteServerRequest, response *IdentityProviderDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the identity provider.
	Get(ctx context.Context, request *IdentityProviderGetServerRequest, response *IdentityProviderGetServerResponse) error
}

// IdentityProviderDeleteServerRequest is the request for the 'delete' method.
type IdentityProviderDeleteServerRequest struct {
}

// IdentityProviderDeleteServerResponse is the response for the 'delete' method.
type IdentityProviderDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *IdentityProviderDeleteServerResponse) Status(value int) *IdentityProviderDeleteServerResponse {
	r.status = value
	return r
}

// IdentityProviderGetServerRequest is the request for the 'get' method.
type IdentityProviderGetServerRequest struct {
}

// IdentityProviderGetServerResponse is the response for the 'get' method.
type IdentityProviderGetServerResponse struct {
	status int
	err    *errors.Error
	body   *IdentityProvider
}

// Body sets the value of the 'body' parameter.
//
//
func (r *IdentityProviderGetServerResponse) Body(value *IdentityProvider) *IdentityProviderGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *IdentityProviderGetServerResponse) Status(value int) *IdentityProviderGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *IdentityProviderGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// IdentityProviderAdapter represents the structs that adapts Requests and Response to internal
// structs.
type IdentityProviderAdapter struct {
	server IdentityProviderServer
	router *mux.Router
}

func NewIdentityProviderAdapter(server IdentityProviderServer, router *mux.Router) *IdentityProviderAdapter {
	adapter := new(IdentityProviderAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodDelete).Path("").HandlerFunc(adapter.handlerDelete)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *IdentityProviderAdapter) readDeleteRequest(r *http.Request) (*IdentityProviderDeleteServerRequest, error) {
	var err error
	result := new(IdentityProviderDeleteServerRequest)
	return result, err
}
func (a *IdentityProviderAdapter) writeDeleteResponse(w http.ResponseWriter, r *IdentityProviderDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *IdentityProviderAdapter) handlerDelete(w http.ResponseWriter, r *http.Request) {
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
	response := new(IdentityProviderDeleteServerResponse)
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
func (a *IdentityProviderAdapter) readGetRequest(r *http.Request) (*IdentityProviderGetServerRequest, error) {
	var err error
	result := new(IdentityProviderGetServerRequest)
	return result, err
}
func (a *IdentityProviderAdapter) writeGetResponse(w http.ResponseWriter, r *IdentityProviderGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *IdentityProviderAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(IdentityProviderGetServerResponse)
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
func (a *IdentityProviderAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
