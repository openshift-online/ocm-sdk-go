//go:build windows

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

package securestore

import (
	"time"

	. "github.com/onsi/ginkgo/v2" // nolint
	. "github.com/onsi/gomega"    // nolint

	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

var _ = Describe("Wincred Keyring", func() {
	const backend = "wincred"

	BeforeEach(func() {
		err := RemoveConfigFromKeyring(backend)
		Expect(err).To(BeNil())
	})

	When("Listing Keyrings", func() {
		It("Lists wincred as a valid keyring", func() {
			backends := AvailableBackends()
			Expect(backends).To(ContainElement(backend))
		})
	})

	When("Using wincred", func() {
		It("Stores/Removes via wincred", func() {
			// Create the token
			accessToken := MakeTokenString("Bearer", 15*time.Minute)

			// Run insert
			err := UpsertConfigToKeyring(backend, []byte(accessToken))

			Expect(err).To(BeNil())

			// Check the content of the keyring
			result, err := GetConfigFromKeyring(backend)
			Expect(result).To(Equal([]byte(accessToken)))
			Expect(err).To(BeNil())

			// Remove the configuration from the keyring
			err = RemoveConfigFromKeyring(backend)
			Expect(err).To(BeNil())

			// Ensure the keyring is empty
			result, err = GetConfigFromKeyring(backend)
			Expect(result).To(Equal([]byte("")))
			Expect(err).To(BeNil())
		})
	})
})
