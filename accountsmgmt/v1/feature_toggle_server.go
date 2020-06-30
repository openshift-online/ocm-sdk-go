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

// FeatureToggleServer represents the interface the manages the 'feature_toggle' resource.
type FeatureToggleServer interface {

	// Post handles a request for the 'post' method.
	//
	// Retrieves the details of the feature toggle by providing query context
	Post(ctx context.Context, request *FeatureTogglePostServerRequest, response *FeatureTogglePostServerResponse) error
}

// FeatureTogglePostServerRequest is the request for the 'post' method.
type FeatureTogglePostServerRequest struct {
	request *FeatureToggleQueryRequest
}

// Request returns the value of the 'request' parameter.
//
//
func (r *FeatureTogglePostServerRequest) Request() *FeatureToggleQueryRequest {
	if r == nil {
		return nil
	}
	return r.request
}

// GetRequest returns the value of the 'request' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *FeatureTogglePostServerRequest) GetRequest() (value *FeatureToggleQueryRequest, ok bool) {
	ok = r != nil && r.request != nil
	if ok {
		value = r.request
	}
	return
}

// FeatureTogglePostServerResponse is the response for the 'post' method.
type FeatureTogglePostServerResponse struct {
	status   int
	err      *errors.Error
	response *FeatureToggle
}

// Response sets the value of the 'response' parameter.
//
//
func (r *FeatureTogglePostServerResponse) Response(value *FeatureToggle) *FeatureTogglePostServerResponse {
	r.response = value
	return r
}

// Status sets the status code.
func (r *FeatureTogglePostServerResponse) Status(value int) *FeatureTogglePostServerResponse {
	r.status = value
	return r
}

// dispatchFeatureToggle navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchFeatureToggle(w http.ResponseWriter, r *http.Request, server FeatureToggleServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "POST":
			adaptFeatureTogglePostRequest(w, r, server)
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

// adaptFeatureTogglePostRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptFeatureTogglePostRequest(w http.ResponseWriter, r *http.Request, server FeatureToggleServer) {
	request := &FeatureTogglePostServerRequest{}
	err := readFeatureTogglePostRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &FeatureTogglePostServerResponse{}
	response.status = 201
	err = server.Post(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeFeatureTogglePostResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
