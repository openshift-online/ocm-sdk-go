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

	// QuotaSummary returns the target 'quota_summary' resource.
	//
	// Reference to the service that returns the summary of the resource quota for this
	// organization.
	QuotaSummary() QuotaSummaryServer

	// ResourceQuota returns the target 'resource_quotas' resource.
	//
	// Reference to the service that manages the resource quotas for this
	// organization.
	ResourceQuota() ResourceQuotasServer
}

// OrganizationGetServerRequest is the request for the 'get' method.
type OrganizationGetServerRequest struct {
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

// Status sets the status code.
func (r *OrganizationGetServerResponse) Status(value int) *OrganizationGetServerResponse {
	r.status = value
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
	body *Organization
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
}

// Status sets the status code.
func (r *OrganizationUpdateServerResponse) Status(value int) *OrganizationUpdateServerResponse {
	r.status = value
	return r
}

// OrganizationAdapter represents the structs that adapts Requests and Response to internal
// structs.
type OrganizationAdapter struct {
	server OrganizationServer
	router *mux.Router
}

func NewOrganizationAdapter(server OrganizationServer, router *mux.Router) *OrganizationAdapter {
	adapter := new(OrganizationAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/quota_summary").HandlerFunc(adapter.quotaSummaryHandler)
	adapter.router.PathPrefix("/resource_quota").HandlerFunc(adapter.resourceQuotaHandler)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	adapter.router.Methods(http.MethodPatch).Path("").HandlerFunc(adapter.handlerUpdate)
	return adapter
}
func (a *OrganizationAdapter) quotaSummaryHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.QuotaSummary()
	targetAdapter := NewQuotaSummaryAdapter(target, a.router.PathPrefix("/quota_summary").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *OrganizationAdapter) resourceQuotaHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.ResourceQuota()
	targetAdapter := NewResourceQuotasAdapter(target, a.router.PathPrefix("/resource_quota").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *OrganizationAdapter) readGetRequest(r *http.Request) (*OrganizationGetServerRequest, error) {
	var err error
	result := new(OrganizationGetServerRequest)
	return result, err
}
func (a *OrganizationAdapter) writeGetResponse(w http.ResponseWriter, r *OrganizationGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *OrganizationAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(OrganizationGetServerResponse)
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
func (a *OrganizationAdapter) readUpdateRequest(r *http.Request) (*OrganizationUpdateServerRequest, error) {
	var err error
	result := new(OrganizationUpdateServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *OrganizationAdapter) writeUpdateResponse(w http.ResponseWriter, r *OrganizationUpdateServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *OrganizationAdapter) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	request, err := a.readUpdateRequest(r)
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
	response := new(OrganizationUpdateServerResponse)
	response.status = http.StatusOK
	err = a.server.Update(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method Update: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeUpdateResponse(w, response)
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
func (a *OrganizationAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
