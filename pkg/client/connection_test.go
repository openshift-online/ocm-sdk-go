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

package client

import (
	"time"

	// nolint
	. "github.com/onsi/ginkgo"
	// nolint
	. "github.com/onsi/gomega"
)

var _ = Describe("Connection", func() {
	It("Can be created with access token", func() {
		accessToken := Token("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Tokens(accessToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with refresh token", func() {
		refreshToken := Token("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Tokens(refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with offline access token", func() {
		offlineToken := Token("Offline", 0)
		connection, err := NewConnectionBuilder().
			Tokens(offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with access and refresh tokens", func() {
		accessToken := Token("Bearer", 5*time.Minute)
		refreshToken := Token("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Tokens(accessToken, refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with access and offline tokens", func() {
		accessToken := Token("Bearer", 5*time.Minute)
		offlineToken := Token("Offline", 10*time.Hour)
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
})
