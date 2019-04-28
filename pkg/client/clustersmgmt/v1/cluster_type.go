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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

import (
	time "time"
)

// ClusterKind is the name of the type used to represent objects
// of type 'cluster'.
const ClusterKind = "Cluster"

// ClusterLinkKind is the name of the type used to represent links
// to objects of type 'cluster'.
const ClusterLinkKind = "ClusterLink"

// ClusterNilKind is the name of the type used to nil references
// to objects of type 'cluster'.
const ClusterNilKind = "ClusterNil"

// Cluster represents the values of the 'cluster' type.
//
// Definition of an _OpenShift_ cluster.
//
// The `cloud_provider` attribute is a reference to the cloud provider. When a
// cluster is retrieved it will be a link to the cloud provider, containing only
// the kind, id and href attributes:
//
// [source,json]
// ----
// {
//   "cloud_provider": {
//     "kind": "CloudProviderLink",
//     "id": "123",
//     "href": "/api/clusters_mgmt/v1/cloud_providers/123"
//   }
// }
// ----
//
// When a cluster is created this is optional, and if used it should contain the
// identifier of the cloud provider to use:
//
// [source,json]
// ----
// {
//   "cloud_provider": {
//     "id": "123",
//   }
// }
// ----
//
// If not included, then the cluster will be created using the default cloud
// provider, which is currently Amazon Web Services.
//
// The region attribute is mandatory when a cluster is created.
//
// The `aws.access_key_id`, `aws.secret_access_key` and `dns.base_domain`
// attributes are mandatory when creation a cluster with your own Amazon Web
// Services account.
type Cluster struct {
	id                *string
	href              *string
	link              bool
	name              *string
	flavour           *Flavour
	console           *ClusterConsole
	multiAZ           *bool
	nodes             *ClusterNodes
	api               *ClusterAPI
	region            *CloudRegion
	displayName       *string
	dns               *DNS
	properties        map[string]string
	state             *ClusterState
	managed           *bool
	memory            *ClusterMetric
	cpu               *ClusterMetric
	storage           *ClusterMetric
	externalID        *string
	aws               *AWS
	network           *Network
	creationTimestamp *time.Time
	cloudProvider     *CloudProvider
	openshiftVersion  *string
	subscription      *Subscription
	groups            *GroupList
	creator           *string
	version           *Version
	identityProviders *IdentityProviderList
}

// Kind returns the name of the type of the object.
func (o *Cluster) Kind() string {
	if o == nil {
		return ClusterNilKind
	}
	if o.link {
		return ClusterLinkKind
	}
	return ClusterKind
}

// ID returns the identifier of the object.
func (o *Cluster) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Cluster) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Cluster) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Cluster) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Cluster) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Name of the cluster. This name is assigned by the user when the
// cluster is created.
func (o *Cluster) Name() string {
	if o != nil && o.name != nil {
		return *o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
// Name of the cluster. This name is assigned by the user when the
// cluster is created.
func (o *Cluster) GetName() (value string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = *o.name
	}
	return
}

// Flavour returns the value of the 'flavour' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to the _flavour_ that was used to create the cluster.
func (o *Cluster) Flavour() *Flavour {
	if o == nil {
		return nil
	}
	return o.flavour
}

// GetFlavour returns the value of the 'flavour' attribute and
// a flag indicating if the attribute has a value.
//
// Link to the _flavour_ that was used to create the cluster.
func (o *Cluster) GetFlavour() (value *Flavour, ok bool) {
	ok = o != nil && o.flavour != nil
	if ok {
		value = o.flavour
	}
	return
}

// Console returns the value of the 'console' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Information about the console of the cluster.
func (o *Cluster) Console() *ClusterConsole {
	if o == nil {
		return nil
	}
	return o.console
}

// GetConsole returns the value of the 'console' attribute and
// a flag indicating if the attribute has a value.
//
// Information about the console of the cluster.
func (o *Cluster) GetConsole() (value *ClusterConsole, ok bool) {
	ok = o != nil && o.console != nil
	if ok {
		value = o.console
	}
	return
}

// MultiAZ returns the value of the 'multi_AZ' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Flag indicating if the cluster should be created with nodes in
// different availability zones or all the nodes in a single one
// randomly selected.
func (o *Cluster) MultiAZ() bool {
	if o != nil && o.multiAZ != nil {
		return *o.multiAZ
	}
	return false
}

// GetMultiAZ returns the value of the 'multi_AZ' attribute and
// a flag indicating if the attribute has a value.
//
// Flag indicating if the cluster should be created with nodes in
// different availability zones or all the nodes in a single one
// randomly selected.
func (o *Cluster) GetMultiAZ() (value bool, ok bool) {
	ok = o != nil && o.multiAZ != nil
	if ok {
		value = *o.multiAZ
	}
	return
}

// Nodes returns the value of the 'nodes' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Information about the nodes of the cluster.
func (o *Cluster) Nodes() *ClusterNodes {
	if o == nil {
		return nil
	}
	return o.nodes
}

// GetNodes returns the value of the 'nodes' attribute and
// a flag indicating if the attribute has a value.
//
// Information about the nodes of the cluster.
func (o *Cluster) GetNodes() (value *ClusterNodes, ok bool) {
	ok = o != nil && o.nodes != nil
	if ok {
		value = o.nodes
	}
	return
}

// API returns the value of the 'API' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Information about the API of the cluster.
func (o *Cluster) API() *ClusterAPI {
	if o == nil {
		return nil
	}
	return o.api
}

// GetAPI returns the value of the 'API' attribute and
// a flag indicating if the attribute has a value.
//
// Information about the API of the cluster.
func (o *Cluster) GetAPI() (value *ClusterAPI, ok bool) {
	ok = o != nil && o.api != nil
	if ok {
		value = o.api
	}
	return
}

// Region returns the value of the 'region' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to the cloud provider region where the cluster is installed.
func (o *Cluster) Region() *CloudRegion {
	if o == nil {
		return nil
	}
	return o.region
}

// GetRegion returns the value of the 'region' attribute and
// a flag indicating if the attribute has a value.
//
// Link to the cloud provider region where the cluster is installed.
func (o *Cluster) GetRegion() (value *CloudRegion, ok bool) {
	ok = o != nil && o.region != nil
	if ok {
		value = o.region
	}
	return
}

// DisplayName returns the value of the 'display_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Name of the cluster for display purposes. It can contain any
// characters, including spaces.
func (o *Cluster) DisplayName() string {
	if o != nil && o.displayName != nil {
		return *o.displayName
	}
	return ""
}

// GetDisplayName returns the value of the 'display_name' attribute and
// a flag indicating if the attribute has a value.
//
// Name of the cluster for display purposes. It can contain any
// characters, including spaces.
func (o *Cluster) GetDisplayName() (value string, ok bool) {
	ok = o != nil && o.displayName != nil
	if ok {
		value = *o.displayName
	}
	return
}

// DNS returns the value of the 'DNS' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// DNS settings of the cluster.
func (o *Cluster) DNS() *DNS {
	if o == nil {
		return nil
	}
	return o.dns
}

// GetDNS returns the value of the 'DNS' attribute and
// a flag indicating if the attribute has a value.
//
// DNS settings of the cluster.
func (o *Cluster) GetDNS() (value *DNS, ok bool) {
	ok = o != nil && o.dns != nil
	if ok {
		value = o.dns
	}
	return
}

// Properties returns the value of the 'properties' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// User defined properties for tagging and querying.
func (o *Cluster) Properties() map[string]string {
	if o == nil {
		return nil
	}
	return o.properties
}

// GetProperties returns the value of the 'properties' attribute and
// a flag indicating if the attribute has a value.
//
// User defined properties for tagging and querying.
func (o *Cluster) GetProperties() (value map[string]string, ok bool) {
	ok = o != nil && o.properties != nil
	if ok {
		value = o.properties
	}
	return
}

// State returns the value of the 'state' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Overall state of the cluster.
func (o *Cluster) State() ClusterState {
	if o != nil && o.state != nil {
		return *o.state
	}
	return ClusterState("")
}

// GetState returns the value of the 'state' attribute and
// a flag indicating if the attribute has a value.
//
// Overall state of the cluster.
func (o *Cluster) GetState() (value ClusterState, ok bool) {
	ok = o != nil && o.state != nil
	if ok {
		value = *o.state
	}
	return
}

// Managed returns the value of the 'managed' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Flag indicating if the cluster is managed (by Red Hat) or
// self-managed by the user.
func (o *Cluster) Managed() bool {
	if o != nil && o.managed != nil {
		return *o.managed
	}
	return false
}

// GetManaged returns the value of the 'managed' attribute and
// a flag indicating if the attribute has a value.
//
// Flag indicating if the cluster is managed (by Red Hat) or
// self-managed by the user.
func (o *Cluster) GetManaged() (value bool, ok bool) {
	ok = o != nil && o.managed != nil
	if ok {
		value = *o.managed
	}
	return
}

// Memory returns the value of the 'memory' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Information about the memory of the cluster.
func (o *Cluster) Memory() *ClusterMetric {
	if o == nil {
		return nil
	}
	return o.memory
}

// GetMemory returns the value of the 'memory' attribute and
// a flag indicating if the attribute has a value.
//
// Information about the memory of the cluster.
func (o *Cluster) GetMemory() (value *ClusterMetric, ok bool) {
	ok = o != nil && o.memory != nil
	if ok {
		value = o.memory
	}
	return
}

// CPU returns the value of the 'CPU' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Information about the CPU of the cluster.
func (o *Cluster) CPU() *ClusterMetric {
	if o == nil {
		return nil
	}
	return o.cpu
}

// GetCPU returns the value of the 'CPU' attribute and
// a flag indicating if the attribute has a value.
//
// Information about the CPU of the cluster.
func (o *Cluster) GetCPU() (value *ClusterMetric, ok bool) {
	ok = o != nil && o.cpu != nil
	if ok {
		value = o.cpu
	}
	return
}

// Storage returns the value of the 'storage' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Information about the storage of the cluster.
func (o *Cluster) Storage() *ClusterMetric {
	if o == nil {
		return nil
	}
	return o.storage
}

// GetStorage returns the value of the 'storage' attribute and
// a flag indicating if the attribute has a value.
//
// Information about the storage of the cluster.
func (o *Cluster) GetStorage() (value *ClusterMetric, ok bool) {
	ok = o != nil && o.storage != nil
	if ok {
		value = o.storage
	}
	return
}

// ExternalID returns the value of the 'external_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// External identifier of the cluster, generated by the installer.
func (o *Cluster) ExternalID() string {
	if o != nil && o.externalID != nil {
		return *o.externalID
	}
	return ""
}

// GetExternalID returns the value of the 'external_ID' attribute and
// a flag indicating if the attribute has a value.
//
// External identifier of the cluster, generated by the installer.
func (o *Cluster) GetExternalID() (value string, ok bool) {
	ok = o != nil && o.externalID != nil
	if ok {
		value = *o.externalID
	}
	return
}

// AWS returns the value of the 'AWS' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Amazon Web Services settings of the cluster.
func (o *Cluster) AWS() *AWS {
	if o == nil {
		return nil
	}
	return o.aws
}

// GetAWS returns the value of the 'AWS' attribute and
// a flag indicating if the attribute has a value.
//
// Amazon Web Services settings of the cluster.
func (o *Cluster) GetAWS() (value *AWS, ok bool) {
	ok = o != nil && o.aws != nil
	if ok {
		value = o.aws
	}
	return
}

// Network returns the value of the 'network' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Network settings of the cluster.
func (o *Cluster) Network() *Network {
	if o == nil {
		return nil
	}
	return o.network
}

// GetNetwork returns the value of the 'network' attribute and
// a flag indicating if the attribute has a value.
//
// Network settings of the cluster.
func (o *Cluster) GetNetwork() (value *Network, ok bool) {
	ok = o != nil && o.network != nil
	if ok {
		value = o.network
	}
	return
}

// CreationTimestamp returns the value of the 'creation_timestamp' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Date and time when the cluster was initially created, using the
// format defined in https://www.ietf.org/rfc/rfc3339.txt[RC3339].
func (o *Cluster) CreationTimestamp() time.Time {
	if o != nil && o.creationTimestamp != nil {
		return *o.creationTimestamp
	}
	return time.Time{}
}

// GetCreationTimestamp returns the value of the 'creation_timestamp' attribute and
// a flag indicating if the attribute has a value.
//
// Date and time when the cluster was initially created, using the
// format defined in https://www.ietf.org/rfc/rfc3339.txt[RC3339].
func (o *Cluster) GetCreationTimestamp() (value time.Time, ok bool) {
	ok = o != nil && o.creationTimestamp != nil
	if ok {
		value = *o.creationTimestamp
	}
	return
}

// CloudProvider returns the value of the 'cloud_provider' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to the cloud provider where the cluster is installed.
func (o *Cluster) CloudProvider() *CloudProvider {
	if o == nil {
		return nil
	}
	return o.cloudProvider
}

// GetCloudProvider returns the value of the 'cloud_provider' attribute and
// a flag indicating if the attribute has a value.
//
// Link to the cloud provider where the cluster is installed.
func (o *Cluster) GetCloudProvider() (value *CloudProvider, ok bool) {
	ok = o != nil && o.cloudProvider != nil
	if ok {
		value = o.cloudProvider
	}
	return
}

// OpenshiftVersion returns the value of the 'openshift_version' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Version of _OpenShift_ installed in the cluster, for example `4.0.0-0.2`.
//
// When retrieving a cluster this will always be reported.
//
// When provisioning a cluster this will be ignored, as the version to
// deploy will be determined internally.
func (o *Cluster) OpenshiftVersion() string {
	if o != nil && o.openshiftVersion != nil {
		return *o.openshiftVersion
	}
	return ""
}

// GetOpenshiftVersion returns the value of the 'openshift_version' attribute and
// a flag indicating if the attribute has a value.
//
// Version of _OpenShift_ installed in the cluster, for example `4.0.0-0.2`.
//
// When retrieving a cluster this will always be reported.
//
// When provisioning a cluster this will be ignored, as the version to
// deploy will be determined internally.
func (o *Cluster) GetOpenshiftVersion() (value string, ok bool) {
	ok = o != nil && o.openshiftVersion != nil
	if ok {
		value = *o.openshiftVersion
	}
	return
}

// Subscription returns the value of the 'subscription' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to the subscription that comes from the account management service when the cluster
// is registered.
func (o *Cluster) Subscription() *Subscription {
	if o == nil {
		return nil
	}
	return o.subscription
}

// GetSubscription returns the value of the 'subscription' attribute and
// a flag indicating if the attribute has a value.
//
// Link to the subscription that comes from the account management service when the cluster
// is registered.
func (o *Cluster) GetSubscription() (value *Subscription, ok bool) {
	ok = o != nil && o.subscription != nil
	if ok {
		value = o.subscription
	}
	return
}

// Groups returns the value of the 'groups' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to the collection of groups of user of the cluster.
func (o *Cluster) Groups() *GroupList {
	if o == nil {
		return nil
	}
	return o.groups
}

// GetGroups returns the value of the 'groups' attribute and
// a flag indicating if the attribute has a value.
//
// Link to the collection of groups of user of the cluster.
func (o *Cluster) GetGroups() (value *GroupList, ok bool) {
	ok = o != nil && o.groups != nil
	if ok {
		value = o.groups
	}
	return
}

// Creator returns the value of the 'creator' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Mail address of the user that created the cluster.
func (o *Cluster) Creator() string {
	if o != nil && o.creator != nil {
		return *o.creator
	}
	return ""
}

// GetCreator returns the value of the 'creator' attribute and
// a flag indicating if the attribute has a value.
//
// Mail address of the user that created the cluster.
func (o *Cluster) GetCreator() (value string, ok bool) {
	ok = o != nil && o.creator != nil
	if ok {
		value = *o.creator
	}
	return
}

// Version returns the value of the 'version' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to the version of _OpenShift_ that will be used to install the cluster.
func (o *Cluster) Version() *Version {
	if o == nil {
		return nil
	}
	return o.version
}

// GetVersion returns the value of the 'version' attribute and
// a flag indicating if the attribute has a value.
//
// Link to the version of _OpenShift_ that will be used to install the cluster.
func (o *Cluster) GetVersion() (value *Version, ok bool) {
	ok = o != nil && o.version != nil
	if ok {
		value = o.version
	}
	return
}

// IdentityProviders returns the value of the 'identity_providers' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to the collection of identity providers of the cluster.
func (o *Cluster) IdentityProviders() *IdentityProviderList {
	if o == nil {
		return nil
	}
	return o.identityProviders
}

// GetIdentityProviders returns the value of the 'identity_providers' attribute and
// a flag indicating if the attribute has a value.
//
// Link to the collection of identity providers of the cluster.
func (o *Cluster) GetIdentityProviders() (value *IdentityProviderList, ok bool) {
	ok = o != nil && o.identityProviders != nil
	if ok {
		value = o.identityProviders
	}
	return
}

// ClusterListKind is the name of the type used to represent list of
// objects of type 'cluster'.
const ClusterListKind = "ClusterList"

// ClusterListLinkKind is the name of the type used to represent links
// to list of objects of type 'cluster'.
const ClusterListLinkKind = "ClusterListLink"

// ClusterNilKind is the name of the type used to nil lists of
// objects of type 'cluster'.
const ClusterListNilKind = "ClusterListNil"

// ClusterList is a list of values of the 'cluster' type.
type ClusterList struct {
	href  *string
	link  bool
	items []*Cluster
}

// Kind returns the name of the type of the object.
func (l *ClusterList) Kind() string {
	if l == nil {
		return ClusterListNilKind
	}
	if l.link {
		return ClusterListLinkKind
	}
	return ClusterListKind
}

// Link returns true iif this is a link.
func (l *ClusterList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *ClusterList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *ClusterList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *ClusterList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterList) Slice() []*Cluster {
	var slice []*Cluster
	if l == nil {
		slice = make([]*Cluster, 0)
	} else {
		slice = make([]*Cluster, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterList) Each(f func(item *Cluster) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterList) Range(f func(index int, item *Cluster) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
