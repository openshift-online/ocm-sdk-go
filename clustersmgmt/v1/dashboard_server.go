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

// DashboardServer represents the interface the manages the 'dashboard' resource.
type DashboardServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the dashboard.
	Get(ctx context.Context, request *DashboardGetServerRequest, response *DashboardGetServerResponse) error
}

// DashboardGetServerRequest is the request for the 'get' method.
type DashboardGetServerRequest struct {
}

// DashboardGetServerResponse is the response for the 'get' method.
type DashboardGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Dashboard
}

// Body sets the value of the 'body' parameter.
//
//
func (r *DashboardGetServerResponse) Body(value *Dashboard) *DashboardGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *DashboardGetServerResponse) Status(value int) *DashboardGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *DashboardGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// DashboardAdapter represents the structs that adapts Requests and Response to internal
// structs.
type DashboardAdapter struct {
	server DashboardServer
	router *mux.Router
}

func NewDashboardAdapter(server DashboardServer, router *mux.Router) *DashboardAdapter {
	adapter := new(DashboardAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *DashboardAdapter) readGetRequest(r *http.Request) (*DashboardGetServerRequest, error) {
	var err error
	result := new(DashboardGetServerRequest)
	return result, err
}
func (a *DashboardAdapter) writeGetResponse(w http.ResponseWriter, r *DashboardGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *DashboardAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(DashboardGetServerResponse)
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
func (a *DashboardAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
