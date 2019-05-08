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

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// ClustersClient is the client of the 'clusters' resource.
//
// Manages the collection of clusters.
type ClustersClient struct {
	transport http.RoundTripper
	path      string
}

// NewClustersClient creates a new client for the 'clusters'
// resource using the given transport to sned the requests and receive the
// responses.
func NewClustersClient(transport http.RoundTripper, path string) *ClustersClient {
	client := new(ClustersClient)
	client.transport = transport
	client.path = path
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of clusters.
func (c *ClustersClient) List() *ClustersListRequest {
	request := new(ClustersListRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Add creates a request for the 'add' method.
//
// Provision a new cluster and add it to the collection of clusters.
//
// See the `register_cluster` method for adding an existing cluster.
func (c *ClustersClient) Add() *ClustersAddRequest {
	request := new(ClustersAddRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Cluster returns the target 'cluster' resource for the given identifier.
//
// Retursn a reference to the service that manages an specific cluster.
func (c *ClustersClient) Cluster(id string) *ClusterClient {
	return NewClusterClient(c.transport, path.Join(c.path, id))
}

// ClustersListRequest is the request for the 'list' method.
type ClustersListRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	page      *int
	size      *int
	search    *string
	total     *int
}

// Parameter adds a query parameter.
func (r *ClustersListRequest) Parameter(name string, value interface{}) *ClustersListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ClustersListRequest) Header(name string, value interface{}) *ClustersListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *ClustersListRequest) Page(value int) *ClustersListRequest {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *ClustersListRequest) Size(value int) *ClustersListRequest {
	r.size = &value
	return r
}

// Search sets the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause
// of an SQL statement, but using the names of the attributes of the cluster
// instead of the names of the columns of a table. For example, in order to
// retrieve all the clusters with a name starting with `my` in the
// `us-east-1` region the value should be:
//
// [source,sql]
// ----
// name like 'my%' and region.id = 'us-east-1'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the
// clusters that the user has permission to see will be returned.
func (r *ClustersListRequest) Search(value string) *ClustersListRequest {
	r.search = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *ClustersListRequest) Total(value int) *ClustersListRequest {
	r.total = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method. If you don't provide a
// context then a new background context will be created.
func (r *ClustersListRequest) Send() (result *ClustersListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *ClustersListRequest) SendContext(ctx context.Context) (result *ClustersListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
	if r.search != nil {
		helpers.AddValue(&query, "search", *r.search)
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
	result = new(ClustersListResponse)
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

// ClustersListResponse is the response for the 'list' method.
type ClustersListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *ClusterList
}

// Status returns the response status code.
func (r *ClustersListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *ClustersListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *ClustersListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *ClustersListResponse) Page() int {
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
func (r *ClustersListResponse) Size() int {
	if r.size != nil {
		return *r.size
	}
	return 0
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *ClustersListResponse) Total() int {
	if r.total != nil {
		return *r.total
	}
	return 0
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of clusters.
func (r *ClustersListResponse) Items() *ClusterList {
	return r.items
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *ClustersListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(clustersListResponseData)
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

// clustersListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type clustersListResponseData struct {
	Page  *int            "json:\"page,omitempty\""
	Size  *int            "json:\"size,omitempty\""
	Total *int            "json:\"total,omitempty\""
	Items clusterListData "json:\"items,omitempty\""
}

// ClustersAddRequest is the request for the 'add' method.
type ClustersAddRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *Cluster
}

// Parameter adds a query parameter.
func (r *ClustersAddRequest) Parameter(name string, value interface{}) *ClustersAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ClustersAddRequest) Header(name string, value interface{}) *ClustersAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Description of the cluster.
func (r *ClustersAddRequest) Body(value *Cluster) *ClustersAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method. If you don't provide a
// context then a new background context will be created.
func (r *ClustersAddRequest) Send() (result *ClustersAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *ClustersAddRequest) SendContext(ctx context.Context) (result *ClustersAddResponse, err error) {
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
	result = new(ClustersAddResponse)
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
func (r *ClustersAddRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ClustersAddResponse is the response for the 'add' method.
type ClustersAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Cluster
}

// Status returns the response status code.
func (r *ClustersAddResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *ClustersAddResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *ClustersAddResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Description of the cluster.
func (r *ClustersAddResponse) Body() *Cluster {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'add' method.
func (r *ClustersAddResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(clusterData)
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
