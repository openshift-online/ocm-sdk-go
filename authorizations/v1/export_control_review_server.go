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

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// ExportControlReviewServer represents the interface the manages the 'export_control_review' resource.
type ExportControlReviewServer interface {

	// Post handles a request for the 'post' method.
	//
	// Screens a user by account user name.
	Post(ctx context.Context, request *ExportControlReviewPostServerRequest, response *ExportControlReviewPostServerResponse) error
}

// ExportControlReviewPostServerRequest is the request for the 'post' method.
type ExportControlReviewPostServerRequest struct {
	request *ExportControlReviewRequest
}

// Request returns the value of the 'request' parameter.
//
//
func (r *ExportControlReviewPostServerRequest) Request() *ExportControlReviewRequest {
	if r == nil {
		return nil
	}
	return r.request
}

// GetRequest returns the value of the 'request' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *ExportControlReviewPostServerRequest) GetRequest() (value *ExportControlReviewRequest, ok bool) {
	ok = r != nil && r.request != nil
	if ok {
		value = r.request
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'post' method.
func (r *ExportControlReviewPostServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(exportControlReviewRequestData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.request, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}

// ExportControlReviewPostServerResponse is the response for the 'post' method.
type ExportControlReviewPostServerResponse struct {
	status   int
	err      *errors.Error
	response *ExportControlReviewResponse
}

// Response sets the value of the 'response' parameter.
//
//
func (r *ExportControlReviewPostServerResponse) Response(value *ExportControlReviewResponse) *ExportControlReviewPostServerResponse {
	r.response = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ExportControlReviewPostServerResponse) SetStatusCode(status int) *ExportControlReviewPostServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'post' method.
func (r *ExportControlReviewPostServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.response.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ExportControlReviewServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type ExportControlReviewServerAdapter struct {
	server ExportControlReviewServer
	router *mux.Router
}

func NewExportControlReviewServerAdapter(server ExportControlReviewServer, router *mux.Router) *ExportControlReviewServerAdapter {
	adapter := new(ExportControlReviewServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("POST").Path("").HandlerFunc(adapter.postHandler)
	return adapter
}
func (a *ExportControlReviewServerAdapter) readExportControlReviewPostServerRequest(r *http.Request) (*ExportControlReviewPostServerRequest, error) {
	var err error
	result := new(ExportControlReviewPostServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *ExportControlReviewServerAdapter) writeExportControlReviewPostServerResponse(w http.ResponseWriter, r *ExportControlReviewPostServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ExportControlReviewServerAdapter) postHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readExportControlReviewPostServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ExportControlReviewPostServerResponse)
	err = a.server.Post(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Post: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeExportControlReviewPostServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ExportControlReviewServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
