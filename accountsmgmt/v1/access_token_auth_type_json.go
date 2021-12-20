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

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalAccessTokenAuth writes a value of the 'access_token_auth' type to the given writer.
func MarshalAccessTokenAuth(object *AccessTokenAuth, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeAccessTokenAuth(object, stream)
	stream.Flush()
	return stream.Error
}

// writeAccessTokenAuth writes a value of the 'access_token_auth' type to the given stream.
func writeAccessTokenAuth(object *AccessTokenAuth, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("auth")
		stream.WriteString(object.auth)
		count++
	}
	present_ = object.bitmap_&2 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("email")
		stream.WriteString(object.email)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalAccessTokenAuth reads a value of the 'access_token_auth' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalAccessTokenAuth(source interface{}) (object *AccessTokenAuth, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readAccessTokenAuth(iterator)
	err = iterator.Error
	return
}

// readAccessTokenAuth reads a value of the 'access_token_auth' type from the given iterator.
func readAccessTokenAuth(iterator *jsoniter.Iterator) *AccessTokenAuth {
	object := &AccessTokenAuth{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "auth":
			value := iterator.ReadString()
			object.auth = value
			object.bitmap_ |= 1
		case "email":
			value := iterator.ReadString()
			object.email = value
			object.bitmap_ |= 2
		default:
			iterator.ReadAny()
		}
	}
	return object
}
