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

// ClusterAlertsFiringBuilder contains the data and logic needed to build 'cluster_alerts_firing' objects.
//
// Counts of different alerts firing inside a cluster.
type ClusterAlertsFiringBuilder struct {
	critical *int
	high     *int
	none     *int
	warning  *int
}

// NewClusterAlertsFiring creates a new builder of 'cluster_alerts_firing' objects.
func NewClusterAlertsFiring() *ClusterAlertsFiringBuilder {
	return new(ClusterAlertsFiringBuilder)
}

// Critical sets the value of the 'critical' attribute to the given value.
//
//
func (b *ClusterAlertsFiringBuilder) Critical(value int) *ClusterAlertsFiringBuilder {
	b.critical = &value
	return b
}

// High sets the value of the 'high' attribute to the given value.
//
//
func (b *ClusterAlertsFiringBuilder) High(value int) *ClusterAlertsFiringBuilder {
	b.high = &value
	return b
}

// None sets the value of the 'none' attribute to the given value.
//
//
func (b *ClusterAlertsFiringBuilder) None(value int) *ClusterAlertsFiringBuilder {
	b.none = &value
	return b
}

// Warning sets the value of the 'warning' attribute to the given value.
//
//
func (b *ClusterAlertsFiringBuilder) Warning(value int) *ClusterAlertsFiringBuilder {
	b.warning = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterAlertsFiringBuilder) Copy(object *ClusterAlertsFiring) *ClusterAlertsFiringBuilder {
	if object == nil {
		return b
	}
	b.critical = object.critical
	b.high = object.high
	b.none = object.none
	b.warning = object.warning
	return b
}

// Build creates a 'cluster_alerts_firing' object using the configuration stored in the builder.
func (b *ClusterAlertsFiringBuilder) Build() (object *ClusterAlertsFiring, err error) {
	object = new(ClusterAlertsFiring)
	object.critical = b.critical
	object.high = b.high
	object.none = b.none
	object.warning = b.warning
	return
}
