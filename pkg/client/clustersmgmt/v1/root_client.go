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
	"net/http"
	"path"
)

// RootClient is the client of the 'root' resource.
//
// Root of the tree of resources of the clusters management service.
type RootClient struct {
	transport http.RoundTripper
	path      string
}

// NewRootClient creates a new client for the 'root'
// resource using the given transport to sned the requests and receive the
// responses.
func NewRootClient(transport http.RoundTripper, path string) *RootClient {
	client := new(RootClient)
	client.transport = transport
	client.path = path
	return client
}

// Clusters returns the target 'clusters' resource.
//
// Reference to the resource that manages the collection of clusters.
func (c *RootClient) Clusters() *ClustersClient {
	return NewClustersClient(c.transport, path.Join(c.path, "clusters"))
}

// Dashboards returns the target 'dashboards' resource.
//
// Reference to the resource that manages the collection of dashboards.
func (c *RootClient) Dashboards() *DashboardsClient {
	return NewDashboardsClient(c.transport, path.Join(c.path, "dashboards"))
}

// Flavours returns the target 'flavours' resource.
//
// Reference to the service that manages the collection of flavours.
func (c *RootClient) Flavours() *FlavoursClient {
	return NewFlavoursClient(c.transport, path.Join(c.path, "flavours"))
}
