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

package sdk

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2/dsl/core"                // nolint
	. "github.com/onsi/gomega"                            // nolint
	. "github.com/openshift-online/ocm-sdk-go/v2/testing" // nolint
)

var _ = Describe("Retry", func() {
	var ctx context.Context
	var token string

	BeforeEach(func() {
		// Create a context:
		ctx = context.Background()

		// Create a token:
		token = MakeTokenString("Bearer", 15*time.Minute)
	})

	Describe("Get", func() {
		It("Retries if protocol error", func() {
			// Create a connection with a transport wrapper that returns an error for
			// the first request and 200 for the second.
			connection, err := NewConnection().
				Logger(logger).
				Tokens(token).
				TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
					return CombineTransports(
						ErrorTransport(errors.New("PROTOCOL_ERROR")),
						JSONTransport(http.StatusOK, "{}"),
					)
				}).
				RetryInterval(10 * time.Millisecond).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Send the request:
			response, err := connection.Get().Path("/mypath").Send(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
		})

		It("Retries for 429", func() {
			// Create a connection with a transport wrapper that returns 429 for the
			// first request and 200 for the second.
			connection, err := NewConnection().
				Logger(logger).
				Tokens(token).
				TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
					return CombineTransports(
						JSONTransport(http.StatusTooManyRequests, "{}"),
						JSONTransport(http.StatusOK, "{}"),
					)
				}).
				RetryInterval(10 * time.Millisecond).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Send the request:
			response, err := connection.Get().Path("/mypath").Send(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
		})

		It("Retries for 503", func() {
			// Create a connection with a transport wrapper that returns 503 for the
			// first request and 200 for the second.
			connection, err := NewConnection().
				Logger(logger).
				Tokens(token).
				TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
					return CombineTransports(
						JSONTransport(http.StatusServiceUnavailable, "{}"),
						JSONTransport(http.StatusOK, "{}"),
					)
				}).
				RetryInterval(10 * time.Millisecond).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Send the request:
			response, err := connection.Get().Path("/mypath").Send(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
		})
	})

	Describe("Delete", func() {
		It("Retries for protocol error", func() {
			// Create a connection with a transport wrapper that returns an error for
			// the first request and 200 for the second.
			connection, err := NewConnection().
				Logger(logger).
				Tokens(token).
				TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
					return CombineTransports(
						ErrorTransport(errors.New("PROTOCOL_ERROR")),
						JSONTransport(http.StatusOK, "{}"),
					)
				}).
				RetryInterval(10 * time.Millisecond).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Send the request:
			response, err := connection.Delete().Path("/mypath").Send(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
		})
	})

	Describe("Post with body", func() {
		It("Retries for protocol error", func() {
			// Create a connection with a transport wrapper that returns an error for
			// the first request and 200 for the second.
			connection, err := NewConnection().
				Logger(logger).
				Tokens(token).
				TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
					return CombineTransports(
						ErrorTransport(errors.New("PROTOCOL_ERROR")),
						JSONTransport(http.StatusOK, "{}"),
					)
				}).
				RetryInterval(10 * time.Millisecond).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Send the request:
			response, err := connection.Post().Path("/mypath").String("{}").Send(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
		})

		It("Retries for 429", func() {
			// Create a connection with a transport wrapper that returns 429 for the
			// first request and 200 for the second.
			connection, err := NewConnection().
				Logger(logger).
				Tokens(token).
				TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
					return CombineTransports(
						JSONTransport(http.StatusTooManyRequests, "{}"),
						JSONTransport(http.StatusOK, "{}"),
					)
				}).
				RetryInterval(10 * time.Millisecond).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Send the request:
			response, err := connection.Post().
				Path("/mypath").
				String(`{}`).
				Send(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
		})

		It("Retries for 503", func() {
			// Create a connection with a transport wrapper that returns 503 for the
			// first request and 200 for the second.
			connection, err := NewConnection().
				Logger(logger).
				Tokens(token).
				TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
					return CombineTransports(
						JSONTransport(http.StatusServiceUnavailable, "{}"),
						JSONTransport(http.StatusOK, "{}"),
					)
				}).
				RetryInterval(10 * time.Millisecond).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Send the request:
			response, err := connection.Post().
				Path("/mypath").
				String(`{}`).
				Send(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
		})
	})

	It("Writes error to the debug log", func() {
		// Create a logger that allows us to inspect the messages written to the log:
		var buffer bytes.Buffer
		logger := MakeLogger(&buffer)

		// Create a connection with a transport wrapper that returns an error for
		// the first request and 200 for the second.
		connection, err := NewConnection().
			Logger(logger).
			Tokens(token).
			TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
				return CombineTransports(
					ErrorTransport(errors.New("PROTOCOL_ERROR")),
					JSONTransport(http.StatusOK, `{}`),
				)
			}).
			RetryInterval(10 * time.Millisecond).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		response, err := connection.Get().Path("/mypath").Send(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(response).ToNot(BeNil())

		// Check that the details of the failed request are in the log:
		messages := buffer.String()
		Expect(messages).To(ContainSubstring("GET"))
		Expect(messages).To(ContainSubstring("mypath"))
		Expect(messages).To(ContainSubstring("failed with protocol error"))
		Expect(messages).To(ContainSubstring("PROTOCOL_ERROR"))
	})

	It("Writes failed response details to the log", func() {
		// Create a logger that allows us to inspect the messages written to the log:
		var buffer bytes.Buffer
		logger = MakeLogger(&buffer)

		// Create a connection with a transport wrapper that returns 503 for the first
		// request and 200 for the second.
		connection, err := NewConnection().
			Logger(logger).
			Tokens(token).
			TransportWrapper(func(_ http.RoundTripper) http.RoundTripper {
				return CombineTransports(
					JSONTransport(http.StatusServiceUnavailable, `{
						"reason": "Something failed",
					}`),
					JSONTransport(http.StatusOK, `{}`),
				)
			}).
			RetryInterval(10 * time.Millisecond).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		response, err := connection.Get().Path("/mypath").Send(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(response).ToNot(BeNil())

		// Check that the details of the failed request are in the log:
		messages := buffer.String()
		Expect(messages).To(ContainSubstring("503"))
		Expect(messages).To(ContainSubstring("Something failed"))
	})
})
