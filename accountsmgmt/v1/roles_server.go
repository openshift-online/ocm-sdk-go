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
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// RolesServer represents the interface the manages the 'roles' resource.
type RolesServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves a list of roles.
	List(ctx context.Context, request *RolesListServerRequest, response *RolesListServerResponse) error

	// Add handles a request for the 'add' method.
	//
	// Creates a new role.
	Add(ctx context.Context, request *RolesAddServerRequest, response *RolesAddServerResponse) error

	// Role returns the target 'role' server for the given identifier.
	//
	// Reference to the service that manages a specific role.
	Role(id string) RoleServer
}

// RolesListServerRequest is the request for the 'list' method.
type RolesListServerRequest struct {
	path  string
	query url.Values
	page  *int
	size  *int
	total *int
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RolesListServerRequest) Page() int {
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
func (r *RolesListServerRequest) GetPage() (value int, ok bool) {
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
func (r *RolesListServerRequest) Size() int {
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
func (r *RolesListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RolesListServerRequest) Total() int {
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
func (r *RolesListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// RolesListServerResponse is the response for the 'list' method.
type RolesListServerResponse struct {
	status int
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *RoleList
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RolesListServerResponse) Page(value int) *RolesListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *RolesListServerResponse) Size(value int) *RolesListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RolesListServerResponse) Total(value int) *RolesListServerResponse {
	r.total = &value
	return r
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of roles.
func (r *RolesListServerResponse) Items(value *RoleList) *RolesListServerResponse {
	r.items = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RolesListServerResponse) SetStatusCode(status int) *RolesListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *RolesListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(rolesListServerResponseData)
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

// rolesListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type rolesListServerResponseData struct {
	Page  *int         "json:\"page,omitempty\""
	Size  *int         "json:\"size,omitempty\""
	Total *int         "json:\"total,omitempty\""
	Items roleListData "json:\"items,omitempty\""
}

// RolesAddServerRequest is the request for the 'add' method.
type RolesAddServerRequest struct {
	path  string
	query url.Values
	body  *Role
}

// Body returns the value of the 'body' parameter.
//
// Role data.
func (r *RolesAddServerRequest) Body() *Role {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Role data.
func (r *RolesAddServerRequest) GetBody() (value *Role, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'add' method.
func (r *RolesAddServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(roleData)
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

// RolesAddServerResponse is the response for the 'add' method.
type RolesAddServerResponse struct {
	status int
	err    *errors.Error
	body   *Role
}

// Body sets the value of the 'body' parameter.
//
// Role data.
func (r *RolesAddServerResponse) Body(value *Role) *RolesAddServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RolesAddServerResponse) SetStatusCode(status int) *RolesAddServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'add' method.
func (r *RolesAddServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RolesServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RolesServerAdapter struct {
	server RolesServer
	router *mux.Router
}

func NewRolesServerAdapter(server RolesServer, router *mux.Router) *RolesServerAdapter {
	adapter := new(RolesServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.roleHandler)
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.listHandler)
	adapter.router.Methods("POST").Path("").HandlerFunc(adapter.addHandler)
	return adapter
}
func (a *RolesServerAdapter) roleHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Role(id)
	targetAdapter := NewRoleServerAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RolesServerAdapter) readRolesListServerRequest(r *http.Request) (*RolesListServerRequest, error) {
	var err error
	result := new(RolesListServerRequest)
	result.query = r.URL.Query()
	result.path = r.URL.Path
	result.page, err = helpers.ParseInteger(result.query, "page")
	if err != nil {
		return nil, err
	}
	result.size, err = helpers.ParseInteger(result.query, "size")
	if err != nil {
		return nil, err
	}
	result.total, err = helpers.ParseInteger(result.query, "total")
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *RolesServerAdapter) writeRolesListServerResponse(w http.ResponseWriter, r *RolesListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RolesServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRolesListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RolesListServerResponse)
	err = a.server.List(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRolesListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RolesServerAdapter) readRolesAddServerRequest(r *http.Request) (*RolesAddServerRequest, error) {
	var err error
	result := new(RolesAddServerRequest)
	result.query = r.URL.Query()
	result.path = r.URL.Path
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *RolesServerAdapter) writeRolesAddServerResponse(w http.ResponseWriter, r *RolesAddServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RolesServerAdapter) addHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRolesAddServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RolesAddServerResponse)
	err = a.server.Add(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Add: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRolesAddServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RolesServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
