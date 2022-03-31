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

func readIncidentDeleteRequest(request *IncidentDeleteServerRequest, r *http.Request) error {
	return nil
}
func writeIncidentDeleteRequest(request *IncidentDeleteRequest, writer io.Writer) error {
	return nil
}
func readIncidentDeleteResponse(response *IncidentDeleteResponse, reader io.Reader) error {
	return nil
}
func writeIncidentDeleteResponse(response *IncidentDeleteServerResponse, w http.ResponseWriter) error {
	return nil
}
func readIncidentGetRequest(request *IncidentGetServerRequest, r *http.Request) error {
	return nil
}
func writeIncidentGetRequest(request *IncidentGetRequest, writer io.Writer) error {
	return nil
}
func readIncidentGetResponse(response *IncidentGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalIncident(reader)
	return err
}
func writeIncidentGetResponse(response *IncidentGetServerResponse, w http.ResponseWriter) error {
	return MarshalIncident(response.body, w)
}
func readIncidentUpdateRequest(request *IncidentUpdateServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalIncident(r.Body)
	return err
}
func writeIncidentUpdateRequest(request *IncidentUpdateRequest, writer io.Writer) error {
	return MarshalIncident(request.body, writer)
}
func readIncidentUpdateResponse(response *IncidentUpdateResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalIncident(reader)
	return err
}
func writeIncidentUpdateResponse(response *IncidentUpdateServerResponse, w http.ResponseWriter) error {
	return MarshalIncident(response.body, w)
}
