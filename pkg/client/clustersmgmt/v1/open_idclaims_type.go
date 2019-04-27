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

// OpenIdclaims represents the values of the 'open_idclaims' type.
//
// _OpenID_ identity provider claims.
type OpenIdclaims struct {
	email             []string
	name              []string
	preferredUsername []string
}

// Email returns the value of the 'email' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of claims to use as the mail address.
func (o *OpenIdclaims) Email() []string {
	if o == nil {
		return nil
	}
	return o.email
}

// GetEmail returns the value of the 'email' attribute and
// a flag indicating if the attribute has a value.
//
// List of claims to use as the mail address.
func (o *OpenIdclaims) GetEmail() (value []string, ok bool) {
	ok = o != nil && o.email != nil
	if ok {
		value = o.email
	}
	return
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of claims to use as the display name.
func (o *OpenIdclaims) Name() []string {
	if o == nil {
		return nil
	}
	return o.name
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
// List of claims to use as the display name.
func (o *OpenIdclaims) GetName() (value []string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = o.name
	}
	return
}

// PreferredUsername returns the value of the 'preferred_username' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of claims to use as the preferred user name when provisioning a user.
func (o *OpenIdclaims) PreferredUsername() []string {
	if o == nil {
		return nil
	}
	return o.preferredUsername
}

// GetPreferredUsername returns the value of the 'preferred_username' attribute and
// a flag indicating if the attribute has a value.
//
// List of claims to use as the preferred user name when provisioning a user.
func (o *OpenIdclaims) GetPreferredUsername() (value []string, ok bool) {
	ok = o != nil && o.preferredUsername != nil
	if ok {
		value = o.preferredUsername
	}
	return
}

// OpenIdclaimsList is a list of values of the 'open_idclaims' type.
type OpenIdclaimsList struct {
	items []*OpenIdclaims
}

// Len returns the length of the list.
func (l *OpenIdclaimsList) Len() int {
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
func (l *OpenIdclaimsList) Slice() []*OpenIdclaims {
	var slice []*OpenIdclaims
	if l == nil {
		slice = make([]*OpenIdclaims, 0)
	} else {
		slice = make([]*OpenIdclaims, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *OpenIdclaimsList) Each(f func(item *OpenIdclaims) bool) {
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
func (l *OpenIdclaimsList) Range(f func(index int, item *OpenIdclaims) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
