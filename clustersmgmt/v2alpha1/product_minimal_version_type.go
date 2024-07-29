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

import (
	time "time"
)

// ProductMinimalVersionKind is the name of the type used to represent objects
// of type 'product_minimal_version'.
const ProductMinimalVersionKind = "ProductMinimalVersion"

// ProductMinimalVersionLinkKind is the name of the type used to represent links
// to objects of type 'product_minimal_version'.
const ProductMinimalVersionLinkKind = "ProductMinimalVersionLink"

// ProductMinimalVersionNilKind is the name of the type used to nil references
// to objects of type 'product_minimal_version'.
const ProductMinimalVersionNilKind = "ProductMinimalVersionNil"

// ProductMinimalVersion represents the values of the 'product_minimal_version' type.
//
// Representation of a product minimal version.
type ProductMinimalVersion struct {
	bitmap_   uint32
	id        string
	href      string
	rosaCli   string
	startDate time.Time
}

// Kind returns the name of the type of the object.
func (o *ProductMinimalVersion) Kind() string {
	if o == nil {
		return ProductMinimalVersionNilKind
	}
	if o.bitmap_&1 != 0 {
		return ProductMinimalVersionLinkKind
	}
	return ProductMinimalVersionKind
}

// Link returns true iif this is a link.
func (o *ProductMinimalVersion) Link() bool {
	return o != nil && o.bitmap_&1 != 0
}

// ID returns the identifier of the object.
func (o *ProductMinimalVersion) ID() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *ProductMinimalVersion) GetID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *ProductMinimalVersion) HREF() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *ProductMinimalVersion) GetHREF() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ProductMinimalVersion) Empty() bool {
	return o == nil || o.bitmap_&^1 == 0
}

// RosaCli returns the value of the 'rosa_cli' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The ROSA CLI minimal version.
func (o *ProductMinimalVersion) RosaCli() string {
	if o != nil && o.bitmap_&8 != 0 {
		return o.rosaCli
	}
	return ""
}

// GetRosaCli returns the value of the 'rosa_cli' attribute and
// a flag indicating if the attribute has a value.
//
// The ROSA CLI minimal version.
func (o *ProductMinimalVersion) GetRosaCli() (value string, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.rosaCli
	}
	return
}

// StartDate returns the value of the 'start_date' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The start date for this minimal version.
func (o *ProductMinimalVersion) StartDate() time.Time {
	if o != nil && o.bitmap_&16 != 0 {
		return o.startDate
	}
	return time.Time{}
}

// GetStartDate returns the value of the 'start_date' attribute and
// a flag indicating if the attribute has a value.
//
// The start date for this minimal version.
func (o *ProductMinimalVersion) GetStartDate() (value time.Time, ok bool) {
	ok = o != nil && o.bitmap_&16 != 0
	if ok {
		value = o.startDate
	}
	return
}

// ProductMinimalVersionListKind is the name of the type used to represent list of objects of
// type 'product_minimal_version'.
const ProductMinimalVersionListKind = "ProductMinimalVersionList"

// ProductMinimalVersionListLinkKind is the name of the type used to represent links to list
// of objects of type 'product_minimal_version'.
const ProductMinimalVersionListLinkKind = "ProductMinimalVersionListLink"

// ProductMinimalVersionNilKind is the name of the type used to nil lists of objects of
// type 'product_minimal_version'.
const ProductMinimalVersionListNilKind = "ProductMinimalVersionListNil"

// ProductMinimalVersionList is a list of values of the 'product_minimal_version' type.
type ProductMinimalVersionList struct {
	href  string
	link  bool
	items []*ProductMinimalVersion
}

// Kind returns the name of the type of the object.
func (l *ProductMinimalVersionList) Kind() string {
	if l == nil {
		return ProductMinimalVersionListNilKind
	}
	if l.link {
		return ProductMinimalVersionListLinkKind
	}
	return ProductMinimalVersionListKind
}

// Link returns true iif this is a link.
func (l *ProductMinimalVersionList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *ProductMinimalVersionList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *ProductMinimalVersionList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *ProductMinimalVersionList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ProductMinimalVersionList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ProductMinimalVersionList) Get(i int) *ProductMinimalVersion {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ProductMinimalVersionList) Slice() []*ProductMinimalVersion {
	var slice []*ProductMinimalVersion
	if l == nil {
		slice = make([]*ProductMinimalVersion, 0)
	} else {
		slice = make([]*ProductMinimalVersion, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ProductMinimalVersionList) Each(f func(item *ProductMinimalVersion) bool) {
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
func (l *ProductMinimalVersionList) Range(f func(index int, item *ProductMinimalVersion) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
