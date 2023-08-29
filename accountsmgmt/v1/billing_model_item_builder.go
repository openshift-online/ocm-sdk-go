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

// BillingModelItemBuilder contains the data and logic needed to build 'billing_model_item' objects.
//
// BillingModelItem that represents a billing model (defined in pkg/api/billing_types.go). Using BillingModelItem to keep backwards compatibility as we already have a BillingModel enum defined.
type BillingModelItemBuilder struct {
	bitmap_     uint32
	href        string
	id          string
	marketplace string
	model       string
}

// NewBillingModelItem creates a new builder of 'billing_model_item' objects.
func NewBillingModelItem() *BillingModelItemBuilder {
	return &BillingModelItemBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *BillingModelItemBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// HREF sets the value of the 'HREF' attribute to the given value.
func (b *BillingModelItemBuilder) HREF(value string) *BillingModelItemBuilder {
	b.href = value
	b.bitmap_ |= 1
	return b
}

// Id sets the value of the 'id' attribute to the given value.
func (b *BillingModelItemBuilder) Id(value string) *BillingModelItemBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// Marketplace sets the value of the 'marketplace' attribute to the given value.
func (b *BillingModelItemBuilder) Marketplace(value string) *BillingModelItemBuilder {
	b.marketplace = value
	b.bitmap_ |= 4
	return b
}

// Model sets the value of the 'model' attribute to the given value.
func (b *BillingModelItemBuilder) Model(value string) *BillingModelItemBuilder {
	b.model = value
	b.bitmap_ |= 8
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *BillingModelItemBuilder) Copy(object *BillingModelItem) *BillingModelItemBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.href = object.href
	b.id = object.id
	b.marketplace = object.marketplace
	b.model = object.model
	return b
}

// Build creates a 'billing_model_item' object using the configuration stored in the builder.
func (b *BillingModelItemBuilder) Build() (object *BillingModelItem, err error) {
	object = new(BillingModelItem)
	object.bitmap_ = b.bitmap_
	object.href = b.href
	object.id = b.id
	object.marketplace = b.marketplace
	object.model = b.model
	return
}
