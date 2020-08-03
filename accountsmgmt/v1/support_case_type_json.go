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

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalSupportCase writes a value of the 'support_case' type to the given writer.
func MarshalSupportCase(object *SupportCase, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeSupportCase(object, stream)
	stream.Flush()
	return stream.Error
}

// writeSupportCase writes a value of the 'support_case' type to the given stream.
func writeSupportCase(object *SupportCase, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if count > 0 {
		stream.WriteMore()
	}
	stream.WriteObjectField("kind")
	if object.link {
		stream.WriteString(SupportCaseLinkKind)
	} else {
		stream.WriteString(SupportCaseKind)
	}
	count++
	if object.id != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(*object.id)
		count++
	}
	if object.href != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(*object.href)
		count++
	}
	if object.clusterUuid != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster_uuid")
		stream.WriteString(*object.clusterUuid)
		count++
	}
	if object.description != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("description")
		stream.WriteString(*object.description)
		count++
	}
	if object.eventStreamId != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("event_stream_id")
		stream.WriteString(*object.eventStreamId)
		count++
	}
	if object.severity != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("severity")
		stream.WriteString(*object.severity)
		count++
	}
	if object.summary != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("summary")
		stream.WriteString(*object.summary)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalSupportCase reads a value of the 'support_case' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalSupportCase(source interface{}) (object *SupportCase, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readSupportCase(iterator)
	err = iterator.Error
	return
}

// readSupportCase reads a value of the 'support_case' type from the given iterator.
func readSupportCase(iterator *jsoniter.Iterator) *SupportCase {
	object := &SupportCase{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			object.link = value == SupportCaseLinkKind
		case "id":
			value := iterator.ReadString()
			object.id = &value
		case "href":
			value := iterator.ReadString()
			object.href = &value
		case "cluster_uuid":
			value := iterator.ReadString()
			object.clusterUuid = &value
		case "description":
			value := iterator.ReadString()
			object.description = &value
		case "event_stream_id":
			value := iterator.ReadString()
			object.eventStreamId = &value
		case "severity":
			value := iterator.ReadString()
			object.severity = &value
		case "summary":
			value := iterator.ReadString()
			object.summary = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
