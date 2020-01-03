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

// gcpFlavourListData is type used internally to marshal and unmarshal lists of objects
// of type 'GCP_flavour'.
type gcpFlavourListData []*gcpFlavourData

// UnmarshalGCPFlavourList reads a list of values of the 'GCP_flavour'
// from the given source, which can be a slice of bytes, a string, an io.Reader or a
// json.Decoder.
func UnmarshalGCPFlavourList(source interface{}) (list *GCPFlavourList, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	var data gcpFlavourListData
	err = decoder.Decode(&data)
	if err != nil {
		return
	}
	list, err = data.unwrap()
	return
}

// wrap is the method used internally to convert a list of values of the
// 'GCP_flavour' value to a JSON document.
func (l *GCPFlavourList) wrap() (data gcpFlavourListData, err error) {
	if l == nil {
		return
	}
	data = make(gcpFlavourListData, len(l.items))
	for i, item := range l.items {
		data[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// list of values of the 'GCP_flavour' type.
func (d gcpFlavourListData) unwrap() (list *GCPFlavourList, err error) {
	if d == nil {
		return
	}
	items := make([]*GCPFlavour, len(d))
	for i, item := range d {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(GCPFlavourList)
	list.items = items
	return
}
