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

// CPUTotalByNodeRolesOSMetricQueryServer represents the interface the manages the 'CPU_total_by_node_roles_OS_metric_query' resource.
type CPUTotalByNodeRolesOSMetricQueryServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the metrics.
	Get(ctx context.Context, request *CPUTotalByNodeRolesOSMetricQueryGetServerRequest, response *CPUTotalByNodeRolesOSMetricQueryGetServerResponse) error
}

// CPUTotalByNodeRolesOSMetricQueryGetServerRequest is the request for the 'get' method.
type CPUTotalByNodeRolesOSMetricQueryGetServerRequest struct {
}

// CPUTotalByNodeRolesOSMetricQueryGetServerResponse is the response for the 'get' method.
type CPUTotalByNodeRolesOSMetricQueryGetServerResponse struct {
	status int
	err    *errors.Error
	body   *CPUTotalsNodeRoleOSMetricNode
}

// Body sets the value of the 'body' parameter.
//
//
func (r *CPUTotalByNodeRolesOSMetricQueryGetServerResponse) Body(value *CPUTotalsNodeRoleOSMetricNode) *CPUTotalByNodeRolesOSMetricQueryGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *CPUTotalByNodeRolesOSMetricQueryGetServerResponse) Status(value int) *CPUTotalByNodeRolesOSMetricQueryGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *CPUTotalByNodeRolesOSMetricQueryGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// CPUTotalByNodeRolesOSMetricQueryAdapter represents the structs that adapts Requests and Response to internal
// structs.
type CPUTotalByNodeRolesOSMetricQueryAdapter struct {
	server CPUTotalByNodeRolesOSMetricQueryServer
	router *mux.Router
}

func NewCPUTotalByNodeRolesOSMetricQueryAdapter(server CPUTotalByNodeRolesOSMetricQueryServer, router *mux.Router) *CPUTotalByNodeRolesOSMetricQueryAdapter {
	adapter := new(CPUTotalByNodeRolesOSMetricQueryAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *CPUTotalByNodeRolesOSMetricQueryAdapter) readGetRequest(r *http.Request) (*CPUTotalByNodeRolesOSMetricQueryGetServerRequest, error) {
	var err error
	result := new(CPUTotalByNodeRolesOSMetricQueryGetServerRequest)
	return result, err
}
func (a *CPUTotalByNodeRolesOSMetricQueryAdapter) writeGetResponse(w http.ResponseWriter, r *CPUTotalByNodeRolesOSMetricQueryGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *CPUTotalByNodeRolesOSMetricQueryAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(CPUTotalByNodeRolesOSMetricQueryGetServerResponse)
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
func (a *CPUTotalByNodeRolesOSMetricQueryAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
