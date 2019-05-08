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
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// RegistryCredentialsClient is the client of the 'registry_credentials' resource.
//
// Manages the collection of registry credentials.
type RegistryCredentialsClient struct {
	transport http.RoundTripper
	path      string
}

// NewRegistryCredentialsClient creates a new client for the 'registry_credentials'
// resource using the given transport to sned the requests and receive the
// responses.
func NewRegistryCredentialsClient(transport http.RoundTripper, path string) *RegistryCredentialsClient {
	client := new(RegistryCredentialsClient)
	client.transport = transport
	client.path = path
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of accounts.
func (c *RegistryCredentialsClient) List() *RegistryCredentialsListRequest {
	request := new(RegistryCredentialsListRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Add creates a request for the 'add' method.
//
// Creates a new registry credential.
func (c *RegistryCredentialsClient) Add() *RegistryCredentialsAddRequest {
	request := new(RegistryCredentialsAddRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// RegistryCredential returns the target 'registry_credential' resource for the given identifier.
//
// Reference to the service that manages an specific registry credential.
func (c *RegistryCredentialsClient) RegistryCredential(id string) *RegistryCredentialClient {
	return NewRegistryCredentialClient(c.transport, path.Join(c.path, id))
}

// RegistryCredentialsListRequest is the request for the 'list' method.
type RegistryCredentialsListRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	page      *int
	size      *int
	total     *int
}

// Parameter adds a query parameter.
func (r *RegistryCredentialsListRequest) Parameter(name string, value interface{}) *RegistryCredentialsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *RegistryCredentialsListRequest) Header(name string, value interface{}) *RegistryCredentialsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RegistryCredentialsListRequest) Page(value int) *RegistryCredentialsListRequest {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *RegistryCredentialsListRequest) Size(value int) *RegistryCredentialsListRequest {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RegistryCredentialsListRequest) Total(value int) *RegistryCredentialsListRequest {
	r.total = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method. If you don't provide a
// context then a new background context will be created.
func (r *RegistryCredentialsListRequest) Send() (result *RegistryCredentialsListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *RegistryCredentialsListRequest) SendContext(ctx context.Context) (result *RegistryCredentialsListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
	if r.total != nil {
		helpers.AddValue(&query, "total", *r.total)
	}
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
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(RegistryCredentialsListResponse)
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

// RegistryCredentialsListResponse is the response for the 'list' method.
type RegistryCredentialsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *RegistryCredentialList
}

// Status returns the response status code.
func (r *RegistryCredentialsListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *RegistryCredentialsListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *RegistryCredentialsListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RegistryCredentialsListResponse) Page() int {
	if r.page != nil {
		return *r.page
	}
	return 0
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *RegistryCredentialsListResponse) Size() int {
	if r.size != nil {
		return *r.size
	}
	return 0
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RegistryCredentialsListResponse) Total() int {
	if r.total != nil {
		return *r.total
	}
	return 0
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of registry credentials.
func (r *RegistryCredentialsListResponse) Items() *RegistryCredentialList {
	return r.items
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *RegistryCredentialsListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(registryCredentialsListResponseData)
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

// registryCredentialsListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type registryCredentialsListResponseData struct {
	Page  *int                       "json:\"page,omitempty\""
	Size  *int                       "json:\"size,omitempty\""
	Total *int                       "json:\"total,omitempty\""
	Items registryCredentialListData "json:\"items,omitempty\""
}

// RegistryCredentialsAddRequest is the request for the 'add' method.
type RegistryCredentialsAddRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *RegistryCredential
}

// Parameter adds a query parameter.
func (r *RegistryCredentialsAddRequest) Parameter(name string, value interface{}) *RegistryCredentialsAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *RegistryCredentialsAddRequest) Header(name string, value interface{}) *RegistryCredentialsAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Registry credential data.
func (r *RegistryCredentialsAddRequest) Body(value *RegistryCredential) *RegistryCredentialsAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method. If you don't provide a
// context then a new background context will be created.
func (r *RegistryCredentialsAddRequest) Send() (result *RegistryCredentialsAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *RegistryCredentialsAddRequest) SendContext(ctx context.Context) (result *RegistryCredentialsAddResponse, err error) {
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
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(RegistryCredentialsAddResponse)
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
func (r *RegistryCredentialsAddRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RegistryCredentialsAddResponse is the response for the 'add' method.
type RegistryCredentialsAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *RegistryCredential
}

// Status returns the response status code.
func (r *RegistryCredentialsAddResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *RegistryCredentialsAddResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *RegistryCredentialsAddResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Registry credential data.
func (r *RegistryCredentialsAddResponse) Body() *RegistryCredential {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'add' method.
func (r *RegistryCredentialsAddResponse) unmarshal(reader io.Reader) error {
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
