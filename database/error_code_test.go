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
	"io"

	. "github.com/onsi/ginkgo/extensions/table" // nolint
	. "github.com/onsi/gomega"                  // nolint

	"github.com/jackc/pgconn"
	"github.com/lib/pq"
)

var _ = DescribeTable(
	"Error code extraction",
	func(err error, expected string) {
		actual := ErrorCode(err)
		Expect(actual).To(Equal(expected))
	},
	Entry(
		"Nil",
		nil,
		"",
	),
	Entry(
		"From 'pq'",
		&pq.Error{
			Code: "123",
		},
		"123",
	),
	Entry(
		"From 'pgconn'",
		&pgconn.PgError{
			Code: "123",
		},
		"123",
	),
	Entry(
		"Non SQL error",
		io.EOF,
		"",
	),
)
