/*
Copyright (c) 2020 Red Hat, Inc.

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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

func writeRolesAddRequest(request *RolesAddRequest, writer io.Writer) error {
	return MarshalRole(request.body, writer)
}
func readRolesAddResponse(response *RolesAddResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalRole(reader)
	return err
}
func writeRolesListRequest(request *RolesListRequest, writer io.Writer) error {
	return nil
}
func readRolesListResponse(response *RolesListResponse, reader io.Reader) error {
	iterator, err := helpers.NewIterator(reader)
	if err != nil {
		return err
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "page":
			value := iterator.ReadInt()
			response.page = &value
		case "size":
			value := iterator.ReadInt()
			response.size = &value
		case "total":
			value := iterator.ReadInt()
			response.total = &value
		case "items":
			items := ReadRoleList(iterator)
			responseItems := &RoleList{}
			responseItems.SetItems(items)
			response.items = responseItems
		default:
			iterator.ReadAny()
		}
	}
	return iterator.Error
}
