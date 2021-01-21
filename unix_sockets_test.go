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

// This file contains tests for the support for Unix sockets.

package sdk

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint

	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Unix sockets", func() {
	var accessToken string
	var refreshToken string
	var oidServer *ghttp.Server
	var oidSocket string
	var oidURL string

	BeforeEach(func() {
		// Create the tokens:
		accessToken = DefaultToken("Bearer", 5*time.Minute)
		refreshToken = DefaultToken("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer, oidSocket = MakeUnixServer()
		oidServer.AppendHandlers(
			ghttp.CombineHandlers(
				RespondWithTokens(accessToken, refreshToken),
			),
		)
		oidURL = "unix://127.0.0.1" + oidSocket
	})

	AfterEach(func() {
		// Stop the OpenID server:
		oidServer.Close()
	})

	Describe("With HTTP", func() {
		var apiServer *ghttp.Server
		var apiURL string
		var apiSocket string
		var connection *Connection

		BeforeEach(func() {
			var err error

			// Create the server:
			apiServer, apiSocket = MakeUnixServer()
			apiURL = "unix://127.0.0.1" + apiSocket

			// Create the connection:
			connection, err = NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				URL(apiURL).
				Tokens(accessToken, refreshToken).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			var err error

			// Close the connection:
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())

			// Close the server:
			apiServer.Close()

			// Remore the temporary files and directories:
			err = os.RemoveAll(filepath.Dir(apiSocket))
			Expect(err).ToNot(HaveOccurred())
		})

		It("Get", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.CombineHandlers(
					RespondWithJSON(http.StatusOK, `{
						"href": "/api/clusters_mgmt"
					}`),
				),
			)

			// Send the request:
			response, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.String()).To(MatchJSON(`{
				"href": "/api/clusters_mgmt"
			}`))
		})
	})

	Describe("With HTTPS", func() {
		var apiServer *ghttp.Server
		var apiURL string
		var apiCA string
		var apiSocket string
		var connection *Connection

		BeforeEach(func() {
			var err error

			// Create the server:
			apiServer, apiCA, apiSocket = MakeUnixTLSServer()
			apiURL = "unix+https://127.0.0.1" + apiSocket

			// Create the connection:
			connection, err = NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				URL(apiURL).
				Tokens(accessToken, refreshToken).
				TrustedCAFile(apiCA).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			var err error

			// Close the connection:
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())

			// Close the server:
			apiServer.Close()

			// Remore the temporary files and directories:
			err = os.RemoveAll(apiCA)
			Expect(err).ToNot(HaveOccurred())
			err = os.RemoveAll(filepath.Dir(apiSocket))
			Expect(err).ToNot(HaveOccurred())
		})

		It("Get", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.CombineHandlers(
					RespondWithJSON(http.StatusOK, `{
						"href": "/api/clusters_mgmt"
					}`),
				),
			)

			// Send the request:
			response, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.String()).To(MatchJSON(`{
				"href": "/api/clusters_mgmt"
			}`))
		})
	})

	Describe("With Unix and TCP simultaneously", func() {
		var unixServer *ghttp.Server
		var unixURL string
		var unixSocket string
		var tcpServer *ghttp.Server
		var tcpURL string
		var connection *Connection

		BeforeEach(func() {
			var err error

			// Create the the main server, using Unix sockets:
			unixServer, unixSocket = MakeUnixServer()
			unixURL = "unix://127.0.0.1" + unixSocket

			// Make the alternative, using TCP:
			tcpServer = MakeTCPServer()
			tcpURL = tcpServer.URL()

			// Create the connection:
			connection, err = NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidURL).
				AlternativeURL("/api/clusters_mgmt", unixURL).
				AlternativeURL("/api/accounts_mgmt", tcpURL).
				Tokens(accessToken, refreshToken).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			var err error

			// Close the connection:
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())

			// Close the servers:
			unixServer.Close()
			tcpServer.Close()

			// Remore the temporary files and directories:
			err = os.RemoveAll(filepath.Dir(unixSocket))
			Expect(err).ToNot(HaveOccurred())
		})

		It("Get", func() {
			// Configure the servers:
			unixServer.AppendHandlers(
				ghttp.CombineHandlers(
					RespondWithJSON(http.StatusOK, `{
						"href": "/api/clusters_mgmt"
					}`),
				),
			)
			tcpServer.AppendHandlers(
				ghttp.CombineHandlers(
					RespondWithJSON(http.StatusOK, `{
						"href": "/api/accounts_mgmt"
					}`),
				),
			)

			// Send a request to the Unix server:
			unixResponse, err := connection.Get().
				Path("/api/clusters_mgmt").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(unixResponse.String()).To(MatchJSON(`{
				"href": "/api/clusters_mgmt"
			}`))

			// Send a request to the TCP server:
			tcpResponse, err := connection.Get().
				Path("/api/accounts_mgmt").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(tcpResponse.String()).To(MatchJSON(`{
				"href": "/api/accounts_mgmt"
			}`))
		})
	})
})
