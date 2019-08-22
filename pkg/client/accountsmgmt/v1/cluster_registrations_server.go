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

// ClusterRegistrationsServer represents the interface the manages the 'cluster_registrations' resource.
type ClusterRegistrationsServer interface {

	// Post handles a request for the 'post' method.
	//
	// Finds or creates a cluster registration with a registry credential
	// token and cluster identifier.
	Post(ctx context.Context, request *ClusterRegistrationsPostServerRequest, response *ClusterRegistrationsPostServerResponse) error
}

// ClusterRegistrationsPostServerRequest is the request for the 'post' method.
type ClusterRegistrationsPostServerRequest struct {
	path    string
	query   url.Values
	request *ClusterRegistrationRequest
}

// Request returns the value of the 'request' parameter.
//
//
func (r *ClusterRegistrationsPostServerRequest) Request() *ClusterRegistrationRequest {
	if r == nil {
		return nil
	}
	return r.request
}

// GetRequest returns the value of the 'request' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *ClusterRegistrationsPostServerRequest) GetRequest() (value *ClusterRegistrationRequest, ok bool) {
	ok = r != nil && r.request != nil
	if ok {
		value = r.request
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'post' method.
func (r *ClusterRegistrationsPostServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(clusterRegistrationRequestData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.request, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}

// ClusterRegistrationsPostServerResponse is the response for the 'post' method.
type ClusterRegistrationsPostServerResponse struct {
	status   int
	err      *errors.Error
	response *ClusterRegistrationResponse
}

// Response sets the value of the 'response' parameter.
//
//
func (r *ClusterRegistrationsPostServerResponse) Response(value *ClusterRegistrationResponse) *ClusterRegistrationsPostServerResponse {
	r.response = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ClusterRegistrationsPostServerResponse) SetStatusCode(status int) *ClusterRegistrationsPostServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'post' method.
func (r *ClusterRegistrationsPostServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.response.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ClusterRegistrationsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type ClusterRegistrationsServerAdapter struct {
	server ClusterRegistrationsServer
	router *mux.Router
}

func NewClusterRegistrationsServerAdapter(server ClusterRegistrationsServer, router *mux.Router) *ClusterRegistrationsServerAdapter {
	adapter := new(ClusterRegistrationsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.HandleFunc("/", adapter.postHandler).Methods("POST")
	return adapter
}
func (a *ClusterRegistrationsServerAdapter) readClusterRegistrationsPostServerRequest(r *http.Request) (*ClusterRegistrationsPostServerRequest, error) {
	result := new(ClusterRegistrationsPostServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *ClusterRegistrationsServerAdapter) writeClusterRegistrationsPostServerResponse(w http.ResponseWriter, r *ClusterRegistrationsPostServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ClusterRegistrationsServerAdapter) postHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readClusterRegistrationsPostServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ClusterRegistrationsPostServerResponse)
	err = a.server.Post(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Post: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeClusterRegistrationsPostServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ClusterRegistrationsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
