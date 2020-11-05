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

// This file contains tests for the alternative URL support.

package sdk

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint

	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Alternative URLs", func() {
	// Tokens used during the tests:
	var accessToken string
	var refreshToken string

	// Servers used during the tests:
	var oidServer *ghttp.Server
	var defaultServer *ghttp.Server
	var alternativeServer *ghttp.Server

	// URLs of the servers:
	var oidURL string
	var defaultURL string
	var alternativeURL string

	BeforeEach(func() {
		// Create the tokens:
		accessToken = DefaultToken("Bearer", 5*time.Minute)
		refreshToken = DefaultToken("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer = MakeServer()
		oidServer.AppendHandlers(
			ghttp.CombineHandlers(
				RespondWithTokens(accessToken, refreshToken),
			),
		)
		oidURL = oidServer.URL()

		// Create the API servers:
		defaultServer = MakeServer()
		defaultURL = defaultServer.URL()
		alternativeServer = MakeServer()
		alternativeURL = alternativeServer.URL()
	})

	AfterEach(func() {
		// Stop the servers:
		oidServer.Close()
		defaultServer.Close()
		alternativeServer.Close()
	})

	Describe("Untyped get", func() {
		It("Honours alternative URL", func() {
			// Configure the alternative server so that it verifies that the request
			// is sent:
			alternativeServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/api/clusters_mgmt"),
					RespondWithJSON(http.StatusOK, "{}"),
				),
			)

			// Create the connection:
			connection, err := NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				Tokens(accessToken, refreshToken).
				URL(defaultURL).
				AlternativeURL("/api/clusters_mgmt", alternativeURL).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer connection.Close()

			// Send the request:
			_, err = connection.Get().
				Path("/api/clusters_mgmt").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Uses default URL", func() {
			// Configure the default server so that it verifies that the request
			// is sent:
			defaultServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/api/clusters_mgmt"),
					RespondWithJSON(http.StatusOK, "{}"),
				),
			)

			// Create the connection:
			connection, err := NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				Tokens(accessToken, refreshToken).
				URL(defaultURL).
				AlternativeURL("/api/accounts_mgmt", alternativeURL).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer connection.Close()

			// Send the request:
			_, err = connection.Get().
				Path("/api/clusters_mgmt").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Uses most specific alternative URL", func() {
			// Configure the default server so that it verifies that the request
			// is sent:
			alternativeServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/api/clusters_mgmt/v1"),
					RespondWithJSON(http.StatusOK, "{}"),
				),
			)

			// Create the connection:
			connection, err := NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				Tokens(accessToken, refreshToken).
				URL(defaultURL).
				AlternativeURL("/api/clusters_mgmt", defaultURL).
				AlternativeURL("/api/clusters_mgmt/v1", alternativeURL).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer connection.Close()

			// Send the request:
			_, err = connection.Get().
				Path("/api/clusters_mgmt/v1").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Typed get", func() {
		It("Honours alternative URL", func() {
			// Configure the alternative server so that it verifies that the request
			// is sent:
			alternativeServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/api/clusters_mgmt/v1"),
					RespondWithJSON(http.StatusOK, "{}"),
				),
			)

			// Create the connection:
			connection, err := NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				Tokens(accessToken, refreshToken).
				URL(defaultURL).
				AlternativeURL("/api/clusters_mgmt", alternativeURL).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer connection.Close()

			// Send the request:
			_, err = connection.ClustersMgmt().V1().Get().Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Uses default URL", func() {
			// Configure the default server so that it verifies that the request
			// is sent:
			defaultServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/api/clusters_mgmt/v1"),
					RespondWithJSON(http.StatusOK, "{}"),
				),
			)

			// Create the connection:
			connection, err := NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				Tokens(accessToken, refreshToken).
				URL(defaultURL).
				AlternativeURL("/api/accounts_mgmt", alternativeURL).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer connection.Close()

			// Send the request:
			_, err = connection.ClustersMgmt().V1().Get().Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Uses most specific alternative URL", func() {
			// Configure the default server so that it verifies that the request
			// is sent:
			alternativeServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/api/clusters_mgmt/v1"),
					RespondWithJSON(http.StatusOK, "{}"),
				),
			)

			// Create the connection:
			connection, err := NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				Tokens(accessToken, refreshToken).
				URL(defaultURL).
				AlternativeURL("/api/clusters_mgmt", defaultURL).
				AlternativeURL("/api/clusters_mgmt/v1", alternativeURL).
				Build()
			Expect(err).ToNot(HaveOccurred())
			defer connection.Close()

			// Send the request:
			_, err = connection.ClustersMgmt().V1().Get().Send()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
