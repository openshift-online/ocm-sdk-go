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

// RootServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RootServerAdapter struct {
	server RootServer
	router *mux.Router
}

func NewRootServerAdapter(server RootServer, router *mux.Router) *RootServerAdapter {
	adapter := new(RootServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/cloud_providers").HandlerFunc(adapter.cloudProvidersHandler)
	adapter.router.PathPrefix("/clusters").HandlerFunc(adapter.clustersHandler)
	adapter.router.PathPrefix("/dashboards").HandlerFunc(adapter.dashboardsHandler)
	adapter.router.PathPrefix("/flavours").HandlerFunc(adapter.flavoursHandler)
	adapter.router.PathPrefix("/versions").HandlerFunc(adapter.versionsHandler)
	return adapter
}
func (a *RootServerAdapter) cloudProvidersHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.CloudProviders()
	targetAdapter := NewCloudProvidersServerAdapter(target, a.router.PathPrefix("/cloud_providers").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) clustersHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Clusters()
	targetAdapter := NewClustersServerAdapter(target, a.router.PathPrefix("/clusters").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) dashboardsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Dashboards()
	targetAdapter := NewDashboardsServerAdapter(target, a.router.PathPrefix("/dashboards").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) flavoursHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Flavours()
	targetAdapter := NewFlavoursServerAdapter(target, a.router.PathPrefix("/flavours").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) versionsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Versions()
	targetAdapter := NewVersionsServerAdapter(target, a.router.PathPrefix("/versions").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
