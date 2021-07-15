/*
Copyright (c) 2021 Red Hat, Inc.

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

package database

import (
	"reflect"
)

// ErrorCode extracts the error code from the given error returned by a database operation. It uses
// reflection to get the value of the `Code` field in order to avoid having to import the `lib/pq`
// and `pgx` packages. It returns an empty string if the passed error is nil or it hasn't the `Code`
// field.
func ErrorCode(err error) string {
	if err == nil {
		return ""
	}
	value := reflect.ValueOf(err)
	if value.Type().Kind() != reflect.Ptr {
		return ""
	}
	elem := value.Elem()
	if elem.Type().Kind() != reflect.Struct {
		return ""
	}
	field := elem.FieldByName("Code")
	if !field.IsValid() {
		return ""
	}
	if field.Type().Kind() != reflect.String {
		return ""
	}
	return field.String()
}
