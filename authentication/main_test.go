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

package authentication

import (
	"os"
	"testing"

	"github.com/renan-campos/ocm-sdk-go/logging"

	. "github.com/onsi/ginkgo/v2/dsl/core"             // nolint
	. "github.com/onsi/gomega"                         // nolint
	. "github.com/renan-campos/ocm-sdk-go/testing" // nolint
)

func TestAuthentication(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication")
}

// Logger used for tests:
var logger logging.Logger

// JSON web key set used for tests:
var keysBytes []byte
var keysFile string

var _ = BeforeSuite(func() {
	var err error

	// Create a temporary file containing the JSON web key set:
	keysBytes = DefaultJWKS()
	keysFD, err := os.CreateTemp("", "jwks-*.json")
	Expect(err).ToNot(HaveOccurred())
	_, err = keysFD.Write(keysBytes)
	Expect(err).ToNot(HaveOccurred())
	err = keysFD.Close()
	Expect(err).ToNot(HaveOccurred())
	keysFile = keysFD.Name()

	// Create the logger that will be used by all the tests:
	logger, err = logging.NewStdLoggerBuilder().
		Streams(GinkgoWriter, GinkgoWriter).
		Debug(true).
		Build()
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	// Delete the temporary files:
	err := os.Remove(keysFile)
	Expect(err).ToNot(HaveOccurred())
})
