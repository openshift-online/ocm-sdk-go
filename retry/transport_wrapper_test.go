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

// This file contains tests for request retrying.

package retry

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/http2"

	. "github.com/onsi/ginkgo"                         // nolint
	. "github.com/onsi/gomega"                         // nolint
	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

var _ = Describe("Creation", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	It("Can't be created without a logger", func() {
		wrapper, err := NewTransportWrapper().
			Build(ctx)
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("logger"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Can be created with positive retry limit", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Limit(10).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(wrapper).ToNot(BeNil())
		err = wrapper.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can be created with zero retry limit", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Limit(0).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(wrapper).ToNot(BeNil())
		err = wrapper.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can't be created with negative retry limit", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Limit(-1).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("limit"))
		Expect(message).To(ContainSubstring("-1"))
		Expect(message).To(ContainSubstring("greater or equal than zero"))
	})

	It("Can be created with positive retry interval", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Interval(5 * time.Second).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(wrapper).ToNot(BeNil())
		err = wrapper.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can't be created with zero retry interval", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Interval(0).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("interval"))
		Expect(message).To(ContainSubstring("0"))
		Expect(message).To(ContainSubstring("greater than zero"))
	})

	It("Can't be created with negative retry interval", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Interval(0).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("interval"))
		Expect(message).To(ContainSubstring("0"))
		Expect(message).To(ContainSubstring("greater than zero"))
	})

	It("Can be created with jitter between zero and one", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Jitter(0.3).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(wrapper).ToNot(BeNil())
		err = wrapper.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can be created with zero jitter", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Jitter(0.0).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(wrapper).ToNot(BeNil())
		err = wrapper.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can't be created with negative jitter", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Jitter(-1).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("jitter"))
		Expect(message).To(ContainSubstring("0"))
		Expect(message).To(ContainSubstring("between zero and one"))
	})

	It("Can't be created with jitter greater than one", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Jitter(2).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("jitter"))
		Expect(message).To(ContainSubstring("2"))
		Expect(message).To(ContainSubstring("between zero and one"))
	})
})

var _ = Describe("Server error", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	When("Retry enabled", func() {
		It("Retries 503 without request body", func() {
			// Create a transport that returns a 503 error for the first request and 200
			// for the second:
			transport := CombineTransports(
				TextTransport(http.StatusServiceUnavailable, `ko`),
				JSONTransport(http.StatusOK, `{ "ok": true }`),
			)

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Interval(100 * time.Millisecond).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Create the client:
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Get("http://api.example.com/mypath")
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body).To(MatchJSON(`{ "ok": true }`))
		})

		It("Retries 503 with request body", func() {
			// Create a transport that returns a 503 error for the first request and 200
			// for the second:
			transport := CombineTransports(
				TextTransport(http.StatusServiceUnavailable, `ko`),
				JSONTransport(http.StatusOK, `{ "ok": true }`),
			)

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Interval(100 * time.Millisecond).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Create the client:
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Post(
				"http://api.example.com/mypath",
				"application/json",
				strings.NewReader(`{}`),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body).To(MatchJSON(`{ "ok": true }`))
		})

		It("Retries 429 without request body", func() {
			// Create a transport that returns a 429 error for the first request and 200
			// for the second:
			transport := CombineTransports(
				JSONTransport(http.StatusTooManyRequests, `{ "ok": false }`),
				JSONTransport(http.StatusOK, `{ "ok": true }`),
			)

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Interval(100 * time.Millisecond).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Create the client:
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Get("http://api.example.com/mypath")
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body).To(MatchJSON(`{ "ok": true }`))
		})

		It("Retries 429 with request body", func() {
			// Create a transport that returns a 429 error for the first request and 200
			// for the second:
			transport := CombineTransports(
				JSONTransport(http.StatusTooManyRequests, `{ "ok": false }`),
				JSONTransport(http.StatusOK, `{ "ok": true }`),
			)

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Interval(100 * time.Millisecond).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Create the client:
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Post(
				"http://api.example.com/mypath",
				"application/json",
				strings.NewReader(`{}`),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body).To(MatchJSON(`{ "ok": true }`))
		})
	})

	When("Retry disabled", func() {
		It("Doesn't retry 503", func() {
			// Create a transport that returns a 503 error for the first request and 200
			// for the second:
			transport := CombineTransports(
				TextTransport(http.StatusServiceUnavailable, `ko`),
				JSONTransport(http.StatusOK, `{ "ok": true }`),
			)

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Limit(0).
				Interval(100 * time.Millisecond).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Create the client:
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Get("http://api.example.com/mypath")
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusServiceUnavailable))
			body, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(body)).To(Equal(`ko`))
		})

		It("Doesn't retry 429", func() {
			// Create a transport that returns a 429 error for the first request and 200
			// for the second:
			transport := CombineTransports(
				JSONTransport(http.StatusTooManyRequests, `{ "ok": false }`),
				JSONTransport(http.StatusOK, `{ "ok": true }`),
			)

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Limit(0).
				Interval(100 * time.Millisecond).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()

			// Create the client:
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Get("http://api.example.com/mypath")
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusTooManyRequests))
			body, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body).To(MatchJSON(`{ "ok": false }`))
		})
	})
})

var _ = Describe("Protocol error", func() {
	var ctx context.Context
	var listener net.Listener
	var address string
	var transport http.RoundTripper

	BeforeEach(func() {
		// Create a context:
		ctx = context.Background()

		// Create a listener:
		listener, address = Listen()

		// Create the basic transport:
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			ForceAttemptHTTP2: true,
		}
	})

	AfterEach(func() {
		// Close the listener:
		err := listener.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	When("Retry is enabled", func() {
		It("Tolerates protocol error", func() {
			var err error

			// Run the HTTP/2 server:
			go func() {
				defer GinkgoRecover()

				// Reject the first connection:
				first := Accept(listener)
				Reject(first)

				// Accept the second connection:
				second := Accept(listener)
				Serve(second, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, err := w.Write([]byte("{}"))
					Expect(err).ToNot(HaveOccurred())
				})
			}()

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Interval(100 * time.Millisecond).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Get(address)
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(response.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body).To(MatchJSON("{}"))
		})

		It("Honours retry limit", func() {
			// Run the HTTP/2 server.
			go func() {
				defer GinkgoRecover()

				// Reject the first four connections:
				conn := Accept(listener)
				Reject(conn)
				conn = Accept(listener)
				Reject(conn)
				conn = Accept(listener)
				Reject(conn)
				conn = Accept(listener)
				Reject(conn)
			}()

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Limit(3).
				Interval(100 * time.Millisecond).
				Jitter(0).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Get(address)
			Expect(err).To(HaveOccurred())
			Expect(response).To(BeNil())
			message := err.Error()
			Expect(message).To(ContainSubstring("PROTOCOL_ERROR"))
		})

		It("Honours retry interval", func() {
			// Run the HTTP/2 server.
			go func() {
				defer GinkgoRecover()

				// Reject the first connection:
				conn := Accept(listener)
				Reject(conn)
				start := time.Now()

				// Reject the second connection an verify that it was sent after
				// waiting the configured interval:
				conn = Accept(listener)
				Reject(conn)
				elapsed := time.Since(start)
				Expect(elapsed).To(BeNumerically(">=", 100*time.Millisecond))

				// Reject the third connection and verify that it was sent after
				// waiting the double of the configured interval:
				conn = Accept(listener)
				Reject(conn)
				elapsed = time.Since(start)
				Expect(elapsed).To(BeNumerically(">=", 200*time.Millisecond))

				// Reject the fourth connection and verify that it was sent after
				// waiting the four times the configured interval:
				conn = Accept(listener)
				Reject(conn)
				elapsed = time.Since(start)
				Expect(elapsed).To(BeNumerically(">=", 400*time.Millisecond))
			}()

			// Wrap the transport setting the jitter to zero so that we can reliably
			// measure retry times:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Limit(3).
				Interval(100 * time.Millisecond).
				Jitter(0).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Get(address)
			Expect(err).To(HaveOccurred())
			Expect(response).To(BeNil())
			message := err.Error()
			Expect(message).To(ContainSubstring("PROTOCOL_ERROR"))
		})
	})

	When("Retry is disabled", func() {
		It("Doesn't tolerate error", func() {
			// Run the HTTP/2 server:
			go func() {
				defer GinkgoRecover()
				conn := Accept(listener)
				Reject(conn)
			}()

			// Wrap the transport:
			wrapper, err := NewTransportWrapper().
				Logger(logger).
				Limit(0).
				Interval(100 * time.Millisecond).
				Jitter(0).
				Build(ctx)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = wrapper.Close()
				Expect(err).ToNot(HaveOccurred())
			}()
			client := &http.Client{
				Transport: wrapper.Wrap(transport),
				Timeout:   5 * time.Second,
			}

			// Send the request:
			response, err := client.Get(address)
			Expect(err).To(HaveOccurred())
			Expect(response).To(BeNil())
			message := err.Error()
			Expect(message).To(ContainSubstring("PROTOCOL_ERROR"))
		})
	})
})

var _ = It("Tolerates connection reset by peer", func() {
	var err error

	// Create a context:
	ctx := context.Background()

	// Create a listener:
	listener, address := Listen()
	defer func() {
		err = listener.Close()
		Expect(err).ToNot(HaveOccurred())
	}()

	// Run the server:
	go func() {
		defer GinkgoRecover()

		// Accept the first connection and close it inmediately. This will trigger
		// the `connection reset by peer` error in the client.
		first := Accept(listener)
		Serve(first, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("{"))
			Expect(err).ToNot(HaveOccurred())
			err = first.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		// Accept the second connection and handle it correctly.
		second := Accept(listener)
		Serve(second, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("{}"))
			Expect(err).ToNot(HaveOccurred())
		})
	}()

	// Wrap the transport:
	wrapper, err := NewTransportWrapper().
		Logger(logger).
		Interval(100 * time.Millisecond).
		Jitter(0).
		Build(ctx)
	Expect(err).ToNot(HaveOccurred())
	defer func() {
		err = wrapper.Close()
		Expect(err).ToNot(HaveOccurred())
	}()
	client := &http.Client{
		Transport: wrapper.Wrap(&http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			ForceAttemptHTTP2: true,
		}),
		Timeout: 5 * time.Second,
	}

	// Send the request:
	response, err := client.Get(address)
	Expect(err).ToNot(HaveOccurred())
	Expect(response).ToNot(BeNil())
	Expect(response.StatusCode).To(Equal(http.StatusOK))
	body, err := ioutil.ReadAll(response.Body)
	Expect(err).ToNot(HaveOccurred())
	Expect(body).To(MatchJSON("{}"))
})

var _ = It("Doesn't change request body object", func() {
	var err error

	// Create a context:
	ctx := context.Background()

	// Prepare the server:
	server := MakeTCPServer()
	defer server.Close()
	server.AppendHandlers(
		RespondWithJSON(http.StatusOK, `{}`),
	)

	// Wrap the transport:
	wrapper, err := NewTransportWrapper().
		Logger(logger).
		Jitter(0).
		Build(ctx)
	Expect(err).ToNot(HaveOccurred())
	defer func() {
		err = wrapper.Close()
		Expect(err).ToNot(HaveOccurred())
	}()
	client := &http.Client{
		Transport: wrapper.Wrap(&http.Transport{}),
		Timeout:   5 * time.Second,
	}

	// Send the request:
	body := ioutil.NopCloser(strings.NewReader(`{}`))
	addr, err := url.Parse(server.URL())
	Expect(err).ToNot(HaveOccurred())
	request := &http.Request{
		Method: http.MethodGet,
		URL:    addr,
		Body:   body,
	}
	_, err = client.Do(request)
	Expect(err).ToNot(HaveOccurred())
	Expect(request.Body).To(Equal(body))
})

var _ = It("Tolerates unepected EOF", func() {
	var err error

	// Create a context:
	ctx := context.Background()

	// Create a listener:
	listener, address := Listen()
	defer func() {
		err = listener.Close()
		Expect(err).ToNot(HaveOccurred())
	}()

	// Run the server:
	go func() {
		defer GinkgoRecover()

		// Accept the first connection and handle it with a server that will close
		// the it after sending only half of the response body. This will trigger
		// the `EOF` error in the client.
		first := Accept(listener)
		Serve(first, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("{"))
			Expect(err).ToNot(HaveOccurred())
			err = first.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		// Accept the second connection and handle it correctly.
		second := Accept(listener)
		Serve(second, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("{}"))
			Expect(err).ToNot(HaveOccurred())
		})
	}()

	// Wrap the transport:
	wrapper, err := NewTransportWrapper().
		Logger(logger).
		Interval(100 * time.Millisecond).
		Jitter(0).
		Build(ctx)
	Expect(err).ToNot(HaveOccurred())
	defer func() {
		err = wrapper.Close()
		Expect(err).ToNot(HaveOccurred())
	}()
	client := &http.Client{
		Transport: wrapper.Wrap(&http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			ForceAttemptHTTP2: true,
		}),
		Timeout: 5 * time.Second,
	}

	// Send the request:
	response, err := client.Get(address)
	Expect(err).ToNot(HaveOccurred())
	Expect(response).ToNot(BeNil())
	Expect(response.StatusCode).To(Equal(http.StatusOK))
	body, err := ioutil.ReadAll(response.Body)
	Expect(err).ToNot(HaveOccurred())
	Expect(body).To(MatchJSON("{}"))
})

// Listen creates an HTTP/2 listener.
func Listen() (listener net.Listener, address string) {
	// Create a TLS listener that will be used to process incoming requests
	// simulating an HTTP/2 server:
	listener, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{
		Certificates: []tls.Certificate{
			LocalhostKeyPair(),
		},
		NextProtos: []string{
			http2.NextProtoTLS,
		},
	})
	Expect(err).ToNot(HaveOccurred())

	// Calculate the listener URL:
	address = "https://" + listener.Addr().String()

	return
}

// Accept accepts an HTTP/2 connection.
func Accept(listener net.Listener) net.Conn {
	// Accept the connection and complete the TLS handshake.
	conn, err := listener.Accept()
	Expect(err).ToNot(HaveOccurred())
	err = conn.(*tls.Conn).Handshake()
	Expect(err).ToNot(HaveOccurred())

	// Return the connection:
	return conn
}

// Reject sends an HTTP/2 go away frame to the given connection and then closes it.
func Reject(conn net.Conn) {
	// Read the HTTP2 connection preface, otherwise the client won't reach the part of
	// the code where the protocol error is detected:
	buffer := make([]byte, len(http2.ClientPreface))
	count, err := conn.Read(buffer)
	Expect(err).ToNot(HaveOccurred())
	Expect(count).To(Equal(len(http2.ClientPreface)))
	Expect(string(buffer)).To(Equal(http2.ClientPreface))

	// Send the go away frame:
	framer := http2.NewFramer(conn, conn)
	err = framer.WriteGoAway(0, http2.ErrCodeStreamClosed, nil)
	Expect(err).ToNot(HaveOccurred())

	// Close the connection:
	err = conn.Close()
	Expect(err).ToNot(HaveOccurred())
}

// Serve handles request received from the given connection using a real HTTP/2 server and the given
// handler function.
func Serve(conn net.Conn, handler http.HandlerFunc) {
	server := &http2.Server{}
	server.ServeConn(conn, &http2.ServeConnOpts{
		Handler: handler,
	})
}
