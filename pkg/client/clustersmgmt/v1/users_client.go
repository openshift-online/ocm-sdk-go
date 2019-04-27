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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	time "time"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// UsersClient is the client of the 'users' resource.
//
// Manages the collection of users of a group.
type UsersClient struct {
	transport http.RoundTripper
	path      string
}

// NewUsersClient creates a new client for the 'users'
// resource using the given transport to sned the requests and receive the
// responses.
func NewUsersClient(transport http.RoundTripper, path string) *UsersClient {
	client := new(UsersClient)
	client.transport = transport
	client.path = path
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of users.
func (c *UsersClient) List() *UsersListRequest {
	request := new(UsersListRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Add creates a request for the 'add' method.
//
// Adds a new user to the group.
func (c *UsersClient) Add() *UsersAddRequest {
	request := new(UsersAddRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// User returns the target 'user' resource for the given identifier.
//
// Reference to the service that manages an specific user.
func (c *UsersClient) User(id string) *UserClient {
	return NewUserClient(c.transport, path.Join(c.path, id))
}

// UsersListRequest is the request for the 'list' method.
type UsersListRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *UsersListRequest) Context(value context.Context) *UsersListRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *UsersListRequest) Timeout(value time.Duration) *UsersListRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *UsersListRequest) Deadline(value time.Time) *UsersListRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *UsersListRequest) Parameter(name string, value interface{}) *UsersListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *UsersListRequest) Header(name string, value interface{}) *UsersListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *UsersListRequest) Send() (result *UsersListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: header,
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(UsersListResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// UsersListResponse is the response for the 'list' method.
type UsersListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *UserList
}

// Status returns the response status code.
func (r *UsersListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *UsersListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *UsersListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *UsersListResponse) Page() int {
	if r.page != nil {
		return *r.page
	}
	return 0
}

// Size returns the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *UsersListResponse) Size() int {
	if r.size != nil {
		return *r.size
	}
	return 0
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *UsersListResponse) Total() int {
	if r.total != nil {
		return *r.total
	}
	return 0
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of users.
func (r *UsersListResponse) Items() *UserList {
	return r.items
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *UsersListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(usersListResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.page = data.Page
	r.size = data.Size
	r.total = data.Total
	r.items, err = data.Items.unwrap()
	if err != nil {
		return err
	}
	return err
}

// usersListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type usersListResponseData struct {
	Page  *int         "json:\"page,omitempty\""
	Size  *int         "json:\"size,omitempty\""
	Total *int         "json:\"total,omitempty\""
	Items userListData "json:\"items,omitempty\""
}

// UsersAddRequest is the request for the 'add' method.
type UsersAddRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
	body      *User
}

// Context sets the context that will be used to send the request.
func (r *UsersAddRequest) Context(value context.Context) *UsersAddRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *UsersAddRequest) Timeout(value time.Duration) *UsersAddRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *UsersAddRequest) Deadline(value time.Time) *UsersAddRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *UsersAddRequest) Parameter(name string, value interface{}) *UsersAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *UsersAddRequest) Header(name string, value interface{}) *UsersAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Description of the user.
func (r *UsersAddRequest) Body(value *User) *UsersAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *UsersAddRequest) Send() (result *UsersAddResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := new(bytes.Buffer)
	err = r.marshal(buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodPost,
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(UsersAddResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'add' method.
func (r *UsersAddRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// UsersAddResponse is the response for the 'add' method.
type UsersAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *User
}

// Status returns the response status code.
func (r *UsersAddResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *UsersAddResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *UsersAddResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Description of the user.
func (r *UsersAddResponse) Body() *User {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'add' method.
func (r *UsersAddResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(userData)
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
