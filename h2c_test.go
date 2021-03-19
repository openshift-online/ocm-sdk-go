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

// This file contains tests for the support for Unix sockets.

package sdk

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"                         // nolint
	. "github.com/onsi/gomega"                         // nolint
	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

var _ = Describe("H2C", func() {
	var (
		accessToken  string
		refreshToken string
		oidServer    *ghttp.Server
	)

	BeforeEach(func() {
		// Create the tokens:
		accessToken = MakeTokenString("Bearer", 5*time.Minute)
		refreshToken = MakeTokenString("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer = MakeTCPServer()
		oidServer.AppendHandlers(
			ghttp.CombineHandlers(
				RespondWithTokens(accessToken, refreshToken),
			),
		)
	})

	AfterEach(func() {
		// Stop the OpenID server:
		oidServer.Close()
	})

	Describe("With TCP", func() {
		var (
			apiServer  *ghttp.Server
			connection *Connection
		)

		BeforeEach(func() {
			var err error

			// Create the API server:
			apiServer = MakeTCPH2CServer()

			// Alter the URL scheme to force use of HTTP/2 without TLS:
			apiURL, err := url.Parse(apiServer.URL())
			Expect(err).ToNot(HaveOccurred())
			apiURL.Scheme = "h2c"

			// Create the connection:
			connection, err = NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidServer.URL()).
				URL(apiURL.String()).
				Tokens(accessToken, refreshToken).
				Build()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			var err error

			// Close the connection:
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())

			// Stop the API server:
			apiServer.Close()
		})

		It("Uses HTTP/2.0", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.CombineHandlers(
					http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						Expect(r.Proto).To(Equal("HTTP/2.0"))
					}),
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

	Describe("With Unix socket", func() {
		var (
			apiServer  *ghttp.Server
			apiSocket  string
			connection *Connection
		)

		BeforeEach(func() {
			var err error

			// Create the API server:
			apiServer, apiSocket = MakeUnixH2CServer()
			apiURL := "unix+h2c://127.0.0.1" + apiSocket

			// Create the connection:
			connection, err = NewConnectionBuilder().
				Logger(logger).
				TokenURL(oidServer.URL()).
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

			// Stop the API server:
			apiServer.Close()

			// Remore the temporary files and directories:
			err = os.RemoveAll(filepath.Dir(apiSocket))
			Expect(err).ToNot(HaveOccurred())
		})

		It("Uses HTTP/2.0", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.CombineHandlers(
					http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						Expect(r.Proto).To(Equal("HTTP/2.0"))
					}),
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
})
