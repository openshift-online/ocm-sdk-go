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

// This file contains tests for the authentication handler.

package authentication

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"

	"github.com/dgrijalva/jwt-go"
)

var _ = Describe("Handler", func() {
	It("Can't be built without a logger", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("logger"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Can't be built without a service identifier", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("service"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Can't be built without a version identifier", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("version"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Can't be built without at least one keys source", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
	})

	It("Can't be built with a keys file that doesn't exist", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile("/does/not/exist").
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("/does/not/exist"))
	})

	It("Can't be built with a malformed keys URL", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysURL("junk").
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("junk"))
	})

	It("Can't be built with a URL that isn't HTTPS", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysURL("http://api.openshift.com/.well-known/jwks.json").
			Next(next).
			Build()
		Expect(err).To(HaveOccurred())
	})

	It("Can be built with one keys file", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can be built with one keys URL", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Try to create the handler:
		_, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysURL("https://api.openshift.com/.well-known/jwks.json").
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Can't be built without a next handler", func() {
		_, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("next"))
		Expect(err.Error()).To(ContainSubstring("mandatory"))
	})

	It("Rejects request without authorization header", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Request doesn't contain the 'Authorization' header"
		}`))
	})

	It("Rejects bad authorization type", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the token:
		bearer := IssueBearer(nil)

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bad "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Authentication type 'Bad' isn't supported"
		}`))
	})

	It("Rejects bad bearer token", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer bad")
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token is malformed"
		}`))
	})

	It("Rejects expired bearer token", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the expired token:
		bearer := IssueBearer(jwt.MapClaims{
			"exp": time.Now().Add(-1 * time.Hour).Unix(),
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token is expired"
		}`))
	})

	It("Rejects expired bearer token", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the expired token:
		bearer := IssueBearer(jwt.MapClaims{
			"exp": time.Now().Add(-1 * time.Hour).Unix(),
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token is expired"
		}`))
	})

	It("Rejects token without type claim", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the not yet valid token:
		bearer := IssueBearer(jwt.MapClaims{
			"typ": nil,
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token doesn't contain required claim 'typ'"
		}`))
	})

	It("Rejects refresh tokens", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the not yet valid token:
		bearer := IssueBearer(jwt.MapClaims{
			"typ": "Refresh",
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token type 'Refresh' isn't supported"
		}`))
	})

	It("Rejects offline tokens", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the not yet valid token:
		bearer := IssueBearer(jwt.MapClaims{
			"typ": "Offline",
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token type 'Offline' isn't supported"
		}`))
	})

	It("Rejects token without issue date", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the not yet valid token:
		bearer := IssueBearer(jwt.MapClaims{
			"iat": nil,
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token doesn't contain required claim 'iat'"
		}`))
	})

	It("Rejects token without expiration date", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the not yet valid token:
		bearer := IssueBearer(jwt.MapClaims{
			"exp": nil,
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token doesn't contain required claim 'exp'"
		}`))
	})

	It("Rejects token issued in the future", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the not yet valid token:
		now := time.Now()
		iat := now.Add(1 * time.Minute)
		exp := iat.Add(1 * time.Minute)
		bearer := IssueBearer(jwt.MapClaims{
			"iat": iat.Unix(),
			"exp": exp.Unix(),
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token was issued in the future"
		}`))
	})

	It("Rejects token that isn't valid yet", func() {
		// Prepare the next handler, which should not be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the not yet valid token:
		iat := time.Now()
		nbf := iat.Add(1 * time.Minute)
		exp := nbf.Add(1 * time.Minute)
		bearer := IssueBearer(jwt.MapClaims{
			"iat": iat.Unix(),
			"nbf": nbf.Unix(),
			"exp": exp.Unix(),
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Bearer token isn't valid yet"
		}`))
	})

	It("Loads keys from file", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the token:
		bearer := IssueBearer(nil)

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Adds token to the request context", func() {
		// Prepare the token:
		bearer := IssueBearer(nil)

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actual, err := BearerFromContext(r.Context())
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(bearer))
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Doesn't require authorization header for public URL", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actual, err := BearerFromContext(r.Context())
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEmpty())
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Public("^/public(/.*)?$").
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request without the authorization header:
		request := httptest.NewRequest(http.MethodGet, "/public", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Doesn't require authorization header for public URL", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actual, err := BearerFromContext(r.Context())
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEmpty())
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Public("^/public(/.*)?$").
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request without the authorization header:
		request := httptest.NewRequest(http.MethodGet, "/public", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Ignores malformed authorization header for public URL", func() {
		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actual, err := BearerFromContext(r.Context())
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEmpty())
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Public("^/public(/.*)?$").
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad token:
		request := httptest.NewRequest(http.MethodGet, "/public", nil)
		request.Header.Set("Authorization", "Bad junk")
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
	})

	It("Ignores expired token for public URL", func() {
		// Prepare the expired token:
		bearer := IssueBearer(jwt.MapClaims{
			"exp": time.Now().Add(-1 * time.Minute).Unix(),
		})

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer, err := BearerFromContext(r.Context())
			Expect(err).ToNot(HaveOccurred())
			Expect(bearer).To(BeEmpty())
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Public("^/public(/.*)?$").
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad token:
		request := httptest.NewRequest(http.MethodGet, "/public", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
	})

	It("Combines multiple public URLs", func() {
		// Prepare the token:
		bearer := IssueBearer(nil)

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actual, err := BearerFromContext(r.Context())
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEmpty())
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Public("^/public(/.*)?$").
			Public("^/open(/.*)?$").
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send a request for one of the public URLs:
		request := httptest.NewRequest(http.MethodGet, "/public", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusOK))

		// Send a request for another of the public URLs:
		request = httptest.NewRequest(http.MethodGet, "/open", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder = httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Doesn't pass ignored token to next handler for public URL", func() {
		// Prepare the token:
		bearer := IssueBearer(nil)

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actual, err := BearerFromContext(r.Context())
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(BeEmpty())
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Public("^/public(/.*)?$").
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad token:
		request := httptest.NewRequest(http.MethodGet, "/public", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Doesn't load insecure keys by default", func() {
		// Prepare the server:
		server := NewTLSServer()
		server.AppendHandlers(
			RespondWith(http.StatusOK, keysBytes),
		)
		server.SetAllowUnhandledRequests(true)
		defer server.Close()

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the token:
		bearer := IssueBearer(nil)

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysURL(server.URL()).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
	})

	It("Loads insecure keys in insecure mode", func() {
		// Prepare the server that will return the keys:
		server := NewTLSServer()
		server.AppendHandlers(
			RespondWith(http.StatusOK, keysBytes),
		)
		defer server.Close()

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the token:
		bearer := IssueBearer(nil)

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysURL(server.URL()).
			KeysInsecure(true).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad type and a good token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify that the request is rejected:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Returns the response of the next handler", func() {
		// Prepare the token:
		bearer := IssueBearer(nil)

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(`{
				"myfield": "myvalue"
			}`))
			Expect(err).ToNot(HaveOccurred())
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send a request:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
		Expect(recorder.Header().Get("Content-Type")).To(Equal("application/json"))
		Expect(recorder.Body).To(MatchJSON(`{
			"myfield": "myvalue"
		}`))
	})

	It("Accepts token if ACL is empty", func() {
		// Prepare the ACL:
		acl, err := ioutil.TempFile("", "acl-*.yml")
		Expect(err).ToNot(HaveOccurred())
		_, err = acl.WriteString("")
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err := os.Remove(acl.Name())
			Expect(err).ToNot(HaveOccurred())
		}()
		err = acl.Close()
		Expect(err).ToNot(HaveOccurred())

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the token:
		bearer := IssueBearer(nil)
		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			ACLFile(acl.Name()).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Accepts token that matches first ACL item", func() {
		// Prepare the ACL:
		acl, err := ioutil.TempFile("", "acl-*.yml")
		Expect(err).ToNot(HaveOccurred())
		_, err = acl.WriteString(`
                        - claim: email
                          pattern: ^.*@example\.com$
                        - claim: sub
                          pattern: ^f:b3f7b485-7184-43c8-8169-37bd6d1fe4aa:myuser$
                `)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err := os.Remove(acl.Name())
			Expect(err).ToNot(HaveOccurred())
		}()
		err = acl.Close()
		Expect(err).ToNot(HaveOccurred())

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the token:
		bearer := IssueBearer(jwt.MapClaims{
			"email": "jdoe@example.com",
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			ACLFile(acl.Name()).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Accepts token that matches second ACL item", func() {
		// Prepare the ACL:
		acl, err := ioutil.TempFile("", "acl-*.yml")
		Expect(err).ToNot(HaveOccurred())
		_, err = acl.WriteString(`
                        - claim: email
                          pattern: ^.*@example\.com$
                        - claim: sub
                          pattern: ^f:b3f7b485-7184-43c8-8169-37bd6d1fe4aa:myuser$
                `)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err := os.Remove(acl.Name())
			Expect(err).ToNot(HaveOccurred())
		}()
		err = acl.Close()
		Expect(err).ToNot(HaveOccurred())

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the token:
		bearer := IssueBearer(jwt.MapClaims{
			"sub": "f:b3f7b485-7184-43c8-8169-37bd6d1fe4aa:myuser",
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			ACLFile(acl.Name()).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Accepts token that matches second ACL file", func() {
		// Prepare the first ACL:
		firstACL, err := ioutil.TempFile("", "acl-*.yml")
		Expect(err).ToNot(HaveOccurred())
		_, err = firstACL.WriteString(`
                        - claim: email
                          pattern: ^.*@example\.com$
                `)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err := os.Remove(firstACL.Name())
			Expect(err).ToNot(HaveOccurred())
		}()
		err = firstACL.Close()
		Expect(err).ToNot(HaveOccurred())

		// Prepare the first ACL:
		secondACL, err := ioutil.TempFile("", "acl-*.yml")
		Expect(err).ToNot(HaveOccurred())
		_, err = secondACL.WriteString(`
                        - claim: sub
                          pattern: ^f:b3f7b485-7184-43c8-8169-37bd6d1fe4aa:myuser$
                `)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err := os.Remove(secondACL.Name())
			Expect(err).ToNot(HaveOccurred())
		}()
		err = secondACL.Close()
		Expect(err).ToNot(HaveOccurred())

		// Prepare the next handler:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Prepare the token:
		bearer := IssueBearer(jwt.MapClaims{
			"sub": "f:b3f7b485-7184-43c8-8169-37bd6d1fe4aa:myuser",
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			ACLFile(firstACL.Name()).
			ACLFile(secondACL.Name()).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("Rejects token that doesn't match the ACL", func() {
		// Prepare the ACL:
		acl, err := ioutil.TempFile("", "acl-*.yml")
		Expect(err).ToNot(HaveOccurred())
		_, err = acl.WriteString(`
                        - claim: email
                          pattern: ^.*@example\.com$
                        - claim: sub
                          pattern: ^f:b3f7b485-7184-43c8-8169-37bd6d1fe4aa:myuser$
                `)
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err := os.Remove(acl.Name())
			Expect(err).ToNot(HaveOccurred())
		}()
		err = acl.Close()
		Expect(err).ToNot(HaveOccurred())

		// Prepare the next handler, which should never be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the token:
		bearer := IssueBearer(jwt.MapClaims{
			"email": "jdoe@hacker.com",
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			ACLFile(acl.Name()).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer "+bearer)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "401",
			"href": "/api/clusters_mgmt/v1/errors/401",
			"code": "CLUSTERS-MGMT-401",
			"reason": "Access denied"
		}`))
	})

	It("Returns expected headers", func() {
		// Prepare the next handler, which should never be called:
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Expect(true).To(BeFalse())
			w.WriteHeader(http.StatusBadRequest)
		})

		// Prepare the handler:
		handler, err := NewHandler().
			Logger(logger).
			Service("clusters_mgmt").
			Version("v1").
			KeysFile(keysFile).
			Next(next).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Send the request with a bad token:
		request := httptest.NewRequest(http.MethodGet, "/private", nil)
		request.Header.Set("Authorization", "Bearer junk")
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Verify the response:
		Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		header := recorder.Header().Get("WWW-Authenticate")
		Expect(header).To(Equal("Bearer realm=\"clusters_mgmt/v1\""))
	})
})
