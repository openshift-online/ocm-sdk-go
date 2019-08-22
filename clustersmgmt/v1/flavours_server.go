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

// FlavoursServer represents the interface the manages the 'flavours' resource.
type FlavoursServer interface {

	// List handles a request for the 'list' method.
	//
	//
	List(ctx context.Context, request *FlavoursListServerRequest, response *FlavoursListServerResponse) error

	// Add handles a request for the 'add' method.
	//
	// Adds a new cluster flavour.
	Add(ctx context.Context, request *FlavoursAddServerRequest, response *FlavoursAddServerResponse) error

	// Flavour returns the target 'flavour' server for the given identifier.
	//
	// Reference to the resource that manages a specific flavour.
	Flavour(id string) FlavourServer
}

// FlavoursListServerRequest is the request for the 'list' method.
type FlavoursListServerRequest struct {
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
func (r *FlavoursListServerRequest) Page() int {
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
func (r *FlavoursListServerRequest) GetPage() (value int, ok bool) {
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
func (r *FlavoursListServerRequest) Size() int {
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
func (r *FlavoursListServerRequest) GetSize() (value int, ok bool) {
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
// The syntax of this parameter is similar to the syntax of the _where_ clause
// of an SQL statement, but using the names of the attributes of the cluster
// instead of the names of the columns of a table. For example, in order to
// retrieve all the flavours with a name starting with `my`the value should be:
//
// [source,sql]
// ----
// name like 'my%'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the
// flavours that the user has permission to see will be returned.
func (r *FlavoursListServerRequest) Search() string {
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
// The syntax of this parameter is similar to the syntax of the _where_ clause
// of an SQL statement, but using the names of the attributes of the cluster
// instead of the names of the columns of a table. For example, in order to
// retrieve all the flavours with a name starting with `my`the value should be:
//
// [source,sql]
// ----
// name like 'my%'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the
// flavours that the user has permission to see will be returned.
func (r *FlavoursListServerRequest) GetSearch() (value string, ok bool) {
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
func (r *FlavoursListServerRequest) Total() int {
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
func (r *FlavoursListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// FlavoursListServerResponse is the response for the 'list' method.
type FlavoursListServerResponse struct {
	status int
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *FlavourList
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *FlavoursListServerResponse) Page(value int) *FlavoursListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *FlavoursListServerResponse) Size(value int) *FlavoursListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *FlavoursListServerResponse) Total(value int) *FlavoursListServerResponse {
	r.total = &value
	return r
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of flavours.
func (r *FlavoursListServerResponse) Items(value *FlavourList) *FlavoursListServerResponse {
	r.items = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *FlavoursListServerResponse) SetStatusCode(status int) *FlavoursListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *FlavoursListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(flavoursListServerResponseData)
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

// flavoursListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type flavoursListServerResponseData struct {
	Page  *int            "json:\"page,omitempty\""
	Size  *int            "json:\"size,omitempty\""
	Total *int            "json:\"total,omitempty\""
	Items flavourListData "json:\"items,omitempty\""
}

// FlavoursAddServerRequest is the request for the 'add' method.
type FlavoursAddServerRequest struct {
	path  string
	query url.Values
	body  *Flavour
}

// Body returns the value of the 'body' parameter.
//
// Details of the cluster flavour.
func (r *FlavoursAddServerRequest) Body() *Flavour {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Details of the cluster flavour.
func (r *FlavoursAddServerRequest) GetBody() (value *Flavour, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'add' method.
func (r *FlavoursAddServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(flavourData)
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

// FlavoursAddServerResponse is the response for the 'add' method.
type FlavoursAddServerResponse struct {
	status int
	err    *errors.Error
	body   *Flavour
}

// Body sets the value of the 'body' parameter.
//
// Details of the cluster flavour.
func (r *FlavoursAddServerResponse) Body(value *Flavour) *FlavoursAddServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *FlavoursAddServerResponse) SetStatusCode(status int) *FlavoursAddServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'add' method.
func (r *FlavoursAddServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// FlavoursServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type FlavoursServerAdapter struct {
	server FlavoursServer
	router *mux.Router
}

func NewFlavoursServerAdapter(server FlavoursServer, router *mux.Router) *FlavoursServerAdapter {
	adapter := new(FlavoursServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}/").HandlerFunc(adapter.flavourHandler)
	adapter.router.HandleFunc("/", adapter.listHandler).Methods("GET")
	adapter.router.HandleFunc("/", adapter.addHandler).Methods("POST")
	return adapter
}
func (a *FlavoursServerAdapter) flavourHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Flavour(id)
	targetAdapter := NewFlavourServerAdapter(target, a.router.PathPrefix("/{id}/").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *FlavoursServerAdapter) readFlavoursListServerRequest(r *http.Request) (*FlavoursListServerRequest, error) {
	result := new(FlavoursListServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *FlavoursServerAdapter) writeFlavoursListServerResponse(w http.ResponseWriter, r *FlavoursListServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *FlavoursServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readFlavoursListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(FlavoursListServerResponse)
	err = a.server.List(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeFlavoursListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *FlavoursServerAdapter) readFlavoursAddServerRequest(r *http.Request) (*FlavoursAddServerRequest, error) {
	result := new(FlavoursAddServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *FlavoursServerAdapter) writeFlavoursAddServerResponse(w http.ResponseWriter, r *FlavoursAddServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *FlavoursServerAdapter) addHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readFlavoursAddServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(FlavoursAddServerResponse)
	err = a.server.Add(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Add: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeFlavoursAddServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *FlavoursServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
