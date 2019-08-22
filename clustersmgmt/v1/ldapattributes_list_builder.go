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

// LdapattributesListBuilder contains the data and logic needed to build
// 'ldapattributes' objects.
type LdapattributesListBuilder struct {
	items []*LdapattributesBuilder
}

// NewLdapattributesList creates a new builder of 'ldapattributes' objects.
func NewLdapattributesList() *LdapattributesListBuilder {
	return new(LdapattributesListBuilder)
}

// Items sets the items of the list.
func (b *LdapattributesListBuilder) Items(values ...*LdapattributesBuilder) *LdapattributesListBuilder {
	b.items = make([]*LdapattributesBuilder, len(values))
	copy(b.items, values)
	return b
}

// Build creates a list of 'ldapattributes' objects using the
// configuration stored in the builder.
func (b *LdapattributesListBuilder) Build() (list *LdapattributesList, err error) {
	items := make([]*Ldapattributes, len(b.items))
	for i, item := range b.items {
		items[i], err = item.Build()
		if err != nil {
			return
		}
	}
	list = new(LdapattributesList)
	list.items = items
	return
}
