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

package v1 // github.com/openshift-online/ocm-sdk-go/webrca/v1

import (
	"context"
	"net/http"
	time "time"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// FollowUpsServer represents the interface the manages the 'follow_ups' resource.
type FollowUpsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of follow-ups
	List(ctx context.Context, request *FollowUpsListServerRequest, response *FollowUpsListServerResponse) error

	// FollowUp returns the target 'follow_up' server for the given identifier.
	//
	//
	FollowUp(id string) FollowUpServer
}

// FollowUpsListServerRequest is the request for the 'list' method.
type FollowUpsListServerRequest struct {
	createdAfter   *time.Time
	createdBefore  *time.Time
	followUpStatus *string
	orderBy        *string
	page           *int
	size           *int
}

// CreatedAfter returns the value of the 'created_after' parameter.
//
//
func (r *FollowUpsListServerRequest) CreatedAfter() time.Time {
	if r != nil && r.createdAfter != nil {
		return *r.createdAfter
	}
	return time.Time{}
}

// GetCreatedAfter returns the value of the 'created_after' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *FollowUpsListServerRequest) GetCreatedAfter() (value time.Time, ok bool) {
	ok = r != nil && r.createdAfter != nil
	if ok {
		value = *r.createdAfter
	}
	return
}

// CreatedBefore returns the value of the 'created_before' parameter.
//
//
func (r *FollowUpsListServerRequest) CreatedBefore() time.Time {
	if r != nil && r.createdBefore != nil {
		return *r.createdBefore
	}
	return time.Time{}
}

// GetCreatedBefore returns the value of the 'created_before' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *FollowUpsListServerRequest) GetCreatedBefore() (value time.Time, ok bool) {
	ok = r != nil && r.createdBefore != nil
	if ok {
		value = *r.createdBefore
	}
	return
}

// FollowUpStatus returns the value of the 'follow_up_status' parameter.
//
//
func (r *FollowUpsListServerRequest) FollowUpStatus() string {
	if r != nil && r.followUpStatus != nil {
		return *r.followUpStatus
	}
	return ""
}

// GetFollowUpStatus returns the value of the 'follow_up_status' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *FollowUpsListServerRequest) GetFollowUpStatus() (value string, ok bool) {
	ok = r != nil && r.followUpStatus != nil
	if ok {
		value = *r.followUpStatus
	}
	return
}

// OrderBy returns the value of the 'order_by' parameter.
//
//
func (r *FollowUpsListServerRequest) OrderBy() string {
	if r != nil && r.orderBy != nil {
		return *r.orderBy
	}
	return ""
}

// GetOrderBy returns the value of the 'order_by' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *FollowUpsListServerRequest) GetOrderBy() (value string, ok bool) {
	ok = r != nil && r.orderBy != nil
	if ok {
		value = *r.orderBy
	}
	return
}

// Page returns the value of the 'page' parameter.
//
//
func (r *FollowUpsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *FollowUpsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
//
func (r *FollowUpsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *FollowUpsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// FollowUpsListServerResponse is the response for the 'list' method.
type FollowUpsListServerResponse struct {
	status int
	err    *errors.Error
	items  *FollowUpList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
//
func (r *FollowUpsListServerResponse) Items(value *FollowUpList) *FollowUpsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
//
func (r *FollowUpsListServerResponse) Page(value int) *FollowUpsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
//
func (r *FollowUpsListServerResponse) Size(value int) *FollowUpsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
//
func (r *FollowUpsListServerResponse) Total(value int) *FollowUpsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *FollowUpsListServerResponse) Status(value int) *FollowUpsListServerResponse {
	r.status = value
	return r
}

// dispatchFollowUps navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchFollowUps(w http.ResponseWriter, r *http.Request, server FollowUpsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptFollowUpsListRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.FollowUp(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchFollowUp(w, r, target, segments[1:])
	}
}

// adaptFollowUpsListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptFollowUpsListRequest(w http.ResponseWriter, r *http.Request, server FollowUpsServer) {
	request := &FollowUpsListServerRequest{}
	err := readFollowUpsListRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &FollowUpsListServerResponse{}
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
	err = writeFollowUpsListResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
