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

// MetricQueriesServer represents the interface the manages the 'metric_queries' resource.
type MetricQueriesServer interface {

	// CPUTotalByNodeRolesOS returns the target 'CPU_total_by_node_roles_OS_metric_query' resource.
	//
	// Reference to the resource that retrieves 24 hour history of the amount of total cpu
	// capacity in the cluster by node role and operating system.
	CPUTotalByNodeRolesOS() CPUTotalByNodeRolesOSMetricQueryServer
}

// MetricQueriesAdapter is an HTTP handler that knows how to translate HTTP requests
// into calls to the methods of an object that implements the MetricQueriesServer
// interface.
type MetricQueriesAdapter struct {
	server MetricQueriesServer
}

// NewMetricQueriesAdapter creates a new adapter that will translate HTTP requests
// into calls to the given server.
func NewMetricQueriesAdapter(server MetricQueriesServer) *MetricQueriesAdapter {
	return &MetricQueriesAdapter{
		server: server,
	}
}

// ServeHTTP is the implementation of the http.Handler interface.
func (a *MetricQueriesAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dispatchMetricQueriesRequest(w, r, a.server, helpers.Segments(r.URL.Path))
}

// dispatchMetricQueriesRequest navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchMetricQueriesRequest(w http.ResponseWriter, r *http.Request, server MetricQueriesServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		default:
			errors.SendMethodNotSupported(w, r)
		}
	} else {
		switch segments[0] {
		case "cpu_total_by_node_roles_os":
			target := server.CPUTotalByNodeRolesOS()
			dispatchCPUTotalByNodeRolesOSMetricQueryRequest(w, r, target, segments[1:])
		default:
			errors.SendNotFound(w, r)
		}
	}
}
