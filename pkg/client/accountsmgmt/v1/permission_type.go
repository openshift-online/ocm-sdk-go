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

// PermissionKind is the name of the type used to represent objects
// of type 'permission'.
const PermissionKind = "Permission"

// PermissionLinkKind is the name of the type used to represent links
// to objects of type 'permission'.
const PermissionLinkKind = "PermissionLink"

// PermissionNilKind is the name of the type used to nil references
// to objects of type 'permission'.
const PermissionNilKind = "PermissionNil"

// Permission represents the values of the 'permission' type.
//
//
type Permission struct {
	id           *string
	href         *string
	link         bool
	action       *Action
	resourceType *string
	roleID       *string
}

// Kind returns the name of the type of the object.
func (o *Permission) Kind() string {
	if o == nil {
		return PermissionNilKind
	}
	if o.link {
		return PermissionLinkKind
	}
	return PermissionKind
}

// ID returns the identifier of the object.
func (o *Permission) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Permission) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Permission) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Permission) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Permission) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Action returns the value of the 'action' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Permission) Action() Action {
	if o != nil && o.action != nil {
		return *o.action
	}
	return Action("")
}

// GetAction returns the value of the 'action' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Permission) GetAction() (value Action, ok bool) {
	ok = o != nil && o.action != nil
	if ok {
		value = *o.action
	}
	return
}

// ResourceType returns the value of the 'resource_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Permission) ResourceType() string {
	if o != nil && o.resourceType != nil {
		return *o.resourceType
	}
	return ""
}

// GetResourceType returns the value of the 'resource_type' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Permission) GetResourceType() (value string, ok bool) {
	ok = o != nil && o.resourceType != nil
	if ok {
		value = *o.resourceType
	}
	return
}

// RoleID returns the value of the 'role_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Permission) RoleID() string {
	if o != nil && o.roleID != nil {
		return *o.roleID
	}
	return ""
}

// GetRoleID returns the value of the 'role_ID' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Permission) GetRoleID() (value string, ok bool) {
	ok = o != nil && o.roleID != nil
	if ok {
		value = *o.roleID
	}
	return
}

// PermissionList is a list of values of the 'permission' type.
type PermissionList struct {
	items []*Permission
}

// Len returns the length of the list.
func (l *PermissionList) Len() int {
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
func (l *PermissionList) Slice() []*Permission {
	var slice []*Permission
	if l == nil {
		slice = make([]*Permission, 0)
	} else {
		slice = make([]*Permission, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *PermissionList) Each(f func(item *Permission) bool) {
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
func (l *PermissionList) Range(f func(index int, item *Permission) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
