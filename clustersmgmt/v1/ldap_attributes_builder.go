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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// LDAPAttributesBuilder contains the data and logic needed to build 'LDAP_attributes' objects.
//
// LDAP attributes used to configure the LDAP identity provider.
type LDAPAttributesBuilder struct {
	email             []string
	id                []string
	name              []string
	preferredUsername []string
}

// NewLDAPAttributes creates a new builder of 'LDAP_attributes' objects.
func NewLDAPAttributes() *LDAPAttributesBuilder {
	return new(LDAPAttributesBuilder)
}

// Email sets the value of the 'email' attribute
// to the given values.
//
//
func (b *LDAPAttributesBuilder) Email(values ...string) *LDAPAttributesBuilder {
	b.email = make([]string, len(values))
	copy(b.email, values)
	return b
}

// ID sets the value of the 'ID' attribute
// to the given values.
//
//
func (b *LDAPAttributesBuilder) ID(values ...string) *LDAPAttributesBuilder {
	b.id = make([]string, len(values))
	copy(b.id, values)
	return b
}

// Name sets the value of the 'name' attribute
// to the given values.
//
//
func (b *LDAPAttributesBuilder) Name(values ...string) *LDAPAttributesBuilder {
	b.name = make([]string, len(values))
	copy(b.name, values)
	return b
}

// PreferredUsername sets the value of the 'preferred_username' attribute
// to the given values.
//
//
func (b *LDAPAttributesBuilder) PreferredUsername(values ...string) *LDAPAttributesBuilder {
	b.preferredUsername = make([]string, len(values))
	copy(b.preferredUsername, values)
	return b
}

// Build creates a 'LDAP_attributes' object using the configuration stored in the builder.
func (b *LDAPAttributesBuilder) Build() (object *LDAPAttributes, err error) {
	object = new(LDAPAttributes)
	if b.email != nil {
		object.email = make([]string, len(b.email))
		copy(object.email, b.email)
	}
	if b.id != nil {
		object.id = make([]string, len(b.id))
		copy(object.id, b.id)
	}
	if b.name != nil {
		object.name = make([]string, len(b.name))
		copy(object.name, b.name)
	}
	if b.preferredUsername != nil {
		object.preferredUsername = make([]string, len(b.preferredUsername))
		copy(object.preferredUsername, b.preferredUsername)
	}
	return
}
