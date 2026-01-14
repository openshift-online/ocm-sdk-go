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

package internal

import (
	"context"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2/dsl/core" // nolint
	. "github.com/onsi/gomega"             // nolint
)

var _ = Describe("Create client selector", func() {
	It("Can't be created without a logger", func() {
		selector, err := NewClientSelector().Build(context.Background())
		Expect(err).To(HaveOccurred())
		Expect(selector).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("logger"))
		Expect(message).To(ContainSubstring("mandatory"))
	})
})

var _ = Describe("Select client", func() {
	var (
		ctx      context.Context
		selector *ClientSelector
	)

	BeforeEach(func() {
		var err error

		// Create a context:
		ctx = context.Background()

		// Create the selector:
		selector, err = NewClientSelector().
			Logger(logger).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(selector).ToNot(BeNil())
	})

	AfterEach(func() {
		// Close the selector:
		err := selector.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Reuses client for same TCP address", func() {
		address, err := ParseServerAddress(ctx, "tcp://my.server.com")
		Expect(err).ToNot(HaveOccurred())
		firstClient, err := selector.Select(ctx, address)
		Expect(err).ToNot(HaveOccurred())
		secondClient, err := selector.Select(ctx, address)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondClient).To(BeIdenticalTo(firstClient))
	})

	It("Doesn't reuse client for different TCP addresses", func() {
		firstAddress, err := ParseServerAddress(ctx, "tcp://my.server.com")
		Expect(err).ToNot(HaveOccurred())
		secondAddress, err := ParseServerAddress(ctx, "tcp://your.server.com")
		Expect(err).ToNot(HaveOccurred())
		firstClient, err := selector.Select(ctx, firstAddress)
		Expect(err).ToNot(HaveOccurred())
		secondClient, err := selector.Select(ctx, secondAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondClient == firstClient).To(BeFalse())
	})

	It("Reuses client for different TCP protocols", func() {
		firstAddress, err := ParseServerAddress(ctx, "http://my.server.com")
		Expect(err).ToNot(HaveOccurred())
		secondAddress, err := ParseServerAddress(ctx, "https://my.server.com")
		Expect(err).ToNot(HaveOccurred())
		firstClient, err := selector.Select(ctx, firstAddress)
		Expect(err).ToNot(HaveOccurred())
		secondClient, err := selector.Select(ctx, secondAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondClient == firstClient).To(BeTrue())
	})

	It("Doesn't resuse client for different Unix sockets", func() {
		firstAddress, err := ParseServerAddress(ctx, "unix://my.server.com/my.socket")
		Expect(err).ToNot(HaveOccurred())
		secondAddress, err := ParseServerAddress(ctx, "unix://my.server.com/your.socket")
		Expect(err).ToNot(HaveOccurred())
		firstClient, err := selector.Select(ctx, firstAddress)
		Expect(err).ToNot(HaveOccurred())
		secondClient, err := selector.Select(ctx, secondAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondClient == firstClient).To(BeFalse())
	})
})

var _ = Describe("Redirect Behavior", func() {
	var (
		ctx                  context.Context
		selector             *ClientSelector
		originServer         *httptest.Server
		responseServer       *httptest.Server
		expectedResponseBody string
	)

	BeforeEach(func() {
		var err error

		// Create a context:
		ctx = context.Background()

		expectedResponseBody = "myServerDotComRedirect"

		responseServer = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", expectedResponseBody)
		}))

		// simulate a redirect to a different domain by responding with a localhost url rather than a 127.0.0.1 url
		redirectURL := strings.Replace(responseServer.URL, "127.0.0.1", "localhost", 1)
		originServer = httptest.NewTLSServer(http.RedirectHandler(redirectURL, http.StatusMovedPermanently))

		cas := x509.NewCertPool()
		cas.AddCert(responseServer.Certificate())
		cas.AddCert(originServer.Certificate())

		// Create the selector:
		selector, err = NewClientSelector().
			TrustedCAs(cas).
			Insecure(true). //need insecure when using "localhost" to connect or you get TLS verification errors
			Logger(logger).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(selector).ToNot(BeNil())
	})

	AfterEach(func() {
		defer responseServer.Close()
		defer originServer.Close()

		// Close the selector:
		err := selector.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Doesn't re-use origin host for redirect", func() {
		address, err := ParseServerAddress(ctx, originServer.URL)
		Expect(err).ToNot(HaveOccurred())

		client, err := selector.Select(ctx, address)
		Expect(err).ToNot(HaveOccurred())

		resp, err := client.Get(originServer.URL)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.TLS.ServerName).To(Equal("localhost"))

		body := make([]byte, len(expectedResponseBody))
		_, _ = resp.Body.Read(body)
		Expect(string(body)).To(Equal("myServerDotComRedirect"))
	})
})
