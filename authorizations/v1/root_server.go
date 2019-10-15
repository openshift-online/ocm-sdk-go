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

	// AccessReview returns the target 'access_review' resource.
	//
	// Reference to the resource that is used to submit access review requests.
	AccessReview() AccessReviewServer

	// ExportControlReview returns the target 'export_control_review' resource.
	//
	// Reference to the resource that is used to submit export control review requests.
	ExportControlReview() ExportControlReviewServer

	// SelfAccessReview returns the target 'self_access_review' resource.
	//
	// Reference to the resource that is used to submit self access review requests.
	SelfAccessReview() SelfAccessReviewServer
}

// RootAdapter represents the structs that adapts Requests and Response to internal
// structs.
type RootAdapter struct {
	server RootServer
	router *mux.Router
}

func NewRootAdapter(server RootServer, router *mux.Router) *RootAdapter {
	adapter := new(RootAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/access_review").HandlerFunc(adapter.accessReviewHandler)
	adapter.router.PathPrefix("/export_control_review").HandlerFunc(adapter.exportControlReviewHandler)
	adapter.router.PathPrefix("/self_access_review").HandlerFunc(adapter.selfAccessReviewHandler)
	return adapter
}
func (a *RootAdapter) accessReviewHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.AccessReview()
	targetAdapter := NewAccessReviewAdapter(target, a.router.PathPrefix("/access_review").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) exportControlReviewHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.ExportControlReview()
	targetAdapter := NewExportControlReviewAdapter(target, a.router.PathPrefix("/export_control_review").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) selfAccessReviewHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.SelfAccessReview()
	targetAdapter := NewSelfAccessReviewAdapter(target, a.router.PathPrefix("/self_access_review").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
