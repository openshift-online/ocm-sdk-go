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

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalActionList writes a list of values of the 'action' type to
// the given writer.
func MarshalActionList(list []Action, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeActionList(list, stream)
	stream.Flush()
	return stream.Error
}

// writeActionList writes a list of value of the 'action' type to
// the given stream.
func writeActionList(list []Action, stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	for i, value := range list {
		if i > 0 {
			stream.WriteMore()
		}
		stream.WriteString(string(value))
	}
	stream.WriteArrayEnd()
}

// UnmarshalActionList reads a list of values of the 'action' type
// from the given source, which can be a slice of bytes, a string or a reader.
func UnmarshalActionList(source interface{}) (items []Action, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	items = readActionList(iterator)
	err = iterator.Error
	return
}

// readActionList reads list of values of the ''action' type from
// the given iterator.
func readActionList(iterator *jsoniter.Iterator) []Action {
	list := []Action{}
	for iterator.ReadArray() {
		text := iterator.ReadString()
		item := Action(text)
		list = append(list, item)
	}
	return list
}
