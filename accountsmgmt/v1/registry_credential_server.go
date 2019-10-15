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

// RegistryCredentialServer represents the interface the manages the 'registry_credential' resource.
type RegistryCredentialServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the registry credential.
	Get(ctx context.Context, request *RegistryCredentialGetServerRequest, response *RegistryCredentialGetServerResponse) error
}

// RegistryCredentialGetServerRequest is the request for the 'get' method.
type RegistryCredentialGetServerRequest struct {
}

// RegistryCredentialGetServerResponse is the response for the 'get' method.
type RegistryCredentialGetServerResponse struct {
	status int
	err    *errors.Error
	body   *RegistryCredential
}

// Body sets the value of the 'body' parameter.
//
//
func (r *RegistryCredentialGetServerResponse) Body(value *RegistryCredential) *RegistryCredentialGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *RegistryCredentialGetServerResponse) Status(value int) *RegistryCredentialGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *RegistryCredentialGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RegistryCredentialAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RegistryCredentialAdapter struct {
	server RegistryCredentialServer
	router *mux.Router
}

func NewRegistryCredentialAdapter(server RegistryCredentialServer, router *mux.Router) *RegistryCredentialAdapter {
	adapter := new(RegistryCredentialAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *RegistryCredentialAdapter) readGetRequest(r *http.Request) (*RegistryCredentialGetServerRequest, error) {
	var err error
	result := new(RegistryCredentialGetServerRequest)
	return result, err
}
func (a *RegistryCredentialAdapter) writeGetResponse(w http.ResponseWriter, r *RegistryCredentialGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RegistryCredentialAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(RegistryCredentialGetServerResponse)
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
func (a *RegistryCredentialAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
