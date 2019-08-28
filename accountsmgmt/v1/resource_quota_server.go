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
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// ResourceQuotaServer represents the interface the manages the 'resource_quota' resource.
type ResourceQuotaServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the resource quota.
	Get(ctx context.Context, request *ResourceQuotaGetServerRequest, response *ResourceQuotaGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	// Updates the resource quota.
	Update(ctx context.Context, request *ResourceQuotaUpdateServerRequest, response *ResourceQuotaUpdateServerResponse) error
}

// ResourceQuotaGetServerRequest is the request for the 'get' method.
type ResourceQuotaGetServerRequest struct {
	path  string
	query url.Values
}

// ResourceQuotaGetServerResponse is the response for the 'get' method.
type ResourceQuotaGetServerResponse struct {
	status int
	err    *errors.Error
	body   *ResourceQuota
}

// Body sets the value of the 'body' parameter.
//
//
func (r *ResourceQuotaGetServerResponse) Body(value *ResourceQuota) *ResourceQuotaGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ResourceQuotaGetServerResponse) SetStatusCode(status int) *ResourceQuotaGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *ResourceQuotaGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ResourceQuotaUpdateServerRequest is the request for the 'update' method.
type ResourceQuotaUpdateServerRequest struct {
	path  string
	query url.Values
	body  *ResourceQuota
}

// Body returns the value of the 'body' parameter.
//
//
func (r *ResourceQuotaUpdateServerRequest) Body() *ResourceQuota {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *ResourceQuotaUpdateServerRequest) GetBody() (value *ResourceQuota, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'update' method.
func (r *ResourceQuotaUpdateServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(resourceQuotaData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}

// ResourceQuotaUpdateServerResponse is the response for the 'update' method.
type ResourceQuotaUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *ResourceQuota
}

// Body sets the value of the 'body' parameter.
//
//
func (r *ResourceQuotaUpdateServerResponse) Body(value *ResourceQuota) *ResourceQuotaUpdateServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ResourceQuotaUpdateServerResponse) SetStatusCode(status int) *ResourceQuotaUpdateServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'update' method.
func (r *ResourceQuotaUpdateServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ResourceQuotaServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type ResourceQuotaServerAdapter struct {
	server ResourceQuotaServer
	router *mux.Router
}

func NewResourceQuotaServerAdapter(server ResourceQuotaServer, router *mux.Router) *ResourceQuotaServerAdapter {
	adapter := new(ResourceQuotaServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	adapter.router.Methods("PATCH").Path("").HandlerFunc(adapter.updateHandler)
	return adapter
}
func (a *ResourceQuotaServerAdapter) readResourceQuotaGetServerRequest(r *http.Request) (*ResourceQuotaGetServerRequest, error) {
	result := new(ResourceQuotaGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *ResourceQuotaServerAdapter) writeResourceQuotaGetServerResponse(w http.ResponseWriter, r *ResourceQuotaGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ResourceQuotaServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readResourceQuotaGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ResourceQuotaGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeResourceQuotaGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ResourceQuotaServerAdapter) readResourceQuotaUpdateServerRequest(r *http.Request) (*ResourceQuotaUpdateServerRequest, error) {
	result := new(ResourceQuotaUpdateServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *ResourceQuotaServerAdapter) writeResourceQuotaUpdateServerResponse(w http.ResponseWriter, r *ResourceQuotaUpdateServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ResourceQuotaServerAdapter) updateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readResourceQuotaUpdateServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ResourceQuotaUpdateServerResponse)
	err = a.server.Update(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Update: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeResourceQuotaUpdateServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ResourceQuotaServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
