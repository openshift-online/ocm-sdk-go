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

// LdapidentityProvider represents the values of the 'ldapidentity_provider' type.
//
// Details for `ldap` identity providers.
type LdapidentityProvider struct {
	ldapattributes *Ldapattributes
	bindDN         *string
	bindPassword   *string
	ca             *string
	url            *string
	insecure       *bool
}

// Ldapattributes returns the value of the 'ldapattributes' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// LDAP attributes used to configure the provider.
func (o *LdapidentityProvider) Ldapattributes() *Ldapattributes {
	if o == nil {
		return nil
	}
	return o.ldapattributes
}

// GetLdapattributes returns the value of the 'ldapattributes' attribute and
// a flag indicating if the attribute has a value.
//
// LDAP attributes used to configure the provider.
func (o *LdapidentityProvider) GetLdapattributes() (value *Ldapattributes, ok bool) {
	ok = o != nil && o.ldapattributes != nil
	if ok {
		value = o.ldapattributes
	}
	return
}

// BindDN returns the value of the 'bind_DN' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Optional distinguished name to use to bind during the search phase.
func (o *LdapidentityProvider) BindDN() string {
	if o != nil && o.bindDN != nil {
		return *o.bindDN
	}
	return ""
}

// GetBindDN returns the value of the 'bind_DN' attribute and
// a flag indicating if the attribute has a value.
//
// Optional distinguished name to use to bind during the search phase.
func (o *LdapidentityProvider) GetBindDN() (value string, ok bool) {
	ok = o != nil && o.bindDN != nil
	if ok {
		value = *o.bindDN
	}
	return
}

// BindPassword returns the value of the 'bind_password' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Optional password to use to bind during the search phase.
func (o *LdapidentityProvider) BindPassword() string {
	if o != nil && o.bindPassword != nil {
		return *o.bindPassword
	}
	return ""
}

// GetBindPassword returns the value of the 'bind_password' attribute and
// a flag indicating if the attribute has a value.
//
// Optional password to use to bind during the search phase.
func (o *LdapidentityProvider) GetBindPassword() (value string, ok bool) {
	ok = o != nil && o.bindPassword != nil
	if ok {
		value = *o.bindPassword
	}
	return
}

// CA returns the value of the 'CA' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Certificate bundle to use to validate server certificates for the configured URL.
func (o *LdapidentityProvider) CA() string {
	if o != nil && o.ca != nil {
		return *o.ca
	}
	return ""
}

// GetCA returns the value of the 'CA' attribute and
// a flag indicating if the attribute has a value.
//
// Certificate bundle to use to validate server certificates for the configured URL.
func (o *LdapidentityProvider) GetCA() (value string, ok bool) {
	ok = o != nil && o.ca != nil
	if ok {
		value = *o.ca
	}
	return
}

// URL returns the value of the 'URL' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// An https://tools.ietf.org/html/rfc2255[RFC 2255] URL which specifies the LDAP host and
// search parameters to use.
func (o *LdapidentityProvider) URL() string {
	if o != nil && o.url != nil {
		return *o.url
	}
	return ""
}

// GetURL returns the value of the 'URL' attribute and
// a flag indicating if the attribute has a value.
//
// An https://tools.ietf.org/html/rfc2255[RFC 2255] URL which specifies the LDAP host and
// search parameters to use.
func (o *LdapidentityProvider) GetURL() (value string, ok bool) {
	ok = o != nil && o.url != nil
	if ok {
		value = *o.url
	}
	return
}

// Insecure returns the value of the 'insecure' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// When `true` no TLS connection is made to the server. When `false` `ldaps://...` URLs
// connect using TLS and `ldap://...` are upgraded to TLS.
func (o *LdapidentityProvider) Insecure() bool {
	if o != nil && o.insecure != nil {
		return *o.insecure
	}
	return false
}

// GetInsecure returns the value of the 'insecure' attribute and
// a flag indicating if the attribute has a value.
//
// When `true` no TLS connection is made to the server. When `false` `ldaps://...` URLs
// connect using TLS and `ldap://...` are upgraded to TLS.
func (o *LdapidentityProvider) GetInsecure() (value bool, ok bool) {
	ok = o != nil && o.insecure != nil
	if ok {
		value = *o.insecure
	}
	return
}

// LdapidentityProviderList is a list of values of the 'ldapidentity_provider' type.
type LdapidentityProviderList struct {
	items []*LdapidentityProvider
}

// Len returns the length of the list.
func (l *LdapidentityProviderList) Len() int {
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
func (l *LdapidentityProviderList) Slice() []*LdapidentityProvider {
	var slice []*LdapidentityProvider
	if l == nil {
		slice = make([]*LdapidentityProvider, 0)
	} else {
		slice = make([]*LdapidentityProvider, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *LdapidentityProviderList) Each(f func(item *LdapidentityProvider) bool) {
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
func (l *LdapidentityProviderList) Range(f func(index int, item *LdapidentityProvider) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
