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

// This file contains tests for the aliases of the types and functoins that have been
// moved to the logging package.

package sdk

import (
	"time"

	// Never import the logging package here, as that will defeat the purpuse of
	// these tests.

	. "github.com/onsi/ginkgo/v2/dsl/core"             // nolint
	. "github.com/onsi/gomega"                         // nolint
	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

var _ = Describe("Deprecated logging", func() {
	Describe("Interface", func() {
		It("Can be declared", func() {
			var logger Logger
			Expect(logger).To(BeNil())
		})
	})

	Describe("Go implementation", func() {
		It("Can be created", func() {
			var logger Logger
			logger, err := NewGoLoggerBuilder().Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(logger).ToNot(BeNil())
		})
	})

	Describe("Std implementation", func() {
		It("Can be created", func() {
			var logger Logger
			logger, err := NewStdLoggerBuilder().Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(logger).ToNot(BeNil())
		})
	})

	Describe("Glog implementation", func() {
		It("Can be created", func() {
			var logger Logger
			logger, err := NewGlogLoggerBuilder().Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(logger).ToNot(BeNil())
		})
	})

	Describe("Connection", func() {
		It("Can be created with deprecated logger", func() {
			// Create the logger:
			var logger Logger
			logger, err := NewGoLoggerBuilder().Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(logger).ToNot(BeNil())

			// Create the connection:
			token := MakeTokenString("Bearer", 5*time.Minute)
			connection, err := NewConnectionBuilder().
				Logger(logger).
				Tokens(token).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = connection.Close()
				Expect(err).ToNot(HaveOccurred())
			}()
			Expect(connection).ToNot(BeNil())
		})
	})
})
