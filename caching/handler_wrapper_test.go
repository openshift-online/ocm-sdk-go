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

package caching

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint
)

var _ = Describe("Handler wrapper", func() {
	var ctx context.Context
	var handler http.Handler

	BeforeEach(func() {
		ctx = context.Background()
	})

	It("Can't be created without a logger", func() {
		wrapper, err := NewHandlerWrapper().Build(ctx)
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("logger"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Can't be created without a cache factory", func() {
		wrapper, err := NewHandlerWrapper().
			Logger(logger).
			CacheFactory(nil).
			Build(ctx)
		Expect(err).To(HaveOccurred())
		Expect(wrapper).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("factory"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Calls wrapped handler", func() {
		// Create the wrapper:
		wrapper, err := NewHandlerWrapper().
			Logger(logger).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Prepare the handler:
		called := false
		handler = wrapper.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.WriteHeader(http.StatusOK)
		}))

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Check that the wrapped handler was called:
		Expect(called).To(BeTrue())
	})

	It("Returns output of wrapped handler", func() {
		// Create the wrapper:
		wrapper, err := NewHandlerWrapper().
			Logger(logger).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Prepare the handler:
		handler = wrapper.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{}`))
			Expect(err).ToNot(HaveOccurred())
		}))

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Check the output:
		Expect(recorder.Code).To(Equal(http.StatusOK))
		Expect(recorder.Body).To(MatchJSON(`{}`))
	})

	It("Creates a memory cache by default", func() {
		// Create the wrapper:
		wrapper, err := NewHandlerWrapper().
			Logger(logger).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Prepare the handler:
		handler = wrapper.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			cache := CacheFromContext(ctx)
			Expect(cache).ToNot(BeNil())
			var memory *MemoryCache
			Expect(cache).To(BeAssignableToTypeOf(memory))
			w.WriteHeader(http.StatusOK)
		}))

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
	})

	It("Uses the provided cache factory", func() {
		// Create the wrapper:
		wrapper, err := NewHandlerWrapper().
			Logger(logger).
			CacheFactory(func(ctx context.Context) (cache Cache, err error) {
				cache, err = NewNopCache().Build(ctx)
				return
			}).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Prepare the handler:
		handler = wrapper.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			cache := CacheFromContext(ctx)
			Expect(cache).ToNot(BeNil())
			var nop *NopCache
			Expect(cache).To(BeAssignableToTypeOf(nop))
			w.WriteHeader(http.StatusOK)
		}))

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
	})

	It("Doesn't create a cache if there is already one in the context", func() {
		// Create the wrapper:
		wrapper, err := NewHandlerWrapper().
			Logger(logger).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Prepare the handler:
		handler = wrapper.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			cache := CacheFromContext(ctx)
			Expect(cache).ToNot(BeNil())
			var nop *NopCache
			Expect(cache).To(BeAssignableToTypeOf(nop))
			w.WriteHeader(http.StatusOK)
		}))

		// Create a nop cache and put it in the context:
		cache, err := NewNopCache().Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		ctx = CacheIntoContext(ctx, cache)

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		request = request.WithContext(ctx)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
	})

	It("Doesn't call the wrapped handler if the cache can't be created", func() {
		// Create the wrapper:
		wrapper, err := NewHandlerWrapper().
			Logger(logger).
			CacheFactory(func(ctx context.Context) (cache Cache, err error) {
				err = fmt.Errorf("myerror")
				return
			}).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Prepare the handler:
		called := false
		handler = wrapper.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
		}))

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Check that the wrapped handler wasn't called:
		Expect(called).To(BeFalse())
	})

	It("Returns a 500 error if the cache can't be created", func() {
		// Create the wrapper:
		wrapper, err := NewHandlerWrapper().
			Logger(logger).
			CacheFactory(func(ctx context.Context) (cache Cache, err error) {
				err = fmt.Errorf("myerror")
				return
			}).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Prepare the handler:
		handler = wrapper.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		// Send the request:
		request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		// Check the response:
		Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
		Expect(recorder.Body).To(MatchJSON(`{
			"kind": "Error",
			"id": "500",
			"reason": "Can't process 'GET' request for path '' due to an internalserver error"
		}`))
	})
})
