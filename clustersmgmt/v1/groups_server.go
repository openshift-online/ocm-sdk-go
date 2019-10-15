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

// GroupsServer represents the interface the manages the 'groups' resource.
type GroupsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of groups.
	List(ctx context.Context, request *GroupsListServerRequest, response *GroupsListServerResponse) error

	// Group returns the target 'group' server for the given identifier.
	//
	// Reference to the service that manages an specific group.
	Group(id string) GroupServer
}

// GroupsListServerRequest is the request for the 'list' method.
type GroupsListServerRequest struct {
}

// GroupsListServerResponse is the response for the 'list' method.
type GroupsListServerResponse struct {
	status int
	err    *errors.Error
	items  *GroupList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of groups.
func (r *GroupsListServerResponse) Items(value *GroupList) *GroupsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *GroupsListServerResponse) Page(value int) *GroupsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *GroupsListServerResponse) Size(value int) *GroupsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *GroupsListServerResponse) Total(value int) *GroupsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *GroupsListServerResponse) Status(value int) *GroupsListServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *GroupsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(groupsListServerResponseData)
	data.Items, err = r.items.wrap()
	if err != nil {
		return err
	}
	data.Page = r.page
	data.Size = r.size
	data.Total = r.total
	err = encoder.Encode(data)
	return err
}

// groupsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type groupsListServerResponseData struct {
	Items groupListData "json:\"items,omitempty\""
	Page  *int          "json:\"page,omitempty\""
	Size  *int          "json:\"size,omitempty\""
	Total *int          "json:\"total,omitempty\""
}

// GroupsAdapter represents the structs that adapts Requests and Response to internal
// structs.
type GroupsAdapter struct {
	server GroupsServer
	router *mux.Router
}

func NewGroupsAdapter(server GroupsServer, router *mux.Router) *GroupsAdapter {
	adapter := new(GroupsAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.groupHandler)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerList)
	return adapter
}
func (a *GroupsAdapter) groupHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Group(id)
	targetAdapter := NewGroupAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *GroupsAdapter) readListRequest(r *http.Request) (*GroupsListServerRequest, error) {
	var err error
	result := new(GroupsListServerRequest)
	return result, err
}
func (a *GroupsAdapter) writeListResponse(w http.ResponseWriter, r *GroupsListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *GroupsAdapter) handlerList(w http.ResponseWriter, r *http.Request) {
	request, err := a.readListRequest(r)
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
	response := new(GroupsListServerResponse)
	response.status = http.StatusOK
	err = a.server.List(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method List: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeListResponse(w, response)
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
func (a *GroupsAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
