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

// EventServer represents the interface the manages the 'event' resource.
type EventServer interface {

	// Delete handles a request for the 'delete' method.
	//
	//
	Delete(ctx context.Context, request *EventDeleteServerRequest, response *EventDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	//
	Get(ctx context.Context, request *EventGetServerRequest, response *EventGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	//
	Update(ctx context.Context, request *EventUpdateServerRequest, response *EventUpdateServerResponse) error

	// Attachments returns the target 'attachments' resource.
	//
	//
	Attachments() AttachmentsServer
}

// EventDeleteServerRequest is the request for the 'delete' method.
type EventDeleteServerRequest struct {
}

// EventDeleteServerResponse is the response for the 'delete' method.
type EventDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *EventDeleteServerResponse) Status(value int) *EventDeleteServerResponse {
	r.status = value
	return r
}

// EventGetServerRequest is the request for the 'get' method.
type EventGetServerRequest struct {
}

// EventGetServerResponse is the response for the 'get' method.
type EventGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Event
}

// Body sets the value of the 'body' parameter.
//
//
func (r *EventGetServerResponse) Body(value *Event) *EventGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *EventGetServerResponse) Status(value int) *EventGetServerResponse {
	r.status = value
	return r
}

// EventUpdateServerRequest is the request for the 'update' method.
type EventUpdateServerRequest struct {
	body *Event
}

// Body returns the value of the 'body' parameter.
//
//
func (r *EventUpdateServerRequest) Body() *Event {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventUpdateServerRequest) GetBody() (value *Event, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// EventUpdateServerResponse is the response for the 'update' method.
type EventUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *Event
}

// Body sets the value of the 'body' parameter.
//
//
func (r *EventUpdateServerResponse) Body(value *Event) *EventUpdateServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *EventUpdateServerResponse) Status(value int) *EventUpdateServerResponse {
	r.status = value
	return r
}

// dispatchEvent navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchEvent(w http.ResponseWriter, r *http.Request, server EventServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptEventDeleteRequest(w, r, server)
			return
		case "GET":
			adaptEventGetRequest(w, r, server)
			return
		case "PATCH":
			adaptEventUpdateRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	case "attachments":
		target := server.Attachments()
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchAttachments(w, r, target, segments[1:])
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptEventDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptEventDeleteRequest(w http.ResponseWriter, r *http.Request, server EventServer) {
	request := &EventDeleteServerRequest{}
	err := readEventDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &EventDeleteServerResponse{}
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
	err = writeEventDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptEventGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptEventGetRequest(w http.ResponseWriter, r *http.Request, server EventServer) {
	request := &EventGetServerRequest{}
	err := readEventGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &EventGetServerResponse{}
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
	err = writeEventGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptEventUpdateRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptEventUpdateRequest(w http.ResponseWriter, r *http.Request, server EventServer) {
	request := &EventUpdateServerRequest{}
	err := readEventUpdateRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &EventUpdateServerResponse{}
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
	err = writeEventUpdateResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
