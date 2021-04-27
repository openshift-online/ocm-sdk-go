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

package v1 // github.com/openshift-online/ocm-sdk-go/jobqueue/v1

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// QueueServer represents the interface the manages the 'queue' resource.
type QueueServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of a job queue by ID.
	Get(ctx context.Context, request *QueueGetServerRequest, response *QueueGetServerResponse) error

	// Pop handles a request for the 'pop' method.
	//
	// POP new job from a job queue
	Pop(ctx context.Context, request *QueuePopServerRequest, response *QueuePopServerResponse) error

	// Push handles a request for the 'push' method.
	//
	// PUSH a new job into job queue
	Push(ctx context.Context, request *QueuePushServerRequest, response *QueuePushServerResponse) error

	// Jobs returns the target 'jobs' resource.
	//
	// jobs' operations (success, failure)
	Jobs() JobsServer
}

// QueueGetServerRequest is the request for the 'get' method.
type QueueGetServerRequest struct {
}

// QueueGetServerResponse is the response for the 'get' method.
type QueueGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Queue
}

// Body sets the value of the 'body' parameter.
//
//
func (r *QueueGetServerResponse) Body(value *Queue) *QueueGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *QueueGetServerResponse) Status(value int) *QueueGetServerResponse {
	r.status = value
	return r
}

// QueuePopServerRequest is the request for the 'pop' method.
type QueuePopServerRequest struct {
}

// QueuePopServerResponse is the response for the 'pop' method.
type QueuePopServerResponse struct {
	status int
	err    *errors.Error
	body   *Job
}

// Body sets the value of the 'body' parameter.
//
//
func (r *QueuePopServerResponse) Body(value *Job) *QueuePopServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *QueuePopServerResponse) Status(value int) *QueuePopServerResponse {
	r.status = value
	return r
}

// QueuePushServerRequest is the request for the 'push' method.
type QueuePushServerRequest struct {
	body *Job
}

// Body returns the value of the 'body' parameter.
//
//
func (r *QueuePushServerRequest) Body() *Job {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *QueuePushServerRequest) GetBody() (value *Job, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// QueuePushServerResponse is the response for the 'push' method.
type QueuePushServerResponse struct {
	status int
	err    *errors.Error
	body   *Job
}

// Body sets the value of the 'body' parameter.
//
//
func (r *QueuePushServerResponse) Body(value *Job) *QueuePushServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *QueuePushServerResponse) Status(value int) *QueuePushServerResponse {
	r.status = value
	return r
}

// dispatchQueue navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchQueue(w http.ResponseWriter, r *http.Request, server QueueServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptQueueGetRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	case "pop":
		if r.Method != "POST" {
			errors.SendMethodNotAllowed(w, r)
			return
		}
		adaptQueuePopRequest(w, r, server)
		return
	case "push":
		if r.Method != "POST" {
			errors.SendMethodNotAllowed(w, r)
			return
		}
		adaptQueuePushRequest(w, r, server)
		return
	case "jobs":
		target := server.Jobs()
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchJobs(w, r, target, segments[1:])
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptQueueGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptQueueGetRequest(w http.ResponseWriter, r *http.Request, server QueueServer) {
	request := &QueueGetServerRequest{}
	err := readQueueGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &QueueGetServerResponse{}
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
	err = writeQueueGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptQueuePopRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptQueuePopRequest(w http.ResponseWriter, r *http.Request, server QueueServer) {
	request := &QueuePopServerRequest{}
	err := readQueuePopRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &QueuePopServerResponse{}
	response.status = 200
	err = server.Pop(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeQueuePopResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptQueuePushRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptQueuePushRequest(w http.ResponseWriter, r *http.Request, server QueueServer) {
	request := &QueuePushServerRequest{}
	err := readQueuePushRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &QueuePushServerResponse{}
	response.status = 200
	err = server.Push(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeQueuePushResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
