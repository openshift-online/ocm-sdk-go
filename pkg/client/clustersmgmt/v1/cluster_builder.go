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

// ClusterBuilder contains the data and logic needed to build 'cluster' objects.
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
type ClusterBuilder struct {
	id                *string
	href              *string
	link              bool
	name              *string
	flavour           *FlavourBuilder
	osVersion         *string
	console           *ClusterConsoleBuilder
	runtimeVersion    *string
	multiAZ           *bool
	nodes             *ClusterNodesBuilder
	api               *ClusterAPIBuilder
	region            *CloudRegionBuilder
	displayName       *string
	dns               *DNSBuilder
	properties        map[string]string
	state             *ClusterState
	managed           *bool
	memory            *ClusterMetricBuilder
	cpu               *ClusterMetricBuilder
	storage           *ClusterMetricBuilder
	externalID        *string
	aws               *AWSBuilder
	network           *NetworkBuilder
	creationTimestamp *time.Time
	cloudProvider     *CloudProviderBuilder
	openshiftVersion  *string
}

// NewCluster creates a new builder of 'cluster' objects.
func NewCluster() *ClusterBuilder {
	return new(ClusterBuilder)
}

// ID sets the identifier of the object.
func (b *ClusterBuilder) ID(value string) *ClusterBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *ClusterBuilder) HREF(value string) *ClusterBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *ClusterBuilder) Link(value bool) *ClusterBuilder {
	b.link = value
	return b
}

// Name sets the value of the 'name' attribute
// to the given value.
//
//
func (b *ClusterBuilder) Name(value string) *ClusterBuilder {
	b.name = &value
	return b
}

// Flavour sets the value of the 'flavour' attribute
// to the given value.
//
// Set of predefined properties of a cluster. For example, a _huge_ flavour can be a cluster
// with 10 infra nodes and 1000 compute nodes.
func (b *ClusterBuilder) Flavour(value *FlavourBuilder) *ClusterBuilder {
	b.flavour = value
	return b
}

// OsVersion sets the value of the 'os_version' attribute
// to the given value.
//
//
func (b *ClusterBuilder) OsVersion(value string) *ClusterBuilder {
	b.osVersion = &value
	return b
}

// Console sets the value of the 'console' attribute
// to the given value.
//
// Information about the console of a cluster.
func (b *ClusterBuilder) Console(value *ClusterConsoleBuilder) *ClusterBuilder {
	b.console = value
	return b
}

// RuntimeVersion sets the value of the 'runtime_version' attribute
// to the given value.
//
//
func (b *ClusterBuilder) RuntimeVersion(value string) *ClusterBuilder {
	b.runtimeVersion = &value
	return b
}

// MultiAZ sets the value of the 'multi_AZ' attribute
// to the given value.
//
//
func (b *ClusterBuilder) MultiAZ(value bool) *ClusterBuilder {
	b.multiAZ = &value
	return b
}

// Nodes sets the value of the 'nodes' attribute
// to the given value.
//
// Counts of different classes of nodes inside a cluster.
func (b *ClusterBuilder) Nodes(value *ClusterNodesBuilder) *ClusterBuilder {
	b.nodes = value
	return b
}

// API sets the value of the 'API' attribute
// to the given value.
//
// Information about the API of a cluster.
func (b *ClusterBuilder) API(value *ClusterAPIBuilder) *ClusterBuilder {
	b.api = value
	return b
}

// Region sets the value of the 'region' attribute
// to the given value.
//
// Description of a region of a cloud provider.
func (b *ClusterBuilder) Region(value *CloudRegionBuilder) *ClusterBuilder {
	b.region = value
	return b
}

// DisplayName sets the value of the 'display_name' attribute
// to the given value.
//
//
func (b *ClusterBuilder) DisplayName(value string) *ClusterBuilder {
	b.displayName = &value
	return b
}

// DNS sets the value of the 'DNS' attribute
// to the given value.
//
// DNS settings of the cluster.
func (b *ClusterBuilder) DNS(value *DNSBuilder) *ClusterBuilder {
	b.dns = value
	return b
}

// Properties sets the value of the 'properties' attribute
// to the given value.
//
//
func (b *ClusterBuilder) Properties(value map[string]string) *ClusterBuilder {
	b.properties = value
	return b
}

// State sets the value of the 'state' attribute
// to the given value.
//
// Overall state of a cluster.
func (b *ClusterBuilder) State(value ClusterState) *ClusterBuilder {
	b.state = &value
	return b
}

// Managed sets the value of the 'managed' attribute
// to the given value.
//
//
func (b *ClusterBuilder) Managed(value bool) *ClusterBuilder {
	b.managed = &value
	return b
}

// Memory sets the value of the 'memory' attribute
// to the given value.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
func (b *ClusterBuilder) Memory(value *ClusterMetricBuilder) *ClusterBuilder {
	b.memory = value
	return b
}

// CPU sets the value of the 'CPU' attribute
// to the given value.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
func (b *ClusterBuilder) CPU(value *ClusterMetricBuilder) *ClusterBuilder {
	b.cpu = value
	return b
}

// Storage sets the value of the 'storage' attribute
// to the given value.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
func (b *ClusterBuilder) Storage(value *ClusterMetricBuilder) *ClusterBuilder {
	b.storage = value
	return b
}

// ExternalID sets the value of the 'external_ID' attribute
// to the given value.
//
//
func (b *ClusterBuilder) ExternalID(value string) *ClusterBuilder {
	b.externalID = &value
	return b
}

// AWS sets the value of the 'AWS' attribute
// to the given value.
//
// _Amazon Web Services_ specific settings of a cluster.
func (b *ClusterBuilder) AWS(value *AWSBuilder) *ClusterBuilder {
	b.aws = value
	return b
}

// Network sets the value of the 'network' attribute
// to the given value.
//
// Network configuration of a cluster.
func (b *ClusterBuilder) Network(value *NetworkBuilder) *ClusterBuilder {
	b.network = value
	return b
}

// CreationTimestamp sets the value of the 'creation_timestamp' attribute
// to the given value.
//
//
func (b *ClusterBuilder) CreationTimestamp(value time.Time) *ClusterBuilder {
	b.creationTimestamp = &value
	return b
}

// CloudProvider sets the value of the 'cloud_provider' attribute
// to the given value.
//
// Cloud provider.
func (b *ClusterBuilder) CloudProvider(value *CloudProviderBuilder) *ClusterBuilder {
	b.cloudProvider = value
	return b
}

// OpenshiftVersion sets the value of the 'openshift_version' attribute
// to the given value.
//
//
func (b *ClusterBuilder) OpenshiftVersion(value string) *ClusterBuilder {
	b.openshiftVersion = &value
	return b
}

// Build creates a 'cluster' object using the configuration stored in the builder.
func (b *ClusterBuilder) Build() (object *Cluster, err error) {
	object = new(Cluster)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.name != nil {
		object.name = b.name
	}
	if b.flavour != nil {
		object.flavour, err = b.flavour.Build()
		if err != nil {
			return
		}
	}
	if b.osVersion != nil {
		object.osVersion = b.osVersion
	}
	if b.console != nil {
		object.console, err = b.console.Build()
		if err != nil {
			return
		}
	}
	if b.runtimeVersion != nil {
		object.runtimeVersion = b.runtimeVersion
	}
	if b.multiAZ != nil {
		object.multiAZ = b.multiAZ
	}
	if b.nodes != nil {
		object.nodes, err = b.nodes.Build()
		if err != nil {
			return
		}
	}
	if b.api != nil {
		object.api, err = b.api.Build()
		if err != nil {
			return
		}
	}
	if b.region != nil {
		object.region, err = b.region.Build()
		if err != nil {
			return
		}
	}
	if b.displayName != nil {
		object.displayName = b.displayName
	}
	if b.dns != nil {
		object.dns, err = b.dns.Build()
		if err != nil {
			return
		}
	}
	if b.properties != nil {
		object.properties = b.properties
	}
	if b.state != nil {
		object.state = b.state
	}
	if b.managed != nil {
		object.managed = b.managed
	}
	if b.memory != nil {
		object.memory, err = b.memory.Build()
		if err != nil {
			return
		}
	}
	if b.cpu != nil {
		object.cpu, err = b.cpu.Build()
		if err != nil {
			return
		}
	}
	if b.storage != nil {
		object.storage, err = b.storage.Build()
		if err != nil {
			return
		}
	}
	if b.externalID != nil {
		object.externalID = b.externalID
	}
	if b.aws != nil {
		object.aws, err = b.aws.Build()
		if err != nil {
			return
		}
	}
	if b.network != nil {
		object.network, err = b.network.Build()
		if err != nil {
			return
		}
	}
	if b.creationTimestamp != nil {
		object.creationTimestamp = b.creationTimestamp
	}
	if b.cloudProvider != nil {
		object.cloudProvider, err = b.cloudProvider.Build()
		if err != nil {
			return
		}
	}
	if b.openshiftVersion != nil {
		object.openshiftVersion = b.openshiftVersion
	}
	return
}
