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

package leadership

import (
	"testing"

	"github.com/go-logr/logr"

	. "github.com/onsi/ginkgo/v2/dsl/core"                // nolint
	. "github.com/onsi/gomega"                            // nolint
	. "github.com/openshift-online/ocm-sdk-go/v2/testing" // nolint

	_ "github.com/lib/pq"
)

func TestLeadership(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Leadership")
}

// logger is the logger that will be used by the tests.
var logger logr.Logger

// dbServer is the database dbServer that will be used to create the databases used by the tests.
var dbServer *DatabaseServer

var _ = BeforeSuite(func() {
	// Create a logger that writes to the Ginkgo stream:
	logger = MakeLogger(GinkgoWriter)

	// Start the database server:
	dbServer = MakeDatabaseServer()
})

var _ = AfterSuite(func() {
	// Stop the database server:
	dbServer.Close()
})
