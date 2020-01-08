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

// machineTypeListData is type used internally to marshal and unmarshal lists of objects
// of type 'machine_type'.
type machineTypeListData []*machineTypeData

// UnmarshalMachineTypeList reads a list of values of the 'machine_type'
// from the given source, which can be a slice of bytes, a string, an io.Reader or a
// json.Decoder.
func UnmarshalMachineTypeList(source interface{}) (list *MachineTypeList, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	var data machineTypeListData
	err = decoder.Decode(&data)
	if err != nil {
		return
	}
	list, err = data.unwrap()
	return
}

// wrap is the method used internally to convert a list of values of the
// 'machine_type' value to a JSON document.
func (l *MachineTypeList) wrap() (data machineTypeListData, err error) {
	if l == nil {
		return
	}
	data = make(machineTypeListData, len(l.items))
	for i, item := range l.items {
		data[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// list of values of the 'machine_type' type.
func (d machineTypeListData) unwrap() (list *MachineTypeList, err error) {
	if d == nil {
		return
	}
	items := make([]*MachineType, len(d))
	for i, item := range d {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(MachineTypeList)
	list.items = items
	return
}

// machineTypeListLinkData is type used internally to marshal and unmarshal links
// to lists of objects of type 'machine_type'.
type machineTypeListLinkData struct {
	Kind  *string            "json:\"kind,omitempty\""
	HREF  *string            "json:\"href,omitempty\""
	Items []*machineTypeData "json:\"items,omitempty\""
}

// wrapLink is the method used internally to convert a list of values of the
// 'machine_type' value to a link.
func (l *MachineTypeList) wrapLink() (data *machineTypeListLinkData, err error) {
	if l == nil {
		return
	}
	items := make([]*machineTypeData, len(l.items))
	for i, item := range l.items {
		items[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	data = new(machineTypeListLinkData)
	data.Items = items
	data.HREF = l.href
	data.Kind = new(string)
	if l.link {
		*data.Kind = MachineTypeListLinkKind
	} else {
		*data.Kind = MachineTypeListKind
	}
	return
}

// unwrapLink is the function used internally to convert a JSON link to a list
// of values of the 'machine_type' type to a list.
func (d *machineTypeListLinkData) unwrapLink() (list *MachineTypeList, err error) {
	if d == nil {
		return
	}
	items := make([]*MachineType, len(d.Items))
	for i, item := range d.Items {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(MachineTypeList)
	list.items = items
	list.href = d.HREF
	if d.Kind != nil {
		list.link = *d.Kind == MachineTypeListLinkKind
	}
	return
}
