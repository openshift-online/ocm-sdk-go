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

// SupportCasesServer represents the interface the manages the 'support_cases' resource.
type SupportCasesServer interface {

	// Add handles a request for the 'add' method.
	//
	// Create a support case related to Hydra
	Add(ctx context.Context, request *SupportCasesAddServerRequest, response *SupportCasesAddServerResponse) error

	// Delete handles a request for the 'delete' method.
	//
	// Close a support case in Hydra.
	Delete(ctx context.Context, request *SupportCasesDeleteServerRequest, response *SupportCasesDeleteServerResponse) error
}

// SupportCasesAddServerRequest is the request for the 'add' method.
type SupportCasesAddServerRequest struct {
	body *SupportCase
}

// Body returns the value of the 'body' parameter.
//
//
func (r *SupportCasesAddServerRequest) Body() *SupportCase {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *SupportCasesAddServerRequest) GetBody() (value *SupportCase, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// SupportCasesAddServerResponse is the response for the 'add' method.
type SupportCasesAddServerResponse struct {
	status int
	err    *errors.Error
	body   *SupportCase
}

// Body sets the value of the 'body' parameter.
//
//
func (r *SupportCasesAddServerResponse) Body(value *SupportCase) *SupportCasesAddServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *SupportCasesAddServerResponse) Status(value int) *SupportCasesAddServerResponse {
	r.status = value
	return r
}

// SupportCasesDeleteServerRequest is the request for the 'delete' method.
type SupportCasesDeleteServerRequest struct {
}

// SupportCasesDeleteServerResponse is the response for the 'delete' method.
type SupportCasesDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *SupportCasesDeleteServerResponse) Status(value int) *SupportCasesDeleteServerResponse {
	r.status = value
	return r
}

// dispatchSupportCases navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchSupportCases(w http.ResponseWriter, r *http.Request, server SupportCasesServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "POST":
			adaptSupportCasesAddRequest(w, r, server)
			return
		case "DELETE":
			adaptSupportCasesDeleteRequest(w, r, server)
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

// adaptSupportCasesAddRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptSupportCasesAddRequest(w http.ResponseWriter, r *http.Request, server SupportCasesServer) {
	request := &SupportCasesAddServerRequest{}
	err := readSupportCasesAddRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &SupportCasesAddServerResponse{}
	response.status = 201
	err = server.Add(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeSupportCasesAddResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptSupportCasesDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptSupportCasesDeleteRequest(w http.ResponseWriter, r *http.Request, server SupportCasesServer) {
	request := &SupportCasesDeleteServerRequest{}
	err := readSupportCasesDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &SupportCasesDeleteServerResponse{}
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
	err = writeSupportCasesDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
