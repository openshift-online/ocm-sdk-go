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

	"github.com/gorilla/mux"
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
	adapter.router.PathPrefix("/skus").HandlerFunc(adapter.skusHandler)
	adapter.router.PathPrefix("/access_token").HandlerFunc(adapter.accessTokenHandler)
	adapter.router.PathPrefix("/accounts").HandlerFunc(adapter.accountsHandler)
	adapter.router.PathPrefix("/cluster_authorizations").HandlerFunc(adapter.clusterAuthorizationsHandler)
	adapter.router.PathPrefix("/cluster_registrations").HandlerFunc(adapter.clusterRegistrationsHandler)
	adapter.router.PathPrefix("/current_account").HandlerFunc(adapter.currentAccountHandler)
	adapter.router.PathPrefix("/organizations").HandlerFunc(adapter.organizationsHandler)
	adapter.router.PathPrefix("/permissions").HandlerFunc(adapter.permissionsHandler)
	adapter.router.PathPrefix("/registries").HandlerFunc(adapter.registriesHandler)
	adapter.router.PathPrefix("/registry_credentials").HandlerFunc(adapter.registryCredentialsHandler)
	adapter.router.PathPrefix("/role_bindings").HandlerFunc(adapter.roleBindingsHandler)
	adapter.router.PathPrefix("/roles").HandlerFunc(adapter.rolesHandler)
	adapter.router.PathPrefix("/subscriptions").HandlerFunc(adapter.subscriptionsHandler)
	return adapter
}
func (a *RootAdapter) skusHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.SKUS()
	targetAdapter := NewSKUSAdapter(target, a.router.PathPrefix("/skus").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) accessTokenHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.AccessToken()
	targetAdapter := NewAccessTokenAdapter(target, a.router.PathPrefix("/access_token").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) accountsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Accounts()
	targetAdapter := NewAccountsAdapter(target, a.router.PathPrefix("/accounts").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) clusterAuthorizationsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.ClusterAuthorizations()
	targetAdapter := NewClusterAuthorizationsAdapter(target, a.router.PathPrefix("/cluster_authorizations").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) clusterRegistrationsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.ClusterRegistrations()
	targetAdapter := NewClusterRegistrationsAdapter(target, a.router.PathPrefix("/cluster_registrations").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) currentAccountHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.CurrentAccount()
	targetAdapter := NewCurrentAccountAdapter(target, a.router.PathPrefix("/current_account").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) organizationsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Organizations()
	targetAdapter := NewOrganizationsAdapter(target, a.router.PathPrefix("/organizations").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) permissionsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Permissions()
	targetAdapter := NewPermissionsAdapter(target, a.router.PathPrefix("/permissions").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) registriesHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Registries()
	targetAdapter := NewRegistriesAdapter(target, a.router.PathPrefix("/registries").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) registryCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.RegistryCredentials()
	targetAdapter := NewRegistryCredentialsAdapter(target, a.router.PathPrefix("/registry_credentials").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) roleBindingsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.RoleBindings()
	targetAdapter := NewRoleBindingsAdapter(target, a.router.PathPrefix("/role_bindings").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) rolesHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Roles()
	targetAdapter := NewRolesAdapter(target, a.router.PathPrefix("/roles").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) subscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Subscriptions()
	targetAdapter := NewSubscriptionsAdapter(target, a.router.PathPrefix("/subscriptions").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
