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

// MarshalSummaryMetrics writes a value of the 'summary_metrics' type to the given writer.
func MarshalSummaryMetrics(object *SummaryMetrics, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeSummaryMetrics(object, stream)
	stream.Flush()
	return stream.Error
}

// writeSummaryMetrics writes a value of the 'summary_metrics' type to the given stream.
func writeSummaryMetrics(object *SummaryMetrics, stream *jsoniter.Stream) {
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
	present_ = object.bitmap_&2 != 0 && object.vector != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("vector")
		writeSummarySampleList(object.vector, stream)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalSummaryMetrics reads a value of the 'summary_metrics' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalSummaryMetrics(source interface{}) (object *SummaryMetrics, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readSummaryMetrics(iterator)
	err = iterator.Error
	return
}

// readSummaryMetrics reads a value of the 'summary_metrics' type from the given iterator.
func readSummaryMetrics(iterator *jsoniter.Iterator) *SummaryMetrics {
	object := &SummaryMetrics{}
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
		case "vector":
			value := readSummarySampleList(iterator)
			object.vector = value
			object.bitmap_ |= 2
		default:
			iterator.ReadAny()
		}
	}
	return object
}
