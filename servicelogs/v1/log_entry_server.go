/*
Copyright (c) 2019 Red Hat, Inc.

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

package v1 // github.com/openshift-online/ocm-sdk-go/servicelogs/v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// LogEntryServer represents the interface the manages the 'log_entry' resource.
type LogEntryServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the log entry.
	Delete(ctx context.Context, request *LogEntryDeleteServerRequest, response *LogEntryDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the log entry.
	Get(ctx context.Context, request *LogEntryGetServerRequest, response *LogEntryGetServerResponse) error
}

// LogEntryDeleteServerRequest is the request for the 'delete' method.
type LogEntryDeleteServerRequest struct {
}

// LogEntryDeleteServerResponse is the response for the 'delete' method.
type LogEntryDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *LogEntryDeleteServerResponse) Status(value int) *LogEntryDeleteServerResponse {
	r.status = value
	return r
}

// LogEntryGetServerRequest is the request for the 'get' method.
type LogEntryGetServerRequest struct {
}

// LogEntryGetServerResponse is the response for the 'get' method.
type LogEntryGetServerResponse struct {
	status int
	err    *errors.Error
	body   *LogEntry
}

// Body sets the value of the 'body' parameter.
//
//
func (r *LogEntryGetServerResponse) Body(value *LogEntry) *LogEntryGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *LogEntryGetServerResponse) Status(value int) *LogEntryGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *LogEntryGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// dispatchLogEntry navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchLogEntry(w http.ResponseWriter, r *http.Request, server LogEntryServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptLogEntryDeleteRequest(w, r, server)
		case "GET":
			adaptLogEntryGetRequest(w, r, server)
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	} else {
		switch segments[0] {
		default:
			errors.SendNotFound(w, r)
			return
		}
	}
}

// readLogEntryDeleteRequest reads the given HTTP requests and translates it
// into an object of type LogEntryDeleteServerRequest.
func readLogEntryDeleteRequest(r *http.Request) (*LogEntryDeleteServerRequest, error) {
	var err error
	result := new(LogEntryDeleteServerRequest)
	return result, err
}

// writeLogEntryDeleteResponse translates the given request object into an
// HTTP response.
func writeLogEntryDeleteResponse(w http.ResponseWriter, r *LogEntryDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}

// adaptLogEntryDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptLogEntryDeleteRequest(w http.ResponseWriter, r *http.Request, server LogEntryServer) {
	request, err := readLogEntryDeleteRequest(r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := new(LogEntryDeleteServerResponse)
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
	err = writeLogEntryDeleteResponse(w, response)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// readLogEntryGetRequest reads the given HTTP requests and translates it
// into an object of type LogEntryGetServerRequest.
func readLogEntryGetRequest(r *http.Request) (*LogEntryGetServerRequest, error) {
	var err error
	result := new(LogEntryGetServerRequest)
	return result, err
}

// writeLogEntryGetResponse translates the given request object into an
// HTTP response.
func writeLogEntryGetResponse(w http.ResponseWriter, r *LogEntryGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}

// adaptLogEntryGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptLogEntryGetRequest(w http.ResponseWriter, r *http.Request, server LogEntryServer) {
	request, err := readLogEntryGetRequest(r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := new(LogEntryGetServerResponse)
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
	err = writeLogEntryGetResponse(w, response)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
