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

// CredentialsServer represents the interface the manages the 'credentials' resource.
type CredentialsServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the credentials of a cluster.
	Get(ctx context.Context, request *CredentialsGetServerRequest, response *CredentialsGetServerResponse) error
}

// CredentialsGetServerRequest is the request for the 'get' method.
type CredentialsGetServerRequest struct {
	path  string
	query url.Values
}

// CredentialsGetServerResponse is the response for the 'get' method.
type CredentialsGetServerResponse struct {
	status int
	err    *errors.Error
	body   *ClusterCredentials
}

// Body sets the value of the 'body' parameter.
//
//
func (r *CredentialsGetServerResponse) Body(value *ClusterCredentials) *CredentialsGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *CredentialsGetServerResponse) SetStatusCode(status int) *CredentialsGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *CredentialsGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// CredentialsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type CredentialsServerAdapter struct {
	server CredentialsServer
	router *mux.Router
}

func NewCredentialsServerAdapter(server CredentialsServer, router *mux.Router) *CredentialsServerAdapter {
	adapter := new(CredentialsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").HandlerFunc(adapter.getHandler)
	return adapter
}
func (a *CredentialsServerAdapter) readCredentialsGetServerRequest(r *http.Request) (*CredentialsGetServerRequest, error) {
	result := new(CredentialsGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *CredentialsServerAdapter) writeCredentialsGetServerResponse(w http.ResponseWriter, r *CredentialsGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *CredentialsServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readCredentialsGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(CredentialsGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeCredentialsGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *CredentialsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
