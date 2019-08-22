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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
)

// RegistryServer represents the interface the manages the 'registry' resource.
type RegistryServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the registry.
	Get(ctx context.Context, request *RegistryGetServerRequest, response *RegistryGetServerResponse) error
}

// RegistryGetServerRequest is the request for the 'get' method.
type RegistryGetServerRequest struct {
	path  string
	query url.Values
}

// RegistryGetServerResponse is the response for the 'get' method.
type RegistryGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Registry
}

// Body sets the value of the 'body' parameter.
//
//
func (r *RegistryGetServerResponse) Body(value *Registry) *RegistryGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RegistryGetServerResponse) SetStatusCode(status int) *RegistryGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *RegistryGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RegistryServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RegistryServerAdapter struct {
	server RegistryServer
	router *mux.Router
}

func NewRegistryServerAdapter(server RegistryServer, router *mux.Router) *RegistryServerAdapter {
	adapter := new(RegistryServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.HandleFunc("/", adapter.getHandler).Methods("GET")
	return adapter
}
func (a *RegistryServerAdapter) readRegistryGetServerRequest(r *http.Request) (*RegistryGetServerRequest, error) {
	result := new(RegistryGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *RegistryServerAdapter) writeRegistryGetServerResponse(w http.ResponseWriter, r *RegistryGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RegistryServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRegistryGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RegistryGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRegistryGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RegistryServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
