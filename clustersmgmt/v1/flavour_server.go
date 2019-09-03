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

// FlavourServer represents the interface the manages the 'flavour' resource.
type FlavourServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the cluster flavour.
	Get(ctx context.Context, request *FlavourGetServerRequest, response *FlavourGetServerResponse) error
}

// FlavourGetServerRequest is the request for the 'get' method.
type FlavourGetServerRequest struct {
	path  string
	query url.Values
}

// FlavourGetServerResponse is the response for the 'get' method.
type FlavourGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Flavour
}

// Body sets the value of the 'body' parameter.
//
//
func (r *FlavourGetServerResponse) Body(value *Flavour) *FlavourGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *FlavourGetServerResponse) SetStatusCode(status int) *FlavourGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *FlavourGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// FlavourServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type FlavourServerAdapter struct {
	server FlavourServer
	router *mux.Router
}

func NewFlavourServerAdapter(server FlavourServer, router *mux.Router) *FlavourServerAdapter {
	adapter := new(FlavourServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	return adapter
}
func (a *FlavourServerAdapter) readFlavourGetServerRequest(r *http.Request) (*FlavourGetServerRequest, error) {
	var err error
	result := new(FlavourGetServerRequest)
	result.query = r.URL.Query()
	result.path = r.URL.Path
	return result, err
}
func (a *FlavourServerAdapter) writeFlavourGetServerResponse(w http.ResponseWriter, r *FlavourGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *FlavourServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readFlavourGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(FlavourGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeFlavourGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *FlavourServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
