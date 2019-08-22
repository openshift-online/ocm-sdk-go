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

// OpenIdurlsBuilder contains the data and logic needed to build 'open_idurls' objects.
//
// _OpenID_ identity provider URLs.
type OpenIdurlsBuilder struct {
	authorize *string
	token     *string
	userInfo  *string
}

// NewOpenIdurls creates a new builder of 'open_idurls' objects.
func NewOpenIdurls() *OpenIdurlsBuilder {
	return new(OpenIdurlsBuilder)
}

// Authorize sets the value of the 'authorize' attribute
// to the given value.
//
//
func (b *OpenIdurlsBuilder) Authorize(value string) *OpenIdurlsBuilder {
	b.authorize = &value
	return b
}

// Token sets the value of the 'token' attribute
// to the given value.
//
//
func (b *OpenIdurlsBuilder) Token(value string) *OpenIdurlsBuilder {
	b.token = &value
	return b
}

// UserInfo sets the value of the 'user_info' attribute
// to the given value.
//
//
func (b *OpenIdurlsBuilder) UserInfo(value string) *OpenIdurlsBuilder {
	b.userInfo = &value
	return b
}

// Build creates a 'open_idurls' object using the configuration stored in the builder.
func (b *OpenIdurlsBuilder) Build() (object *OpenIdurls, err error) {
	object = new(OpenIdurls)
	if b.authorize != nil {
		object.authorize = b.authorize
	}
	if b.token != nil {
		object.token = b.token
	}
	if b.userInfo != nil {
		object.userInfo = b.userInfo
	}
	return
}
