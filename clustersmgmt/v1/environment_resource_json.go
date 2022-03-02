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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

func readEnvironmentGetRequest(request *EnvironmentGetServerRequest, r *http.Request) error {
	return nil
}
func writeEnvironmentGetRequest(request *EnvironmentGetRequest, writer io.Writer) error {
	return nil
}
func readEnvironmentGetResponse(response *EnvironmentGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalEnvironment(reader)
	return err
}
func writeEnvironmentGetResponse(response *EnvironmentGetServerResponse, w http.ResponseWriter) error {
	return MarshalEnvironment(response.body, w)
}
func readEnvironmentPatchRequest(request *EnvironmentPatchServerRequest, r *http.Request) error {
	iterator, err := helpers.NewIterator(r.Body)
	if err != nil {
		return err
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "body":
			value := readEnvironment(iterator)
			request.body = value
		default:
			iterator.ReadAny()
		}
	}
	err = iterator.Error
	if err != nil {
		return err
	}
	return nil
}
func writeEnvironmentPatchRequest(request *EnvironmentPatchRequest, writer io.Writer) error {
	count := 0
	stream := helpers.NewStream(writer)
	stream.WriteObjectStart()
	if request.body != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("body")
		writeEnvironment(request.body, stream)
		count++
	}
	stream.WriteObjectEnd()
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}
func readEnvironmentPatchResponse(response *EnvironmentPatchResponse, reader io.Reader) error {
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
		case "body":
			value := readEnvironment(iterator)
			response.body = value
		default:
			iterator.ReadAny()
		}
	}
	return iterator.Error
}
func writeEnvironmentPatchResponse(response *EnvironmentPatchServerResponse, w http.ResponseWriter) error {
	count := 0
	stream := helpers.NewStream(w)
	stream.WriteObjectStart()
	if response.body != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("body")
		writeEnvironment(response.body, stream)
		count++
	}
	stream.WriteObjectEnd()
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}
