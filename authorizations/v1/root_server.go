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

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RootServer represents the interface the manages the 'root' resource.
type RootServer interface {

	// ExportControlReview returns the target 'export_control_review' resource.
	//
	// Reference to the resource that is used to submit export control review requests.
	ExportControlReview() ExportControlReviewServer
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
	adapter.router.PathPrefix("/export_control_review").HandlerFunc(adapter.exportControlReviewHandler)
	return adapter
}
func (a *RootServerAdapter) exportControlReviewHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.ExportControlReview()
	targetAdapter := NewExportControlReviewServerAdapter(target, a.router.PathPrefix("/export_control_review").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
