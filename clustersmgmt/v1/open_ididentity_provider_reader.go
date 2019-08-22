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

import (
	"github.com/openshift-online/uhc-sdk-go/helpers"
)

// openIdidentityProviderData is the data structure used internally to marshal and unmarshal
// objects of type 'open_ididentity_provider'.
type openIdidentityProviderData struct {
	CA                       *string           "json:\"ca,omitempty\""
	Claims                   *openIdclaimsData "json:\"claims,omitempty\""
	ClientID                 *string           "json:\"client_id,omitempty\""
	ClientSecret             *string           "json:\"client_secret,omitempty\""
	ExtraAuthorizeParameters map[string]string "json:\"extra_authorize_parameters,omitempty\""
	ExtraScopes              []string          "json:\"extra_scopes,omitempty\""
	URLS                     *openIdurlsData   "json:\"urls,omitempty\""
}

// MarshalOpenIdidentityProvider writes a value of the 'open_ididentity_provider' to the given target,
// which can be a writer or a JSON encoder.
func MarshalOpenIdidentityProvider(object *OpenIdidentityProvider, target interface{}) error {
	encoder, err := helpers.NewEncoder(target)
	if err != nil {
		return err
	}
	data, err := object.wrap()
	if err != nil {
		return err
	}
	return encoder.Encode(data)
}

// wrap is the method used internally to convert a value of the 'open_ididentity_provider'
// value to a JSON document.
func (o *OpenIdidentityProvider) wrap() (data *openIdidentityProviderData, err error) {
	if o == nil {
		return
	}
	data = new(openIdidentityProviderData)
	data.CA = o.ca
	data.Claims, err = o.claims.wrap()
	if err != nil {
		return
	}
	data.ClientID = o.clientID
	data.ClientSecret = o.clientSecret
	data.ExtraAuthorizeParameters = o.extraAuthorizeParameters
	data.ExtraScopes = o.extraScopes
	data.URLS, err = o.urls.wrap()
	if err != nil {
		return
	}
	return
}

// UnmarshalOpenIdidentityProvider reads a value of the 'open_ididentity_provider' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalOpenIdidentityProvider(source interface{}) (object *OpenIdidentityProvider, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(openIdidentityProviderData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'open_ididentity_provider' type.
func (d *openIdidentityProviderData) unwrap() (object *OpenIdidentityProvider, err error) {
	if d == nil {
		return
	}
	object = new(OpenIdidentityProvider)
	object.ca = d.CA
	object.claims, err = d.Claims.unwrap()
	if err != nil {
		return
	}
	object.clientID = d.ClientID
	object.clientSecret = d.ClientSecret
	object.extraAuthorizeParameters = d.ExtraAuthorizeParameters
	object.extraScopes = d.ExtraScopes
	object.urls, err = d.URLS.unwrap()
	if err != nil {
		return
	}
	return
}
