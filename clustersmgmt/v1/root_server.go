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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// RootServer represents the interface the manages the 'root' resource.
type RootServer interface {

	// CloudProviders returns the target 'cloud_providers' resource.
	//
	// Reference to the resource that manages the collection of cloud providers.
	CloudProviders() CloudProvidersServer

	// Clusters returns the target 'clusters' resource.
	//
	// Reference to the resource that manages the collection of clusters.
	Clusters() ClustersServer

	// Dashboards returns the target 'dashboards' resource.
	//
	// Reference to the resource that manages the collection of dashboards.
	Dashboards() DashboardsServer

	// Flavours returns the target 'flavours' resource.
	//
	// Reference to the service that manages the collection of flavours.
	Flavours() FlavoursServer

	// Versions returns the target 'versions' resource.
	//
	// Reference to the resource that manage the collection of versions.
	Versions() VersionsServer
}

// RootAdapter is an HTTP handler that knows how to translate HTTP requests
// into calls to the methods of an object that implements the RootServer
// interface.
type RootAdapter struct {
	server RootServer
}

// NewRootAdapter creates a new adapter that will translate HTTP requests
// into calls to the given server.
func NewRootAdapter(server RootServer) *RootAdapter {
	return &RootAdapter{
		server: server,
	}
}

// ServeHTTP is the implementation of the http.Handler interface.
func (a *RootAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dispatchRootRequest(w, r, a.server, helpers.Segments(r.URL.Path))
}

// dispatchRootRequest navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchRootRequest(w http.ResponseWriter, r *http.Request, server RootServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		default:
			errors.SendMethodNotSupported(w, r)
		}
	} else {
		switch segments[0] {
		case "cloud_providers":
			target := server.CloudProviders()
			dispatchCloudProvidersRequest(w, r, target, segments[1:])
		case "clusters":
			target := server.Clusters()
			dispatchClustersRequest(w, r, target, segments[1:])
		case "dashboards":
			target := server.Dashboards()
			dispatchDashboardsRequest(w, r, target, segments[1:])
		case "flavours":
			target := server.Flavours()
			dispatchFlavoursRequest(w, r, target, segments[1:])
		case "versions":
			target := server.Versions()
			dispatchVersionsRequest(w, r, target, segments[1:])
		default:
			errors.SendNotFound(w, r)
		}
	}
}
