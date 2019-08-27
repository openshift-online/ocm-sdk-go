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

// CloudProvider represents the values of the 'cloud_provider' type.
//
// Cloud provider.
type CloudProvider struct {
	name        *string
	displayName *string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *CloudProvider) Empty() bool {
	return o == nil || (o.name == nil &&
		o.displayName == nil &&
		true)
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Human friendly identifier of the cloud provider, for example `aws`.
func (o *CloudProvider) Name() string {
	if o != nil && o.name != nil {
		return *o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
// Human friendly identifier of the cloud provider, for example `aws`.
func (o *CloudProvider) GetName() (value string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = *o.name
	}
	return
}

// DisplayName returns the value of the 'display_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Name of the cloud provider for display purposes. It can contain any characters,
// including spaces.
func (o *CloudProvider) DisplayName() string {
	if o != nil && o.displayName != nil {
		return *o.displayName
	}
	return ""
}

// GetDisplayName returns the value of the 'display_name' attribute and
// a flag indicating if the attribute has a value.
//
// Name of the cloud provider for display purposes. It can contain any characters,
// including spaces.
func (o *CloudProvider) GetDisplayName() (value string, ok bool) {
	ok = o != nil && o.displayName != nil
	if ok {
		value = *o.displayName
	}
	return
}

// CloudProviderList is a list of values of the 'cloud_provider' type.
type CloudProviderList struct {
	items []*CloudProvider
}

// Len returns the length of the list.
func (l *CloudProviderList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *CloudProviderList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *CloudProviderList) Get(i int) *CloudProvider {
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
func (l *CloudProviderList) Slice() []*CloudProvider {
	var slice []*CloudProvider
	if l == nil {
		slice = make([]*CloudProvider, 0)
	} else {
		slice = make([]*CloudProvider, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *CloudProviderList) Each(f func(item *CloudProvider) bool) {
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
func (l *CloudProviderList) Range(f func(index int, item *CloudProvider) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
