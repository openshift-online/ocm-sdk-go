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
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	time "time"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// DashboardClient is the client of the 'dashboard' resource.
//
// Manages a specific dashboard.
type DashboardClient struct {
	transport http.RoundTripper
	path      string
}

// NewDashboardClient creates a new client for the 'dashboard'
// resource using the given transport to sned the requests and receive the
// responses.
func NewDashboardClient(transport http.RoundTripper, path string) *DashboardClient {
	client := new(DashboardClient)
	client.transport = transport
	client.path = path
	return client
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the dashboard.
func (c *DashboardClient) Get() *DashboardGetRequest {
	request := new(DashboardGetRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// DashboardGetRequest is the request for the 'get' method.
type DashboardGetRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *DashboardGetRequest) Context(value context.Context) *DashboardGetRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *DashboardGetRequest) Timeout(value time.Duration) *DashboardGetRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *DashboardGetRequest) Deadline(value time.Time) *DashboardGetRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *DashboardGetRequest) Parameter(name string, value interface{}) *DashboardGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *DashboardGetRequest) Header(name string, value interface{}) *DashboardGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *DashboardGetRequest) Send() (result *DashboardGetResponse, err error) {
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
	result = new(DashboardGetResponse)
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

// DashboardGetResponse is the response for the 'get' method.
type DashboardGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Dashboard
}

// Status returns the response status code.
func (r *DashboardGetResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *DashboardGetResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *DashboardGetResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *DashboardGetResponse) Body() *Dashboard {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'get' method.
func (r *DashboardGetResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(dashboardData)
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
