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

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *CPUTotalByNodeRolesOSMetricQueryGetServerResponse) SetStatusCode(status int) *CPUTotalByNodeRolesOSMetricQueryGetServerResponse {
	r.status = status
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

// CPUTotalByNodeRolesOSMetricQueryServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type CPUTotalByNodeRolesOSMetricQueryServerAdapter struct {
	server CPUTotalByNodeRolesOSMetricQueryServer
	router *mux.Router
}

func NewCPUTotalByNodeRolesOSMetricQueryServerAdapter(server CPUTotalByNodeRolesOSMetricQueryServer, router *mux.Router) *CPUTotalByNodeRolesOSMetricQueryServerAdapter {
	adapter := new(CPUTotalByNodeRolesOSMetricQueryServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	return adapter
}
func (a *CPUTotalByNodeRolesOSMetricQueryServerAdapter) readCPUTotalByNodeRolesOSMetricQueryGetServerRequest(r *http.Request) (*CPUTotalByNodeRolesOSMetricQueryGetServerRequest, error) {
	var err error
	result := new(CPUTotalByNodeRolesOSMetricQueryGetServerRequest)
	return result, err
}
func (a *CPUTotalByNodeRolesOSMetricQueryServerAdapter) writeCPUTotalByNodeRolesOSMetricQueryGetServerResponse(w http.ResponseWriter, r *CPUTotalByNodeRolesOSMetricQueryGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *CPUTotalByNodeRolesOSMetricQueryServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readCPUTotalByNodeRolesOSMetricQueryGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(CPUTotalByNodeRolesOSMetricQueryGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeCPUTotalByNodeRolesOSMetricQueryGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *CPUTotalByNodeRolesOSMetricQueryServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
