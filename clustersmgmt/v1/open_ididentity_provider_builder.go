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

// OpenIdidentityProviderBuilder contains the data and logic needed to build 'open_ididentity_provider' objects.
//
// Details for `openid` identity providers.
type OpenIdidentityProviderBuilder struct {
	ca                       *string
	claims                   *OpenIdclaimsBuilder
	clientID                 *string
	clientSecret             *string
	extraAuthorizeParameters map[string]string
	extraScopes              []string
	urls                     *OpenIdurlsBuilder
}

// NewOpenIdidentityProvider creates a new builder of 'open_ididentity_provider' objects.
func NewOpenIdidentityProvider() *OpenIdidentityProviderBuilder {
	return new(OpenIdidentityProviderBuilder)
}

// CA sets the value of the 'CA' attribute
// to the given value.
//
//
func (b *OpenIdidentityProviderBuilder) CA(value string) *OpenIdidentityProviderBuilder {
	b.ca = &value
	return b
}

// Claims sets the value of the 'claims' attribute
// to the given value.
//
// _OpenID_ identity provider claims.
func (b *OpenIdidentityProviderBuilder) Claims(value *OpenIdclaimsBuilder) *OpenIdidentityProviderBuilder {
	b.claims = value
	return b
}

// ClientID sets the value of the 'client_ID' attribute
// to the given value.
//
//
func (b *OpenIdidentityProviderBuilder) ClientID(value string) *OpenIdidentityProviderBuilder {
	b.clientID = &value
	return b
}

// ClientSecret sets the value of the 'client_secret' attribute
// to the given value.
//
//
func (b *OpenIdidentityProviderBuilder) ClientSecret(value string) *OpenIdidentityProviderBuilder {
	b.clientSecret = &value
	return b
}

// ExtraAuthorizeParameters sets the value of the 'extra_authorize_parameters' attribute
// to the given value.
//
//
func (b *OpenIdidentityProviderBuilder) ExtraAuthorizeParameters(value map[string]string) *OpenIdidentityProviderBuilder {
	b.extraAuthorizeParameters = value
	return b
}

// ExtraScopes sets the value of the 'extra_scopes' attribute
// to the given values.
//
//
func (b *OpenIdidentityProviderBuilder) ExtraScopes(values ...string) *OpenIdidentityProviderBuilder {
	b.extraScopes = make([]string, len(values))
	copy(b.extraScopes, values)
	return b
}

// URLS sets the value of the 'URLS' attribute
// to the given value.
//
// _OpenID_ identity provider URLs.
func (b *OpenIdidentityProviderBuilder) URLS(value *OpenIdurlsBuilder) *OpenIdidentityProviderBuilder {
	b.urls = value
	return b
}

// Build creates a 'open_ididentity_provider' object using the configuration stored in the builder.
func (b *OpenIdidentityProviderBuilder) Build() (object *OpenIdidentityProvider, err error) {
	object = new(OpenIdidentityProvider)
	if b.ca != nil {
		object.ca = b.ca
	}
	if b.claims != nil {
		object.claims, err = b.claims.Build()
		if err != nil {
			return
		}
	}
	if b.clientID != nil {
		object.clientID = b.clientID
	}
	if b.clientSecret != nil {
		object.clientSecret = b.clientSecret
	}
	if b.extraAuthorizeParameters != nil {
		object.extraAuthorizeParameters = b.extraAuthorizeParameters
	}
	if b.extraScopes != nil {
		object.extraScopes = make([]string, len(b.extraScopes))
		copy(object.extraScopes, b.extraScopes)
	}
	if b.urls != nil {
		object.urls, err = b.urls.Build()
		if err != nil {
			return
		}
	}
	return
}
