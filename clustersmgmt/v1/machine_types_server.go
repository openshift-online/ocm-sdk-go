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
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MachineTypesServer represents the interface the manages the 'machine_types' resource.
type MachineTypesServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of machine types.
	List(ctx context.Context, request *MachineTypesListServerRequest, response *MachineTypesListServerResponse) error
}

// MachineTypesListServerRequest is the request for the 'list' method.
type MachineTypesListServerRequest struct {
	order  *string
	page   *int
	search *string
	size   *int
}

// Order returns the value of the 'order' parameter.
//
// Order criteria.
//
// The syntax of this parameter is similar to the syntax of the _order by_ clause of
// a SQL statement, but using the names of the attributes of the machine type
// instead of the names of the columns of a table. For example, in order to sort the
// machine types descending by name identifier the value should be:
//
// [source,sql]
// ----
// name desc
// ----
//
// If the parameter isn't provided, or if the value is empty, then the order of the
// results is undefined.
func (r *MachineTypesListServerRequest) Order() string {
	if r != nil && r.order != nil {
		return *r.order
	}
	return ""
}

// GetOrder returns the value of the 'order' parameter and
// a flag indicating if the parameter has a value.
//
// Order criteria.
//
// The syntax of this parameter is similar to the syntax of the _order by_ clause of
// a SQL statement, but using the names of the attributes of the machine type
// instead of the names of the columns of a table. For example, in order to sort the
// machine types descending by name identifier the value should be:
//
// [source,sql]
// ----
// name desc
// ----
//
// If the parameter isn't provided, or if the value is empty, then the order of the
// results is undefined.
func (r *MachineTypesListServerRequest) GetOrder() (value string, ok bool) {
	ok = r != nil && r.order != nil
	if ok {
		value = *r.order
	}
	return
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *MachineTypesListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
func (r *MachineTypesListServerRequest) GetPage() (value int, ok bool) {
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
// The syntax of this parameter is similar to the syntax of the _where_ clause of a
// SQL statement, but using the names of the attributes of the machine type
// instead of the names of the columns of a table. For example, in order to retrieve
// all the machine types with a name starting with `A` the value should be:
//
// [source,sql]
// ----
// name like 'A%'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the machine
// types that the user has permission to see will be returned.
func (r *MachineTypesListServerRequest) Search() string {
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
// The syntax of this parameter is similar to the syntax of the _where_ clause of a
// SQL statement, but using the names of the attributes of the machine type
// instead of the names of the columns of a table. For example, in order to retrieve
// all the machine types with a name starting with `A` the value should be:
//
// [source,sql]
// ----
// name like 'A%'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the machine
// types that the user has permission to see will be returned.
func (r *MachineTypesListServerRequest) GetSearch() (value string, ok bool) {
	ok = r != nil && r.search != nil
	if ok {
		value = *r.search
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
func (r *MachineTypesListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
func (r *MachineTypesListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// MachineTypesListServerResponse is the response for the 'list' method.
type MachineTypesListServerResponse struct {
	status int
	err    *errors.Error
	items  *MachineTypeList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of cloud providers.
func (r *MachineTypesListServerResponse) Items(value *MachineTypeList) *MachineTypesListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *MachineTypesListServerResponse) Page(value int) *MachineTypesListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
func (r *MachineTypesListServerResponse) Size(value int) *MachineTypesListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *MachineTypesListServerResponse) Total(value int) *MachineTypesListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *MachineTypesListServerResponse) Status(value int) *MachineTypesListServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *MachineTypesListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(machineTypesListServerResponseData)
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

// machineTypesListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type machineTypesListServerResponseData struct {
	Items machineTypeListData "json:\"items,omitempty\""
	Page  *int                "json:\"page,omitempty\""
	Size  *int                "json:\"size,omitempty\""
	Total *int                "json:\"total,omitempty\""
}

// dispatchMachineTypes navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchMachineTypes(w http.ResponseWriter, r *http.Request, server MachineTypesServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptMachineTypesListRequest(w, r, server)
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	} else {
		switch segments[0] {
		default:
			errors.SendNotFound(w, r)
			return
		}
	}
}

// readMachineTypesListRequest reads the given HTTP requests and translates it
// into an object of type MachineTypesListServerRequest.
func readMachineTypesListRequest(r *http.Request) (*MachineTypesListServerRequest, error) {
	var err error
	result := new(MachineTypesListServerRequest)
	query := r.URL.Query()
	result.order, err = helpers.ParseString(query, "order")
	if err != nil {
		return nil, err
	}
	result.page, err = helpers.ParseInteger(query, "page")
	if err != nil {
		return nil, err
	}
	if result.page == nil {
		result.page = helpers.NewInteger(1)
	}
	result.search, err = helpers.ParseString(query, "search")
	if err != nil {
		return nil, err
	}
	result.size, err = helpers.ParseInteger(query, "size")
	if err != nil {
		return nil, err
	}
	if result.size == nil {
		result.size = helpers.NewInteger(100)
	}
	return result, err
}

// writeMachineTypesListResponse translates the given request object into an
// HTTP response.
func writeMachineTypesListResponse(w http.ResponseWriter, r *MachineTypesListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}

// adaptMachineTypesListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptMachineTypesListRequest(w http.ResponseWriter, r *http.Request, server MachineTypesServer) {
	request, err := readMachineTypesListRequest(r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := new(MachineTypesListServerResponse)
	response.status = 200
	err = server.List(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeMachineTypesListResponse(w, response)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
