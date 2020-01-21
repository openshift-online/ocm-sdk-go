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

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint

	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Methods", func() {
	// Servers used during the tests:
	var oidServer *ghttp.Server
	var apiServer *ghttp.Server

	// Logger used during the testss:
	var logger Logger

	// Connection used during the tests:
	var connection *Connection

	jsonHeader := http.Header{"Content-Type": []string{"application/json"}}

	BeforeEach(func() {
		var err error

		// Create the tokens:
		accessToken := DefaultToken("Bearer", 5*time.Minute)
		refreshToken := DefaultToken("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer = ghttp.NewServer()
		oidServer.AppendHandlers(
			ghttp.CombineHandlers(
				RespondWithTokens(accessToken, refreshToken),
			),
		)

		// Create the API server:
		apiServer = ghttp.NewServer()

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
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/mypath"),
					ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
				),
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
				ghttp.CombineHandlers(
					ghttp.VerifyHeaderKV("Accept", "application/json"),
					ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
				),
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
				ghttp.CombineHandlers(
					ghttp.VerifyFormKV("myparameter", "myvalue"),
					ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
				),
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
				ghttp.CombineHandlers(
					ghttp.VerifyFormKV("myparameter", "myvalue"),
					ghttp.VerifyFormKV("yourparameter", "yourvalue"),
					ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
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
				ghttp.CombineHandlers(
					ghttp.VerifyHeaderKV("myheader", "myvalue"),
					ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
				),
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
				ghttp.CombineHandlers(
					ghttp.VerifyHeaderKV("myheader", "myvalue"),
					ghttp.VerifyHeaderKV("yourheader", "yourvalue"),
					ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
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
				ghttp.RespondWith(http.StatusOK, `{"test":"mybody"}`, jsonHeader),
			)

			// Send the request:
			response, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusOK))
			Expect(response.String()).To(Equal(`{"test":"mybody"}`))
			Expect(response.Bytes()).To(Equal([]byte(`{"test":"mybody"}`)))
		})

		It("Receives status code 200", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
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
				ghttp.RespondWith(http.StatusBadRequest, nil, jsonHeader),
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
				ghttp.RespondWith(http.StatusInternalServerError, nil, jsonHeader),
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
				ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
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
				ghttp.RespondWith(http.StatusOK, nil, jsonHeader),
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

	Describe("JSON header", func() {
		It("It should ignore letter case", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.RespondWith(http.StatusOK, nil, http.Header{"cOnTeNt-TyPe": []string{"AppLicaTion/JSON"}}),
			)

			// Send the request:
			response, err := connection.Get().Path("/mypath").Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusOK))
		})

		It("It should error if not json", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.RespondWith(http.StatusOK, "test", http.Header{"Content-Type": []string{"application/html"}}),
			)

			// Send the request:
			response, err := connection.Get().Path("/mypath").Send()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("expected JSON"))
			Expect(err.Error()).To(ContainSubstring("request:"))
			Expect(err.Error()).To(ContainSubstring("response status:"))
			Expect(err.Error()).To(ContainSubstring("response body: test"))
			Expect(response).To(BeNil())
		})

		It("It should trim response", func() {
			// Configure the server:
			longText := `<textarea class="Playground-input js-playgroundCodeEl" ` +
				`spellcheck="false" aria-label="Try Go">` +
				"// You can edit this code! // Click here and start typing. " +
				`package main import "fmt" ` +
				`func main() { fmt.Println("Hello, 世界") } </textarea>`
			longText = longText + longText + longText
			apiServer.AppendHandlers(
				ghttp.RespondWith(http.StatusOK, longText, http.Header{"Content-Type": []string{"application/html"}}),
			)

			// Send the request:
			response, err := connection.Get().Path("/mypath").Send()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(`<textarea class="Playground-input`))
			Expect(len(err.Error()) < 400).To(BeTrue())
			Expect(response).To(BeNil())
		})
	})
})
