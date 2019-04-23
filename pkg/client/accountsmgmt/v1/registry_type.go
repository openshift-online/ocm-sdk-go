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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

// RegistryKind is the name of the type used to represent objects
// of type 'registry'.
const RegistryKind = "Registry"

// RegistryLinkKind is the name of the type used to represent links
// to objects of type 'registry'.
const RegistryLinkKind = "RegistryLink"

// RegistryNilKind is the name of the type used to nil references
// to objects of type 'registry'.
const RegistryNilKind = "RegistryNil"

// Registry represents the values of the 'registry' type.
//
//
type Registry struct {
	id         *string
	href       *string
	link       bool
	name       *string
	url        *string
	teamName   *string
	orgName    *string
	type_      *string
	cloudAlias *bool
}

// Kind returns the name of the type of the object.
func (o *Registry) Kind() string {
	if o == nil {
		return RegistryNilKind
	}
	if o.link {
		return RegistryLinkKind
	}
	return RegistryKind
}

// ID returns the identifier of the object.
func (o *Registry) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Registry) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Registry) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Registry) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Registry) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Registry) Name() string {
	if o != nil && o.name != nil {
		return *o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Registry) GetName() (value string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = *o.name
	}
	return
}

// URL returns the value of the 'URL' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Registry) URL() string {
	if o != nil && o.url != nil {
		return *o.url
	}
	return ""
}

// GetURL returns the value of the 'URL' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Registry) GetURL() (value string, ok bool) {
	ok = o != nil && o.url != nil
	if ok {
		value = *o.url
	}
	return
}

// TeamName returns the value of the 'team_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Registry) TeamName() string {
	if o != nil && o.teamName != nil {
		return *o.teamName
	}
	return ""
}

// GetTeamName returns the value of the 'team_name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Registry) GetTeamName() (value string, ok bool) {
	ok = o != nil && o.teamName != nil
	if ok {
		value = *o.teamName
	}
	return
}

// OrgName returns the value of the 'org_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Registry) OrgName() string {
	if o != nil && o.orgName != nil {
		return *o.orgName
	}
	return ""
}

// GetOrgName returns the value of the 'org_name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Registry) GetOrgName() (value string, ok bool) {
	ok = o != nil && o.orgName != nil
	if ok {
		value = *o.orgName
	}
	return
}

// Type returns the value of the 'type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Registry) Type() string {
	if o != nil && o.type_ != nil {
		return *o.type_
	}
	return ""
}

// GetType returns the value of the 'type' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Registry) GetType() (value string, ok bool) {
	ok = o != nil && o.type_ != nil
	if ok {
		value = *o.type_
	}
	return
}

// CloudAlias returns the value of the 'cloud_alias' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Registry) CloudAlias() bool {
	if o != nil && o.cloudAlias != nil {
		return *o.cloudAlias
	}
	return false
}

// GetCloudAlias returns the value of the 'cloud_alias' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Registry) GetCloudAlias() (value bool, ok bool) {
	ok = o != nil && o.cloudAlias != nil
	if ok {
		value = *o.cloudAlias
	}
	return
}

// RegistryList is a list of values of the 'registry' type.
type RegistryList struct {
	items []*Registry
}

// Len returns the length of the list.
func (l *RegistryList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *RegistryList) Slice() []*Registry {
	var slice []*Registry
	if l == nil {
		slice = make([]*Registry, 0)
	} else {
		slice = make([]*Registry, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *RegistryList) Each(f func(item *Registry) bool) {
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
func (l *RegistryList) Range(f func(index int, item *Registry) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
