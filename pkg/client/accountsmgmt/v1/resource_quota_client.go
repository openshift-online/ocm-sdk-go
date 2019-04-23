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
	time "time"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// ResourceQuotaClient is the client of the 'resource_quota' resource.
//
// Manages a specific resource quota.
type ResourceQuotaClient struct {
	transport http.RoundTripper
	path      string
}

// NewResourceQuotaClient creates a new client for the 'resource_quota'
// resource using the given transport to sned the requests and receive the
// responses.
func NewResourceQuotaClient(transport http.RoundTripper, path string) *ResourceQuotaClient {
	client := new(ResourceQuotaClient)
	client.transport = transport
	client.path = path
	return client
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the resource quota.
func (c *ResourceQuotaClient) Get() *ResourceQuotaGetRequest {
	request := new(ResourceQuotaGetRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Update creates a request for the 'update' method.
//
// Updates the resource quota.
func (c *ResourceQuotaClient) Update() *ResourceQuotaUpdateRequest {
	request := new(ResourceQuotaUpdateRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// ResourceQuotaGetRequest is the request for the 'get' method.
type ResourceQuotaGetRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *ResourceQuotaGetRequest) Context(value context.Context) *ResourceQuotaGetRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *ResourceQuotaGetRequest) Timeout(value time.Duration) *ResourceQuotaGetRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *ResourceQuotaGetRequest) Deadline(value time.Time) *ResourceQuotaGetRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *ResourceQuotaGetRequest) Parameter(name string, value interface{}) *ResourceQuotaGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ResourceQuotaGetRequest) Header(name string, value interface{}) *ResourceQuotaGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *ResourceQuotaGetRequest) Send() (result *ResourceQuotaGetResponse, err error) {
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
	result = new(ResourceQuotaGetResponse)
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

// ResourceQuotaGetResponse is the response for the 'get' method.
type ResourceQuotaGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *ResourceQuota
}

// Status returns the response status code.
func (r *ResourceQuotaGetResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *ResourceQuotaGetResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *ResourceQuotaGetResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *ResourceQuotaGetResponse) Body() *ResourceQuota {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'get' method.
func (r *ResourceQuotaGetResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(resourceQuotaData)
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

// ResourceQuotaUpdateRequest is the request for the 'update' method.
type ResourceQuotaUpdateRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
	body      *ResourceQuota
}

// Context sets the context that will be used to send the request.
func (r *ResourceQuotaUpdateRequest) Context(value context.Context) *ResourceQuotaUpdateRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *ResourceQuotaUpdateRequest) Timeout(value time.Duration) *ResourceQuotaUpdateRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *ResourceQuotaUpdateRequest) Deadline(value time.Time) *ResourceQuotaUpdateRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *ResourceQuotaUpdateRequest) Parameter(name string, value interface{}) *ResourceQuotaUpdateRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ResourceQuotaUpdateRequest) Header(name string, value interface{}) *ResourceQuotaUpdateRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *ResourceQuotaUpdateRequest) Body(value *ResourceQuota) *ResourceQuotaUpdateRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *ResourceQuotaUpdateRequest) Send() (result *ResourceQuotaUpdateResponse, err error) {
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
		Method: http.MethodPatch,
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(ResourceQuotaUpdateResponse)
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
// 'update' method.
func (r *ResourceQuotaUpdateRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ResourceQuotaUpdateResponse is the response for the 'update' method.
type ResourceQuotaUpdateResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *ResourceQuota
}

// Status returns the response status code.
func (r *ResourceQuotaUpdateResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *ResourceQuotaUpdateResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *ResourceQuotaUpdateResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *ResourceQuotaUpdateResponse) Body() *ResourceQuota {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'update' method.
func (r *ResourceQuotaUpdateResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(resourceQuotaData)
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
