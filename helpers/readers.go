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

package helpers // github.com/openshift-online/ocm-sdk-go/helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// NewEncoder creates a new JSON encoder from the given target. The target can be a
// a writer or a JSON encoder.
func NewEncoder(target interface{}) (encoder *json.Encoder, err error) {
	switch output := target.(type) {
	case io.Writer:
		encoder = json.NewEncoder(output)
	case *json.Encoder:
		encoder = output
	default:
		err = fmt.Errorf(
			"expected writer or JSON decoder, but got %T",
			output,
		)
	}
	return
}

// NewDecoder creates a new JSON decoder from the given source. The source can be a
// slice of bytes, a string, a reader or a JSON decoder.
func NewDecoder(source interface{}) (decoder *json.Decoder, err error) {
	switch input := source.(type) {
	case []byte:
		decoder = json.NewDecoder(bytes.NewBuffer(input))
	case string:
		decoder = json.NewDecoder(bytes.NewBufferString(input))
	case io.Reader:
		decoder = json.NewDecoder(input)
	case *json.Decoder:
		decoder = input
	default:
		err = fmt.Errorf(
			"expected bytes, string, reader or JSON decoder, but got %T",
			input,
		)
	}
	return
}
