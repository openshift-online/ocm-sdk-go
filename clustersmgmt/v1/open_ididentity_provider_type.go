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

// OpenIdidentityProvider represents the values of the 'open_ididentity_provider' type.
//
// Details for `openid` identity providers.
type OpenIdidentityProvider struct {
	ca                       *string
	claims                   *OpenIdclaims
	clientID                 *string
	clientSecret             *string
	extraAuthorizeParameters map[string]string
	extraScopes              []string
	urls                     *OpenIdurls
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *OpenIdidentityProvider) Empty() bool {
	return o == nil || (o.ca == nil &&
		o.claims == nil &&
		o.clientID == nil &&
		o.clientSecret == nil &&
		o.extraAuthorizeParameters == nil &&
		o.extraScopes == nil &&
		o.urls == nil &&
		true)
}

// CA returns the value of the 'CA' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Certificate bunde to use to validate server certificates for the configured URL.
func (o *OpenIdidentityProvider) CA() string {
	if o != nil && o.ca != nil {
		return *o.ca
	}
	return ""
}

// GetCA returns the value of the 'CA' attribute and
// a flag indicating if the attribute has a value.
//
// Certificate bunde to use to validate server certificates for the configured URL.
func (o *OpenIdidentityProvider) GetCA() (value string, ok bool) {
	ok = o != nil && o.ca != nil
	if ok {
		value = *o.ca
	}
	return
}

// Claims returns the value of the 'claims' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Claims used to configure the provider.
func (o *OpenIdidentityProvider) Claims() *OpenIdclaims {
	if o == nil {
		return nil
	}
	return o.claims
}

// GetClaims returns the value of the 'claims' attribute and
// a flag indicating if the attribute has a value.
//
// Claims used to configure the provider.
func (o *OpenIdidentityProvider) GetClaims() (value *OpenIdclaims, ok bool) {
	ok = o != nil && o.claims != nil
	if ok {
		value = o.claims
	}
	return
}

// ClientID returns the value of the 'client_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Identifier of a client registered with the _OpenID_ provider.
func (o *OpenIdidentityProvider) ClientID() string {
	if o != nil && o.clientID != nil {
		return *o.clientID
	}
	return ""
}

// GetClientID returns the value of the 'client_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Identifier of a client registered with the _OpenID_ provider.
func (o *OpenIdidentityProvider) GetClientID() (value string, ok bool) {
	ok = o != nil && o.clientID != nil
	if ok {
		value = *o.clientID
	}
	return
}

// ClientSecret returns the value of the 'client_secret' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Client secret.
func (o *OpenIdidentityProvider) ClientSecret() string {
	if o != nil && o.clientSecret != nil {
		return *o.clientSecret
	}
	return ""
}

// GetClientSecret returns the value of the 'client_secret' attribute and
// a flag indicating if the attribute has a value.
//
// Client secret.
func (o *OpenIdidentityProvider) GetClientSecret() (value string, ok bool) {
	ok = o != nil && o.clientSecret != nil
	if ok {
		value = *o.clientSecret
	}
	return
}

// ExtraAuthorizeParameters returns the value of the 'extra_authorize_parameters' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Optional map of extra parameters to add to the authorization token request.
func (o *OpenIdidentityProvider) ExtraAuthorizeParameters() map[string]string {
	if o == nil {
		return nil
	}
	return o.extraAuthorizeParameters
}

// GetExtraAuthorizeParameters returns the value of the 'extra_authorize_parameters' attribute and
// a flag indicating if the attribute has a value.
//
// Optional map of extra parameters to add to the authorization token request.
func (o *OpenIdidentityProvider) GetExtraAuthorizeParameters() (value map[string]string, ok bool) {
	ok = o != nil && o.extraAuthorizeParameters != nil
	if ok {
		value = o.extraAuthorizeParameters
	}
	return
}

// ExtraScopes returns the value of the 'extra_scopes' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Optional list of scopes to request, in addition to the `openid` scope, during the
// authorization token request.
func (o *OpenIdidentityProvider) ExtraScopes() []string {
	if o == nil {
		return nil
	}
	return o.extraScopes
}

// GetExtraScopes returns the value of the 'extra_scopes' attribute and
// a flag indicating if the attribute has a value.
//
// Optional list of scopes to request, in addition to the `openid` scope, during the
// authorization token request.
func (o *OpenIdidentityProvider) GetExtraScopes() (value []string, ok bool) {
	ok = o != nil && o.extraScopes != nil
	if ok {
		value = o.extraScopes
	}
	return
}

// URLS returns the value of the 'URLS' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// URLs of the provider.
func (o *OpenIdidentityProvider) URLS() *OpenIdurls {
	if o == nil {
		return nil
	}
	return o.urls
}

// GetURLS returns the value of the 'URLS' attribute and
// a flag indicating if the attribute has a value.
//
// URLs of the provider.
func (o *OpenIdidentityProvider) GetURLS() (value *OpenIdurls, ok bool) {
	ok = o != nil && o.urls != nil
	if ok {
		value = o.urls
	}
	return
}

// OpenIdidentityProviderList is a list of values of the 'open_ididentity_provider' type.
type OpenIdidentityProviderList struct {
	items []*OpenIdidentityProvider
}

// Len returns the length of the list.
func (l *OpenIdidentityProviderList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *OpenIdidentityProviderList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *OpenIdidentityProviderList) Get(i int) *OpenIdidentityProvider {
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
func (l *OpenIdidentityProviderList) Slice() []*OpenIdidentityProvider {
	var slice []*OpenIdidentityProvider
	if l == nil {
		slice = make([]*OpenIdidentityProvider, 0)
	} else {
		slice = make([]*OpenIdidentityProvider, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *OpenIdidentityProviderList) Each(f func(item *OpenIdidentityProvider) bool) {
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
func (l *OpenIdidentityProviderList) Range(f func(index int, item *OpenIdidentityProvider) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
