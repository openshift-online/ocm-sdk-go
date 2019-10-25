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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// RootServer represents the interface the manages the 'root' resource.
type RootServer interface {

	// SKUS returns the target 'SKUS' resource.
	//
	// Reference to the resource that manages the collection of
	// SKUS
	SKUS() SKUSServer

	// AccessToken returns the target 'access_token' resource.
	//
	// Reference to the resource that manages generates access tokens.
	AccessToken() AccessTokenServer

	// Accounts returns the target 'accounts' resource.
	//
	// Reference to the resource that manages the collection of accounts.
	Accounts() AccountsServer

	// ClusterAuthorizations returns the target 'cluster_authorizations' resource.
	//
	// Reference to the resource that manages cluster authorizations.
	ClusterAuthorizations() ClusterAuthorizationsServer

	// ClusterRegistrations returns the target 'cluster_registrations' resource.
	//
	// Reference to the resource that manages cluster registrations.
	ClusterRegistrations() ClusterRegistrationsServer

	// CurrentAccount returns the target 'current_account' resource.
	//
	// Reference to the resource that manages the current authenticated
	// acount.
	CurrentAccount() CurrentAccountServer

	// Organizations returns the target 'organizations' resource.
	//
	// Reference to the resource that manages the collection of
	// organizations.
	Organizations() OrganizationsServer

	// Permissions returns the target 'permissions' resource.
	//
	// Reference to the resource that manages the collection of permissions.
	Permissions() PermissionsServer

	// Registries returns the target 'registries' resource.
	//
	// Reference to the resource that manages the collection of registries.
	Registries() RegistriesServer

	// RegistryCredentials returns the target 'registry_credentials' resource.
	//
	// Reference to the resource that manages the collection of registry
	// credentials.
	RegistryCredentials() RegistryCredentialsServer

	// RoleBindings returns the target 'role_bindings' resource.
	//
	// Reference to the resource that manages the collection of role
	// bindings.
	RoleBindings() RoleBindingsServer

	// Roles returns the target 'roles' resource.
	//
	// Reference to the resource that manages the collection of roles.
	Roles() RolesServer

	// Subscriptions returns the target 'subscriptions' resource.
	//
	// Reference to the resource that manages the collection of
	// subscriptions.
	Subscriptions() SubscriptionsServer
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
		case "skus":
			target := server.SKUS()
			dispatchSKUSRequest(w, r, target, segments[1:])
		case "access_token":
			target := server.AccessToken()
			dispatchAccessTokenRequest(w, r, target, segments[1:])
		case "accounts":
			target := server.Accounts()
			dispatchAccountsRequest(w, r, target, segments[1:])
		case "cluster_authorizations":
			target := server.ClusterAuthorizations()
			dispatchClusterAuthorizationsRequest(w, r, target, segments[1:])
		case "cluster_registrations":
			target := server.ClusterRegistrations()
			dispatchClusterRegistrationsRequest(w, r, target, segments[1:])
		case "current_account":
			target := server.CurrentAccount()
			dispatchCurrentAccountRequest(w, r, target, segments[1:])
		case "organizations":
			target := server.Organizations()
			dispatchOrganizationsRequest(w, r, target, segments[1:])
		case "permissions":
			target := server.Permissions()
			dispatchPermissionsRequest(w, r, target, segments[1:])
		case "registries":
			target := server.Registries()
			dispatchRegistriesRequest(w, r, target, segments[1:])
		case "registry_credentials":
			target := server.RegistryCredentials()
			dispatchRegistryCredentialsRequest(w, r, target, segments[1:])
		case "role_bindings":
			target := server.RoleBindings()
			dispatchRoleBindingsRequest(w, r, target, segments[1:])
		case "roles":
			target := server.Roles()
			dispatchRolesRequest(w, r, target, segments[1:])
		case "subscriptions":
			target := server.Subscriptions()
			dispatchSubscriptionsRequest(w, r, target, segments[1:])
		default:
			errors.SendNotFound(w, r)
		}
	}
}
