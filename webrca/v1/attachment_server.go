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

// AttachmentServer represents the interface the manages the 'attachment' resource.
type AttachmentServer interface {

	// Delete handles a request for the 'delete' method.
	//
	//
	Delete(ctx context.Context, request *AttachmentDeleteServerRequest, response *AttachmentDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	//
	Get(ctx context.Context, request *AttachmentGetServerRequest, response *AttachmentGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	//
	Update(ctx context.Context, request *AttachmentUpdateServerRequest, response *AttachmentUpdateServerResponse) error
}

// AttachmentDeleteServerRequest is the request for the 'delete' method.
type AttachmentDeleteServerRequest struct {
}

// AttachmentDeleteServerResponse is the response for the 'delete' method.
type AttachmentDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *AttachmentDeleteServerResponse) Status(value int) *AttachmentDeleteServerResponse {
	r.status = value
	return r
}

// AttachmentGetServerRequest is the request for the 'get' method.
type AttachmentGetServerRequest struct {
}

// AttachmentGetServerResponse is the response for the 'get' method.
type AttachmentGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Attachment
}

// Body sets the value of the 'body' parameter.
//
//
func (r *AttachmentGetServerResponse) Body(value *Attachment) *AttachmentGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *AttachmentGetServerResponse) Status(value int) *AttachmentGetServerResponse {
	r.status = value
	return r
}

// AttachmentUpdateServerRequest is the request for the 'update' method.
type AttachmentUpdateServerRequest struct {
	body *Attachment
}

// Body returns the value of the 'body' parameter.
//
//
func (r *AttachmentUpdateServerRequest) Body() *Attachment {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *AttachmentUpdateServerRequest) GetBody() (value *Attachment, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// AttachmentUpdateServerResponse is the response for the 'update' method.
type AttachmentUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *Attachment
}

// Body sets the value of the 'body' parameter.
//
//
func (r *AttachmentUpdateServerResponse) Body(value *Attachment) *AttachmentUpdateServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *AttachmentUpdateServerResponse) Status(value int) *AttachmentUpdateServerResponse {
	r.status = value
	return r
}

// dispatchAttachment navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchAttachment(w http.ResponseWriter, r *http.Request, server AttachmentServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptAttachmentDeleteRequest(w, r, server)
			return
		case "GET":
			adaptAttachmentGetRequest(w, r, server)
			return
		case "PATCH":
			adaptAttachmentUpdateRequest(w, r, server)
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

// adaptAttachmentDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptAttachmentDeleteRequest(w http.ResponseWriter, r *http.Request, server AttachmentServer) {
	request := &AttachmentDeleteServerRequest{}
	err := readAttachmentDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &AttachmentDeleteServerResponse{}
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
	err = writeAttachmentDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptAttachmentGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptAttachmentGetRequest(w http.ResponseWriter, r *http.Request, server AttachmentServer) {
	request := &AttachmentGetServerRequest{}
	err := readAttachmentGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &AttachmentGetServerResponse{}
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
	err = writeAttachmentGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptAttachmentUpdateRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptAttachmentUpdateRequest(w http.ResponseWriter, r *http.Request, server AttachmentServer) {
	request := &AttachmentUpdateServerRequest{}
	err := readAttachmentUpdateRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &AttachmentUpdateServerResponse{}
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
	err = writeAttachmentUpdateResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
