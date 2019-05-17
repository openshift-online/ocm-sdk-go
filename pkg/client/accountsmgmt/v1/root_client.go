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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

import (
	"net/http"
	"path"
)

// RootClient is the client of the 'root' resource.
//
// Root of the tree of resources of the clusters management service.
type RootClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewRootClient creates a new client for the 'root'
// resource using the given transport to sned the requests and receive the
// responses.
func NewRootClient(transport http.RoundTripper, path string, metric string) *RootClient {
	client := new(RootClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// Accounts returns the target 'accounts' resource.
//
// Reference to the resource that manages the collection of accounts.
func (c *RootClient) Accounts() *AccountsClient {
	return NewAccountsClient(
		c.transport,
		path.Join(c.path, "accounts"),
		path.Join(c.metric, "accounts"),
	)
}

// CurrentAccount returns the target 'current_account' resource.
//
// Reference to the resource that manages the current authenticated
// acount.
func (c *RootClient) CurrentAccount() *CurrentAccountClient {
	return NewCurrentAccountClient(
		c.transport,
		path.Join(c.path, "current_account"),
		path.Join(c.metric, "current_account"),
	)
}

// Organizations returns the target 'organizations' resource.
//
// Reference to the resource that manages the collection of
// organizations.
func (c *RootClient) Organizations() *OrganizationsClient {
	return NewOrganizationsClient(
		c.transport,
		path.Join(c.path, "organizations"),
		path.Join(c.metric, "organizations"),
	)
}

// AccessToken returns the target 'access_token' resource.
//
// Reference to the resource that manages generates access tokens.
func (c *RootClient) AccessToken() *AccessTokenClient {
	return NewAccessTokenClient(
		c.transport,
		path.Join(c.path, "access_token"),
		path.Join(c.metric, "access_token"),
	)
}

// Permissions returns the target 'permissions' resource.
//
// Reference to the resource that manages the collection of permissions.
func (c *RootClient) Permissions() *PermissionsClient {
	return NewPermissionsClient(
		c.transport,
		path.Join(c.path, "permissions"),
		path.Join(c.metric, "permissions"),
	)
}

// Registries returns the target 'registries' resource.
//
// Reference to the resource that manages the collection of registries.
func (c *RootClient) Registries() *RegistriesClient {
	return NewRegistriesClient(
		c.transport,
		path.Join(c.path, "registries"),
		path.Join(c.metric, "registries"),
	)
}

// RegistryCredentials returns the target 'registry_credentials' resource.
//
// Reference to the resource that manages the collection of registry
// credentials.
func (c *RootClient) RegistryCredentials() *RegistryCredentialsClient {
	return NewRegistryCredentialsClient(
		c.transport,
		path.Join(c.path, "registry_credentials"),
		path.Join(c.metric, "registry_credentials"),
	)
}

// ClusterAuthorizations returns the target 'cluster_authorizations' resource.
//
// Reference to the resource that manages cluster authorizations.
func (c *RootClient) ClusterAuthorizations() *ClusterAuthorizationsClient {
	return NewClusterAuthorizationsClient(
		c.transport,
		path.Join(c.path, "cluster_authorizations"),
		path.Join(c.metric, "cluster_authorizations"),
	)
}

// ClusterRegistrations returns the target 'cluster_registrations' resource.
//
// Reference to the resource that manages cluster registrations.
func (c *RootClient) ClusterRegistrations() *ClusterRegistrationsClient {
	return NewClusterRegistrationsClient(
		c.transport,
		path.Join(c.path, "cluster_registrations"),
		path.Join(c.metric, "cluster_registrations"),
	)
}

// Roles returns the target 'roles' resource.
//
// Reference to the resource that manages the collection of roles.
func (c *RootClient) Roles() *RolesClient {
	return NewRolesClient(
		c.transport,
		path.Join(c.path, "roles"),
		path.Join(c.metric, "roles"),
	)
}

// RoleBindings returns the target 'role_bindings' resource.
//
// Reference to the resource that manages the collection of role
// bindings.
func (c *RootClient) RoleBindings() *RoleBindingsClient {
	return NewRoleBindingsClient(
		c.transport,
		path.Join(c.path, "role_bindings"),
		path.Join(c.metric, "role_bindings"),
	)
}

// Subscriptions returns the target 'subscriptions' resource.
//
// Reference to the resource that manages the collection of
// subscriptions.
func (c *RootClient) Subscriptions() *SubscriptionsClient {
	return NewSubscriptionsClient(
		c.transport,
		path.Join(c.path, "subscriptions"),
		path.Join(c.metric, "subscriptions"),
	)
}
