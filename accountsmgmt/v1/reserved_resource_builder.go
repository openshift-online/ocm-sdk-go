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

// ReservedResourceBuilder contains the data and logic needed to build 'reserved_resource' objects.
//
//
type ReservedResourceBuilder struct {
	resourceName         *string
	resourceType         *string
	byoc                 *bool
	availabilityZoneType *string
	count                *int
}

// NewReservedResource creates a new builder of 'reserved_resource' objects.
func NewReservedResource() *ReservedResourceBuilder {
	return new(ReservedResourceBuilder)
}

// ResourceName sets the value of the 'resource_name' attribute
// to the given value.
//
//
func (b *ReservedResourceBuilder) ResourceName(value string) *ReservedResourceBuilder {
	b.resourceName = &value
	return b
}

// ResourceType sets the value of the 'resource_type' attribute
// to the given value.
//
//
func (b *ReservedResourceBuilder) ResourceType(value string) *ReservedResourceBuilder {
	b.resourceType = &value
	return b
}

// BYOC sets the value of the 'BYOC' attribute
// to the given value.
//
//
func (b *ReservedResourceBuilder) BYOC(value bool) *ReservedResourceBuilder {
	b.byoc = &value
	return b
}

// AvailabilityZoneType sets the value of the 'availability_zone_type' attribute
// to the given value.
//
//
func (b *ReservedResourceBuilder) AvailabilityZoneType(value string) *ReservedResourceBuilder {
	b.availabilityZoneType = &value
	return b
}

// Count sets the value of the 'count' attribute
// to the given value.
//
//
func (b *ReservedResourceBuilder) Count(value int) *ReservedResourceBuilder {
	b.count = &value
	return b
}

// Build creates a 'reserved_resource' object using the configuration stored in the builder.
func (b *ReservedResourceBuilder) Build() (object *ReservedResource, err error) {
	object = new(ReservedResource)
	if b.resourceName != nil {
		object.resourceName = b.resourceName
	}
	if b.resourceType != nil {
		object.resourceType = b.resourceType
	}
	if b.byoc != nil {
		object.byoc = b.byoc
	}
	if b.availabilityZoneType != nil {
		object.availabilityZoneType = b.availabilityZoneType
	}
	if b.count != nil {
		object.count = b.count
	}
	return
}
