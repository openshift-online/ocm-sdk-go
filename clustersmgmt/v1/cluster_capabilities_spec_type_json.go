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

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalClusterCapabilitiesSpec writes a value of the 'cluster_capabilities_spec' type to the given writer.
func MarshalClusterCapabilitiesSpec(object *ClusterCapabilitiesSpec, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteClusterCapabilitiesSpec(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteClusterCapabilitiesSpec writes a value of the 'cluster_capabilities_spec' type to the given stream.
func WriteClusterCapabilitiesSpec(object *ClusterCapabilitiesSpec, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0 && object.disabled != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("disabled")
		WriteOptionalClusterCapabilityList(object.disabled, stream)
	}
	stream.WriteObjectEnd()
}

// UnmarshalClusterCapabilitiesSpec reads a value of the 'cluster_capabilities_spec' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalClusterCapabilitiesSpec(source interface{}) (object *ClusterCapabilitiesSpec, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadClusterCapabilitiesSpec(iterator)
	err = iterator.Error
	return
}

// ReadClusterCapabilitiesSpec reads a value of the 'cluster_capabilities_spec' type from the given iterator.
func ReadClusterCapabilitiesSpec(iterator *jsoniter.Iterator) *ClusterCapabilitiesSpec {
	object := &ClusterCapabilitiesSpec{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "disabled":
			value := ReadOptionalClusterCapabilityList(iterator)
			object.disabled = value
			object.bitmap_ |= 1
		default:
			iterator.ReadAny()
		}
	}
	return object
}
