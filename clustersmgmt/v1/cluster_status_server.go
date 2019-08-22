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

// ClusterStatusServer represents the interface the manages the 'cluster_status' resource.
type ClusterStatusServer interface {

	// Get handles a request for the 'get' method.
	//
	//
	Get(ctx context.Context, request *ClusterStatusGetServerRequest, response *ClusterStatusGetServerResponse) error
}

// ClusterStatusGetServerRequest is the request for the 'get' method.
type ClusterStatusGetServerRequest struct {
	path  string
	query url.Values
}

// ClusterStatusGetServerResponse is the response for the 'get' method.
type ClusterStatusGetServerResponse struct {
	status  int
	err     *errors.Error
	status_ *ClusterStatus
}

// Status_ sets the value of the 'status' parameter.
//
//
func (r *ClusterStatusGetServerResponse) Status_(value *ClusterStatus) *ClusterStatusGetServerResponse {
	r.status_ = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ClusterStatusGetServerResponse) SetStatusCode(status int) *ClusterStatusGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *ClusterStatusGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.status_.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ClusterStatusServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type ClusterStatusServerAdapter struct {
	server ClusterStatusServer
	router *mux.Router
}

func NewClusterStatusServerAdapter(server ClusterStatusServer, router *mux.Router) *ClusterStatusServerAdapter {
	adapter := new(ClusterStatusServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.HandleFunc("/", adapter.getHandler).Methods("GET")
	return adapter
}
func (a *ClusterStatusServerAdapter) readClusterStatusGetServerRequest(r *http.Request) (*ClusterStatusGetServerRequest, error) {
	result := new(ClusterStatusGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *ClusterStatusServerAdapter) writeClusterStatusGetServerResponse(w http.ResponseWriter, r *ClusterStatusGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ClusterStatusServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readClusterStatusGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ClusterStatusGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeClusterStatusGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ClusterStatusServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
