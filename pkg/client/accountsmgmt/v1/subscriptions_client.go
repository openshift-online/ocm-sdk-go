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
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// SubscriptionsClient is the client of the 'subscriptions' resource.
//
// Manages the collection of subscriptions.
type SubscriptionsClient struct {
	transport http.RoundTripper
	path      string
}

// NewSubscriptionsClient creates a new client for the 'subscriptions'
// resource using the given transport to sned the requests and receive the
// responses.
func NewSubscriptionsClient(transport http.RoundTripper, path string) *SubscriptionsClient {
	client := new(SubscriptionsClient)
	client.transport = transport
	client.path = path
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves a list of subscriptions.
func (c *SubscriptionsClient) List() *SubscriptionsListRequest {
	request := new(SubscriptionsListRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Subscription returns the target 'subscription' resource for the given identifier.
//
// Reference to the service that manages a specific subscription.
func (c *SubscriptionsClient) Subscription(id string) *SubscriptionClient {
	return NewSubscriptionClient(c.transport, path.Join(c.path, id))
}

// SubscriptionsListRequest is the request for the 'list' method.
type SubscriptionsListRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	page      *int
	size      *int
	total     *int
}

// Parameter adds a query parameter.
func (r *SubscriptionsListRequest) Parameter(name string, value interface{}) *SubscriptionsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *SubscriptionsListRequest) Header(name string, value interface{}) *SubscriptionsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *SubscriptionsListRequest) Page(value int) *SubscriptionsListRequest {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *SubscriptionsListRequest) Size(value int) *SubscriptionsListRequest {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *SubscriptionsListRequest) Total(value int) *SubscriptionsListRequest {
	r.total = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method. If you don't provide a
// context then a new background context will be created.
func (r *SubscriptionsListRequest) Send() (result *SubscriptionsListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *SubscriptionsListRequest) SendContext(ctx context.Context) (result *SubscriptionsListResponse, err error) {
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
	result = new(SubscriptionsListResponse)
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

// SubscriptionsListResponse is the response for the 'list' method.
type SubscriptionsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *SubscriptionList
}

// Status returns the response status code.
func (r *SubscriptionsListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *SubscriptionsListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *SubscriptionsListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *SubscriptionsListResponse) Page() int {
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
func (r *SubscriptionsListResponse) Size() int {
	if r.size != nil {
		return *r.size
	}
	return 0
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *SubscriptionsListResponse) Total() int {
	if r.total != nil {
		return *r.total
	}
	return 0
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of subscriptions.
func (r *SubscriptionsListResponse) Items() *SubscriptionList {
	return r.items
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *SubscriptionsListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(subscriptionsListResponseData)
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

// subscriptionsListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type subscriptionsListResponseData struct {
	Page  *int                 "json:\"page,omitempty\""
	Size  *int                 "json:\"size,omitempty\""
	Total *int                 "json:\"total,omitempty\""
	Items subscriptionListData "json:\"items,omitempty\""
}
