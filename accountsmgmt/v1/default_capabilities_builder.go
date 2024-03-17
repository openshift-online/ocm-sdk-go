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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// DefaultCapabilitiesBuilder contains the data and logic needed to build 'default_capabilities' objects.
//
// Default Capability model that represents default internal labels with a key-value pair.
type DefaultCapabilitiesBuilder struct {
	bitmap_ uint32
	name    string
	value   string
}

// NewDefaultCapabilities creates a new builder of 'default_capabilities' objects.
func NewDefaultCapabilities() *DefaultCapabilitiesBuilder {
	return &DefaultCapabilitiesBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *DefaultCapabilitiesBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Name sets the value of the 'name' attribute to the given value.
func (b *DefaultCapabilitiesBuilder) Name(value string) *DefaultCapabilitiesBuilder {
	b.name = value
	b.bitmap_ |= 1
	return b
}

// Value sets the value of the 'value' attribute to the given value.
func (b *DefaultCapabilitiesBuilder) Value(value string) *DefaultCapabilitiesBuilder {
	b.value = value
	b.bitmap_ |= 2
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *DefaultCapabilitiesBuilder) Copy(object *DefaultCapabilities) *DefaultCapabilitiesBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.name = object.name
	b.value = object.value
	return b
}

// Build creates a 'default_capabilities' object using the configuration stored in the builder.
func (b *DefaultCapabilitiesBuilder) Build() (object *DefaultCapabilities, err error) {
	object = new(DefaultCapabilities)
	object.bitmap_ = b.bitmap_
	object.name = b.name
	object.value = b.value
	return
}
