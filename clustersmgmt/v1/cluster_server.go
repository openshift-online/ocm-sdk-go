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

// ClusterServer represents the interface the manages the 'cluster' resource.
type ClusterServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the cluster.
	Delete(ctx context.Context, request *ClusterDeleteServerRequest, response *ClusterDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the cluster.
	Get(ctx context.Context, request *ClusterGetServerRequest, response *ClusterGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	// Updates the cluster.
	Update(ctx context.Context, request *ClusterUpdateServerRequest, response *ClusterUpdateServerResponse) error

	// Credentials returns the target 'credentials' resource.
	//
	// Reference to the resource that manages the credentials of the cluster.
	Credentials() CredentialsServer

	// Groups returns the target 'groups' resource.
	//
	// Reference to the resource that manages the collection of groups.
	Groups() GroupsServer

	// IdentityProviders returns the target 'identity_providers' resource.
	//
	// Reference to the resource that manages the collection of identity providers.
	IdentityProviders() IdentityProvidersServer

	// Logs returns the target 'logs' resource.
	//
	// Reference to the resource that manages the collection of logs of the cluster.
	Logs() LogsServer

	// MetricQueries returns the target 'metric_queries' resource.
	//
	// Reference to the resource that manages metrics queries for the cluster.
	MetricQueries() MetricQueriesServer

	// Status returns the target 'cluster_status' resource.
	//
	// Reference to the resource that manages the detailed status of the cluster.
	Status() ClusterStatusServer
}

// ClusterDeleteServerRequest is the request for the 'delete' method.
type ClusterDeleteServerRequest struct {
}

// ClusterDeleteServerResponse is the response for the 'delete' method.
type ClusterDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ClusterDeleteServerResponse) SetStatusCode(status int) *ClusterDeleteServerResponse {
	r.status = status
	return r
}

// ClusterGetServerRequest is the request for the 'get' method.
type ClusterGetServerRequest struct {
}

// ClusterGetServerResponse is the response for the 'get' method.
type ClusterGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Cluster
}

// Body sets the value of the 'body' parameter.
//
//
func (r *ClusterGetServerResponse) Body(value *Cluster) *ClusterGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ClusterGetServerResponse) SetStatusCode(status int) *ClusterGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *ClusterGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ClusterUpdateServerRequest is the request for the 'update' method.
type ClusterUpdateServerRequest struct {
	body *Cluster
}

// Body returns the value of the 'body' parameter.
//
//
func (r *ClusterUpdateServerRequest) Body() *Cluster {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *ClusterUpdateServerRequest) GetBody() (value *Cluster, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'update' method.
func (r *ClusterUpdateServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(clusterData)
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

// ClusterUpdateServerResponse is the response for the 'update' method.
type ClusterUpdateServerResponse struct {
	status int
	err    *errors.Error
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ClusterUpdateServerResponse) SetStatusCode(status int) *ClusterUpdateServerResponse {
	r.status = status
	return r
}

// ClusterServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type ClusterServerAdapter struct {
	server ClusterServer
	router *mux.Router
}

func NewClusterServerAdapter(server ClusterServer, router *mux.Router) *ClusterServerAdapter {
	adapter := new(ClusterServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/credentials").HandlerFunc(adapter.credentialsHandler)
	adapter.router.PathPrefix("/groups").HandlerFunc(adapter.groupsHandler)
	adapter.router.PathPrefix("/identity_providers").HandlerFunc(adapter.identityProvidersHandler)
	adapter.router.PathPrefix("/logs").HandlerFunc(adapter.logsHandler)
	adapter.router.PathPrefix("/metric_queries").HandlerFunc(adapter.metricQueriesHandler)
	adapter.router.PathPrefix("/status").HandlerFunc(adapter.statusHandler)
	adapter.router.Methods("DELETE").Path("").HandlerFunc(adapter.deleteHandler)
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	adapter.router.Methods("PATCH").Path("").HandlerFunc(adapter.updateHandler)
	return adapter
}
func (a *ClusterServerAdapter) credentialsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Credentials()
	targetAdapter := NewCredentialsServerAdapter(target, a.router.PathPrefix("/credentials").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterServerAdapter) groupsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Groups()
	targetAdapter := NewGroupsServerAdapter(target, a.router.PathPrefix("/groups").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterServerAdapter) identityProvidersHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.IdentityProviders()
	targetAdapter := NewIdentityProvidersServerAdapter(target, a.router.PathPrefix("/identity_providers").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterServerAdapter) logsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Logs()
	targetAdapter := NewLogsServerAdapter(target, a.router.PathPrefix("/logs").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterServerAdapter) metricQueriesHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.MetricQueries()
	targetAdapter := NewMetricQueriesServerAdapter(target, a.router.PathPrefix("/metric_queries").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterServerAdapter) statusHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Status()
	targetAdapter := NewClusterStatusServerAdapter(target, a.router.PathPrefix("/status").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterServerAdapter) readClusterDeleteServerRequest(r *http.Request) (*ClusterDeleteServerRequest, error) {
	var err error
	result := new(ClusterDeleteServerRequest)
	return result, err
}
func (a *ClusterServerAdapter) writeClusterDeleteServerResponse(w http.ResponseWriter, r *ClusterDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *ClusterServerAdapter) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readClusterDeleteServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ClusterDeleteServerResponse)
	err = a.server.Delete(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Delete: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeClusterDeleteServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ClusterServerAdapter) readClusterGetServerRequest(r *http.Request) (*ClusterGetServerRequest, error) {
	var err error
	result := new(ClusterGetServerRequest)
	return result, err
}
func (a *ClusterServerAdapter) writeClusterGetServerResponse(w http.ResponseWriter, r *ClusterGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ClusterServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readClusterGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ClusterGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeClusterGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ClusterServerAdapter) readClusterUpdateServerRequest(r *http.Request) (*ClusterUpdateServerRequest, error) {
	var err error
	result := new(ClusterUpdateServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *ClusterServerAdapter) writeClusterUpdateServerResponse(w http.ResponseWriter, r *ClusterUpdateServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *ClusterServerAdapter) updateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readClusterUpdateServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ClusterUpdateServerResponse)
	err = a.server.Update(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Update: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeClusterUpdateServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ClusterServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
