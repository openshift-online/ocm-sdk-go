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

// AccessReviewServer represents the interface the manages the 'access_review' resource.
type AccessReviewServer interface {

	// Post handles a request for the 'post' method.
	//
	// Reviews a user's access to a resource
	Post(ctx context.Context, request *AccessReviewPostServerRequest, response *AccessReviewPostServerResponse) error
}

// AccessReviewPostServerRequest is the request for the 'post' method.
type AccessReviewPostServerRequest struct {
	request *AccessReviewRequest
}

// Request returns the value of the 'request' parameter.
//
//
func (r *AccessReviewPostServerRequest) Request() *AccessReviewRequest {
	if r == nil {
		return nil
	}
	return r.request
}

// GetRequest returns the value of the 'request' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *AccessReviewPostServerRequest) GetRequest() (value *AccessReviewRequest, ok bool) {
	ok = r != nil && r.request != nil
	if ok {
		value = r.request
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'post' method.
func (r *AccessReviewPostServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(accessReviewRequestData)
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

// AccessReviewPostServerResponse is the response for the 'post' method.
type AccessReviewPostServerResponse struct {
	status   int
	err      *errors.Error
	response *AccessReviewResponse
}

// Response sets the value of the 'response' parameter.
//
//
func (r *AccessReviewPostServerResponse) Response(value *AccessReviewResponse) *AccessReviewPostServerResponse {
	r.response = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *AccessReviewPostServerResponse) SetStatusCode(status int) *AccessReviewPostServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'post' method.
func (r *AccessReviewPostServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.response.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// AccessReviewServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type AccessReviewServerAdapter struct {
	server AccessReviewServer
	router *mux.Router
}

func NewAccessReviewServerAdapter(server AccessReviewServer, router *mux.Router) *AccessReviewServerAdapter {
	adapter := new(AccessReviewServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("POST").Path("").HandlerFunc(adapter.postHandler)
	return adapter
}
func (a *AccessReviewServerAdapter) readAccessReviewPostServerRequest(r *http.Request) (*AccessReviewPostServerRequest, error) {
	var err error
	result := new(AccessReviewPostServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *AccessReviewServerAdapter) writeAccessReviewPostServerResponse(w http.ResponseWriter, r *AccessReviewPostServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *AccessReviewServerAdapter) postHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readAccessReviewPostServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(AccessReviewPostServerResponse)
	err = a.server.Post(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Post: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeAccessReviewPostServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *AccessReviewServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
