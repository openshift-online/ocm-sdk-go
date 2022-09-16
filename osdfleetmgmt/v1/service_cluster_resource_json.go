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

package v1 // github.com/renan-campos/ocm-sdk-go/osdfleetmgmt/v1

import "io"

func writeServiceClusterDeleteRequest(request *ServiceClusterDeleteRequest, writer io.Writer) error {
	return nil
}
func readServiceClusterDeleteResponse(response *ServiceClusterDeleteResponse, reader io.Reader) error {
	return nil
}
func writeServiceClusterGetRequest(request *ServiceClusterGetRequest, writer io.Writer) error {
	return nil
}
func readServiceClusterGetResponse(response *ServiceClusterGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalServiceCluster(reader)
	return err
}
func writeServiceClusterPostRequest(request *ServiceClusterPostRequest, writer io.Writer) error {
	return MarshalServiceClusterRequestPayload(request.request, writer)
}
func readServiceClusterPostResponse(response *ServiceClusterPostResponse, reader io.Reader) error {
	var err error
	response.response, err = UnmarshalServiceCluster(reader)
	return err
}