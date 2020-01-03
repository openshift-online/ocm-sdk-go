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

import (
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// gcpFlavourData is the data structure used internally to marshal and unmarshal
// objects of type 'GCP_flavour'.
type gcpFlavourData struct {
	ComputeInstanceType *string "json:\"compute_instance_type,omitempty\""
	InfraInstanceType   *string "json:\"infra_instance_type,omitempty\""
	MasterInstanceType  *string "json:\"master_instance_type,omitempty\""
}

// MarshalGCPFlavour writes a value of the 'GCP_flavour' to the given target,
// which can be a writer or a JSON encoder.
func MarshalGCPFlavour(object *GCPFlavour, target interface{}) error {
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

// wrap is the method used internally to convert a value of the 'GCP_flavour'
// value to a JSON document.
func (o *GCPFlavour) wrap() (data *gcpFlavourData, err error) {
	if o == nil {
		return
	}
	data = new(gcpFlavourData)
	data.ComputeInstanceType = o.computeInstanceType
	data.InfraInstanceType = o.infraInstanceType
	data.MasterInstanceType = o.masterInstanceType
	return
}

// UnmarshalGCPFlavour reads a value of the 'GCP_flavour' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalGCPFlavour(source interface{}) (object *GCPFlavour, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(gcpFlavourData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'GCP_flavour' type.
func (d *gcpFlavourData) unwrap() (object *GCPFlavour, err error) {
	if d == nil {
		return
	}
	object = new(GCPFlavour)
	object.computeInstanceType = d.ComputeInstanceType
	object.infraInstanceType = d.InfraInstanceType
	object.masterInstanceType = d.MasterInstanceType
	return
}
