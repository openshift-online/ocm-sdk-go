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

package v1 // github.com/openshift-online/ocm-sdk-go/webrca/v1

import (
	"io"
	"net/http"
)

func readEventDeleteRequest(request *EventDeleteServerRequest, r *http.Request) error {
	return nil
}
func writeEventDeleteRequest(request *EventDeleteRequest, writer io.Writer) error {
	return nil
}
func readEventDeleteResponse(response *EventDeleteResponse, reader io.Reader) error {
	return nil
}
func writeEventDeleteResponse(response *EventDeleteServerResponse, w http.ResponseWriter) error {
	return nil
}
func readEventGetRequest(request *EventGetServerRequest, r *http.Request) error {
	return nil
}
func writeEventGetRequest(request *EventGetRequest, writer io.Writer) error {
	return nil
}
func readEventGetResponse(response *EventGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalEvent(reader)
	return err
}
func writeEventGetResponse(response *EventGetServerResponse, w http.ResponseWriter) error {
	return MarshalEvent(response.body, w)
}
func readEventUpdateRequest(request *EventUpdateServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalEvent(r.Body)
	return err
}
func writeEventUpdateRequest(request *EventUpdateRequest, writer io.Writer) error {
	return MarshalEvent(request.body, writer)
}
func readEventUpdateResponse(response *EventUpdateResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalEvent(reader)
	return err
}
func writeEventUpdateResponse(response *EventUpdateServerResponse, w http.ResponseWriter) error {
	return MarshalEvent(response.body, w)
}
