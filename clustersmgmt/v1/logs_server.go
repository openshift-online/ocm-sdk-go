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

package v1 // github.com/openshift-online/uhc-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/uhc-sdk-go/errors"
)

// LogsServer represents the interface the manages the 'logs' resource.
type LogsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of clusters.
	List(ctx context.Context, request *LogsListServerRequest, response *LogsListServerResponse) error

	// Log returns the target 'log' server for the given identifier.
	//
	// Returns a reference to the service that manages an specific log.
	Log(id string) LogServer
}

// LogsListServerRequest is the request for the 'list' method.
type LogsListServerRequest struct {
	path  string
	query url.Values
}

// LogsListServerResponse is the response for the 'list' method.
type LogsListServerResponse struct {
	status int
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *LogList
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *LogsListServerResponse) Page(value int) *LogsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *LogsListServerResponse) Size(value int) *LogsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *LogsListServerResponse) Total(value int) *LogsListServerResponse {
	r.total = &value
	return r
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of logs.
func (r *LogsListServerResponse) Items(value *LogList) *LogsListServerResponse {
	r.items = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *LogsListServerResponse) SetStatusCode(status int) *LogsListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *LogsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(logsListServerResponseData)
	data.Page = r.page
	data.Size = r.size
	data.Total = r.total
	data.Items, err = r.items.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// logsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type logsListServerResponseData struct {
	Page  *int        "json:\"page,omitempty\""
	Size  *int        "json:\"size,omitempty\""
	Total *int        "json:\"total,omitempty\""
	Items logListData "json:\"items,omitempty\""
}

// LogsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type LogsServerAdapter struct {
	server LogsServer
	router *mux.Router
}

func NewLogsServerAdapter(server LogsServer, router *mux.Router) *LogsServerAdapter {
	adapter := new(LogsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.logHandler)
	adapter.router.Methods("GET").HandlerFunc(adapter.listHandler)
	return adapter
}
func (a *LogsServerAdapter) logHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Log(id)
	targetAdapter := NewLogServerAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *LogsServerAdapter) readLogsListServerRequest(r *http.Request) (*LogsListServerRequest, error) {
	result := new(LogsListServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *LogsServerAdapter) writeLogsListServerResponse(w http.ResponseWriter, r *LogsListServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *LogsServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readLogsListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(LogsListServerResponse)
	err = a.server.List(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeLogsListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *LogsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
