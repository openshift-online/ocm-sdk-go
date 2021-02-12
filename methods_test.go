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
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint

	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Methods", func() {
	// Servers used during the tests:
	var oidServer *ghttp.Server
	var apiServer *ghttp.Server

	// Names of the temporary files containing the CAs for the servers:
	var oidCA string
	var apiCA string

	// Connection used during the tests:
	var connection *Connection

	BeforeEach(func() {
		var err error

		// Create the tokens:
		accessToken := DefaultToken("Bearer", 5*time.Minute)
		refreshToken := DefaultToken("Refresh", 10*time.Hour)

		// Create the OpenID server:
		oidServer, oidCA = MakeTCPTLSServer()
		oidServer.AppendHandlers(
			ghttp.CombineHandlers(
				RespondWithTokens(accessToken, refreshToken),
			),
		)

		// Create the API server:
		apiServer, apiCA = MakeTCPTLSServer()

		// Metrics subsystem - value doesn't matter but configuring it enables
		// prometheus exporting, exercising the counter increment functionality
		// (e.g. will catch inconsistent labels).
		metricsSubsystem := "test_subsystem"

		// Create the connection:
		connection, err = NewConnectionBuilder().
			Logger(logger).
			MetricsSubsystem(metricsSubsystem).
			TokenURL(oidServer.URL()).
			URL(apiServer.URL()).
			Tokens(refreshToken).
			TrustedCAFile(oidCA).
			TrustedCAFile(apiCA).
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

		// Remove the temporary CA files:
		err = os.Remove(oidCA)
		Expect(err).ToNot(HaveOccurred())
		err = os.Remove(apiCA)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Get", func() {
		It("Sends path", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/mypath"),
					RespondWithJSON(http.StatusOK, ""),
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
					RespondWithJSON(http.StatusOK, ""),
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
					RespondWithJSON(http.StatusOK, ""),
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
					RespondWithJSON(http.StatusOK, ""),
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
					RespondWithJSON(http.StatusOK, ""),
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
					RespondWithJSON(http.StatusOK, ""),
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
				RespondWithJSON(http.StatusOK, `{"test":"mybody"}`),
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
				RespondWithJSON(http.StatusOK, ""),
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
				RespondWithJSON(http.StatusBadRequest, ""),
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
				RespondWithJSON(http.StatusInternalServerError, ""),
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

		It("Honors cookies", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.CombineHandlers(
					RespondWithCookie("mycookie", "myvalue"),
					RespondWithJSONTemplate(http.StatusOK, "{}"),
				),
				ghttp.CombineHandlers(
					VerifyCookie("mycookie", "myvalue"),
					RespondWithJSONTemplate(http.StatusOK, "{}"),
				),
			)

			// Send first request. The server will respond setting a cookie.
			_, err := connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())

			// Send second request, which should include the cookie returned by the
			// server in the first response.
			_, err = connection.Get().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Wraps deadline exceeded error", func() {
			// Configure the server so that it introduces an artificial delay:
			apiServer.AppendHandlers(
				ghttp.CombineHandlers(
					http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						time.Sleep(10 * time.Millisecond)
					}),
					RespondWithJSON(http.StatusOK, ""),
				),
			)

			// Send the request with a timeout smaller than the artificial delay
			// introduced by the server so that a deadline exceeded error will be
			// created and returned:
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Millisecond)
			_, err := connection.Get().
				Path("/mypath").
				SendContext(ctx)
			Expect(err).To(HaveOccurred())
			Expect(errors.Is(err, context.DeadlineExceeded)).To(BeTrue())
		})
	})

	Describe("Post", func() {
		It("Accepts empty body", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				RespondWithJSON(http.StatusOK, ""),
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
				RespondWithJSON(http.StatusOK, ""),
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

	Describe("Put", func() {
		It("Accepts empty body", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				RespondWithJSON(http.StatusOK, ""),
			)

			// Send the request:
			response, err := connection.Put().
				Path("/mypath").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusOK))
		})
	})

	When("Server doesn't return JSON content type", func() {
		It("It should ignore letter case", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.RespondWith(
					http.StatusOK, nil, http.Header{
						"cOnTeNt-TyPe": []string{
							"AppLicaTion/JSON",
						},
					},
				),
			)

			// Send the request:
			response, err := connection.Get().
				Path("/api/clusters_mgmt/v1/clusters").
				Send()
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.Status()).To(Equal(http.StatusOK))
		})

		It("Adds complete content to error message if it is short", func() {
			// Configure the server:
			apiServer.AppendHandlers(
				ghttp.RespondWith(
					http.StatusBadGateway,
					`Service not available`,
					http.Header{
						"Content-Type": []string{
							"text/plain",
						},
					},
				),
			)

			// Try to get the access token:
			_, err := connection.Get().
				Path("/api/clusters_mgmt/v1/clusters").
				Send()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("text/plain"))
			Expect(message).To(ContainSubstring("Service not available"))
		})

		It("Extracts and summarizes text if it's a long html", func() {
			// Calculate a long message:
			content := gatewayError

			// Configure the server:
			apiServer.AppendHandlers(
				RespondWithContent(http.StatusBadGateway, "text/html", content),
			)

			// Try to get the access token:
			_, err := connection.Get().
				Path("/api/clusters_mgmt/v1/clusters").
				Send()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("text/html"))
			Expect(message).To(ContainSubstring("Application is not available"))
			Expect(message).To(ContainSubstring("..."))
		})

		It("Summary shows html entities in a readable form", func() {
			content := errorWithHTMLEntities

			// Configure the server:
			apiServer.AppendHandlers(
				RespondWithContent(http.StatusBadGateway, "text/html", content),
			)

			// Try to get the access token:
			_, err := connection.Get().
				Path("/api/clusters_mgmt/v1/clusters").
				Send()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("text/html"))
			Expect(message).NotTo(ContainSubstring("tag was not removed"))
			Expect(message).To(ContainSubstring(
				`You don't have permission to access "http://sso.redhat.com/AK_PM_VPATH0/" ` +
					`on this server. Reference #18.3500e8ac.1601993172.3a9c59e`))
			Expect(message).To(ContainSubstring(`< > " & € ∭`))
			// Sufficiently short to log Akamai reference number without shortening.
			Expect(message).NotTo(ContainSubstring("..."))
		})
	})
})

const gatewayError = `
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1">

  <style type="text/css">
  /*!
   * Bootstrap v3.3.5 (http://getbootstrap.com)
   * Copyright 2011-2015 Twitter, Inc.
   * Licensed under MIT (https://github.com/twbs/bootstrap/blob/master/LICENSE)
   */
  /*! normalize.css v3.0.3 | MIT License | github.com/necolas/normalize.css */
  html {
    font-family: sans-serif;
    -ms-text-size-adjust: 100%;
    -webkit-text-size-adjust: 100%;
  }
  body {
    margin: 0;
  }
  h1 {
    font-size: 1.7em;
    font-weight: 400;
    line-height: 1.3;
    margin: 0.68em 0;
  }
  * {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
  *:before,
  *:after {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
  html {
    -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
  }
  body {
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
    line-height: 1.66666667;
    font-size: 13px;
    color: #333333;
    background-color: #ffffff;
    margin: 2em 1em;
  }
  p {
    margin: 0 0 10px;
    font-size: 13px;
  }
  .alert.alert-info {
    padding: 15px;
    margin-bottom: 20px;
    border: 1px solid transparent;
    background-color: #f5f5f5;
    border-color: #8b8d8f;
    color: #363636;
    margin-top: 30px;
  }
  .alert p {
    padding-left: 35px;
  }
  a {
    color: #0088ce;
  }

  ul {
    position: relative;
    padding-left: 51px;
  }
  p.info {
    position: relative;
    font-size: 15px;
    margin-bottom: 10px;
  }
  p.info:before, p.info:after {
    content: "";
    position: absolute;
    top: 9%;
    left: 0;
  }
  p.info:before {
    content: "i";
    left: 3px;
    width: 20px;
    height: 20px;
    font-family: serif;
    font-size: 15px;
    font-weight: bold;
    line-height: 21px;
    text-align: center;
    color: #fff;
    background: #4d5258;
    border-radius: 16px;
  }

  @media (min-width: 768px) {
    body {
      margin: 4em 3em;
    }
    h1 {
      font-size: 2.15em;}
  }

  </style>
  </head>
  <body>
    <div>
      <h1>Application is not available</h1>
      <p>The application is currently not serving requests at this endpoint.
		It may not have been started or is still starting.</p>

      <div class="alert alert-info">
        <p class="info">
          Possible reasons you are seeing this page:
        </p>
        <ul>
          <li>
            <strong>The host doesn't exist.</strong>
            Make sure the hostname was typed correctly and that a route matching this hostname exists.
          </li>
          <li>
            <strong>The host exists, but doesn't have a matching path.</strong>
            Check if the URL path was typed correctly and that the route was created using the desired path.
          </li>
          <li>
            <strong>Route and path matches, but all pods are down.</strong>
            Make sure that the resources exposed by this route (pods, services, deployment configs, etc)
			have at least one pod running.
          </li>
        </ul>
      </div>
    </div>
  </body>
</html>
`

// The text in body is a real response from Akamai blocking/rate-limiting our access to SSO.
// I don't have the original HTML so the head & tags are made up.
const errorWithHTMLEntities = `
<html>
  <head>
  <title>Access Denied</title>
  <script>
   if(2 < 3) alert("2 &lt; 3 but more imporantly &lt;script&gt; tag was not removed!");
  </script>
  </head>
  <body>
   <h1>Access Denied</h1>
   <p>You don't have permission to access
   "http&#58;&#47;&#47;sso&#46;redhat&#46;com&#47;AK&#95;PM&#95;VPATH0&#47;" on this server.
   Reference&#32;&#35;18&#46;3500e8ac&#46;1601993172&#46;3a9c59e</p>
   &lt; &gt; &quot; &amp; &euro; &tint;
  </body>
</html>
`
