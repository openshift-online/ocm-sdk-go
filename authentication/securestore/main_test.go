/*
Copyright (c) 2024 Red Hat, Inc.

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

// This file contains tests for the functions that extract authentication and authorization
// information from contexts.

package securestore

import (
	"testing"

	. "github.com/onsi/ginkgo/v2/dsl/core" // nolint
	. "github.com/onsi/gomega"             // nolint
)

func TestSecurestore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication.Securestore")
}

var _ = Describe("Validation", func() {
	It("Validates an invalid backend", func() {
		err := ValidateBackend("invalid")
		Expect(err).To(Equal(ErrKeyringInvalid))

		err = UpsertConfigToKeyring("invalid", []byte("test"))
		Expect(err).To(Equal(ErrKeyringInvalid))

		bytes, err := GetConfigFromKeyring("invalid")
		Expect(err).To(Equal(ErrKeyringInvalid))
		Expect(bytes).To(BeNil())

		err = RemoveConfigFromKeyring("invalid")
		Expect(err).To(Equal(ErrKeyringInvalid))
	})

	It("Validates an empty backend", func() {
		err := ValidateBackend("")
		Expect(err).To(Equal(ErrKeyringInvalid))

		err = UpsertConfigToKeyring("", []byte("test"))
		Expect(err).To(Equal(ErrKeyringInvalid))

		bytes, err := GetConfigFromKeyring("")
		Expect(err).To(Equal(ErrKeyringInvalid))
		Expect(bytes).To(BeNil())

		err = RemoveConfigFromKeyring("")
		Expect(err).To(Equal(ErrKeyringInvalid))
	})

})
