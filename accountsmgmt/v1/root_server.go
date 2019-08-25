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

package v1 // github.com/openshift-online/uhc-sdk-go/accountsmgmt/v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RootServer represents the interface the manages the 'root' resource.
type RootServer interface {

	// Accounts returns the target 'accounts' resource.
	//
	// Reference to the resource that manages the collection of accounts.
	Accounts() AccountsServer

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

	// AccessToken returns the target 'access_token' resource.
	//
	// Reference to the resource that manages generates access tokens.
	AccessToken() AccessTokenServer

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

	// ClusterAuthorizations returns the target 'cluster_authorizations' resource.
	//
	// Reference to the resource that manages cluster authorizations.
	ClusterAuthorizations() ClusterAuthorizationsServer

	// ClusterRegistrations returns the target 'cluster_registrations' resource.
	//
	// Reference to the resource that manages cluster registrations.
	ClusterRegistrations() ClusterRegistrationsServer

	// Roles returns the target 'roles' resource.
	//
	// Reference to the resource that manages the collection of roles.
	Roles() RolesServer

	// RoleBindings returns the target 'role_bindings' resource.
	//
	// Reference to the resource that manages the collection of role
	// bindings.
	RoleBindings() RoleBindingsServer

	// Subscriptions returns the target 'subscriptions' resource.
	//
	// Reference to the resource that manages the collection of
	// subscriptions.
	Subscriptions() SubscriptionsServer
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
	adapter.router.PathPrefix("/accounts").HandlerFunc(adapter.accountsHandler)
	adapter.router.PathPrefix("/current_account").HandlerFunc(adapter.currentAccountHandler)
	adapter.router.PathPrefix("/organizations").HandlerFunc(adapter.organizationsHandler)
	adapter.router.PathPrefix("/access_token").HandlerFunc(adapter.accessTokenHandler)
	adapter.router.PathPrefix("/permissions").HandlerFunc(adapter.permissionsHandler)
	adapter.router.PathPrefix("/registries").HandlerFunc(adapter.registriesHandler)
	adapter.router.PathPrefix("/registry_credentials").HandlerFunc(adapter.registryCredentialsHandler)
	adapter.router.PathPrefix("/cluster_authorizations").HandlerFunc(adapter.clusterAuthorizationsHandler)
	adapter.router.PathPrefix("/cluster_registrations").HandlerFunc(adapter.clusterRegistrationsHandler)
	adapter.router.PathPrefix("/roles").HandlerFunc(adapter.rolesHandler)
	adapter.router.PathPrefix("/role_bindings").HandlerFunc(adapter.roleBindingsHandler)
	adapter.router.PathPrefix("/subscriptions").HandlerFunc(adapter.subscriptionsHandler)
	return adapter
}
func (a *RootServerAdapter) accountsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Accounts()
	targetAdapter := NewAccountsServerAdapter(target, a.router.PathPrefix("/accounts").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) currentAccountHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.CurrentAccount()
	targetAdapter := NewCurrentAccountServerAdapter(target, a.router.PathPrefix("/current_account").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) organizationsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Organizations()
	targetAdapter := NewOrganizationsServerAdapter(target, a.router.PathPrefix("/organizations").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) accessTokenHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.AccessToken()
	targetAdapter := NewAccessTokenServerAdapter(target, a.router.PathPrefix("/access_token").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) permissionsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Permissions()
	targetAdapter := NewPermissionsServerAdapter(target, a.router.PathPrefix("/permissions").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) registriesHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Registries()
	targetAdapter := NewRegistriesServerAdapter(target, a.router.PathPrefix("/registries").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) registryCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.RegistryCredentials()
	targetAdapter := NewRegistryCredentialsServerAdapter(target, a.router.PathPrefix("/registry_credentials").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) clusterAuthorizationsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.ClusterAuthorizations()
	targetAdapter := NewClusterAuthorizationsServerAdapter(target, a.router.PathPrefix("/cluster_authorizations").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) clusterRegistrationsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.ClusterRegistrations()
	targetAdapter := NewClusterRegistrationsServerAdapter(target, a.router.PathPrefix("/cluster_registrations").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) rolesHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Roles()
	targetAdapter := NewRolesServerAdapter(target, a.router.PathPrefix("/roles").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) roleBindingsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.RoleBindings()
	targetAdapter := NewRoleBindingsServerAdapter(target, a.router.PathPrefix("/role_bindings").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) subscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	target := a.server.Subscriptions()
	targetAdapter := NewSubscriptionsServerAdapter(target, a.router.PathPrefix("/subscriptions").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *RootServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
