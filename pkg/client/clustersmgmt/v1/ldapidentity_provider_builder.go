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

// LdapidentityProviderBuilder contains the data and logic needed to build 'ldapidentity_provider' objects.
//
// Details for `ldap` identity providers.
type LdapidentityProviderBuilder struct {
	ldapattributes *LdapattributesBuilder
	bindDN         *string
	bindPassword   *string
	ca             *string
	url            *string
	insecure       *bool
}

// NewLdapidentityProvider creates a new builder of 'ldapidentity_provider' objects.
func NewLdapidentityProvider() *LdapidentityProviderBuilder {
	return new(LdapidentityProviderBuilder)
}

// Ldapattributes sets the value of the 'ldapattributes' attribute
// to the given value.
//
// LDAP attributes used to configure the LDAP identity provider.
func (b *LdapidentityProviderBuilder) Ldapattributes(value *LdapattributesBuilder) *LdapidentityProviderBuilder {
	b.ldapattributes = value
	return b
}

// BindDN sets the value of the 'bind_DN' attribute
// to the given value.
//
//
func (b *LdapidentityProviderBuilder) BindDN(value string) *LdapidentityProviderBuilder {
	b.bindDN = &value
	return b
}

// BindPassword sets the value of the 'bind_password' attribute
// to the given value.
//
//
func (b *LdapidentityProviderBuilder) BindPassword(value string) *LdapidentityProviderBuilder {
	b.bindPassword = &value
	return b
}

// CA sets the value of the 'CA' attribute
// to the given value.
//
//
func (b *LdapidentityProviderBuilder) CA(value string) *LdapidentityProviderBuilder {
	b.ca = &value
	return b
}

// URL sets the value of the 'URL' attribute
// to the given value.
//
//
func (b *LdapidentityProviderBuilder) URL(value string) *LdapidentityProviderBuilder {
	b.url = &value
	return b
}

// Insecure sets the value of the 'insecure' attribute
// to the given value.
//
//
func (b *LdapidentityProviderBuilder) Insecure(value bool) *LdapidentityProviderBuilder {
	b.insecure = &value
	return b
}

// Build creates a 'ldapidentity_provider' object using the configuration stored in the builder.
func (b *LdapidentityProviderBuilder) Build() (object *LdapidentityProvider, err error) {
	object = new(LdapidentityProvider)
	if b.ldapattributes != nil {
		object.ldapattributes, err = b.ldapattributes.Build()
		if err != nil {
			return
		}
	}
	if b.bindDN != nil {
		object.bindDN = b.bindDN
	}
	if b.bindPassword != nil {
		object.bindPassword = b.bindPassword
	}
	if b.ca != nil {
		object.ca = b.ca
	}
	if b.url != nil {
		object.url = b.url
	}
	if b.insecure != nil {
		object.insecure = b.insecure
	}
	return
}
