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

package v2alpha1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v2alpha1

// WifTemplateBuilder contains the data and logic needed to build 'wif_template' objects.
//
// Definition of an wif_template resource.
type WifTemplateBuilder struct {
	bitmap_         uint32
	id              string
	href            string
	serviceAccounts []*WifServiceAccountBuilder
}

// NewWifTemplate creates a new builder of 'wif_template' objects.
func NewWifTemplate() *WifTemplateBuilder {
	return &WifTemplateBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *WifTemplateBuilder) Link(value bool) *WifTemplateBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *WifTemplateBuilder) ID(value string) *WifTemplateBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *WifTemplateBuilder) HREF(value string) *WifTemplateBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *WifTemplateBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// ServiceAccounts sets the value of the 'service_accounts' attribute to the given values.
func (b *WifTemplateBuilder) ServiceAccounts(values ...*WifServiceAccountBuilder) *WifTemplateBuilder {
	b.serviceAccounts = make([]*WifServiceAccountBuilder, len(values))
	copy(b.serviceAccounts, values)
	b.bitmap_ |= 8
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *WifTemplateBuilder) Copy(object *WifTemplate) *WifTemplateBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	if object.serviceAccounts != nil {
		b.serviceAccounts = make([]*WifServiceAccountBuilder, len(object.serviceAccounts))
		for i, v := range object.serviceAccounts {
			b.serviceAccounts[i] = NewWifServiceAccount().Copy(v)
		}
	} else {
		b.serviceAccounts = nil
	}
	return b
}

// Build creates a 'wif_template' object using the configuration stored in the builder.
func (b *WifTemplateBuilder) Build() (object *WifTemplate, err error) {
	object = new(WifTemplate)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	if b.serviceAccounts != nil {
		object.serviceAccounts = make([]*WifServiceAccount, len(b.serviceAccounts))
		for i, v := range b.serviceAccounts {
			object.serviceAccounts[i], err = v.Build()
			if err != nil {
				return
			}
		}
	}
	return
}
