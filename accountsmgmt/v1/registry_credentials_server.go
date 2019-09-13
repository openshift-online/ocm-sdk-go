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
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// RegistryCredentialsServer represents the interface the manages the 'registry_credentials' resource.
type RegistryCredentialsServer interface {

	// Add handles a request for the 'add' method.
	//
	// Creates a new registry credential.
	Add(ctx context.Context, request *RegistryCredentialsAddServerRequest, response *RegistryCredentialsAddServerResponse) error

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of accounts.
	List(ctx context.Context, request *RegistryCredentialsListServerRequest, response *RegistryCredentialsListServerResponse) error

	// RegistryCredential returns the target 'registry_credential' server for the given identifier.
	//
	// Reference to the service that manages an specific registry credential.
	RegistryCredential(id string) RegistryCredentialServer
}

// RegistryCredentialsAddServerRequest is the request for the 'add' method.
type RegistryCredentialsAddServerRequest struct {
	body *RegistryCredential
}

// Body returns the value of the 'body' parameter.
//
// Registry credential data.
func (r *RegistryCredentialsAddServerRequest) Body() *RegistryCredential {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Registry credential data.
func (r *RegistryCredentialsAddServerRequest) GetBody() (value *RegistryCredential, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'add' method.
func (r *RegistryCredentialsAddServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(registryCredentialData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}

// RegistryCredentialsAddServerResponse is the response for the 'add' method.
type RegistryCredentialsAddServerResponse struct {
	status int
	err    *errors.Error
	body   *RegistryCredential
}

// Body sets the value of the 'body' parameter.
//
// Registry credential data.
func (r *RegistryCredentialsAddServerResponse) Body(value *RegistryCredential) *RegistryCredentialsAddServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RegistryCredentialsAddServerResponse) SetStatusCode(status int) *RegistryCredentialsAddServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'add' method.
func (r *RegistryCredentialsAddServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RegistryCredentialsListServerRequest is the request for the 'list' method.
type RegistryCredentialsListServerRequest struct {
	page  *int
	size  *int
	total *int
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RegistryCredentialsListServerRequest) Page() int {
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
func (r *RegistryCredentialsListServerRequest) GetPage() (value int, ok bool) {
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
func (r *RegistryCredentialsListServerRequest) Size() int {
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
func (r *RegistryCredentialsListServerRequest) GetSize() (value int, ok bool) {
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
func (r *RegistryCredentialsListServerRequest) Total() int {
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
func (r *RegistryCredentialsListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// RegistryCredentialsListServerResponse is the response for the 'list' method.
type RegistryCredentialsListServerResponse struct {
	status int
	err    *errors.Error
	items  *RegistryCredentialList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of registry credentials.
func (r *RegistryCredentialsListServerResponse) Items(value *RegistryCredentialList) *RegistryCredentialsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RegistryCredentialsListServerResponse) Page(value int) *RegistryCredentialsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *RegistryCredentialsListServerResponse) Size(value int) *RegistryCredentialsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RegistryCredentialsListServerResponse) Total(value int) *RegistryCredentialsListServerResponse {
	r.total = &value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *RegistryCredentialsListServerResponse) SetStatusCode(status int) *RegistryCredentialsListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *RegistryCredentialsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(registryCredentialsListServerResponseData)
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

// registryCredentialsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type registryCredentialsListServerResponseData struct {
	Items registryCredentialListData "json:\"items,omitempty\""
	Page  *int                       "json:\"page,omitempty\""
	Size  *int                       "json:\"size,omitempty\""
	Total *int                       "json:\"total,omitempty\""
}

// RegistryCredentialsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RegistryCredentialsServerAdapter struct {
	server RegistryCredentialsServer
	router *mux.Router
}

func NewRegistryCredentialsServerAdapter(server RegistryCredentialsServer, router *mux.Router) *RegistryCredentialsServerAdapter {
	adapter := new(RegistryCredentialsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.registryCredentialHandler)
	adapter.router.Methods("POST").Path("").HandlerFunc(adapter.addHandler)
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.listHandler)
	return adapter
}
func (a *RegistryCredentialsServerAdapter) registryCredentialHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.RegistryCredential(id)
	targetAdapter := NewRegistryCredentialServerAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RegistryCredentialsServerAdapter) readRegistryCredentialsAddServerRequest(r *http.Request) (*RegistryCredentialsAddServerRequest, error) {
	var err error
	result := new(RegistryCredentialsAddServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *RegistryCredentialsServerAdapter) writeRegistryCredentialsAddServerResponse(w http.ResponseWriter, r *RegistryCredentialsAddServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RegistryCredentialsServerAdapter) addHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRegistryCredentialsAddServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RegistryCredentialsAddServerResponse)
	err = a.server.Add(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Add: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRegistryCredentialsAddServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RegistryCredentialsServerAdapter) readRegistryCredentialsListServerRequest(r *http.Request) (*RegistryCredentialsListServerRequest, error) {
	var err error
	result := new(RegistryCredentialsListServerRequest)
	query := r.URL.Query()
	result.page, err = helpers.ParseInteger(query, "page")
	if err != nil {
		return nil, err
	}
	result.size, err = helpers.ParseInteger(query, "size")
	if err != nil {
		return nil, err
	}
	result.total, err = helpers.ParseInteger(query, "total")
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *RegistryCredentialsServerAdapter) writeRegistryCredentialsListServerResponse(w http.ResponseWriter, r *RegistryCredentialsListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *RegistryCredentialsServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readRegistryCredentialsListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(RegistryCredentialsListServerResponse)
	err = a.server.List(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeRegistryCredentialsListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *RegistryCredentialsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
