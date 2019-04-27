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

// OpenIdclaimsBuilder contains the data and logic needed to build 'open_idclaims' objects.
//
// _OpenID_ identity provider claims.
type OpenIdclaimsBuilder struct {
	email             []string
	name              []string
	preferredUsername []string
}

// NewOpenIdclaims creates a new builder of 'open_idclaims' objects.
func NewOpenIdclaims() *OpenIdclaimsBuilder {
	return new(OpenIdclaimsBuilder)
}

// Email sets the value of the 'email' attribute
// to the given values.
//
//
func (b *OpenIdclaimsBuilder) Email(values ...string) *OpenIdclaimsBuilder {
	b.email = make([]string, len(values))
	copy(b.email, values)
	return b
}

// Name sets the value of the 'name' attribute
// to the given values.
//
//
func (b *OpenIdclaimsBuilder) Name(values ...string) *OpenIdclaimsBuilder {
	b.name = make([]string, len(values))
	copy(b.name, values)
	return b
}

// PreferredUsername sets the value of the 'preferred_username' attribute
// to the given values.
//
//
func (b *OpenIdclaimsBuilder) PreferredUsername(values ...string) *OpenIdclaimsBuilder {
	b.preferredUsername = make([]string, len(values))
	copy(b.preferredUsername, values)
	return b
}

// Build creates a 'open_idclaims' object using the configuration stored in the builder.
func (b *OpenIdclaimsBuilder) Build() (object *OpenIdclaims, err error) {
	object = new(OpenIdclaims)
	if b.email != nil {
		object.email = make([]string, len(b.email))
		copy(object.email, b.email)
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
