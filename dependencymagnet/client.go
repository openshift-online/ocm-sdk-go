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

package dependencymagnet // github.com/openshift-online/ocm-sdk-go/dependencymagnet

import (
	"net/http"
)

// Client is the client for service 'dependencymagnet'.
type Client struct {
	transport http.RoundTripper
	path      string
}

// NewClient creates a new client for the service 'dependencymagnet' using the
// given transport to send the requests and receive the responses.
func NewClient(transport http.RoundTripper, path string) *Client {
	client := new(Client)
	client.transport = transport
	client.path = path
	return client
}
