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

// This file contains tests for the follow redirects support.

package sdk

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo/v2/dsl/core"                // nolint
	. "github.com/onsi/gomega"                            // nolint
	. "github.com/openshift-online/ocm-sdk-go/v2/testing" // nolint
)

var _ = Describe("Redirect", func() {
	// Tokens used during the tests:
	var accessToken string
	var refreshToken string

	// Servers used during the tests:
	var oidServer *ghttp.Server
	var redirectServer *ghttp.Server
	var realServer *ghttp.Server

	// Names of the temporary files containing the CAs for the servers:
	var oidCA string
	var redirectCA string
	var realCA string

	// URLs of the servers:
	var oidURL string
	var redirectURL string
	var realURL string

	// Connection used for the tests:
	var connection *Connection

	BeforeEach(func() {
		var err error

		// Create the tokens:
		accessToken = MakeTokenString("Bearer", 5*time.Minute)
		refreshToken = MakeTokenString("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer, oidCA = MakeTCPTLSServer()
		oidServer.AppendHandlers(
			ghttp.CombineHandlers(
				RespondWithAccessAndRefreshTokens(accessToken, refreshToken),
			),
		)
		oidURL = oidServer.URL()

		// Create the API servers:
		redirectServer, redirectCA = MakeTCPTLSServer()
		redirectURL = redirectServer.URL()
		realServer, realCA = MakeTCPTLSServer()
		realURL = realServer.URL()

		// Configure the real server so that it verifies that the request is
		// received:
		realServer.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest(http.MethodGet, "/api/clusters_mgmt/v1"),
				RespondWithJSON(http.StatusOK, "{}"),
			),
		)

		// Create the connection:
		connection, err = NewConnection().
			Logger(logger).
			TokenURL(oidURL).
			Tokens(accessToken, refreshToken).
			URL(redirectURL).
			TrustedCAFile(oidCA).
			TrustedCAFile(redirectCA).
			TrustedCAFile(realCA).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		var err error

		// Close the connection:
		err = connection.Close()
		Expect(err).ToNot(HaveOccurred())

		// Stop the servers:
		oidServer.Close()
		redirectServer.Close()
		realServer.Close()

		// Remove the temporary CA files:
		err = os.Remove(oidCA)
		Expect(err).ToNot(HaveOccurred())
		err = os.Remove(redirectCA)
		Expect(err).ToNot(HaveOccurred())
		err = os.Remove(realCA)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Untyped get", func() {
		It("Honours permanent redirect", func() {
			// Configure the redirect server so that it redirects to the real server:
			redirectServer.AppendHandlers(
				RespondWithPermanentRedirect(realURL),
			)

			// Send the request:
			_, err := connection.Get().
				Path("/api/clusters_mgmt/v1").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Honours temporary redirect", func() {
			// Configure the redirect server so that it redirects to the real server:
			redirectServer.AppendHandlers(
				RespondWithTemporaryRedirect(realURL),
			)

			// Send the request:
			_, err := connection.Get().
				Path("/api/clusters_mgmt/v1").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Typed get", func() {
		It("Honours permanent redirect", func() {
			// Configure the redirect server so that it redirects to the rea server:
			redirectServer.AppendHandlers(
				RespondWithPermanentRedirect(realURL),
			)

			// Send the request:
			_, err := connection.ClustersMgmt().V1().Get().Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Honours temporary redirect", func() {
			// Configure the redirect server so that it redirects to the real server:
			redirectServer.AppendHandlers(
				RespondWithTemporaryRedirect(realURL),
			)

			// Send the request:
			_, err := connection.ClustersMgmt().V1().Get().Send()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

// RespondWithPermanentRedirect responds with a permanent redirect to the given target URL,
// changing the scheme, host and port number but preserving the path of the original request.
func RespondWithPermanentRedirect(target string) http.HandlerFunc {
	return RespondWithRedirect(http.StatusPermanentRedirect, target)
}

// RespondWithTemporaryRedirect responds with a permanent redirect to the given target URL,
// changing the scheme, host and port number but preserving the path of the original request.
func RespondWithTemporaryRedirect(target string) http.HandlerFunc {
	return RespondWithRedirect(http.StatusTemporaryRedirect, target)
}

// RespondTemporaryRedirect responds with a redirect with the given code to the given target URL,
// changing the scheme, host and port number but preserving the path of the original request.
func RespondWithRedirect(code int, target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parsed, err := url.Parse(target)
		Expect(err).ToNot(HaveOccurred())
		location := r.URL
		location.Scheme = parsed.Scheme
		location.Host = parsed.Host
		w.Header().Set("Location", location.String())
		w.WriteHeader(code)
	}
}
