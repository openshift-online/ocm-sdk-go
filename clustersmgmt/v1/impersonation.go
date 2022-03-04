/*
Copyright (c) 2022 Red Hat, Inc.

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
package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import "github.com/openshift-online/ocm-sdk-go/helpers"

// ImpersonateUser is the name of the request header that is used to indicate that the caller wants
// to impersonate another user.
const ImpersonateUserHeader = "Impersonate-User"

// Impersonation can be done by adding a "Impersonate-User" entry to the header.
// However, use of the sdk may look cleaner by
// adding and using the impersonation methods defined below.

// Call this method to impersonate a user when adding a cluster
func (r *ClustersAddRequest) Impersonate(username string) *ClustersAddRequest {
	helpers.AddHeader(&r.header, ImpersonateUserHeader, username)
	return r
}
