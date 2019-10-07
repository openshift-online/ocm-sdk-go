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

	"github.com/gorilla/mux"
)

// MetricQueriesServer represents the interface the manages the 'metric_queries' resource.
type MetricQueriesServer interface {

	// CPUTotalByNodeRolesOS returns the target 'CPU_total_by_node_roles_OS_metric_query' resource.
	//
	// Reference to the resource that retrieves 24 hour history of the amount of total cpu
	// capacity in the cluster by node role and operating system.
	CPUTotalByNodeRolesOS() CPUTotalByNodeRolesOSMetricQueryServer
}

// MetricQueriesServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type MetricQueriesServerAdapter struct {
	server MetricQueriesServer
	router *mux.Router
}

func NewMetricQueriesServerAdapter(server MetricQueriesServer, router *mux.Router) *MetricQueriesServerAdapter {
	adapter := new(MetricQueriesServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/cpu_total_by_node_roles_os").HandlerFunc(adapter.cpuTotalByNodeRolesOSHandler)
	return adapter
}
func (a *MetricQueriesServerAdapter) cpuTotalByNodeRolesOSHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.CPUTotalByNodeRolesOS()
	targetAdapter := NewCPUTotalByNodeRolesOSMetricQueryServerAdapter(target, a.router.PathPrefix("/cpu_total_by_node_roles_os").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *MetricQueriesServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
