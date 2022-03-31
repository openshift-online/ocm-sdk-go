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

// IncidentServer represents the interface the manages the 'incident' resource.
type IncidentServer interface {

	// Delete handles a request for the 'delete' method.
	//
	//
	Delete(ctx context.Context, request *IncidentDeleteServerRequest, response *IncidentDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	//
	Get(ctx context.Context, request *IncidentGetServerRequest, response *IncidentGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	//
	Update(ctx context.Context, request *IncidentUpdateServerRequest, response *IncidentUpdateServerResponse) error

	// Events returns the target 'events' resource.
	//
	//
	Events() EventsServer

	// FollowUps returns the target 'follow_ups' resource.
	//
	//
	FollowUps() FollowUpsServer

	// Notifications returns the target 'notifications' resource.
	//
	//
	Notifications() NotificationsServer
}

// IncidentDeleteServerRequest is the request for the 'delete' method.
type IncidentDeleteServerRequest struct {
}

// IncidentDeleteServerResponse is the response for the 'delete' method.
type IncidentDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *IncidentDeleteServerResponse) Status(value int) *IncidentDeleteServerResponse {
	r.status = value
	return r
}

// IncidentGetServerRequest is the request for the 'get' method.
type IncidentGetServerRequest struct {
}

// IncidentGetServerResponse is the response for the 'get' method.
type IncidentGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Incident
}

// Body sets the value of the 'body' parameter.
//
//
func (r *IncidentGetServerResponse) Body(value *Incident) *IncidentGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *IncidentGetServerResponse) Status(value int) *IncidentGetServerResponse {
	r.status = value
	return r
}

// IncidentUpdateServerRequest is the request for the 'update' method.
type IncidentUpdateServerRequest struct {
	body *Incident
}

// Body returns the value of the 'body' parameter.
//
//
func (r *IncidentUpdateServerRequest) Body() *Incident {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentUpdateServerRequest) GetBody() (value *Incident, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// IncidentUpdateServerResponse is the response for the 'update' method.
type IncidentUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *Incident
}

// Body sets the value of the 'body' parameter.
//
//
func (r *IncidentUpdateServerResponse) Body(value *Incident) *IncidentUpdateServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *IncidentUpdateServerResponse) Status(value int) *IncidentUpdateServerResponse {
	r.status = value
	return r
}

// dispatchIncident navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchIncident(w http.ResponseWriter, r *http.Request, server IncidentServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptIncidentDeleteRequest(w, r, server)
			return
		case "GET":
			adaptIncidentGetRequest(w, r, server)
			return
		case "PATCH":
			adaptIncidentUpdateRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	case "events":
		target := server.Events()
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchEvents(w, r, target, segments[1:])
	case "follow_ups":
		target := server.FollowUps()
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchFollowUps(w, r, target, segments[1:])
	case "notifications":
		target := server.Notifications()
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchNotifications(w, r, target, segments[1:])
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptIncidentDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptIncidentDeleteRequest(w http.ResponseWriter, r *http.Request, server IncidentServer) {
	request := &IncidentDeleteServerRequest{}
	err := readIncidentDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &IncidentDeleteServerResponse{}
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
	err = writeIncidentDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptIncidentGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptIncidentGetRequest(w http.ResponseWriter, r *http.Request, server IncidentServer) {
	request := &IncidentGetServerRequest{}
	err := readIncidentGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &IncidentGetServerResponse{}
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
	err = writeIncidentGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptIncidentUpdateRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptIncidentUpdateRequest(w http.ResponseWriter, r *http.Request, server IncidentServer) {
	request := &IncidentUpdateServerRequest{}
	err := readIncidentUpdateRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &IncidentUpdateServerResponse{}
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
	err = writeIncidentUpdateResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
