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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// GroupKind is the name of the type used to represent objects
// of type 'group'.
const GroupKind = "Group"

// GroupLinkKind is the name of the type used to represent links
// to objects of type 'group'.
const GroupLinkKind = "GroupLink"

// GroupNilKind is the name of the type used to nil references
// to objects of type 'group'.
const GroupNilKind = "GroupNil"

// Group represents the values of the 'group' type.
//
// Representation of a group of users.
type Group struct {
	id    *string
	href  *string
	link  bool
	users *UserList
}

// Kind returns the name of the type of the object.
func (o *Group) Kind() string {
	if o == nil {
		return GroupNilKind
	}
	if o.link {
		return GroupLinkKind
	}
	return GroupKind
}

// ID returns the identifier of the object.
func (o *Group) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Group) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Group) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Group) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Group) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Users returns the value of the 'users' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of users of the group.
func (o *Group) Users() *UserList {
	if o == nil {
		return nil
	}
	return o.users
}

// GetUsers returns the value of the 'users' attribute and
// a flag indicating if the attribute has a value.
//
// List of users of the group.
func (o *Group) GetUsers() (value *UserList, ok bool) {
	ok = o != nil && o.users != nil
	if ok {
		value = o.users
	}
	return
}

// GroupList is a list of values of the 'group' type.
type GroupList struct {
	items []*Group
}

// Len returns the length of the list.
func (l *GroupList) Len() int {
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
func (l *GroupList) Slice() []*Group {
	var slice []*Group
	if l == nil {
		slice = make([]*Group, 0)
	} else {
		slice = make([]*Group, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *GroupList) Each(f func(item *Group) bool) {
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
func (l *GroupList) Range(f func(index int, item *Group) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
