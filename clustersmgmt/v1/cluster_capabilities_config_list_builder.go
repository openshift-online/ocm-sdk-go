/*
Copyright (c) 2020 Red Hat, Inc.

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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterCapabilitiesConfigListBuilder contains the data and logic needed to build
// 'cluster_capabilities_config' objects.
type ClusterCapabilitiesConfigListBuilder struct {
	items []*ClusterCapabilitiesConfigBuilder
}

// NewClusterCapabilitiesConfigList creates a new builder of 'cluster_capabilities_config' objects.
func NewClusterCapabilitiesConfigList() *ClusterCapabilitiesConfigListBuilder {
	return new(ClusterCapabilitiesConfigListBuilder)
}

// Items sets the items of the list.
func (b *ClusterCapabilitiesConfigListBuilder) Items(values ...*ClusterCapabilitiesConfigBuilder) *ClusterCapabilitiesConfigListBuilder {
	b.items = make([]*ClusterCapabilitiesConfigBuilder, len(values))
	copy(b.items, values)
	return b
}

// Empty returns true if the list is empty.
func (b *ClusterCapabilitiesConfigListBuilder) Empty() bool {
	return b == nil || len(b.items) == 0
}

// Copy copies the items of the given list into this builder, discarding any previous items.
func (b *ClusterCapabilitiesConfigListBuilder) Copy(list *ClusterCapabilitiesConfigList) *ClusterCapabilitiesConfigListBuilder {
	if list == nil || list.items == nil {
		b.items = nil
	} else {
		b.items = make([]*ClusterCapabilitiesConfigBuilder, len(list.items))
		for i, v := range list.items {
			b.items[i] = NewClusterCapabilitiesConfig().Copy(v)
		}
	}
	return b
}

// Build creates a list of 'cluster_capabilities_config' objects using the
// configuration stored in the builder.
func (b *ClusterCapabilitiesConfigListBuilder) Build() (list *ClusterCapabilitiesConfigList, err error) {
	items := make([]*ClusterCapabilitiesConfig, len(b.items))
	for i, item := range b.items {
		items[i], err = item.Build()
		if err != nil {
			return
		}
	}
	list = new(ClusterCapabilitiesConfigList)
	list.items = items
	return
}
