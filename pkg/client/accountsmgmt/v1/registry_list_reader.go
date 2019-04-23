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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

import (
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// registryListData is type used internally to marshal and unmarshal lists of objects
// of type 'registry'.
type registryListData []*registryData

// UnmarshalRegistryList reads a list of values of the 'registry'
// from the given source, which can be a slice of bytes, a string, an io.Reader or a
// json.Decoder.
func UnmarshalRegistryList(source interface{}) (list *RegistryList, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	var data registryListData
	err = decoder.Decode(&data)
	if err != nil {
		return
	}
	list, err = data.unwrap()
	return
}

// wrap is the method used internally to convert a list of values of the
// 'registry' value to a JSON document.
func (o *RegistryList) wrap() (data registryListData, err error) {
	if o == nil {
		return
	}
	data = make(registryListData, len(o.items))
	for i, item := range o.items {
		data[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// list of values of the 'registry' type.
func (d registryListData) unwrap() (list *RegistryList, err error) {
	if d == nil {
		return
	}
	items := make([]*Registry, len(d))
	for i, item := range d {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(RegistryList)
	list.items = items
	return
}
