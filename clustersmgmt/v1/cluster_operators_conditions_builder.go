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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterOperatorsConditionsBuilder contains the data and logic needed to build 'cluster_operators_conditions' objects.
//
// Counts of operators conditions inside a cluster.
type ClusterOperatorsConditionsBuilder struct {
	available *int
	degraded  *int
	failing   *int
}

// NewClusterOperatorsConditions creates a new builder of 'cluster_operators_conditions' objects.
func NewClusterOperatorsConditions() *ClusterOperatorsConditionsBuilder {
	return new(ClusterOperatorsConditionsBuilder)
}

// Available sets the value of the 'available' attribute to the given value.
//
//
func (b *ClusterOperatorsConditionsBuilder) Available(value int) *ClusterOperatorsConditionsBuilder {
	b.available = &value
	return b
}

// Degraded sets the value of the 'degraded' attribute to the given value.
//
//
func (b *ClusterOperatorsConditionsBuilder) Degraded(value int) *ClusterOperatorsConditionsBuilder {
	b.degraded = &value
	return b
}

// Failing sets the value of the 'failing' attribute to the given value.
//
//
func (b *ClusterOperatorsConditionsBuilder) Failing(value int) *ClusterOperatorsConditionsBuilder {
	b.failing = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterOperatorsConditionsBuilder) Copy(object *ClusterOperatorsConditions) *ClusterOperatorsConditionsBuilder {
	if object == nil {
		return b
	}
	b.available = object.available
	b.degraded = object.degraded
	b.failing = object.failing
	return b
}

// Build creates a 'cluster_operators_conditions' object using the configuration stored in the builder.
func (b *ClusterOperatorsConditionsBuilder) Build() (object *ClusterOperatorsConditions, err error) {
	object = new(ClusterOperatorsConditions)
	object.available = b.available
	object.degraded = b.degraded
	object.failing = b.failing
	return
}
