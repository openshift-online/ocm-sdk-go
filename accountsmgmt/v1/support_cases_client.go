/*
Copyright (c) 2020 Red Hat, Inc.

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
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// SupportCasesClient is the client of the 'support_cases' resource.
//
// Manages the support cases endpoint
type SupportCasesClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewSupportCasesClient creates a new client for the 'support_cases'
// resource using the given transport to send the requests and receive the
// responses.
func NewSupportCasesClient(transport http.RoundTripper, path string, metric string) *SupportCasesClient {
	return &SupportCasesClient{
		transport: transport,
		path:      path,
		metric:    metric,
	}
}

// Add creates a request for the 'add' method.
//
// Create a support case related to Hydra
func (c *SupportCasesClient) Add() *SupportCasesAddRequest {
	return &SupportCasesAddRequest{
		transport: c.transport,
		path:      c.path,
		metric:    c.metric,
	}
}

// Delete creates a request for the 'delete' method.
//
// Close a support case in Hydra.
func (c *SupportCasesClient) Delete() *SupportCasesDeleteRequest {
	return &SupportCasesDeleteRequest{
		transport: c.transport,
		path:      c.path,
		metric:    c.metric,
	}
}

// SupportCasesAddRequest is the request for the 'add' method.
type SupportCasesAddRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	body      *SupportCase
}

// Parameter adds a query parameter.
func (r *SupportCasesAddRequest) Parameter(name string, value interface{}) *SupportCasesAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *SupportCasesAddRequest) Header(name string, value interface{}) *SupportCasesAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *SupportCasesAddRequest) Body(value *SupportCase) *SupportCasesAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *SupportCasesAddRequest) Send() (result *SupportCasesAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *SupportCasesAddRequest) SendContext(ctx context.Context) (result *SupportCasesAddResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	buffer := &bytes.Buffer{}
	err = writeSupportCasesAddRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
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
	result = &SupportCasesAddResponse{}
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
	err = readSupportCasesAddResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'add' method.
func (r *SupportCasesAddRequest) marshal(writer io.Writer) error {
	stream := helpers.NewStream(writer)
	r.stream(stream)
	return stream.Error
}
func (r *SupportCasesAddRequest) stream(stream *jsoniter.Stream) {
}

// SupportCasesAddResponse is the response for the 'add' method.
type SupportCasesAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *SupportCase
}

// Status returns the response status code.
func (r *SupportCasesAddResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *SupportCasesAddResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *SupportCasesAddResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *SupportCasesAddResponse) Body() *SupportCase {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *SupportCasesAddResponse) GetBody() (value *SupportCase, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// SupportCasesDeleteRequest is the request for the 'delete' method.
type SupportCasesDeleteRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *SupportCasesDeleteRequest) Parameter(name string, value interface{}) *SupportCasesDeleteRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *SupportCasesDeleteRequest) Header(name string, value interface{}) *SupportCasesDeleteRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *SupportCasesDeleteRequest) Send() (result *SupportCasesDeleteResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *SupportCasesDeleteRequest) SendContext(ctx context.Context) (result *SupportCasesDeleteResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "DELETE",
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
	result = &SupportCasesDeleteResponse{}
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
	return
}

// SupportCasesDeleteResponse is the response for the 'delete' method.
type SupportCasesDeleteResponse struct {
	status int
	header http.Header
	err    *errors.Error
}

// Status returns the response status code.
func (r *SupportCasesDeleteResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *SupportCasesDeleteResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *SupportCasesDeleteResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}
