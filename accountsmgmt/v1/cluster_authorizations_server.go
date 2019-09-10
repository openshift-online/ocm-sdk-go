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

// ClusterAuthorizationsServer represents the interface the manages the 'cluster_authorizations' resource.
type ClusterAuthorizationsServer interface {

	// Post handles a request for the 'post' method.
	//
	// Authorizes new cluster creation against an existing subscription.
	Post(ctx context.Context, request *ClusterAuthorizationsPostServerRequest, response *ClusterAuthorizationsPostServerResponse) error
}

// ClusterAuthorizationsPostServerRequest is the request for the 'post' method.
type ClusterAuthorizationsPostServerRequest struct {
	request *ClusterAuthorizationRequest
}

// Request returns the value of the 'request' parameter.
//
//
func (r *ClusterAuthorizationsPostServerRequest) Request() *ClusterAuthorizationRequest {
	if r == nil {
		return nil
	}
	return r.request
}

// GetRequest returns the value of the 'request' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *ClusterAuthorizationsPostServerRequest) GetRequest() (value *ClusterAuthorizationRequest, ok bool) {
	ok = r != nil && r.request != nil
	if ok {
		value = r.request
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'post' method.
func (r *ClusterAuthorizationsPostServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(clusterAuthorizationRequestData)
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

// ClusterAuthorizationsPostServerResponse is the response for the 'post' method.
type ClusterAuthorizationsPostServerResponse struct {
	status   int
	err      *errors.Error
	response *ClusterAuthorizationResponse
}

// Response sets the value of the 'response' parameter.
//
//
func (r *ClusterAuthorizationsPostServerResponse) Response(value *ClusterAuthorizationResponse) *ClusterAuthorizationsPostServerResponse {
	r.response = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ClusterAuthorizationsPostServerResponse) SetStatusCode(status int) *ClusterAuthorizationsPostServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'post' method.
func (r *ClusterAuthorizationsPostServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.response.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ClusterAuthorizationsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type ClusterAuthorizationsServerAdapter struct {
	server ClusterAuthorizationsServer
	router *mux.Router
}

func NewClusterAuthorizationsServerAdapter(server ClusterAuthorizationsServer, router *mux.Router) *ClusterAuthorizationsServerAdapter {
	adapter := new(ClusterAuthorizationsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("POST").Path("").HandlerFunc(adapter.postHandler)
	return adapter
}
func (a *ClusterAuthorizationsServerAdapter) readClusterAuthorizationsPostServerRequest(r *http.Request) (*ClusterAuthorizationsPostServerRequest, error) {
	var err error
	result := new(ClusterAuthorizationsPostServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *ClusterAuthorizationsServerAdapter) writeClusterAuthorizationsPostServerResponse(w http.ResponseWriter, r *ClusterAuthorizationsPostServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ClusterAuthorizationsServerAdapter) postHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readClusterAuthorizationsPostServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ClusterAuthorizationsPostServerResponse)
	err = a.server.Post(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Post: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeClusterAuthorizationsPostServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ClusterAuthorizationsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
