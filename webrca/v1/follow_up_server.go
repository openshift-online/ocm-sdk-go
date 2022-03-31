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

// FollowUpServer represents the interface the manages the 'follow_up' resource.
type FollowUpServer interface {

	// Delete handles a request for the 'delete' method.
	//
	//
	Delete(ctx context.Context, request *FollowUpDeleteServerRequest, response *FollowUpDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	//
	Get(ctx context.Context, request *FollowUpGetServerRequest, response *FollowUpGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	//
	Update(ctx context.Context, request *FollowUpUpdateServerRequest, response *FollowUpUpdateServerResponse) error
}

// FollowUpDeleteServerRequest is the request for the 'delete' method.
type FollowUpDeleteServerRequest struct {
}

// FollowUpDeleteServerResponse is the response for the 'delete' method.
type FollowUpDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *FollowUpDeleteServerResponse) Status(value int) *FollowUpDeleteServerResponse {
	r.status = value
	return r
}

// FollowUpGetServerRequest is the request for the 'get' method.
type FollowUpGetServerRequest struct {
}

// FollowUpGetServerResponse is the response for the 'get' method.
type FollowUpGetServerResponse struct {
	status int
	err    *errors.Error
	body   *FollowUp
}

// Body sets the value of the 'body' parameter.
//
//
func (r *FollowUpGetServerResponse) Body(value *FollowUp) *FollowUpGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *FollowUpGetServerResponse) Status(value int) *FollowUpGetServerResponse {
	r.status = value
	return r
}

// FollowUpUpdateServerRequest is the request for the 'update' method.
type FollowUpUpdateServerRequest struct {
	body *FollowUp
}

// Body returns the value of the 'body' parameter.
//
//
func (r *FollowUpUpdateServerRequest) Body() *FollowUp {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *FollowUpUpdateServerRequest) GetBody() (value *FollowUp, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// FollowUpUpdateServerResponse is the response for the 'update' method.
type FollowUpUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *FollowUp
}

// Body sets the value of the 'body' parameter.
//
//
func (r *FollowUpUpdateServerResponse) Body(value *FollowUp) *FollowUpUpdateServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *FollowUpUpdateServerResponse) Status(value int) *FollowUpUpdateServerResponse {
	r.status = value
	return r
}

// dispatchFollowUp navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchFollowUp(w http.ResponseWriter, r *http.Request, server FollowUpServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptFollowUpDeleteRequest(w, r, server)
			return
		case "GET":
			adaptFollowUpGetRequest(w, r, server)
			return
		case "PATCH":
			adaptFollowUpUpdateRequest(w, r, server)
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

// adaptFollowUpDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptFollowUpDeleteRequest(w http.ResponseWriter, r *http.Request, server FollowUpServer) {
	request := &FollowUpDeleteServerRequest{}
	err := readFollowUpDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &FollowUpDeleteServerResponse{}
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
	err = writeFollowUpDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptFollowUpGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptFollowUpGetRequest(w http.ResponseWriter, r *http.Request, server FollowUpServer) {
	request := &FollowUpGetServerRequest{}
	err := readFollowUpGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &FollowUpGetServerResponse{}
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
	err = writeFollowUpGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptFollowUpUpdateRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptFollowUpUpdateRequest(w http.ResponseWriter, r *http.Request, server FollowUpServer) {
	request := &FollowUpUpdateServerRequest{}
	err := readFollowUpUpdateRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &FollowUpUpdateServerResponse{}
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
	err = writeFollowUpUpdateResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
