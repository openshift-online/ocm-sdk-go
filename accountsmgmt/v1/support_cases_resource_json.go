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
	"net/http"
)

func readSupportCasesAddRequest(request *SupportCasesAddServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalSupportCase(r.Body)
	return err
}
func writeSupportCasesAddRequest(request *SupportCasesAddRequest, writer io.Writer) error {
	return MarshalSupportCase(request.body, writer)
}
func readSupportCasesAddResponse(response *SupportCasesAddResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalSupportCase(reader)
	return err
}
func writeSupportCasesAddResponse(response *SupportCasesAddServerResponse, w http.ResponseWriter) error {
	return MarshalSupportCase(response.body, w)
}
func readSupportCasesDeleteRequest(request *SupportCasesDeleteServerRequest, r *http.Request) error {
	return nil
}
func writeSupportCasesDeleteRequest(request *SupportCasesDeleteRequest, writer io.Writer) error {
	return nil
}
func readSupportCasesDeleteResponse(response *SupportCasesDeleteResponse, reader io.Reader) error {
	return nil
}
func writeSupportCasesDeleteResponse(response *SupportCasesDeleteServerResponse, w http.ResponseWriter) error {
	return nil
}
