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

// This file contains tests for the connection.

package sdk

import (
	"context"
	"net/http"
	"time"

	"github.com/onsi/gomega/gbytes"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint
)

var _ = Describe("Connection", func() {
	It("Can be created with access token", func() {
		accessToken := DefaultToken("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with refresh token", func() {
		refreshToken := DefaultToken("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with offline access token", func() {
		offlineToken := DefaultToken("Offline", 0)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with access and refresh tokens", func() {
		accessToken := DefaultToken("Bearer", 5*time.Minute)
		refreshToken := DefaultToken("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken, refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with access and offline tokens", func() {
		accessToken := DefaultToken("Bearer", 5*time.Minute)
		offlineToken := DefaultToken("Offline", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken, offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with user name and password", func() {
		connection, err := NewConnectionBuilder().
			Logger(logger).
			User("myuser", "mypassword").
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with client identifier and secret", func() {
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Client("myclientid", "myclientsecret").
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Selects default OpenID server with default access token", func() {
		accessToken := DefaultToken("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		tokenURL := connection.TokenURL()
		Expect(tokenURL).To(Equal(DefaultTokenURL))
		clientID, clientSecret := connection.Client()
		Expect(clientID).To(Equal(DefaultClientID))
		Expect(clientSecret).To(Equal(DefaultClientSecret))
	})

	It("Selects default OpenID server with default refresh token", func() {
		refreshToken := DefaultToken("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		tokenURL := connection.TokenURL()
		Expect(tokenURL).To(Equal(DefaultTokenURL))
		clientID, clientSecret := connection.Client()
		Expect(clientID).To(Equal(DefaultClientID))
		Expect(clientSecret).To(Equal(DefaultClientSecret))
	})

	It("Selects default OpenID server with default offline access token", func() {
		offlineToken := DefaultToken("Offline", 0)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		tokenURL := connection.TokenURL()
		Expect(tokenURL).To(Equal(DefaultTokenURL))
		clientID, clientSecret := connection.Client()
		Expect(clientID).To(Equal(DefaultClientID))
		Expect(clientSecret).To(Equal(DefaultClientSecret))
	})

	It("Honours explicitly provided OpenID server with user name and password", func() {
		connection, err := NewConnectionBuilder().
			Logger(logger).
			User("myuser", "mypassword").
			TokenURL(DefaultTokenURL).
			Client(DefaultClientID, DefaultClientSecret).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		tokenURL := connection.TokenURL()
		Expect(tokenURL).To(Equal(DefaultTokenURL))
		clientID, clientSecret := connection.Client()
		Expect(clientID).To(Equal(DefaultClientID))
		Expect(clientSecret).To(Equal(DefaultClientSecret))
	})

	It("Use transport wrapper", func() {
		// Create a connection:
		transport := NewTestTransport()
		connection, err := NewConnectionBuilder().
			Logger(logger).
			User("test", "test").
			TransportWrapper(func(wrapped http.RoundTripper) http.RoundTripper {
				return transport
			}).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()

		// Try to get the tokens using a explicit and short timeout to make the test run
		// faster (by default it takes up to 15 seconds) but give it enough time to retry
		// a few times:
		ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
		_, _, err = connection.TokensContext(ctx)

		// Check that the transport was called at least three times:
		Expect(transport.called).To(BeNumerically(">=", 3))
		Expect(err).To(HaveOccurred())
	})
})

type TestTransport struct {
	called int
}

func (t *TestTransport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	t.called++
	header := http.Header{}
	header.Add("Content-type", "application/json")
	response = &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     header,
		Body:       gbytes.BufferWithBytes([]byte("{}")),
	}
	return response, nil
}

func NewTestTransport() *TestTransport {
	return &TestTransport{called: 0}
}
