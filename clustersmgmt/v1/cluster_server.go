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

// Status sets the status code.
func (r *ClusterDeleteServerResponse) Status(value int) *ClusterDeleteServerResponse {
	r.status = value
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

// Status sets the status code.
func (r *ClusterGetServerResponse) Status(value int) *ClusterGetServerResponse {
	r.status = value
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

// Status sets the status code.
func (r *ClusterUpdateServerResponse) Status(value int) *ClusterUpdateServerResponse {
	r.status = value
	return r
}

// ClusterAdapter represents the structs that adapts Requests and Response to internal
// structs.
type ClusterAdapter struct {
	server ClusterServer
	router *mux.Router
}

func NewClusterAdapter(server ClusterServer, router *mux.Router) *ClusterAdapter {
	adapter := new(ClusterAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/credentials").HandlerFunc(adapter.credentialsHandler)
	adapter.router.PathPrefix("/groups").HandlerFunc(adapter.groupsHandler)
	adapter.router.PathPrefix("/identity_providers").HandlerFunc(adapter.identityProvidersHandler)
	adapter.router.PathPrefix("/logs").HandlerFunc(adapter.logsHandler)
	adapter.router.PathPrefix("/metric_queries").HandlerFunc(adapter.metricQueriesHandler)
	adapter.router.PathPrefix("/status").HandlerFunc(adapter.statusHandler)
	adapter.router.Methods(http.MethodDelete).Path("").HandlerFunc(adapter.handlerDelete)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	adapter.router.Methods(http.MethodPatch).Path("").HandlerFunc(adapter.handlerUpdate)
	return adapter
}
func (a *ClusterAdapter) credentialsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Credentials()
	targetAdapter := NewCredentialsAdapter(target, a.router.PathPrefix("/credentials").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterAdapter) groupsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Groups()
	targetAdapter := NewGroupsAdapter(target, a.router.PathPrefix("/groups").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterAdapter) identityProvidersHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.IdentityProviders()
	targetAdapter := NewIdentityProvidersAdapter(target, a.router.PathPrefix("/identity_providers").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterAdapter) logsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Logs()
	targetAdapter := NewLogsAdapter(target, a.router.PathPrefix("/logs").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterAdapter) metricQueriesHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.MetricQueries()
	targetAdapter := NewMetricQueriesAdapter(target, a.router.PathPrefix("/metric_queries").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterAdapter) statusHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Status()
	targetAdapter := NewClusterStatusAdapter(target, a.router.PathPrefix("/status").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClusterAdapter) readDeleteRequest(r *http.Request) (*ClusterDeleteServerRequest, error) {
	var err error
	result := new(ClusterDeleteServerRequest)
	return result, err
}
func (a *ClusterAdapter) writeDeleteResponse(w http.ResponseWriter, r *ClusterDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *ClusterAdapter) handlerDelete(w http.ResponseWriter, r *http.Request) {
	request, err := a.readDeleteRequest(r)
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
	response := new(ClusterDeleteServerResponse)
	response.status = http.StatusOK
	err = a.server.Delete(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method Delete: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeDeleteResponse(w, response)
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
func (a *ClusterAdapter) readGetRequest(r *http.Request) (*ClusterGetServerRequest, error) {
	var err error
	result := new(ClusterGetServerRequest)
	return result, err
}
func (a *ClusterAdapter) writeGetResponse(w http.ResponseWriter, r *ClusterGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ClusterAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
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
	response := new(ClusterGetServerResponse)
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
func (a *ClusterAdapter) readUpdateRequest(r *http.Request) (*ClusterUpdateServerRequest, error) {
	var err error
	result := new(ClusterUpdateServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *ClusterAdapter) writeUpdateResponse(w http.ResponseWriter, r *ClusterUpdateServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}
func (a *ClusterAdapter) handlerUpdate(w http.ResponseWriter, r *http.Request) {
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
	response := new(ClusterUpdateServerResponse)
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
func (a *ClusterAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
