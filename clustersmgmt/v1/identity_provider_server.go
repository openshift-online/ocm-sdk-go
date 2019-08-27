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
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// IdentityProviderServer represents the interface the manages the 'identity_provider' resource.
type IdentityProviderServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the identity provider.
	Get(ctx context.Context, request *IdentityProviderGetServerRequest, response *IdentityProviderGetServerResponse) error

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the identity provider.
	Delete(ctx context.Context, request *IdentityProviderDeleteServerRequest, response *IdentityProviderDeleteServerResponse) error
}

// IdentityProviderGetServerRequest is the request for the 'get' method.
type IdentityProviderGetServerRequest struct {
	path  string
	query url.Values
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

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *IdentityProviderGetServerResponse) SetStatusCode(status int) *IdentityProviderGetServerResponse {
	r.status = status
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

// IdentityProviderDeleteServerRequest is the request for the 'delete' method.
type IdentityProviderDeleteServerRequest struct {
	path  string
	query url.Values
}

// IdentityProviderDeleteServerResponse is the response for the 'delete' method.
type IdentityProviderDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *IdentityProviderDeleteServerResponse) SetStatusCode(status int) *IdentityProviderDeleteServerResponse {
	r.status = status
	return r
}

// IdentityProviderServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type IdentityProviderServerAdapter struct {
	server IdentityProviderServer
	router *mux.Router
}

func NewIdentityProviderServerAdapter(server IdentityProviderServer, router *mux.Router) *IdentityProviderServerAdapter {
	adapter := new(IdentityProviderServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").HandlerFunc(adapter.getHandler)
	adapter.router.Methods("DELETE").HandlerFunc(adapter.deleteHandler)
	return adapter
}
func (a *IdentityProviderServerAdapter) readIdentityProviderGetServerRequest(r *http.Request) (*IdentityProviderGetServerRequest, error) {
	result := new(IdentityProviderGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *IdentityProviderServerAdapter) writeIdentityProviderGetServerResponse(w http.ResponseWriter, r *IdentityProviderGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *IdentityProviderServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readIdentityProviderGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(IdentityProviderGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeIdentityProviderGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *IdentityProviderServerAdapter) readIdentityProviderDeleteServerRequest(r *http.Request) (*IdentityProviderDeleteServerRequest, error) {
	result := new(IdentityProviderDeleteServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *IdentityProviderServerAdapter) writeIdentityProviderDeleteServerResponse(w http.ResponseWriter, r *IdentityProviderDeleteServerResponse) error {
	w.WriteHeader(r.status)
	return nil
}
func (a *IdentityProviderServerAdapter) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readIdentityProviderDeleteServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(IdentityProviderDeleteServerResponse)
	err = a.server.Delete(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Delete: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeIdentityProviderDeleteServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *IdentityProviderServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
