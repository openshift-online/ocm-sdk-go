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

// OrganizationServer represents the interface the manages the 'organization' resource.
type OrganizationServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the organization.
	Get(ctx context.Context, request *OrganizationGetServerRequest, response *OrganizationGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	// Updates the organization.
	Update(ctx context.Context, request *OrganizationUpdateServerRequest, response *OrganizationUpdateServerResponse) error

	// ResourceQuota returns the target 'resource_quotas' resource.
	//
	// Reference to the service that manages the resource quotas for this
	// organization.
	ResourceQuota() ResourceQuotasServer

	// QuotaSummary returns the target 'quota_summary' resource.
	//
	// Reference to the service that returns the summary of the resource quota for this
	// organization.
	QuotaSummary() QuotaSummaryServer
}

// OrganizationGetServerRequest is the request for the 'get' method.
type OrganizationGetServerRequest struct {
	path  string
	query url.Values
}

// OrganizationGetServerResponse is the response for the 'get' method.
type OrganizationGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Organization
}

// Body sets the value of the 'body' parameter.
//
//
func (r *OrganizationGetServerResponse) Body(value *Organization) *OrganizationGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *OrganizationGetServerResponse) SetStatusCode(status int) *OrganizationGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *OrganizationGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// OrganizationUpdateServerRequest is the request for the 'update' method.
type OrganizationUpdateServerRequest struct {
	path  string
	query url.Values
	body  *Organization
}

// Body returns the value of the 'body' parameter.
//
//
func (r *OrganizationUpdateServerRequest) Body() *Organization {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *OrganizationUpdateServerRequest) GetBody() (value *Organization, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'update' method.
func (r *OrganizationUpdateServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(organizationData)
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

// OrganizationUpdateServerResponse is the response for the 'update' method.
type OrganizationUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *Organization
}

// Body sets the value of the 'body' parameter.
//
//
func (r *OrganizationUpdateServerResponse) Body(value *Organization) *OrganizationUpdateServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *OrganizationUpdateServerResponse) SetStatusCode(status int) *OrganizationUpdateServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'update' method.
func (r *OrganizationUpdateServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// OrganizationServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type OrganizationServerAdapter struct {
	server OrganizationServer
	router *mux.Router
}

func NewOrganizationServerAdapter(server OrganizationServer, router *mux.Router) *OrganizationServerAdapter {
	adapter := new(OrganizationServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/resource_quota").HandlerFunc(adapter.resourceQuotaHandler)
	adapter.router.PathPrefix("/quota_summary").HandlerFunc(adapter.quotaSummaryHandler)
	adapter.router.Methods("GET").HandlerFunc(adapter.getHandler)
	adapter.router.Methods("PATCH").HandlerFunc(adapter.updateHandler)
	return adapter
}
func (a *OrganizationServerAdapter) resourceQuotaHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.ResourceQuota()
	targetAdapter := NewResourceQuotasServerAdapter(target, a.router.PathPrefix("/resource_quota").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *OrganizationServerAdapter) quotaSummaryHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.QuotaSummary()
	targetAdapter := NewQuotaSummaryServerAdapter(target, a.router.PathPrefix("/quota_summary").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *OrganizationServerAdapter) readOrganizationGetServerRequest(r *http.Request) (*OrganizationGetServerRequest, error) {
	result := new(OrganizationGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *OrganizationServerAdapter) writeOrganizationGetServerResponse(w http.ResponseWriter, r *OrganizationGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *OrganizationServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readOrganizationGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(OrganizationGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeOrganizationGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *OrganizationServerAdapter) readOrganizationUpdateServerRequest(r *http.Request) (*OrganizationUpdateServerRequest, error) {
	result := new(OrganizationUpdateServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *OrganizationServerAdapter) writeOrganizationUpdateServerResponse(w http.ResponseWriter, r *OrganizationUpdateServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *OrganizationServerAdapter) updateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readOrganizationUpdateServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(OrganizationUpdateServerResponse)
	err = a.server.Update(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Update: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeOrganizationUpdateServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *OrganizationServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
