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

// WifTemplateKind is the name of the type used to represent objects
// of type 'wif_template'.
const WifTemplateKind = "WifTemplate"

// WifTemplateLinkKind is the name of the type used to represent links
// to objects of type 'wif_template'.
const WifTemplateLinkKind = "WifTemplateLink"

// WifTemplateNilKind is the name of the type used to nil references
// to objects of type 'wif_template'.
const WifTemplateNilKind = "WifTemplateNil"

// WifTemplate represents the values of the 'wif_template' type.
//
// Definition of an wif_template resource.
type WifTemplate struct {
	bitmap_         uint32
	id              string
	href            string
	serviceAccounts []*WifServiceAccount
}

// Kind returns the name of the type of the object.
func (o *WifTemplate) Kind() string {
	if o == nil {
		return WifTemplateNilKind
	}
	if o.bitmap_&1 != 0 {
		return WifTemplateLinkKind
	}
	return WifTemplateKind
}

// Link returns true iif this is a link.
func (o *WifTemplate) Link() bool {
	return o != nil && o.bitmap_&1 != 0
}

// ID returns the identifier of the object.
func (o *WifTemplate) ID() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *WifTemplate) GetID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *WifTemplate) HREF() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *WifTemplate) GetHREF() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *WifTemplate) Empty() bool {
	return o == nil || o.bitmap_&^1 == 0
}

// ServiceAccounts returns the value of the 'service_accounts' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The list of service accounts and their associated roles that this template
// would require to be configured on the user's GCP project.
func (o *WifTemplate) ServiceAccounts() []*WifServiceAccount {
	if o != nil && o.bitmap_&8 != 0 {
		return o.serviceAccounts
	}
	return nil
}

// GetServiceAccounts returns the value of the 'service_accounts' attribute and
// a flag indicating if the attribute has a value.
//
// The list of service accounts and their associated roles that this template
// would require to be configured on the user's GCP project.
func (o *WifTemplate) GetServiceAccounts() (value []*WifServiceAccount, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.serviceAccounts
	}
	return
}

// WifTemplateListKind is the name of the type used to represent list of objects of
// type 'wif_template'.
const WifTemplateListKind = "WifTemplateList"

// WifTemplateListLinkKind is the name of the type used to represent links to list
// of objects of type 'wif_template'.
const WifTemplateListLinkKind = "WifTemplateListLink"

// WifTemplateNilKind is the name of the type used to nil lists of objects of
// type 'wif_template'.
const WifTemplateListNilKind = "WifTemplateListNil"

// WifTemplateList is a list of values of the 'wif_template' type.
type WifTemplateList struct {
	href  string
	link  bool
	items []*WifTemplate
}

// Kind returns the name of the type of the object.
func (l *WifTemplateList) Kind() string {
	if l == nil {
		return WifTemplateListNilKind
	}
	if l.link {
		return WifTemplateListLinkKind
	}
	return WifTemplateListKind
}

// Link returns true iif this is a link.
func (l *WifTemplateList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *WifTemplateList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *WifTemplateList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *WifTemplateList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *WifTemplateList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *WifTemplateList) Get(i int) *WifTemplate {
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
func (l *WifTemplateList) Slice() []*WifTemplate {
	var slice []*WifTemplate
	if l == nil {
		slice = make([]*WifTemplate, 0)
	} else {
		slice = make([]*WifTemplate, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *WifTemplateList) Each(f func(item *WifTemplate) bool) {
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
func (l *WifTemplateList) Range(f func(index int, item *WifTemplate) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
