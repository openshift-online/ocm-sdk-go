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

// OpenIdurls represents the values of the 'open_idurls' type.
//
// _OpenID_ identity provider URLs.
type OpenIdurls struct {
	authorize *string
	token     *string
	userInfo  *string
}

// Authorize returns the value of the 'authorize' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Authorization endpoint described in the _OpenID_ specification. Must use HTTPS.
func (o *OpenIdurls) Authorize() string {
	if o != nil && o.authorize != nil {
		return *o.authorize
	}
	return ""
}

// GetAuthorize returns the value of the 'authorize' attribute and
// a flag indicating if the attribute has a value.
//
// Authorization endpoint described in the _OpenID_ specification. Must use HTTPS.
func (o *OpenIdurls) GetAuthorize() (value string, ok bool) {
	ok = o != nil && o.authorize != nil
	if ok {
		value = *o.authorize
	}
	return
}

// Token returns the value of the 'token' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Token endpoint described in the _OpenID_ specification. Must use HTTPS.
func (o *OpenIdurls) Token() string {
	if o != nil && o.token != nil {
		return *o.token
	}
	return ""
}

// GetToken returns the value of the 'token' attribute and
// a flag indicating if the attribute has a value.
//
// Token endpoint described in the _OpenID_ specification. Must use HTTPS.
func (o *OpenIdurls) GetToken() (value string, ok bool) {
	ok = o != nil && o.token != nil
	if ok {
		value = *o.token
	}
	return
}

// UserInfo returns the value of the 'user_info' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// User information endpoint described in the _OpenID_ specification. Must use HTTPS.
func (o *OpenIdurls) UserInfo() string {
	if o != nil && o.userInfo != nil {
		return *o.userInfo
	}
	return ""
}

// GetUserInfo returns the value of the 'user_info' attribute and
// a flag indicating if the attribute has a value.
//
// User information endpoint described in the _OpenID_ specification. Must use HTTPS.
func (o *OpenIdurls) GetUserInfo() (value string, ok bool) {
	ok = o != nil && o.userInfo != nil
	if ok {
		value = *o.userInfo
	}
	return
}

// OpenIdurlsList is a list of values of the 'open_idurls' type.
type OpenIdurlsList struct {
	items []*OpenIdurls
}

// Len returns the length of the list.
func (l *OpenIdurlsList) Len() int {
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
func (l *OpenIdurlsList) Slice() []*OpenIdurls {
	var slice []*OpenIdurls
	if l == nil {
		slice = make([]*OpenIdurls, 0)
	} else {
		slice = make([]*OpenIdurls, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *OpenIdurlsList) Each(f func(item *OpenIdurls) bool) {
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
func (l *OpenIdurlsList) Range(f func(index int, item *OpenIdurls) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
