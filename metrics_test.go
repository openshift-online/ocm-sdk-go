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

// This file contains tests for the Prometheus metrics.

package sdk

import (
	"net/http"
	"time"

	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo/v2/dsl/core"                // nolint
	. "github.com/onsi/gomega"                            // nolint
	. "github.com/openshift-online/ocm-sdk-go/v2/testing" // nolint
)

var _ = Describe("Metrics enabled", func() {
	// Servers used during the tests:
	var oidServer *ghttp.Server
	var apiServer *ghttp.Server
	var metricsServer *MetricsServer

	// Connection used during the tests:
	var connection *Connection

	BeforeEach(func() {
		var err error

		// Create the tokens:
		accessToken := MakeTokenString("Bearer", 5*time.Minute)
		refreshToken := MakeTokenString("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer = MakeTCPServer()
		oidServer.AppendHandlers(
			ghttp.CombineHandlers(
				RespondWithAccessAndRefreshTokens(accessToken, refreshToken),
			),
		)

		// Create the API server:
		apiServer = MakeTCPServer()
		apiServer.AppendHandlers(
			RespondWithJSON(http.StatusOK, ""),
		)

		// Create the metrics server:
		metricsServer = NewMetricsServer()

		// Create the connection:
		connection, err = NewConnectionBuilder().
			Logger(logger).
			URL(apiServer.URL()).
			TokenURL(oidServer.URL() + "/my_token").
			Tokens(refreshToken).
			MetricsSubsystem("my").
			MetricsRegisterer(metricsServer.Registry()).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		// Stop the servers:
		oidServer.Close()
		apiServer.Close()
		metricsServer.Close()

		// Close the connection:
		err := connection.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Generates request count for raw request", func() {
		// Send the request:
		_, err := connection.Get().
			Path("/api/clusters_mgmt/v1/clusters/123").
			Send()
		Expect(err).ToNot(HaveOccurred())

		// Verify the metrics:
		metrics := metricsServer.Metrics()
		Expect(metrics).To(MatchLine(`^my_request_count\{.*path="/api/clusters_mgmt/v1/clusters/-".*\} .*$`))
	})

	It("Generates request count for type safe request", func() {
		// Send the request:
		_, err := connection.ClustersMgmt().V1().Clusters().Cluster("123").Get().
			Send()
		Expect(err).ToNot(HaveOccurred())

		// Verify the metrics:
		metrics := metricsServer.Metrics()
		Expect(metrics).To(MatchLine(`^my_request_count\{.*path="/api/clusters_mgmt/v1/clusters/-".*\} .*$`))
	})

	It("Generates token request count", func() {
		// Send the request:
		_, err := connection.ClustersMgmt().V1().Clusters().Cluster("123").Get().
			Send()
		Expect(err).ToNot(HaveOccurred())

		// Verify the metrics:
		metrics := metricsServer.Metrics()
		Expect(metrics).To(MatchLine(`^my_request_count\{.*path="/my_token".*\} .*$`))
		Expect(metrics).To(MatchLine(`^my_token_request_count\{attempt="1",code="200"\} .*$`))
	})

	It("Generates token request duration", func() {
		// Send the request:
		_, err := connection.ClustersMgmt().V1().Clusters().Cluster("123").Get().
			Send()
		Expect(err).ToNot(HaveOccurred())

		// Verify the metrics:
		metrics := metricsServer.Metrics()
		Expect(metrics).To(MatchLine(`^my_request_duration_bucket\{.*path="/my_token".*,le="0.1"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_request_duration_bucket\{.*path="/my_token".*,le="1"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_request_duration_bucket\{.*path="/my_token".*,le="10"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_request_duration_bucket\{.*path="/my_token".*,le="30"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_request_duration_bucket\{.*path="/my_token".*,le="\+Inf"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_request_duration_count\{.*path="/my_token".*\} .*$`))
		Expect(metrics).To(MatchLine(`^my_request_duration_sum\{.*path="/my_token"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_token_request_duration_bucket\{attempt="1",code="200",le="0.1"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_token_request_duration_bucket\{attempt="1",code="200",le="1"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_token_request_duration_bucket\{attempt="1",code="200",le="10"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_token_request_duration_bucket\{attempt="1",code="200",le="30"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_token_request_duration_bucket\{attempt="1",code="200",le="\+Inf"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_token_request_duration_count\{attempt="1",code="200"\} .*$`))
		Expect(metrics).To(MatchLine(`^my_token_request_duration_sum\{attempt="1",code="200"\} .*$`))
	})
})

var _ = Describe("Metrics disabled", func() {
	// Servers used during the tests:
	var oidServer *ghttp.Server
	var apiServer *ghttp.Server
	var metricsServer *MetricsServer

	// Connection used during the tests:
	var connection *Connection

	BeforeEach(func() {
		var err error

		// Create the tokens:
		accessToken := MakeTokenString("Bearer", 5*time.Minute)
		refreshToken := MakeTokenString("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer = MakeTCPServer()
		oidServer.AppendHandlers(
			ghttp.CombineHandlers(
				RespondWithAccessAndRefreshTokens(accessToken, refreshToken),
			),
		)

		// Create the API server:
		apiServer = MakeTCPServer()
		apiServer.AppendHandlers(
			RespondWithJSON(http.StatusOK, ""),
		)

		// Create the metrics server:
		metricsServer = NewMetricsServer()

		// Create the connection:
		connection, err = NewConnectionBuilder().
			Logger(logger).
			URL(apiServer.URL()).
			TokenURL(oidServer.URL()).
			Tokens(refreshToken).
			MetricsRegisterer(metricsServer.Registry()).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		// Stop the servers:
		oidServer.Close()
		apiServer.Close()
		metricsServer.Close()

		// Close the connection:
		err := connection.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Doesn't generate metrics for raw request", func() {
		// Send the request:
		_, err := connection.Get().
			Path("/api/clusters_mgmt/v1/clusters/123").
			Send()
		Expect(err).ToNot(HaveOccurred())

		// Verify the metrics:
		metrics := metricsServer.Metrics()
		Expect(metrics).To(ConsistOf(""))
	})

	It("Doesn't generate metrics for type safe request", func() {
		// Send the request:
		_, err := connection.ClustersMgmt().V1().Clusters().Cluster("123").Get().
			Send()
		Expect(err).ToNot(HaveOccurred())

		// Verify the metrics:
		metrics := metricsServer.Metrics()
		Expect(metrics).To(ConsistOf(""))
	})
})
