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
			Tokens(accessToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with refresh token", func() {
		refreshToken := DefaultToken("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Tokens(refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with offline access token", func() {
		offlineToken := DefaultToken("Offline", 0)
		connection, err := NewConnectionBuilder().
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
			Tokens(accessToken, offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with user name and password", func() {
		connection, err := NewConnectionBuilder().
			User("myuser", "mypassword").
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with client identifier and secret", func() {
		connection, err := NewConnectionBuilder().
			Client("myclientid", "myclientsecret").
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Selects default OpenID server with default access token", func() {
		accessToken := DefaultToken("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
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
		transport := NewTestTransport()
		connection, err := NewConnectionBuilder().
			User("test", "test").
			TransportWrapper(func(wrapped http.RoundTripper) http.RoundTripper {
				return transport
			}).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		_, _, err = connection.Tokens()
		Expect(transport.called > 2).To(BeTrue()) // it means the retry was called
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
