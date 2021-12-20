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
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalProvisionShard writes a value of the 'provision_shard' type to the given writer.
func MarshalProvisionShard(object *ProvisionShard, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeProvisionShard(object, stream)
	stream.Flush()
	return stream.Error
}

// writeProvisionShard writes a value of the 'provision_shard' type to the given stream.
func writeProvisionShard(object *ProvisionShard, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if object.bitmap_&1 != 0 {
		stream.WriteString(ProvisionShardLinkKind)
	} else {
		stream.WriteString(ProvisionShardKind)
	}
	count++
	if object.bitmap_&2 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if object.bitmap_&4 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = object.bitmap_&8 != 0 && object.awsAccountOperatorConfig != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("aws_account_operator_config")
		writeServerConfig(object.awsAccountOperatorConfig, stream)
		count++
	}
	present_ = object.bitmap_&16 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("aws_base_domain")
		stream.WriteString(object.awsBaseDomain)
		count++
	}
	present_ = object.bitmap_&32 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("gcp_base_domain")
		stream.WriteString(object.gcpBaseDomain)
		count++
	}
	present_ = object.bitmap_&64 != 0 && object.gcpProjectOperator != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("gcp_project_operator")
		writeServerConfig(object.gcpProjectOperator, stream)
		count++
	}
	present_ = object.bitmap_&128 != 0 && object.hiveConfig != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("hive_config")
		writeServerConfig(object.hiveConfig, stream)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalProvisionShard reads a value of the 'provision_shard' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalProvisionShard(source interface{}) (object *ProvisionShard, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readProvisionShard(iterator)
	err = iterator.Error
	return
}

// readProvisionShard reads a value of the 'provision_shard' type from the given iterator.
func readProvisionShard(iterator *jsoniter.Iterator) *ProvisionShard {
	object := &ProvisionShard{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == ProvisionShardLinkKind {
				object.bitmap_ |= 1
			}
		case "id":
			object.id = iterator.ReadString()
			object.bitmap_ |= 2
		case "href":
			object.href = iterator.ReadString()
			object.bitmap_ |= 4
		case "aws_account_operator_config":
			value := readServerConfig(iterator)
			object.awsAccountOperatorConfig = value
			object.bitmap_ |= 8
		case "aws_base_domain":
			value := iterator.ReadString()
			object.awsBaseDomain = value
			object.bitmap_ |= 16
		case "gcp_base_domain":
			value := iterator.ReadString()
			object.gcpBaseDomain = value
			object.bitmap_ |= 32
		case "gcp_project_operator":
			value := readServerConfig(iterator)
			object.gcpProjectOperator = value
			object.bitmap_ |= 64
		case "hive_config":
			value := readServerConfig(iterator)
			object.hiveConfig = value
			object.bitmap_ |= 128
		default:
			iterator.ReadAny()
		}
	}
	return object
}
