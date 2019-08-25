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

// DashboardsServer represents the interface the manages the 'dashboards' resource.
type DashboardsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves a list of dashboards.
	List(ctx context.Context, request *DashboardsListServerRequest, response *DashboardsListServerResponse) error

	// Dashboard returns the target 'dashboard' server for the given identifier.
	//
	// Reference to the resource that manages a specific dashboard.
	Dashboard(id string) DashboardServer
}

// DashboardsListServerRequest is the request for the 'list' method.
type DashboardsListServerRequest struct {
	path   string
	query  url.Values
	page   *int
	size   *int
	search *string
	total  *int
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *DashboardsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *DashboardsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *DashboardsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *DashboardsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Search returns the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the
// _where_ clause of an SQL statement, but using the names of
// the attributes of the dashboard instead of the names of the
// columns of a table. For example, in order to retrieve all the
// dashboards with a name starting with `my` the value should be:
//
// [source,sql]
// ----
// name like 'my%'
// ----
//
// If the parameter isn't provided, or if the value is empty,
// then all the dashboards that the user has permission to see
// will be returned.
func (r *DashboardsListServerRequest) Search() string {
	if r != nil && r.search != nil {
		return *r.search
	}
	return ""
}

// GetSearch returns the value of the 'search' parameter and
// a flag indicating if the parameter has a value.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the
// _where_ clause of an SQL statement, but using the names of
// the attributes of the dashboard instead of the names of the
// columns of a table. For example, in order to retrieve all the
// dashboards with a name starting with `my` the value should be:
//
// [source,sql]
// ----
// name like 'my%'
// ----
//
// If the parameter isn't provided, or if the value is empty,
// then all the dashboards that the user has permission to see
// will be returned.
func (r *DashboardsListServerRequest) GetSearch() (value string, ok bool) {
	ok = r != nil && r.search != nil
	if ok {
		value = *r.search
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *DashboardsListServerRequest) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *DashboardsListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// DashboardsListServerResponse is the response for the 'list' method.
type DashboardsListServerResponse struct {
	status int
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *DashboardList
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *DashboardsListServerResponse) Page(value int) *DashboardsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *DashboardsListServerResponse) Size(value int) *DashboardsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *DashboardsListServerResponse) Total(value int) *DashboardsListServerResponse {
	r.total = &value
	return r
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of dashboards.
func (r *DashboardsListServerResponse) Items(value *DashboardList) *DashboardsListServerResponse {
	r.items = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *DashboardsListServerResponse) SetStatusCode(status int) *DashboardsListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *DashboardsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(dashboardsListServerResponseData)
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

// dashboardsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type dashboardsListServerResponseData struct {
	Page  *int              "json:\"page,omitempty\""
	Size  *int              "json:\"size,omitempty\""
	Total *int              "json:\"total,omitempty\""
	Items dashboardListData "json:\"items,omitempty\""
}

// DashboardsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type DashboardsServerAdapter struct {
	server DashboardsServer
	router *mux.Router
}

func NewDashboardsServerAdapter(server DashboardsServer, router *mux.Router) *DashboardsServerAdapter {
	adapter := new(DashboardsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.dashboardHandler)
	adapter.router.Methods("GET").HandlerFunc(adapter.listHandler)
	return adapter
}
func (a *DashboardsServerAdapter) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Dashboard(id)
	targetAdapter := NewDashboardServerAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *DashboardsServerAdapter) readDashboardsListServerRequest(r *http.Request) (*DashboardsListServerRequest, error) {
	result := new(DashboardsListServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *DashboardsServerAdapter) writeDashboardsListServerResponse(w http.ResponseWriter, r *DashboardsListServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *DashboardsServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readDashboardsListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(DashboardsListServerResponse)
	err = a.server.List(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeDashboardsListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *DashboardsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
