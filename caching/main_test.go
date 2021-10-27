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

package caching

import (
	"testing"

	"github.com/openshift-online/ocm-sdk-go/logging"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint
)

func TestCaching(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Caching")
}

// Logger used for tests:
var logger logging.Logger

var _ = BeforeSuite(func() {
	var err error

	// Create the logger that will be used by all the tests:
	logger, err = logging.NewStdLoggerBuilder().
		Streams(GinkgoWriter, GinkgoWriter).
		Debug(true).
		Build()
	Expect(err).ToNot(HaveOccurred())
})
