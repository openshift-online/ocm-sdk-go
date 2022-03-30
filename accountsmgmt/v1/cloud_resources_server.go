/*
Copyright (c) 2020 Red Hat, Inc.

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
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// CloudResourcesServer represents the interface the manages the 'cloud_resources' resource.
type CloudResourcesServer interface {

	// Add handles a request for the 'add' method.
	//
	// Creates a new cloud resource
	Add(ctx context.Context, request *CloudResourcesAddServerRequest, response *CloudResourcesAddServerResponse) error

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of cloud resources.
	List(ctx context.Context, request *CloudResourcesListServerRequest, response *CloudResourcesListServerResponse) error

	// CloudResource returns the target 'cloud_resource' server for the given identifier.
	//
	// Reference to the service that manages a specific cloud resource.
	CloudResource(id string) CloudResourceServer
}

// CloudResourcesAddServerRequest is the request for the 'add' method.
type CloudResourcesAddServerRequest struct {
	body *CloudResource
}

// Body returns the value of the 'body' parameter.
//
//
func (r *CloudResourcesAddServerRequest) Body() *CloudResource {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *CloudResourcesAddServerRequest) GetBody() (value *CloudResource, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// CloudResourcesAddServerResponse is the response for the 'add' method.
type CloudResourcesAddServerResponse struct {
	status int
	err    *errors.Error
	body   *CloudResource
}

// Body sets the value of the 'body' parameter.
//
//
func (r *CloudResourcesAddServerResponse) Body(value *CloudResource) *CloudResourcesAddServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *CloudResourcesAddServerResponse) Status(value int) *CloudResourcesAddServerResponse {
	r.status = value
	return r
}

// CloudResourcesListServerRequest is the request for the 'list' method.
type CloudResourcesListServerRequest struct {
	page   *int
	search *string
	size   *int
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *CloudResourcesListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
func (r *CloudResourcesListServerRequest) GetPage() (value int, ok bool) {
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
// of an SQL statement, but using the names of the attributes of the cloud resource
// instead of the names of the columns of a table.
//
// If the parameter isn't provided, or if the value is empty, then all the
// items that the user has permission to see will be returned.
func (r *CloudResourcesListServerRequest) Search() string {
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
// of an SQL statement, but using the names of the attributes of the cloud resource
// instead of the names of the columns of a table.
//
// If the parameter isn't provided, or if the value is empty, then all the
// items that the user has permission to see will be returned.
func (r *CloudResourcesListServerRequest) GetSearch() (value string, ok bool) {
	ok = r != nil && r.search != nil
	if ok {
		value = *r.search
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
func (r *CloudResourcesListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
func (r *CloudResourcesListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// CloudResourcesListServerResponse is the response for the 'list' method.
type CloudResourcesListServerResponse struct {
	status int
	err    *errors.Error
	items  *CloudResourceList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of cloud resources.
func (r *CloudResourcesListServerResponse) Items(value *CloudResourceList) *CloudResourcesListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *CloudResourcesListServerResponse) Page(value int) *CloudResourcesListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
func (r *CloudResourcesListServerResponse) Size(value int) *CloudResourcesListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *CloudResourcesListServerResponse) Total(value int) *CloudResourcesListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *CloudResourcesListServerResponse) Status(value int) *CloudResourcesListServerResponse {
	r.status = value
	return r
}

// dispatchCloudResources navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchCloudResources(w http.ResponseWriter, r *http.Request, server CloudResourcesServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "POST":
			adaptCloudResourcesAddRequest(w, r, server)
			return
		case "GET":
			adaptCloudResourcesListRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.CloudResource(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchCloudResource(w, r, target, segments[1:])
	}
}

// adaptCloudResourcesAddRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptCloudResourcesAddRequest(w http.ResponseWriter, r *http.Request, server CloudResourcesServer) {
	request := &CloudResourcesAddServerRequest{}
	err := readCloudResourcesAddRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &CloudResourcesAddServerResponse{}
	response.status = 201
	err = server.Add(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeCloudResourcesAddResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptCloudResourcesListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptCloudResourcesListRequest(w http.ResponseWriter, r *http.Request, server CloudResourcesServer) {
	request := &CloudResourcesListServerRequest{}
	err := readCloudResourcesListRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &CloudResourcesListServerResponse{}
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
	err = writeCloudResourcesListResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
