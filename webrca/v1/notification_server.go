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

// NotificationServer represents the interface the manages the 'notification' resource.
type NotificationServer interface {

	// Delete handles a request for the 'delete' method.
	//
	//
	Delete(ctx context.Context, request *NotificationDeleteServerRequest, response *NotificationDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	//
	Get(ctx context.Context, request *NotificationGetServerRequest, response *NotificationGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	//
	Update(ctx context.Context, request *NotificationUpdateServerRequest, response *NotificationUpdateServerResponse) error
}

// NotificationDeleteServerRequest is the request for the 'delete' method.
type NotificationDeleteServerRequest struct {
}

// NotificationDeleteServerResponse is the response for the 'delete' method.
type NotificationDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *NotificationDeleteServerResponse) Status(value int) *NotificationDeleteServerResponse {
	r.status = value
	return r
}

// NotificationGetServerRequest is the request for the 'get' method.
type NotificationGetServerRequest struct {
}

// NotificationGetServerResponse is the response for the 'get' method.
type NotificationGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Notification
}

// Body sets the value of the 'body' parameter.
//
//
func (r *NotificationGetServerResponse) Body(value *Notification) *NotificationGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *NotificationGetServerResponse) Status(value int) *NotificationGetServerResponse {
	r.status = value
	return r
}

// NotificationUpdateServerRequest is the request for the 'update' method.
type NotificationUpdateServerRequest struct {
	body *Notification
}

// Body returns the value of the 'body' parameter.
//
//
func (r *NotificationUpdateServerRequest) Body() *Notification {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *NotificationUpdateServerRequest) GetBody() (value *Notification, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// NotificationUpdateServerResponse is the response for the 'update' method.
type NotificationUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *Notification
}

// Body sets the value of the 'body' parameter.
//
//
func (r *NotificationUpdateServerResponse) Body(value *Notification) *NotificationUpdateServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *NotificationUpdateServerResponse) Status(value int) *NotificationUpdateServerResponse {
	r.status = value
	return r
}

// dispatchNotification navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchNotification(w http.ResponseWriter, r *http.Request, server NotificationServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptNotificationDeleteRequest(w, r, server)
			return
		case "GET":
			adaptNotificationGetRequest(w, r, server)
			return
		case "PATCH":
			adaptNotificationUpdateRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptNotificationDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptNotificationDeleteRequest(w http.ResponseWriter, r *http.Request, server NotificationServer) {
	request := &NotificationDeleteServerRequest{}
	err := readNotificationDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &NotificationDeleteServerResponse{}
	response.status = 204
	err = server.Delete(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeNotificationDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptNotificationGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptNotificationGetRequest(w http.ResponseWriter, r *http.Request, server NotificationServer) {
	request := &NotificationGetServerRequest{}
	err := readNotificationGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &NotificationGetServerResponse{}
	response.status = 200
	err = server.Get(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeNotificationGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptNotificationUpdateRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptNotificationUpdateRequest(w http.ResponseWriter, r *http.Request, server NotificationServer) {
	request := &NotificationUpdateServerRequest{}
	err := readNotificationUpdateRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &NotificationUpdateServerResponse{}
	response.status = 200
	err = server.Update(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeNotificationUpdateResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
