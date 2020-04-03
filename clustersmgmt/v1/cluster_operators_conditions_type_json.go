/*
Copyright (c) 2019 Red Hat, Inc.

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

// MarshalClusterOperatorsConditions writes a value of the 'cluster_operators_conditions' type to the given writer.
func MarshalClusterOperatorsConditions(object *ClusterOperatorsConditions, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeClusterOperatorsConditions(object, stream)
	stream.Flush()
	return stream.Error
}

// writeClusterOperatorsConditions writes a value of the 'cluster_operators_conditions' type to the given stream.
func writeClusterOperatorsConditions(object *ClusterOperatorsConditions, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.available != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("available")
		stream.WriteInt(*object.available)
		count++
	}
	if object.degraded != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("degraded")
		stream.WriteInt(*object.degraded)
		count++
	}
	if object.failing != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("failing")
		stream.WriteInt(*object.failing)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalClusterOperatorsConditions reads a value of the 'cluster_operators_conditions' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalClusterOperatorsConditions(source interface{}) (object *ClusterOperatorsConditions, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readClusterOperatorsConditions(iterator)
	err = iterator.Error
	return
}

// readClusterOperatorsConditions reads a value of the 'cluster_operators_conditions' type from the given iterator.
func readClusterOperatorsConditions(iterator *jsoniter.Iterator) *ClusterOperatorsConditions {
	object := &ClusterOperatorsConditions{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "available":
			value := iterator.ReadInt()
			object.available = &value
		case "degraded":
			value := iterator.ReadInt()
			object.degraded = &value
		case "failing":
			value := iterator.ReadInt()
			object.failing = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
