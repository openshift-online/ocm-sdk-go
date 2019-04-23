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

// RoleKind is the name of the type used to represent objects
// of type 'role'.
const RoleKind = "Role"

// RoleLinkKind is the name of the type used to represent links
// to objects of type 'role'.
const RoleLinkKind = "RoleLink"

// RoleNilKind is the name of the type used to nil references
// to objects of type 'role'.
const RoleNilKind = "RoleNil"

// Role represents the values of the 'role' type.
//
//
type Role struct {
	id          *string
	href        *string
	link        bool
	name        *string
	permissions *PermissionList
}

// Kind returns the name of the type of the object.
func (o *Role) Kind() string {
	if o == nil {
		return RoleNilKind
	}
	if o.link {
		return RoleLinkKind
	}
	return RoleKind
}

// ID returns the identifier of the object.
func (o *Role) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Role) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Role) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Role) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Role) GetHREF() (value string, ok bool) {
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
func (o *Role) Name() string {
	if o != nil && o.name != nil {
		return *o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Role) GetName() (value string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = *o.name
	}
	return
}

// Permissions returns the value of the 'permissions' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Role) Permissions() *PermissionList {
	if o == nil {
		return nil
	}
	return o.permissions
}

// GetPermissions returns the value of the 'permissions' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Role) GetPermissions() (value *PermissionList, ok bool) {
	ok = o != nil && o.permissions != nil
	if ok {
		value = o.permissions
	}
	return
}

// RoleList is a list of values of the 'role' type.
type RoleList struct {
	items []*Role
}

// Len returns the length of the list.
func (l *RoleList) Len() int {
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
func (l *RoleList) Slice() []*Role {
	var slice []*Role
	if l == nil {
		slice = make([]*Role, 0)
	} else {
		slice = make([]*Role, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *RoleList) Each(f func(item *Role) bool) {
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
func (l *RoleList) Range(f func(index int, item *Role) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
