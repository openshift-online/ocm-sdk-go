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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// GroupServer represents the interface the manages the 'group' resource.
type GroupServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the group.
	Get(ctx context.Context, request *GroupGetServerRequest, response *GroupGetServerResponse) error

	// Users returns the target 'users' resource.
	//
	// Reference to the resource that manages the collection of users.
	Users() UsersServer
}

// GroupGetServerRequest is the request for the 'get' method.
type GroupGetServerRequest struct {
}

// GroupGetServerResponse is the response for the 'get' method.
type GroupGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Group
}

// Body sets the value of the 'body' parameter.
//
//
func (r *GroupGetServerResponse) Body(value *Group) *GroupGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *GroupGetServerResponse) Status(value int) *GroupGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *GroupGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// GroupAdapter represents the structs that adapts Requests and Response to internal
// structs.
type GroupAdapter struct {
	server GroupServer
	router *mux.Router
}

func NewGroupAdapter(server GroupServer, router *mux.Router) *GroupAdapter {
	adapter := new(GroupAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/users").HandlerFunc(adapter.usersHandler)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerGet)
	return adapter
}
func (a *GroupAdapter) usersHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Users()
	targetAdapter := NewUsersAdapter(target, a.router.PathPrefix("/users").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *GroupAdapter) readGetRequest(r *http.Request) (*GroupGetServerRequest, error) {
	var err error
	result := new(GroupGetServerRequest)
	return result, err
}
func (a *GroupAdapter) writeGetResponse(w http.ResponseWriter, r *GroupGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *GroupAdapter) handlerGet(w http.ResponseWriter, r *http.Request) {
	request, err := a.readGetRequest(r)
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
	response := new(GroupGetServerResponse)
	response.status = http.StatusOK
	err = a.server.Get(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method Get: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeGetResponse(w, response)
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
func (a *GroupAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
