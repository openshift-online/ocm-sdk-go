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

package v1 // github.com/openshift-online/ocm-sdk-go/jobqueue/v1

import (
	"io"
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

func readQueueGetRequest(request *QueueGetServerRequest, r *http.Request) error {
	return nil
}
func writeQueueGetRequest(request *QueueGetRequest, writer io.Writer) error {
	return nil
}
func readQueueGetResponse(response *QueueGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalQueue(reader)
	return err
}
func writeQueueGetResponse(response *QueueGetServerResponse, w http.ResponseWriter) error {
	return MarshalQueue(response.body, w)
}
func readQueuePopRequest(request *QueuePopServerRequest, r *http.Request) error {
	return nil
}
func writeQueuePopRequest(request *QueuePopRequest, writer io.Writer) error {
	return nil
}
func readQueuePopResponse(response *QueuePopResponse, reader io.Reader) error {
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
			value := readJob(iterator)
			response.body = value
		default:
			iterator.ReadAny()
		}
	}
	return iterator.Error
}
func writeQueuePopResponse(response *QueuePopServerResponse, w http.ResponseWriter) error {
	count := 0
	stream := helpers.NewStream(w)
	stream.WriteObjectStart()
	if response.body != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("body")
		writeJob(response.body, stream)
		count++
	}
	stream.WriteObjectEnd()
	stream.Flush()
	return stream.Error
}
func readQueuePushRequest(request *QueuePushServerRequest, r *http.Request) error {
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
			value := readJob(iterator)
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
func writeQueuePushRequest(request *QueuePushRequest, writer io.Writer) error {
	count := 0
	stream := helpers.NewStream(writer)
	stream.WriteObjectStart()
	if request.body != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("body")
		writeJob(request.body, stream)
		count++
	}
	stream.WriteObjectEnd()
	stream.Flush()
	return stream.Error
}
func readQueuePushResponse(response *QueuePushResponse, reader io.Reader) error {
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
			value := readJob(iterator)
			response.body = value
		default:
			iterator.ReadAny()
		}
	}
	return iterator.Error
}
func writeQueuePushResponse(response *QueuePushServerResponse, w http.ResponseWriter) error {
	count := 0
	stream := helpers.NewStream(w)
	stream.WriteObjectStart()
	if response.body != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("body")
		writeJob(response.body, stream)
		count++
	}
	stream.WriteObjectEnd()
	stream.Flush()
	return stream.Error
}
