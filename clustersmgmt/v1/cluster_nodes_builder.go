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

package v1 // github.com/openshift-online/uhc-sdk-go/clustersmgmt/v1

// ClusterNodesBuilder contains the data and logic needed to build 'cluster_nodes' objects.
//
// Counts of different classes of nodes inside a cluster.
type ClusterNodesBuilder struct {
	total   *int
	master  *int
	infra   *int
	compute *int
}

// NewClusterNodes creates a new builder of 'cluster_nodes' objects.
func NewClusterNodes() *ClusterNodesBuilder {
	return new(ClusterNodesBuilder)
}

// Total sets the value of the 'total' attribute
// to the given value.
//
//
func (b *ClusterNodesBuilder) Total(value int) *ClusterNodesBuilder {
	b.total = &value
	return b
}

// Master sets the value of the 'master' attribute
// to the given value.
//
//
func (b *ClusterNodesBuilder) Master(value int) *ClusterNodesBuilder {
	b.master = &value
	return b
}

// Infra sets the value of the 'infra' attribute
// to the given value.
//
//
func (b *ClusterNodesBuilder) Infra(value int) *ClusterNodesBuilder {
	b.infra = &value
	return b
}

// Compute sets the value of the 'compute' attribute
// to the given value.
//
//
func (b *ClusterNodesBuilder) Compute(value int) *ClusterNodesBuilder {
	b.compute = &value
	return b
}

// Build creates a 'cluster_nodes' object using the configuration stored in the builder.
func (b *ClusterNodesBuilder) Build() (object *ClusterNodes, err error) {
	object = new(ClusterNodes)
	if b.total != nil {
		object.total = b.total
	}
	if b.master != nil {
		object.master = b.master
	}
	if b.infra != nil {
		object.infra = b.infra
	}
	if b.compute != nil {
		object.compute = b.compute
	}
	return
}
