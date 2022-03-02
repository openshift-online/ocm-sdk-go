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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// EnvironmentServer represents the interface the manages the 'environment' resource.
type EnvironmentServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the environment.
	Get(ctx context.Context, request *EnvironmentGetServerRequest, response *EnvironmentGetServerResponse) error

	// Patch handles a request for the 'patch' method.
	//
	// Updates the environment.
	//
	// Attributes that can be updated are:
	//
	// - `last_upgrade_available_check`
	// - `last_limited_support_check`
	Patch(ctx context.Context, request *EnvironmentPatchServerRequest, response *EnvironmentPatchServerResponse) error
}

// EnvironmentGetServerRequest is the request for the 'get' method.
type EnvironmentGetServerRequest struct {
}

// EnvironmentGetServerResponse is the response for the 'get' method.
type EnvironmentGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Environment
}

// Body sets the value of the 'body' parameter.
//
//
func (r *EnvironmentGetServerResponse) Body(value *Environment) *EnvironmentGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *EnvironmentGetServerResponse) Status(value int) *EnvironmentGetServerResponse {
	r.status = value
	return r
}

// EnvironmentPatchServerRequest is the request for the 'patch' method.
type EnvironmentPatchServerRequest struct {
	body *Environment
}

// Body returns the value of the 'body' parameter.
//
//
func (r *EnvironmentPatchServerRequest) Body() *Environment {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EnvironmentPatchServerRequest) GetBody() (value *Environment, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// EnvironmentPatchServerResponse is the response for the 'patch' method.
type EnvironmentPatchServerResponse struct {
	status int
	err    *errors.Error
	body   *Environment
}

// Body sets the value of the 'body' parameter.
//
//
func (r *EnvironmentPatchServerResponse) Body(value *Environment) *EnvironmentPatchServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *EnvironmentPatchServerResponse) Status(value int) *EnvironmentPatchServerResponse {
	r.status = value
	return r
}

// dispatchEnvironment navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchEnvironment(w http.ResponseWriter, r *http.Request, server EnvironmentServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptEnvironmentGetRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	case "patch":
		if r.Method != "POST" {
			errors.SendMethodNotAllowed(w, r)
			return
		}
		adaptEnvironmentPatchRequest(w, r, server)
		return
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptEnvironmentGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptEnvironmentGetRequest(w http.ResponseWriter, r *http.Request, server EnvironmentServer) {
	request := &EnvironmentGetServerRequest{}
	err := readEnvironmentGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &EnvironmentGetServerResponse{}
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
	err = writeEnvironmentGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptEnvironmentPatchRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptEnvironmentPatchRequest(w http.ResponseWriter, r *http.Request, server EnvironmentServer) {
	request := &EnvironmentPatchServerRequest{}
	err := readEnvironmentPatchRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &EnvironmentPatchServerResponse{}
	response.status = 200
	err = server.Patch(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeEnvironmentPatchResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
