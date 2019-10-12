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

// This file contains tests for the methods that request tokens.

package sdk

import (
	"net/http"
	"time"

	// nolint
	. "github.com/onsi/ginkgo"
	// nolint
	. "github.com/onsi/gomega"
	// nolint
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Methods", func() {
	// Servers used during the tests:
	var oidServer *Server
	var apiServer *Server

	// Logger used during the testss:
	var logger Logger

	// Connection used during the tests:
	var connection *Connection

	BeforeEach(func() {
		var err error

		// Create the tokens:
		accessToken := DefaultToken("Bearer", 5*time.Minute)
		refreshToken := DefaultToken("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer = NewServer()
		oidServer.AppendHandlers(
			CombineHandlers(
				RespondWithTokens(accessToken, refreshToken),
			),
		)

		// Create the API server:
		apiServer = NewServer()

		// Create the logger:
		logger, err = NewStdLoggerBuilder().
			Streams(GinkgoWriter, GinkgoWriter).
			Debug(true).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Create the connection:
		connection, err = NewConnectionBuilder().
			Logger(logger).
			TokenURL(oidServer.URL()).
			URL(apiServer.URL()).
			Tokens(refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		// Stop the servers:
		oidServer.Close()
		apiServer.Close()

		// Close the connection:
		err := connection.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Get", func() {
		It("Sends path", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				VerifyRequest(http.MethodGet, "/mypath"),
			)

			// Send the request:
			_, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Sends accept header", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				VerifyHeaderKV("Accept", "application/json"),
			)

			// Send the request:
			_, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Sends one query parameter", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				VerifyFormKV("myparameter", "myvalue"),
			)

			// Send the request:
			_, err := connection.Get().
				Path("/mypath").
				Parameter("myparameter", "myvalue").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Sends two query parameters", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				CombineHandlers(
					VerifyFormKV("myparameter", "myvalue"),
					VerifyFormKV("yourparameter", "yourvalue"),
				),
			)

			// Send the request:
			_, err := connection.Get().
				Path("/mypath").
				Parameter("myparameter", "myvalue").
				Parameter("yourparameter", "yourvalue").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Sends one header", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				VerifyHeaderKV("myheader", "myvalue"),
			)

			// Send the request:
			_, err := connection.Get().
				Path("/mypath").
				Header("myheader", "myvalue").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Sends two headers", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				CombineHandlers(
					VerifyHeaderKV("myheader", "myvalue"),
					VerifyHeaderKV("yourheader", "yourvalue"),
				),
			)

			// Send the request:
			_, err := connection.Get().
				Path("/mypath").
				Header("myheader", "myvalue").
				Header("yourheader", "yourvalue").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Receives body", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				RespondWith(http.StatusOK, "mybody"),
			)

			// Send the request:
			response, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusOK))
			Expect(response.String()).To(Equal("mybody"))
			Expect(response.Bytes()).To(Equal([]byte("mybody")))
		})

		It("Receives status code 200", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				RespondWith(http.StatusOK, nil),
			)

			// Send the request:
			response, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusOK))
		})

		It("Receives status code 400", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				RespondWith(http.StatusBadRequest, nil),
			)

			// Send the request:
			response, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusBadRequest))
		})

		It("Receives status code 500", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				RespondWith(http.StatusInternalServerError, nil),
			)

			// Send the request:
			response, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusInternalServerError))
		})

		It("Fails if no path is given", func() {
			response, err := connection.Get().
				Send()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("path"))
			Expect(err.Error()).To(ContainSubstring("mandatory"))
			Expect(response).To(BeNil())
		})

	})

	Describe("Post", func() {
		It("Accepts empty body", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				RespondWith(http.StatusOK, nil),
			)

			// Send the request:
			response, err := connection.Post().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusOK))
		})
	})

	Describe("Patch", func() {
		It("Accepts empty body", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				RespondWith(http.StatusOK, nil),
			)

			// Send the request:
			response, err := connection.Patch().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusOK))
		})
	})
})
