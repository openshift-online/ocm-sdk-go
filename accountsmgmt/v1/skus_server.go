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

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// SKUSServer represents the interface the manages the 'SKUS' resource.
type SKUSServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves a list of SKUS.
	List(ctx context.Context, request *SKUSListServerRequest, response *SKUSListServerResponse) error

	// SKU returns the target 'SKU' server for the given identifier.
	//
	// Reference to the service that manages a specific SKU.
	SKU(id string) SKUServer
}

// SKUSListServerRequest is the request for the 'list' method.
type SKUSListServerRequest struct {
	page   *int
	search *string
	size   *int
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *SKUSListServerRequest) Page() int {
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
func (r *SKUSListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Search returns the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause
// of an SQL statement, but using the names of the attributes of the SKU
// instead of the names of the columns of a table. For example, in order to
// retrieve SKUS large sized resources:
//
// [source,sql]
// ----
// resource_name like '%large'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the
// items that the user has permission to see will be returned.
func (r *SKUSListServerRequest) Search() string {
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
// of an SQL statement, but using the names of the attributes of the SKU
// instead of the names of the columns of a table. For example, in order to
// retrieve SKUS large sized resources:
//
// [source,sql]
// ----
// resource_name like '%large'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the
// items that the user has permission to see will be returned.
func (r *SKUSListServerRequest) GetSearch() (value string, ok bool) {
	ok = r != nil && r.search != nil
	if ok {
		value = *r.search
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *SKUSListServerRequest) Size() int {
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
func (r *SKUSListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// SKUSListServerResponse is the response for the 'list' method.
type SKUSListServerResponse struct {
	status int
	err    *errors.Error
	items  *SKUList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of SKUS.
func (r *SKUSListServerResponse) Items(value *SKUList) *SKUSListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *SKUSListServerResponse) Page(value int) *SKUSListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *SKUSListServerResponse) Size(value int) *SKUSListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *SKUSListServerResponse) Total(value int) *SKUSListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *SKUSListServerResponse) Status(value int) *SKUSListServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *SKUSListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(skusListServerResponseData)
	data.Items, err = r.items.wrap()
	if err != nil {
		return err
	}
	data.Page = r.page
	data.Size = r.size
	data.Total = r.total
	err = encoder.Encode(data)
	return err
}

// skusListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type skusListServerResponseData struct {
	Items skuListData "json:\"items,omitempty\""
	Page  *int        "json:\"page,omitempty\""
	Size  *int        "json:\"size,omitempty\""
	Total *int        "json:\"total,omitempty\""
}

// SKUSAdapter represents the structs that adapts Requests and Response to internal
// structs.
type SKUSAdapter struct {
	server SKUSServer
	router *mux.Router
}

func NewSKUSAdapter(server SKUSServer, router *mux.Router) *SKUSAdapter {
	adapter := new(SKUSAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.skuHandler)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerList)
	return adapter
}
func (a *SKUSAdapter) skuHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.SKU(id)
	targetAdapter := NewSKUAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *SKUSAdapter) readListRequest(r *http.Request) (*SKUSListServerRequest, error) {
	var err error
	result := new(SKUSListServerRequest)
	query := r.URL.Query()
	result.page, err = helpers.ParseInteger(query, "page")
	if err != nil {
		return nil, err
	}
	result.search, err = helpers.ParseString(query, "search")
	if err != nil {
		return nil, err
	}
	result.size, err = helpers.ParseInteger(query, "size")
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *SKUSAdapter) writeListResponse(w http.ResponseWriter, r *SKUSListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *SKUSAdapter) handlerList(w http.ResponseWriter, r *http.Request) {
	request, err := a.readListRequest(r)
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
	response := new(SKUSListServerResponse)
	response.status = http.StatusOK
	err = a.server.List(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method List: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeListResponse(w, response)
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
func (a *SKUSAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
