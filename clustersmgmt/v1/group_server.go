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

package v1 // github.com/openshift-online/uhc-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/uhc-sdk-go/errors"
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
	path  string
	query url.Values
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

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *GroupGetServerResponse) SetStatusCode(status int) *GroupGetServerResponse {
	r.status = status
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

// GroupServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type GroupServerAdapter struct {
	server GroupServer
	router *mux.Router
}

func NewGroupServerAdapter(server GroupServer, router *mux.Router) *GroupServerAdapter {
	adapter := new(GroupServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/users/").HandlerFunc(adapter.usersHandler)
	adapter.router.HandleFunc("/", adapter.getHandler).Methods("GET")
	return adapter
}
func (a *GroupServerAdapter) usersHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Users()
	targetAdapter := NewUsersServerAdapter(target, a.router.PathPrefix("/users/").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *GroupServerAdapter) readGroupGetServerRequest(r *http.Request) (*GroupGetServerRequest, error) {
	result := new(GroupGetServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *GroupServerAdapter) writeGroupGetServerResponse(w http.ResponseWriter, r *GroupGetServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *GroupServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readGroupGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(GroupGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeGroupGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *GroupServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
