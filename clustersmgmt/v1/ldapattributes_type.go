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

package v1 // github.com/openshift-online/uhc-sdk-go/clustersmgmt/v1

// Ldapattributes represents the values of the 'ldapattributes' type.
//
// LDAP attributes used to configure the LDAP identity provider.
type Ldapattributes struct {
	email             []string
	id                []string
	name              []string
	preferredUsername []string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *Ldapattributes) Empty() bool {
	return o == nil || (o.email == nil &&
		o.id == nil &&
		o.name == nil &&
		o.preferredUsername == nil &&
		true)
}

// Email returns the value of the 'email' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of attributes to use as the mail address.
func (o *Ldapattributes) Email() []string {
	if o == nil {
		return nil
	}
	return o.email
}

// GetEmail returns the value of the 'email' attribute and
// a flag indicating if the attribute has a value.
//
// List of attributes to use as the mail address.
func (o *Ldapattributes) GetEmail() (value []string, ok bool) {
	ok = o != nil && o.email != nil
	if ok {
		value = o.email
	}
	return
}

// ID returns the value of the 'ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of attributes to use as the identity.
func (o *Ldapattributes) ID() []string {
	if o == nil {
		return nil
	}
	return o.id
}

// GetID returns the value of the 'ID' attribute and
// a flag indicating if the attribute has a value.
//
// List of attributes to use as the identity.
func (o *Ldapattributes) GetID() (value []string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = o.id
	}
	return
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of attributes to use as the display name.
func (o *Ldapattributes) Name() []string {
	if o == nil {
		return nil
	}
	return o.name
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
// List of attributes to use as the display name.
func (o *Ldapattributes) GetName() (value []string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = o.name
	}
	return
}

// PreferredUsername returns the value of the 'preferred_username' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of attributes to use as the preferred user name when provisioning a user.
func (o *Ldapattributes) PreferredUsername() []string {
	if o == nil {
		return nil
	}
	return o.preferredUsername
}

// GetPreferredUsername returns the value of the 'preferred_username' attribute and
// a flag indicating if the attribute has a value.
//
// List of attributes to use as the preferred user name when provisioning a user.
func (o *Ldapattributes) GetPreferredUsername() (value []string, ok bool) {
	ok = o != nil && o.preferredUsername != nil
	if ok {
		value = o.preferredUsername
	}
	return
}

// LdapattributesList is a list of values of the 'ldapattributes' type.
type LdapattributesList struct {
	items []*Ldapattributes
}

// Len returns the length of the list.
func (l *LdapattributesList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *LdapattributesList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *LdapattributesList) Get(i int) *Ldapattributes {
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
func (l *LdapattributesList) Slice() []*Ldapattributes {
	var slice []*Ldapattributes
	if l == nil {
		slice = make([]*Ldapattributes, 0)
	} else {
		slice = make([]*Ldapattributes, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *LdapattributesList) Each(f func(item *Ldapattributes) bool) {
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
func (l *LdapattributesList) Range(f func(index int, item *Ldapattributes) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
