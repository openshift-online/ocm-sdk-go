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

// CloudResourceServer represents the interface the manages the 'cloud_resource' resource.
type CloudResourceServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the cloud resource.
	Delete(ctx context.Context, request *CloudResourceDeleteServerRequest, response *CloudResourceDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the cloud resource.
	Get(ctx context.Context, request *CloudResourceGetServerRequest, response *CloudResourceGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	// Updates the cloud resource.
	Update(ctx context.Context, request *CloudResourceUpdateServerRequest, response *CloudResourceUpdateServerResponse) error
}

// CloudResourceDeleteServerRequest is the request for the 'delete' method.
type CloudResourceDeleteServerRequest struct {
}

// CloudResourceDeleteServerResponse is the response for the 'delete' method.
type CloudResourceDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *CloudResourceDeleteServerResponse) Status(value int) *CloudResourceDeleteServerResponse {
	r.status = value
	return r
}

// CloudResourceGetServerRequest is the request for the 'get' method.
type CloudResourceGetServerRequest struct {
}

// CloudResourceGetServerResponse is the response for the 'get' method.
type CloudResourceGetServerResponse struct {
	status int
	err    *errors.Error
	body   *CloudResource
}

// Body sets the value of the 'body' parameter.
//
//
func (r *CloudResourceGetServerResponse) Body(value *CloudResource) *CloudResourceGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *CloudResourceGetServerResponse) Status(value int) *CloudResourceGetServerResponse {
	r.status = value
	return r
}

// CloudResourceUpdateServerRequest is the request for the 'update' method.
type CloudResourceUpdateServerRequest struct {
	body *CloudResource
}

// Body returns the value of the 'body' parameter.
//
//
func (r *CloudResourceUpdateServerRequest) Body() *CloudResource {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *CloudResourceUpdateServerRequest) GetBody() (value *CloudResource, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// CloudResourceUpdateServerResponse is the response for the 'update' method.
type CloudResourceUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *CloudResource
}

// Body sets the value of the 'body' parameter.
//
//
func (r *CloudResourceUpdateServerResponse) Body(value *CloudResource) *CloudResourceUpdateServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *CloudResourceUpdateServerResponse) Status(value int) *CloudResourceUpdateServerResponse {
	r.status = value
	return r
}

// dispatchCloudResource navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchCloudResource(w http.ResponseWriter, r *http.Request, server CloudResourceServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptCloudResourceDeleteRequest(w, r, server)
			return
		case "GET":
			adaptCloudResourceGetRequest(w, r, server)
			return
		case "PATCH":
			adaptCloudResourceUpdateRequest(w, r, server)
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

// adaptCloudResourceDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptCloudResourceDeleteRequest(w http.ResponseWriter, r *http.Request, server CloudResourceServer) {
	request := &CloudResourceDeleteServerRequest{}
	err := readCloudResourceDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &CloudResourceDeleteServerResponse{}
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
	err = writeCloudResourceDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptCloudResourceGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptCloudResourceGetRequest(w http.ResponseWriter, r *http.Request, server CloudResourceServer) {
	request := &CloudResourceGetServerRequest{}
	err := readCloudResourceGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &CloudResourceGetServerResponse{}
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
	err = writeCloudResourceGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptCloudResourceUpdateRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptCloudResourceUpdateRequest(w http.ResponseWriter, r *http.Request, server CloudResourceServer) {
	request := &CloudResourceUpdateServerRequest{}
	err := readCloudResourceUpdateRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &CloudResourceUpdateServerResponse{}
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
	err = writeCloudResourceUpdateResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
