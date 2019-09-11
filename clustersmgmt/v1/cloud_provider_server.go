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

// CloudProviderServer represents the interface the manages the 'cloud_provider' resource.
type CloudProviderServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the cloud provider.
	Get(ctx context.Context, request *CloudProviderGetServerRequest, response *CloudProviderGetServerResponse) error

	// Regions returns the target 'cloud_regions' resource.
	//
	// Reference to the resource that manages the collection of regions for
	// this cloud provider.
	Regions() CloudRegionsServer
}

// CloudProviderGetServerRequest is the request for the 'get' method.
type CloudProviderGetServerRequest struct {
}

// CloudProviderGetServerResponse is the response for the 'get' method.
type CloudProviderGetServerResponse struct {
	status int
	err    *errors.Error
	body   *CloudProvider
}

// Body sets the value of the 'body' parameter.
//
//
func (r *CloudProviderGetServerResponse) Body(value *CloudProvider) *CloudProviderGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *CloudProviderGetServerResponse) SetStatusCode(status int) *CloudProviderGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *CloudProviderGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// CloudProviderServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type CloudProviderServerAdapter struct {
	server CloudProviderServer
	router *mux.Router
}

func NewCloudProviderServerAdapter(server CloudProviderServer, router *mux.Router) *CloudProviderServerAdapter {
	adapter := new(CloudProviderServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/regions").HandlerFunc(adapter.regionsHandler)
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	return adapter
}
func (a *CloudProviderServerAdapter) regionsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Regions()
	targetAdapter := NewCloudRegionsServerAdapter(target, a.router.PathPrefix("/regions").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *CloudProviderServerAdapter) readCloudProviderGetServerRequest(r *http.Request) (*CloudProviderGetServerRequest, error) {
	var err error
	result := new(CloudProviderGetServerRequest)
	return result, err
}
func (a *CloudProviderServerAdapter) writeCloudProviderGetServerResponse(w http.ResponseWriter, r *CloudProviderGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *CloudProviderServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readCloudProviderGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(CloudProviderGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeCloudProviderGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *CloudProviderServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
