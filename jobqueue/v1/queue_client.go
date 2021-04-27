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

package v1 // github.com/openshift-online/ocm-sdk-go/jobqueue/v1

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// QueueClient is the client of the 'queue' resource.
//
// Manages a specific job queue.
type QueueClient struct {
	transport http.RoundTripper
	path      string
}

// NewQueueClient creates a new client for the 'queue'
// resource using the given transport to send the requests and receive the
// responses.
func NewQueueClient(transport http.RoundTripper, path string) *QueueClient {
	return &QueueClient{
		transport: transport,
		path:      path,
	}
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of a job queue by ID.
func (c *QueueClient) Get() *QueueGetRequest {
	return &QueueGetRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Pop creates a request for the 'pop' method.
//
// POP new job from a job queue
func (c *QueueClient) Pop() *QueuePopRequest {
	return &QueuePopRequest{
		transport: c.transport,
		path:      path.Join(c.path, "pop"),
	}
}

// Push creates a request for the 'push' method.
//
// PUSH a new job into job queue
func (c *QueueClient) Push() *QueuePushRequest {
	return &QueuePushRequest{
		transport: c.transport,
		path:      path.Join(c.path, "push"),
	}
}

// Jobs returns the target 'jobs' resource.
//
// jobs' operations (success, failure)
func (c *QueueClient) Jobs() *JobsClient {
	return NewJobsClient(
		c.transport,
		path.Join(c.path, "jobs"),
	)
}

// QueuePollRequest is the request for the Poll method.
type QueuePollRequest struct {
	request    *QueueGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *QueuePollRequest) Parameter(name string, value interface{}) *QueuePollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *QueuePollRequest) Header(name string, value interface{}) *QueuePollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *QueuePollRequest) Interval(value time.Duration) *QueuePollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *QueuePollRequest) Status(value int) *QueuePollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *QueuePollRequest) Predicate(value func(*QueueGetResponse) bool) *QueuePollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*QueueGetResponse))
	})
	return r
}

// StartContext starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *QueuePollRequest) StartContext(ctx context.Context) (response *QueuePollResponse, err error) {
	result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &QueuePollResponse{
			response: result.(*QueueGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *QueuePollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.SendContext(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// QueuePollResponse is the response for the Poll method.
type QueuePollResponse struct {
	response *QueueGetResponse
}

// Status returns the response status code.
func (r *QueuePollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *QueuePollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *QueuePollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
//
//
func (r *QueuePollResponse) Body() *Queue {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *QueuePollResponse) GetBody() (value *Queue, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *QueueClient) Poll() *QueuePollRequest {
	return &QueuePollRequest{
		request: c.Get(),
	}
}

// QueueGetRequest is the request for the 'get' method.
type QueueGetRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *QueueGetRequest) Parameter(name string, value interface{}) *QueueGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *QueueGetRequest) Header(name string, value interface{}) *QueueGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *QueueGetRequest) Send() (result *QueueGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *QueueGetRequest) SendContext(ctx context.Context) (result *QueueGetResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "GET",
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
	result = &QueueGetResponse{}
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
	err = readQueueGetResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// QueueGetResponse is the response for the 'get' method.
type QueueGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Queue
}

// Status returns the response status code.
func (r *QueueGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *QueueGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *QueueGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *QueueGetResponse) Body() *Queue {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *QueueGetResponse) GetBody() (value *Queue, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// QueuePopRequest is the request for the 'pop' method.
type QueuePopRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *QueuePopRequest) Parameter(name string, value interface{}) *QueuePopRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *QueuePopRequest) Header(name string, value interface{}) *QueuePopRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *QueuePopRequest) Send() (result *QueuePopResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *QueuePopRequest) SendContext(ctx context.Context) (result *QueuePopResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
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
	result = &QueuePopResponse{}
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
	err = readQueuePopResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// QueuePopResponse is the response for the 'pop' method.
type QueuePopResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Job
}

// Status returns the response status code.
func (r *QueuePopResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *QueuePopResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *QueuePopResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *QueuePopResponse) Body() *Job {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *QueuePopResponse) GetBody() (value *Job, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// QueuePushRequest is the request for the 'push' method.
type QueuePushRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *Job
}

// Parameter adds a query parameter.
func (r *QueuePushRequest) Parameter(name string, value interface{}) *QueuePushRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *QueuePushRequest) Header(name string, value interface{}) *QueuePushRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *QueuePushRequest) Body(value *Job) *QueuePushRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *QueuePushRequest) Send() (result *QueuePushResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *QueuePushRequest) SendContext(ctx context.Context) (result *QueuePushResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeQueuePushRequest(r, buffer)
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
	result = &QueuePushResponse{}
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
	err = readQueuePushResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'push' method.
func (r *QueuePushRequest) marshal(writer io.Writer) error {
	stream := helpers.NewStream(writer)
	r.stream(stream)
	return stream.Error
}
func (r *QueuePushRequest) stream(stream *jsoniter.Stream) {
}

// QueuePushResponse is the response for the 'push' method.
type QueuePushResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Job
}

// Status returns the response status code.
func (r *QueuePushResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *QueuePushResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *QueuePushResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *QueuePushResponse) Body() *Job {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *QueuePushResponse) GetBody() (value *Job, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}
