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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
)

// SubscriptionsServer represents the interface the manages the 'subscriptions' resource.
type SubscriptionsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves a list of subscriptions.
	List(request *SubscriptionsListServerRequest, response *SubscriptionsListServerResponse) error

	// Subscription returns the target 'subscription' server for the given identifier.
	//
	// Reference to the service that manages a specific subscription.
	Subscription(id string) SubscriptionServer
}

// SubscriptionsListServerRequest is the request for the 'list' method.
type SubscriptionsListServerRequest struct {
	path  string
	query url.Values
	page  *int
	size  *int
	total *int
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *SubscriptionsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *SubscriptionsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *SubscriptionsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *SubscriptionsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *SubscriptionsListServerRequest) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *SubscriptionsListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// SubscriptionsListServerResponse is the response for the 'list' method.
type SubscriptionsListServerResponse struct {
	status int
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *SubscriptionList
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *SubscriptionsListServerResponse) Page(value int) *SubscriptionsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *SubscriptionsListServerResponse) Size(value int) *SubscriptionsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *SubscriptionsListServerResponse) Total(value int) *SubscriptionsListServerResponse {
	r.total = &value
	return r
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of subscriptions.
func (r *SubscriptionsListServerResponse) Items(value *SubscriptionList) *SubscriptionsListServerResponse {
	r.items = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *SubscriptionsListServerResponse) SetStatusCode(status int) *SubscriptionsListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *SubscriptionsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(subscriptionsListServerResponseData)
	data.Page = r.page
	data.Size = r.size
	data.Total = r.total
	data.Items, err = r.items.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// subscriptionsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type subscriptionsListServerResponseData struct {
	Page  *int                 "json:\"page,omitempty\""
	Size  *int                 "json:\"size,omitempty\""
	Total *int                 "json:\"total,omitempty\""
	Items subscriptionListData "json:\"items,omitempty\""
}

// SubscriptionsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type SubscriptionsServerAdapter struct {
	server SubscriptionsServer
	router *mux.Router
}

func NewSubscriptionsServerAdapter(server SubscriptionsServer, router *mux.Router) *SubscriptionsServerAdapter {
	adapter := new(SubscriptionsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}/").HandlerFunc(adapter.subscriptionHandler)
	adapter.router.HandleFunc("/", adapter.listHandler).Methods("GET")
	return adapter
}
func (a *SubscriptionsServerAdapter) subscriptionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Subscription(id)
	targetAdapter := NewSubscriptionServerAdapter(target, a.router.PathPrefix("/{id}/").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *SubscriptionsServerAdapter) readSubscriptionsListServerRequest(r *http.Request) (*SubscriptionsListServerRequest, error) {
	result := new(SubscriptionsListServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	return result, nil
}
func (a *SubscriptionsServerAdapter) writeSubscriptionsListServerResponse(w http.ResponseWriter, r *SubscriptionsListServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *SubscriptionsServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readSubscriptionsListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(SubscriptionsListServerResponse)
	err = a.server.List(req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeSubscriptionsListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *SubscriptionsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
