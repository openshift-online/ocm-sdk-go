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

// CredentialsClient is the client of the 'credentials' resource.
//
// Manages the credentials of a cluster.
type CredentialsClient struct {
	transport http.RoundTripper
	path      string
}

// NewCredentialsClient creates a new client for the 'credentials'
// resource using the given transport to sned the requests and receive the
// responses.
func NewCredentialsClient(transport http.RoundTripper, path string) *CredentialsClient {
	client := new(CredentialsClient)
	client.transport = transport
	client.path = path
	return client
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the credentials of a cluster.
func (c *CredentialsClient) Get() *CredentialsGetRequest {
	request := new(CredentialsGetRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// CredentialsGetRequest is the request for the 'get' method.
type CredentialsGetRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *CredentialsGetRequest) Context(value context.Context) *CredentialsGetRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *CredentialsGetRequest) Timeout(value time.Duration) *CredentialsGetRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *CredentialsGetRequest) Deadline(value time.Time) *CredentialsGetRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *CredentialsGetRequest) Parameter(name string, value interface{}) *CredentialsGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *CredentialsGetRequest) Header(name string, value interface{}) *CredentialsGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *CredentialsGetRequest) Send() (result *CredentialsGetResponse, err error) {
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
	result = new(CredentialsGetResponse)
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

// CredentialsGetResponse is the response for the 'get' method.
type CredentialsGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *ClusterCredentials
}

// Status returns the response status code.
func (r *CredentialsGetResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *CredentialsGetResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *CredentialsGetResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *CredentialsGetResponse) Body() *ClusterCredentials {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'get' method.
func (r *CredentialsGetResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(clusterCredentialsData)
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
