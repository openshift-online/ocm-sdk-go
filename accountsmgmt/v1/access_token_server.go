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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// AccessTokenServer represents the interface the manages the 'access_token' resource.
type AccessTokenServer interface {

	// Post handles a request for the 'post' method.
	//
	// Returns access token generated from registries in docker format.
	Post(ctx context.Context, request *AccessTokenPostServerRequest, response *AccessTokenPostServerResponse) error
}

// AccessTokenPostServerRequest is the request for the 'post' method.
type AccessTokenPostServerRequest struct {
}

// AccessTokenPostServerResponse is the response for the 'post' method.
type AccessTokenPostServerResponse struct {
	status int
	err    *errors.Error
	body   *AccessToken
}

// Body sets the value of the 'body' parameter.
//
//
func (r *AccessTokenPostServerResponse) Body(value *AccessToken) *AccessTokenPostServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *AccessTokenPostServerResponse) Status(value int) *AccessTokenPostServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'post' method.
func (r *AccessTokenPostServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// AccessTokenAdapter represents the structs that adapts Requests and Response to internal
// structs.
type AccessTokenAdapter struct {
	server AccessTokenServer
	router *mux.Router
}

func NewAccessTokenAdapter(server AccessTokenServer, router *mux.Router) *AccessTokenAdapter {
	adapter := new(AccessTokenAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods(http.MethodPost).Path("").HandlerFunc(adapter.handlerPost)
	return adapter
}
func (a *AccessTokenAdapter) readPostRequest(r *http.Request) (*AccessTokenPostServerRequest, error) {
	var err error
	result := new(AccessTokenPostServerRequest)
	return result, err
}
func (a *AccessTokenAdapter) writePostResponse(w http.ResponseWriter, r *AccessTokenPostServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *AccessTokenAdapter) handlerPost(w http.ResponseWriter, r *http.Request) {
	request, err := a.readPostRequest(r)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to read request from client: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
		return
	}
	response := new(AccessTokenPostServerResponse)
	response.status = http.StatusOK
	err = a.server.Post(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method Post: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writePostResponse(w, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to write response for client: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
}
func (a *AccessTokenAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
