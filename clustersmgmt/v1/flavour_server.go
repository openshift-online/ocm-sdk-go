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

// FlavourServer represents the interface the manages the 'flavour' resource.
type FlavourServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the cluster flavour.
	Get(ctx context.Context, request *FlavourGetServerRequest, response *FlavourGetServerResponse) error
}

// FlavourGetServerRequest is the request for the 'get' method.
type FlavourGetServerRequest struct {
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

// Status sets the status code.
func (r *FlavourGetServerResponse) Status(value int) *FlavourGetServerResponse {
	r.status = value
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

// FlavourAdapter represents the structs that adapts Requests and Response to internal
// structs.
type FlavourAdapter struct {
	server FlavourServer
	router *mux.Router
}

func NewFlavourAdapter(server FlavourServer, router *mux.Router) *FlavourAdapter {
	adapter := new(FlavourAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *FlavourAdapter) readGetRequest(r *http.Request) (*FlavourGetServerRequest, error) {
	var err error
	result := new(FlavourGetServerRequest)
	return result, err
}
func (a *FlavourAdapter) writeGetResponse(w http.ResponseWriter, r *FlavourGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *FlavourAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(FlavourGetServerResponse)
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
func (a *FlavourAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
