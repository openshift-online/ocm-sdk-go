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

package v1 // github.com/openshift-online/uhc-sdk-go/accountsmgmt/v1

import (
	"github.com/openshift-online/uhc-sdk-go/helpers"
)

// quotaSummaryData is the data structure used internally to marshal and unmarshal
// objects of type 'quota_summary'.
type quotaSummaryData struct {
	OrganizationID       *string "json:\"organization_id,omitempty\""
	ResourceName         *string "json:\"resource_name,omitempty\""
	ResourceType         *string "json:\"resource_type,omitempty\""
	BYOC                 *bool   "json:\"byoc,omitempty\""
	AvailabilityZoneType *string "json:\"availability_zone_type,omitempty\""
	Allowed              *int    "json:\"allowed,omitempty\""
	Reserved             *int    "json:\"reserved,omitempty\""
}

// MarshalQuotaSummary writes a value of the 'quota_summary' to the given target,
// which can be a writer or a JSON encoder.
func MarshalQuotaSummary(object *QuotaSummary, target interface{}) error {
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

// wrap is the method used internally to convert a value of the 'quota_summary'
// value to a JSON document.
func (o *QuotaSummary) wrap() (data *quotaSummaryData, err error) {
	if o == nil {
		return
	}
	data = new(quotaSummaryData)
	data.OrganizationID = o.organizationID
	data.ResourceName = o.resourceName
	data.ResourceType = o.resourceType
	data.BYOC = o.byoc
	data.AvailabilityZoneType = o.availabilityZoneType
	data.Allowed = o.allowed
	data.Reserved = o.reserved
	return
}

// UnmarshalQuotaSummary reads a value of the 'quota_summary' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalQuotaSummary(source interface{}) (object *QuotaSummary, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(quotaSummaryData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'quota_summary' type.
func (d *quotaSummaryData) unwrap() (object *QuotaSummary, err error) {
	if d == nil {
		return
	}
	object = new(QuotaSummary)
	object.organizationID = d.OrganizationID
	object.resourceName = d.ResourceName
	object.resourceType = d.ResourceType
	object.byoc = d.BYOC
	object.availabilityZoneType = d.AvailabilityZoneType
	object.allowed = d.Allowed
	object.reserved = d.Reserved
	return
}
