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

// AttachmentsServer represents the interface the manages the 'attachments' resource.
type AttachmentsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of attachments
	List(ctx context.Context, request *AttachmentsListServerRequest, response *AttachmentsListServerResponse) error

	// Attachment returns the target 'attachment' server for the given identifier.
	//
	//
	Attachment(id string) AttachmentServer
}

// AttachmentsListServerRequest is the request for the 'list' method.
type AttachmentsListServerRequest struct {
	page *int
	size *int
}

// Page returns the value of the 'page' parameter.
//
//
func (r *AttachmentsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *AttachmentsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
//
func (r *AttachmentsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *AttachmentsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// AttachmentsListServerResponse is the response for the 'list' method.
type AttachmentsListServerResponse struct {
	status int
	err    *errors.Error
	items  *AttachmentList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
//
func (r *AttachmentsListServerResponse) Items(value *AttachmentList) *AttachmentsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
//
func (r *AttachmentsListServerResponse) Page(value int) *AttachmentsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
//
func (r *AttachmentsListServerResponse) Size(value int) *AttachmentsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
//
func (r *AttachmentsListServerResponse) Total(value int) *AttachmentsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *AttachmentsListServerResponse) Status(value int) *AttachmentsListServerResponse {
	r.status = value
	return r
}

// dispatchAttachments navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchAttachments(w http.ResponseWriter, r *http.Request, server AttachmentsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptAttachmentsListRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.Attachment(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchAttachment(w, r, target, segments[1:])
	}
}

// adaptAttachmentsListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptAttachmentsListRequest(w http.ResponseWriter, r *http.Request, server AttachmentsServer) {
	request := &AttachmentsListServerRequest{}
	err := readAttachmentsListRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &AttachmentsListServerResponse{}
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
	err = writeAttachmentsListResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
