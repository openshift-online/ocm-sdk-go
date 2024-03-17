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

// MarshalDefaultCapabilities writes a value of the 'default_capabilities' type to the given writer.
func MarshalDefaultCapabilities(object *DefaultCapabilities, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeDefaultCapabilities(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeDefaultCapabilities writes a value of the 'default_capabilities' type to the given stream.
func writeDefaultCapabilities(object *DefaultCapabilities, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("name")
		stream.WriteString(object.name)
		count++
	}
	present_ = object.bitmap_&2 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("value")
		stream.WriteString(object.value)
	}
	stream.WriteObjectEnd()
}

// UnmarshalDefaultCapabilities reads a value of the 'default_capabilities' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalDefaultCapabilities(source interface{}) (object *DefaultCapabilities, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readDefaultCapabilities(iterator)
	err = iterator.Error
	return
}

// readDefaultCapabilities reads a value of the 'default_capabilities' type from the given iterator.
func readDefaultCapabilities(iterator *jsoniter.Iterator) *DefaultCapabilities {
	object := &DefaultCapabilities{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "name":
			value := iterator.ReadString()
			object.name = value
			object.bitmap_ |= 1
		case "value":
			value := iterator.ReadString()
			object.value = value
			object.bitmap_ |= 2
		default:
			iterator.ReadAny()
		}
	}
	return object
}
