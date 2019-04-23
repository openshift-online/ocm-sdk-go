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

package errors // github.com/openshift-online/uhc-sdk-go/pkg/client/errors

import (
	"fmt"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// Error kind is the name of the type used to represent errors.
const ErrorKind = "Error"

// ErrorNilKind is the name of the type used to nil errors.
const ErrorNilKind = "ErrorNil"

// Error represents errors.
type Error struct {
	id     *string
	href   *string
	code   *string
	reason *string
}

// Kind returns the name of the type of the error.
func (e *Error) Kind() string {
	if e == nil {
		return ErrorNilKind
	}
	return ErrorKind
}

// ID returns the identifier of the error.
func (e *Error) ID() string {
	if e != nil && e.id != nil {
		return *e.id
	}
	return ""
}

// GetID returns the identifier of the error and a flag indicating if the
// identifier has a value.
func (e *Error) GetID() (value string, ok bool) {
	ok = e != nil && e.id != nil
	if ok {
		value = *e.id
	}
	return
}

// HREF returns the link to the error.
func (e *Error) HREF() string {
	if e != nil && e.href != nil {
		return *e.href
	}
	return ""
}

// GetHREF returns the link of the error and a flag indicating if the
// link has a value.
func (e *Error) GetHREF() (value string, ok bool) {
	ok = e != nil && e.href != nil
	if ok {
		value = *e.href
	}
	return
}

// Code returns the code of the error.
func (e *Error) Code() string {
	if e != nil && e.code != nil {
		return *e.code
	}
	return ""
}

// GetCode returns the link of the error and a flag indicating if the
// code has a value.
func (e *Error) GetCode() (value string, ok bool) {
	ok = e != nil && e.code != nil
	if ok {
		value = *e.code
	}
	return
}

// Reason returns the reason of the error.
func (e *Error) Reason() string {
	if e != nil && e.reason != nil {
		return *e.reason
	}
	return ""
}

// GetReason returns the link of the error and a flag indicating if the
// reason has a value.
func (e *Error) GetReason() (value string, ok bool) {
	ok = e != nil && e.reason != nil
	if ok {
		value = *e.reason
	}
	return
}

// Error is the implementation of the error interface.
func (e *Error) Error() string {
	if e.reason != nil {
		return *e.reason
	}
	if e.code != nil {
		return *e.code
	}
	if e.id != nil {
		return *e.id
	}
	return "unknown error"
}

// UnmarshalError reads an error from the given which can be an slice of bytes, a
// string, a reader or a JSON decoder.
func UnmarshalError(source interface{}) (object *Error, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(errorData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// errorData is the data structure used internally to marshal and unmarshal errors.
type errorData struct {
	Kind   *string "json:\"kind,omitempty\""
	ID     *string "json:\"id,omitempty\""
	HREF   *string "json:\"href,omitempty\""
	Code   *string "json:\"code,omitempty\""
	Reason *string "json:\"reason,omitempty\""
}

// unwrap is the method used internally to convert the JSON unmarshalled data to an
// error.
func (d *errorData) unwrap() (object *Error, err error) {
	if d == nil {
		return
	}
	object = new(Error)
	if d.Kind != nil && *d.Kind != ErrorKind {
		err = fmt.Errorf(
			"expected kind '%s' but got '%s'",
			ErrorKind, *d.Kind,
		)
		return
	}
	object.id = d.ID
	object.href = d.HREF
	object.code = d.Code
	object.reason = d.Reason
	return
}
