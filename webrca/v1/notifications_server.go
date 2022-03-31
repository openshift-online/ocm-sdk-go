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

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// NotificationsServer represents the interface the manages the 'notifications' resource.
type NotificationsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of notifications
	List(ctx context.Context, request *NotificationsListServerRequest, response *NotificationsListServerResponse) error

	// Notification returns the target 'notification' server for the given identifier.
	//
	//
	Notification(id string) NotificationServer
}

// NotificationsListServerRequest is the request for the 'list' method.
type NotificationsListServerRequest struct {
	checked *bool
	page    *int
	size    *int
}

// Checked returns the value of the 'checked' parameter.
//
//
func (r *NotificationsListServerRequest) Checked() bool {
	if r != nil && r.checked != nil {
		return *r.checked
	}
	return false
}

// GetChecked returns the value of the 'checked' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *NotificationsListServerRequest) GetChecked() (value bool, ok bool) {
	ok = r != nil && r.checked != nil
	if ok {
		value = *r.checked
	}
	return
}

// Page returns the value of the 'page' parameter.
//
//
func (r *NotificationsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *NotificationsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
//
func (r *NotificationsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *NotificationsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// NotificationsListServerResponse is the response for the 'list' method.
type NotificationsListServerResponse struct {
	status int
	err    *errors.Error
	items  *NotificationList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
//
func (r *NotificationsListServerResponse) Items(value *NotificationList) *NotificationsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
//
func (r *NotificationsListServerResponse) Page(value int) *NotificationsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
//
func (r *NotificationsListServerResponse) Size(value int) *NotificationsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
//
func (r *NotificationsListServerResponse) Total(value int) *NotificationsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *NotificationsListServerResponse) Status(value int) *NotificationsListServerResponse {
	r.status = value
	return r
}

// dispatchNotifications navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchNotifications(w http.ResponseWriter, r *http.Request, server NotificationsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptNotificationsListRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.Notification(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchNotification(w, r, target, segments[1:])
	}
}

// adaptNotificationsListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptNotificationsListRequest(w http.ResponseWriter, r *http.Request, server NotificationsServer) {
	request := &NotificationsListServerRequest{}
	err := readNotificationsListRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &NotificationsListServerResponse{}
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
	err = writeNotificationsListResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
