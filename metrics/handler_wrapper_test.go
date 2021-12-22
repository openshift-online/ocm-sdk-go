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

// This file contains tests for the metrics transport wrapper.

package metrics

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"    // nolint
	. "github.com/onsi/gomega"       // nolint
	. "github.com/onsi/gomega/ghttp" // nolint

	. "github.com/openshift-online/ocm-sdk-go/testing"
)

var _ = Describe("Create", func() {
	It("Can't be created without a subsystem", func() {
		wrapper, err := NewHandlerWrapper().
			Build()
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("subsystem"))
		Expect(message).To(ContainSubstring("mandatory"))
	})
})

var _ = Describe("Metrics", func() {
	var (
		server  *MetricsServer
		wrapper *HandlerWrapper
		handler http.Handler
	)

	BeforeEach(func() {
		var err error

		// Start the metrics server:
		server = NewMetricsServer()

		// Create the wrapper:
		wrapper, err = NewHandlerWrapper().
			Path("/my/path").
			Subsystem("my").
			Registerer(server.Registry()).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		// Stop the metrics server:
		server.Close()
	})

	// Get sends a GET request to the API server.
	var Send = func(method, path string) {
		request := httptest.NewRequest(method, "http://localhost"+path, nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
	}

	It("Calls wrapped handler", func() {
		// Prepare the handler:
		called := false
		handler = wrapper.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
		}))

		// Send the request:
		Send(http.MethodGet, "/api")

		// Check that the wrapped handler was called:
		Expect(called).To(BeTrue())
	})

	Describe("Request count", func() {
		It("Honours subsystem", func() {
			// Prepare the handler:
			handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

			// Send the request:
			Send(http.MethodGet, "/api")

			// Verify the metrics:
			metrics := server.Metrics()
			Expect(metrics).To(MatchLine(`^my_request_count\{.*\} .*$`))
		})

		DescribeTable(
			"Counts correctly",
			func(count int) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

				// Send the requests:
				for i := 0; i < count; i++ {
					Send(http.MethodGet, "/api")
				}

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_count\{.*\} %d$`, count))
			},
			Entry("One", 1),
			Entry("Two", 2),
			Entry("Trhee", 3),
		)

		DescribeTable(
			"Includes method label",
			func(method string) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

				// Send the requests:
				Send(method, "/api")

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_count\{.*method="%s".*\} .*$`, method))
			},
			Entry("GET", http.MethodGet),
			Entry("POST", http.MethodPost),
			Entry("PATCH", http.MethodPatch),
			Entry("DELETE", http.MethodDelete),
		)

		DescribeTable(
			"Includes path label",
			func(path, label string) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

				// Send the requests:
				Send(http.MethodGet, path)

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_count\{.*path="%s".*\} .*$`, label))
			},
			Entry(
				"Empty",
				"",
				"/-",
			),
			Entry(
				"One slash",
				"/",
				"/-",
			),
			Entry(
				"Two slashes",
				"//",
				"/-",
			),
			Entry(
				"Tree slashes",
				"///",
				"/-",
			),
			Entry(
				"API root",
				"/api",
				"/api",
			),
			Entry(
				"API root with trailing slash",
				"/api/",
				"/api",
			),
			Entry(
				"Unknown root",
				"/junk/",
				"/-",
			),
			Entry(
				"Service root",
				"/api/clusters_mgmt",
				"/api/clusters_mgmt",
			),
			Entry(
				"Unknown service root",
				"/api/junk",
				"/-",
			),
			Entry(
				"Version root",
				"/api/clusters_mgmt/v1",
				"/api/clusters_mgmt/v1",
			),
			Entry(
				"Unknown version root",
				"/api/junk/v1",
				"/-",
			),
			Entry(
				"Collection",
				"/api/clusters_mgmt/v1/clusters",
				"/api/clusters_mgmt/v1/clusters",
			),
			Entry(
				"Unknown collection",
				"/api/clusters_mgmt/v1/junk",
				"/-",
			),
			Entry(
				"Collection item",
				"/api/clusters_mgmt/v1/clusters/123",
				"/api/clusters_mgmt/v1/clusters/-",
			),
			Entry(
				"Collection item action",
				"/api/clusters_mgmt/v1/clusters/123/hibernate",
				"/api/clusters_mgmt/v1/clusters/-/hibernate",
			),
			Entry(
				"Unknown collection item action",
				"/api/clusters_mgmt/v1/clusters/123/junk",
				"/-",
			),
			Entry(
				"Subcollection",
				"/api/clusters_mgmt/v1/clusters/123/groups",
				"/api/clusters_mgmt/v1/clusters/-/groups",
			),
			Entry(
				"Unknown subcollection",
				"/api/clusters_mgmt/v1/clusters/123/junks",
				"/-",
			),
			Entry(
				"Subcollection item",
				"/api/clusters_mgmt/v1/clusters/123/groups/456",
				"/api/clusters_mgmt/v1/clusters/-/groups/-",
			),
			Entry(
				"Too long",
				"/api/clusters_mgmt/v1/clusters/123/groups/456/junk",
				"/-",
			),
			Entry(
				"Explicitly specified path",
				"/my/path",
				"/my/path",
			),
			Entry(
				"Unknown path",
				"/your/path",
				"/-",
			),
		)

		DescribeTable(
			"Includes code label",
			func(code int) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(code, nil))

				// Send the requests:
				Send(http.MethodGet, "/api")

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_count\{.*code="%d".*\} .*$`, code))
			},
			Entry("200", http.StatusOK),
			Entry("201", http.StatusCreated),
			Entry("202", http.StatusAccepted),
			Entry("401", http.StatusUnauthorized),
			Entry("404", http.StatusNotFound),
			Entry("500", http.StatusInternalServerError),
		)

		DescribeTable(
			"Includes API service label",
			func(path, label string) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

				// Send the requests:
				Send(http.MethodGet, path)

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_count\{.*apiservice="%s".*\} .*$`, label))
			},
			Entry(
				"Empty",
				"",
				"",
			),
			Entry(
				"Root",
				"/",
				"",
			),
			Entry(
				"Clusters root",
				"/api/clusters_mgmt",
				"ocm-clusters-service",
			),
			Entry(
				"Clusters version",
				"/api/clusters_mgmt/v1",
				"ocm-clusters-service",
			),
			Entry(
				"Clusters collection",
				"/api/clusters_mgmt/v1/clusters",
				"ocm-clusters-service",
			),
			Entry(
				"Clusters item",
				"/api/clusters_mgmt/v1/clusters/123",
				"ocm-clusters-service",
			),
			Entry(
				"Accounts root",
				"/api/accounts_mgmt",
				"ocm-accounts-service",
			),
			Entry(
				"Accounts version",
				"/api/accounts_mgmt/v1",
				"ocm-accounts-service",
			),
			Entry(
				"Accounts collection",
				"/api/accounts_mgmt/v1/accounts",
				"ocm-accounts-service",
			),
			Entry(
				"Accounts item",
				"/api/accounts_mgmt/v1/accounts/123",
				"ocm-accounts-service",
			),
			Entry(
				"Logs root",
				"/api/service_logs",
				"ocm-logs-service",
			),
			Entry(
				"Logs version",
				"/api/service_logs/v1",
				"ocm-logs-service",
			),
			Entry(
				"Logs collection",
				"/api/service_logs/v1/accounts",
				"ocm-logs-service",
			),
			Entry(
				"Logs item",
				"/api/service_logs/v1/accounts/123",
				"ocm-logs-service",
			),
		)
	})

	Describe("Request duration", func() {
		It("Honours subsystem", func() {
			// Prepare the handler:
			handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

			// Send the request:
			Send(http.MethodGet, "/api")

			// Verify the metrics:
			metrics := server.Metrics()
			Expect(metrics).To(MatchLine(`^my_request_duration_bucket\{.*\} .*$`))
			Expect(metrics).To(MatchLine(`^my_request_duration_sum\{.*\} .*$`))
			Expect(metrics).To(MatchLine(`^my_request_duration_count\{.*\} .*$`))
		})

		It("Honours buckets", func() {
			// Prepare the handler:
			handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

			// Send the request:
			Send(http.MethodGet, "/api")

			// Verify the metrics:
			metrics := server.Metrics()
			Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*,le="0.1"\} .*$`))
			Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*,le="1"\} .*$`))
			Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*,le="10"\} .*$`))
			Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*,le="30"\} .*$`))
			Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*,le="\+Inf"\} .*$`))
		})

		DescribeTable(
			"Counts correctly",
			func(count int) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

				// Send the requests:
				for i := 0; i < count; i++ {
					Send(http.MethodGet, "/api")
				}

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_duration_count\{.*\} %d$`, count))
			},
			Entry("One", 1),
			Entry("Two", 2),
			Entry("Trhee", 3),
		)

		DescribeTable(
			"Includes method label",
			func(method string) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

				// Send the requests:
				Send(method, "/api")

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*method="%s".*\} .*$`, method))
				Expect(metrics).To(MatchLine(`^\w+_request_duration_sum\{.*method="%s".*\} .*$`, method))
				Expect(metrics).To(MatchLine(`^\w+_request_duration_count\{.*method="%s".*\} .*$`, method))
			},
			Entry("GET", http.MethodGet),
			Entry("POST", http.MethodPost),
			Entry("PATCH", http.MethodPatch),
			Entry("DELETE", http.MethodDelete),
		)

		DescribeTable(
			"Includes path label",
			func(path, label string) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

				// Send the requests:
				Send(http.MethodGet, path)

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*path="%s".*\} .*$`, label))
				Expect(metrics).To(MatchLine(`^\w+_request_duration_sum\{.*path="%s".*\} .*$`, label))
				Expect(metrics).To(MatchLine(`^\w+_request_duration_count\{.*path="%s".*\} .*$`, label))
			},
			Entry(
				"Empty",
				"",
				"/-",
			),
			Entry(
				"One slash",
				"/",
				"/-",
			),
			Entry(
				"Two slashes",
				"//",
				"/-",
			),
			Entry(
				"Tree slashes",
				"///",
				"/-",
			),
			Entry(
				"API root",
				"/api",
				"/api",
			),
			Entry(
				"API root with trailing slash",
				"/api/",
				"/api",
			),
			Entry(
				"Unknown root",
				"/junk/",
				"/-",
			),
			Entry(
				"Service root",
				"/api/clusters_mgmt",
				"/api/clusters_mgmt",
			),
			Entry(
				"Unknown service root",
				"/api/junk",
				"/-",
			),
			Entry(
				"Version root",
				"/api/clusters_mgmt/v1",
				"/api/clusters_mgmt/v1",
			),
			Entry(
				"Unknown version root",
				"/api/junk/v1",
				"/-",
			),
			Entry(
				"Collection",
				"/api/clusters_mgmt/v1/clusters",
				"/api/clusters_mgmt/v1/clusters",
			),
			Entry(
				"Unknown collection",
				"/api/clusters_mgmt/v1/junk",
				"/-",
			),
			Entry(
				"Collection item",
				"/api/clusters_mgmt/v1/clusters/123",
				"/api/clusters_mgmt/v1/clusters/-",
			),
			Entry(
				"Collection item action",
				"/api/clusters_mgmt/v1/clusters/123/hibernate",
				"/api/clusters_mgmt/v1/clusters/-/hibernate",
			),
			Entry(
				"Unknown collection item action",
				"/api/clusters_mgmt/v1/clusters/123/junk",
				"/-",
			),
			Entry(
				"Subcollection",
				"/api/clusters_mgmt/v1/clusters/123/groups",
				"/api/clusters_mgmt/v1/clusters/-/groups",
			),
			Entry(
				"Unknown subcollection",
				"/api/clusters_mgmt/v1/clusters/123/junks",
				"/-",
			),
			Entry(
				"Subcollection item",
				"/api/clusters_mgmt/v1/clusters/123/groups/456",
				"/api/clusters_mgmt/v1/clusters/-/groups/-",
			),
			Entry(
				"Too long",
				"/api/clusters_mgmt/v1/clusters/123/groups/456/junk",
				"/-",
			),
			Entry(
				"Explicitly specified path",
				"/my/path",
				"/my/path",
			),
			Entry(
				"Unknown path",
				"/your/path",
				"/-",
			),
		)

		DescribeTable(
			"Includes code label",
			func(code int) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(code, nil))

				// Send the requests:
				Send(http.MethodGet, "/api")

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*code="%d".*\} .*$`, code))
				Expect(metrics).To(MatchLine(`^\w+_request_duration_sum\{.*code="%d".*\} .*$`, code))
				Expect(metrics).To(MatchLine(`^\w+_request_duration_count\{.*code="%d".*\} .*$`, code))
			},
			Entry("200", http.StatusOK),
			Entry("201", http.StatusCreated),
			Entry("202", http.StatusAccepted),
			Entry("401", http.StatusUnauthorized),
			Entry("404", http.StatusNotFound),
			Entry("500", http.StatusInternalServerError),
		)

		DescribeTable(
			"Includes API service label",
			func(path, label string) {
				// Prepare the handler:
				handler = wrapper.Wrap(RespondWith(http.StatusOK, nil))

				// Send the requests:
				Send(http.MethodGet, path)

				// Verify the metrics:
				metrics := server.Metrics()
				Expect(metrics).To(MatchLine(`^\w+_request_duration_bucket\{.*apiservice="%s".*\} .*$`, label))
				Expect(metrics).To(MatchLine(`^\w+_request_duration_sum\{.*apiservice="%s".*\} .*$`, label))
				Expect(metrics).To(MatchLine(`^\w+_request_duration_count\{.*apiservice="%s".*\} .*$`, label))
			},
			Entry(
				"Empty",
				"",
				"",
			),
			Entry(
				"Root",
				"/",
				"",
			),
			Entry(
				"Clusters root",
				"/api/clusters_mgmt",
				"ocm-clusters-service",
			),
			Entry(
				"Clusters version",
				"/api/clusters_mgmt/v1",
				"ocm-clusters-service",
			),
			Entry(
				"Clusters collection",
				"/api/clusters_mgmt/v1/clusters",
				"ocm-clusters-service",
			),
			Entry(
				"Clusters item",
				"/api/clusters_mgmt/v1/clusters/123",
				"ocm-clusters-service",
			),
			Entry(
				"Accounts root",
				"/api/accounts_mgmt",
				"ocm-accounts-service",
			),
			Entry(
				"Accounts version",
				"/api/accounts_mgmt/v1",
				"ocm-accounts-service",
			),
			Entry(
				"Accounts collection",
				"/api/accounts_mgmt/v1/accounts",
				"ocm-accounts-service",
			),
			Entry(
				"Accounts item",
				"/api/accounts_mgmt/v1/accounts/123",
				"ocm-accounts-service",
			),
			Entry(
				"Logs root",
				"/api/service_logs",
				"ocm-logs-service",
			),
			Entry(
				"Logs version",
				"/api/service_logs/v1",
				"ocm-logs-service",
			),
			Entry(
				"Logs collection",
				"/api/service_logs/v1/accounts",
				"ocm-logs-service",
			),
			Entry(
				"Logs item",
				"/api/service_logs/v1/accounts/123",
				"ocm-logs-service",
			),
		)
	})
})
