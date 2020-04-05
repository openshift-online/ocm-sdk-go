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

// MarshalClusterAlertsFiring writes a value of the 'cluster_alerts_firing' type to the given writer.
func MarshalClusterAlertsFiring(object *ClusterAlertsFiring, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeClusterAlertsFiring(object, stream)
	stream.Flush()
	return stream.Error
}

// writeClusterAlertsFiring writes a value of the 'cluster_alerts_firing' type to the given stream.
func writeClusterAlertsFiring(object *ClusterAlertsFiring, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.critical != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("critical")
		stream.WriteInt(*object.critical)
		count++
	}
	if object.high != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("high")
		stream.WriteInt(*object.high)
		count++
	}
	if object.none != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("none")
		stream.WriteInt(*object.none)
		count++
	}
	if object.warning != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("warning")
		stream.WriteInt(*object.warning)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalClusterAlertsFiring reads a value of the 'cluster_alerts_firing' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalClusterAlertsFiring(source interface{}) (object *ClusterAlertsFiring, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readClusterAlertsFiring(iterator)
	err = iterator.Error
	return
}

// readClusterAlertsFiring reads a value of the 'cluster_alerts_firing' type from the given iterator.
func readClusterAlertsFiring(iterator *jsoniter.Iterator) *ClusterAlertsFiring {
	object := &ClusterAlertsFiring{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "critical":
			value := iterator.ReadInt()
			object.critical = &value
		case "high":
			value := iterator.ReadInt()
			object.high = &value
		case "none":
			value := iterator.ReadInt()
			object.none = &value
		case "warning":
			value := iterator.ReadInt()
			object.warning = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
