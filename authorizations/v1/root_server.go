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

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
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
		case "access_review":
			target := server.AccessReview()
			dispatchAccessReviewRequest(w, r, target, segments[1:])
		case "export_control_review":
			target := server.ExportControlReview()
			dispatchExportControlReviewRequest(w, r, target, segments[1:])
		case "self_access_review":
			target := server.SelfAccessReview()
			dispatchSelfAccessReviewRequest(w, r, target, segments[1:])
		default:
			errors.SendNotFound(w, r)
		}
	}
}
