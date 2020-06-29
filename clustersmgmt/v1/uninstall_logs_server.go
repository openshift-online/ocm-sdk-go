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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// UninstallLogsServer represents the interface the manages the 'uninstall_logs' resource.
type UninstallLogsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of install logs.
	List(ctx context.Context, request *UninstallLogsListServerRequest, response *UninstallLogsListServerResponse) error

	// Log returns the target 'log' server for the given identifier.
	//
	// Returns a reference to the service that manages an specific log.
	Log(id string) LogServer
}

// UninstallLogsListServerRequest is the request for the 'list' method.
type UninstallLogsListServerRequest struct {
	page *int
	size *int
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *UninstallLogsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
func (r *UninstallLogsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *UninstallLogsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Number of items contained in the returned page.
func (r *UninstallLogsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// UninstallLogsListServerResponse is the response for the 'list' method.
type UninstallLogsListServerResponse struct {
	status int
	err    *errors.Error
	items  *LogList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of install logs.
func (r *UninstallLogsListServerResponse) Items(value *LogList) *UninstallLogsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *UninstallLogsListServerResponse) Page(value int) *UninstallLogsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *UninstallLogsListServerResponse) Size(value int) *UninstallLogsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *UninstallLogsListServerResponse) Total(value int) *UninstallLogsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *UninstallLogsListServerResponse) Status(value int) *UninstallLogsListServerResponse {
	r.status = value
	return r
}

// dispatchUninstallLogs navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchUninstallLogs(w http.ResponseWriter, r *http.Request, server UninstallLogsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptUninstallLogsListRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.Log(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchLog(w, r, target, segments[1:])
	}
}

// adaptUninstallLogsListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptUninstallLogsListRequest(w http.ResponseWriter, r *http.Request, server UninstallLogsServer) {
	request := &UninstallLogsListServerRequest{}
	err := readUninstallLogsListRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &UninstallLogsListServerResponse{}
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
	err = writeUninstallLogsListResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
