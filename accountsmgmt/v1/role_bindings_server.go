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

// RoleBindingsServer represents the interface the manages the 'role_bindings' resource.
type RoleBindingsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves a list of role bindings.
	List(ctx context.Context, request *RoleBindingsListServerRequest, response *RoleBindingsListServerResponse) error

	// Add handles a request for the 'add' method.
	//
	// Creates a new role binding.
	Add(ctx context.Context, request *RoleBindingsAddServerRequest, response *RoleBindingsAddServerResponse) error

	// RoleBinding returns the target 'role_binding' server for the given identifier.
	//
	// Reference to the service that manages a specific role binding.
	RoleBinding(id string) RoleBindingServer
}

// RoleBindingsListServerRequest is the request for the 'list' method.
type RoleBindingsListServerRequest struct {
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
func (r *RoleBindingsListServerRequest) Page() int {
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
func (r *RoleBindingsListServerRequest) GetPage() (value int, ok bool) {
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
func (r *RoleBindingsListServerRequest) Size() int {
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
func (r *RoleBindingsListServerRequest) GetSize() (value int, ok bool) {
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
func (r *RoleBindingsListServerRequest) Total() int {
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
func (r *RoleBindingsListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// RoleBindingsListServerResponse is the response for the 'list' method.
type RoleBindingsListServerResponse struct {
	status int
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *RoleBindingList
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RoleBindingsListServerResponse) Page(value int) *RoleBindingsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *RoleBindingsListServerResponse) Size(value int) *RoleBindingsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RoleBindingsListServerResponse) Total(value int) *RoleBindingsListServerResponse {
	r.total = &value
	return r
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of role bindings.
func (r *RoleBindingsListServerResponse) Items(value *RoleBindingList) *RoleBindingsListServerResponse {
	r.items = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RoleBindingsListServerResponse) SetStatusCode(status int) *RoleBindingsListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *RoleBindingsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(roleBindingsListServerResponseData)
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

// roleBindingsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type roleBindingsListServerResponseData struct {
	Page  *int                "json:\"page,omitempty\""
	Size  *int                "json:\"size,omitempty\""
	Total *int                "json:\"total,omitempty\""
	Items roleBindingListData "json:\"items,omitempty\""
}

// RoleBindingsAddServerRequest is the request for the 'add' method.
type RoleBindingsAddServerRequest struct {
	path  string
	query url.Values
	body  *RoleBinding
}

// Body returns the value of the 'body' parameter.
//
// Role binding data.
func (r *RoleBindingsAddServerRequest) Body() *RoleBinding {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Role binding data.
func (r *RoleBindingsAddServerRequest) GetBody() (value *RoleBinding, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'add' method.
func (r *RoleBindingsAddServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(roleBindingData)
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

// RoleBindingsAddServerResponse is the response for the 'add' method.
type RoleBindingsAddServerResponse struct {
	status int
	err    *errors.Error
	body   *RoleBinding
}

// Body sets the value of the 'body' parameter.
//
// Role binding data.
func (r *RoleBindingsAddServerResponse) Body(value *RoleBinding) *RoleBindingsAddServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RoleBindingsAddServerResponse) SetStatusCode(status int) *RoleBindingsAddServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'add' method.
func (r *RoleBindingsAddServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RoleBindingsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RoleBindingsServerAdapter struct {
	server RoleBindingsServer
	router *mux.Router
}

func NewRoleBindingsServerAdapter(server RoleBindingsServer, router *mux.Router) *RoleBindingsServerAdapter {
	adapter := new(RoleBindingsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.roleBindingHandler)
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.listHandler)
	adapter.router.Methods("POST").Path("").HandlerFunc(adapter.addHandler)
	return adapter
}
func (a *RoleBindingsServerAdapter) roleBindingHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.RoleBinding(id)
	targetAdapter := NewRoleBindingServerAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RoleBindingsServerAdapter) readRoleBindingsListServerRequest(r *http.Request) (*RoleBindingsListServerRequest, error) {
	result := new(RoleBindingsListServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *RoleBindingsServerAdapter) writeRoleBindingsListServerResponse(w http.ResponseWriter, r *RoleBindingsListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RoleBindingsServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRoleBindingsListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RoleBindingsListServerResponse)
	err = a.server.List(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRoleBindingsListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RoleBindingsServerAdapter) readRoleBindingsAddServerRequest(r *http.Request) (*RoleBindingsAddServerRequest, error) {
	result := new(RoleBindingsAddServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *RoleBindingsServerAdapter) writeRoleBindingsAddServerResponse(w http.ResponseWriter, r *RoleBindingsAddServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RoleBindingsServerAdapter) addHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRoleBindingsAddServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RoleBindingsAddServerResponse)
	err = a.server.Add(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Add: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRoleBindingsAddServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RoleBindingsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
