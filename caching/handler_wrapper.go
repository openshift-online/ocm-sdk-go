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

// This file contains the implementations of a wrapper that wraps HTTP handlers that attatch an
// empty cache to the request context. This cache can then be used to store object that will be
// frequently used for the duration of the request.

package caching

import (
	"context"
	"fmt"
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/logging"
)

// HandlerWrapperBuilder contains the data and logic needed to build a new chacing handler wrapper.
// Don't create objects of this type directly; use the NewHandlerWrapper function instead.
type HandlerWrapperBuilder struct {
	logger       logging.Logger
	cacheFactory CacheFactory
}

// HandlerWrapper contains the data and logic needed to wrap an HTTP handler.
type HandlerWrapper struct {
	logger       logging.Logger
	cacheFactory CacheFactory
}

// handler is an HTTP handler that adds a new cache to the request context.
type handler struct {
	logger       logging.Logger
	cacheFactory CacheFactory
	handler      http.Handler
}

// Make sure that we implement the interface:
var _ http.Handler = (*handler)(nil)

// NewHandlerWrapper creates a new builder that can then be used to configure and create a new
// caching handler wrapper.
func NewHandlerWrapper() *HandlerWrapperBuilder {
	return &HandlerWrapperBuilder{
		cacheFactory: defaultHandlerWrapperCacheFactory,
	}
}

// Logger sets the logger that the handlers will use to write messages to the log.
func (b *HandlerWrapperBuilder) Logger(value logging.Logger) *HandlerWrapperBuilder {
	b.logger = value
	return b
}

// CacheFactory sets a function that the handlers will use to create the cache. For example, to
// configure the handlers so that they create memory caches:
//
//	wrapper, err := caching.NewHandlerWrapper().
//		CacheFactory(func (ctx context.Context) (cache Cache, err error) {
//			cache, err = caching.NewMemoryCache().Build(ctx)
//			return
//		}).
//		Build(ctx)
//	if err != nil {
//		...
//	}
//
// Note that this is just an example, and is not required as the handlers will create memory caches
// by default.
func (b *HandlerWrapperBuilder) CacheFactory(value CacheFactory) *HandlerWrapperBuilder {
	b.cacheFactory = value
	return b
}

// Build uses the information stored in the builder to create a new handler wrapper.
func (b *HandlerWrapperBuilder) Build(ctx context.Context) (result *HandlerWrapper, err error) {
	// Check parameters:
	if b.logger == nil {
		err = fmt.Errorf("logger is mandatory")
		return
	}
	if b.cacheFactory == nil {
		err = fmt.Errorf("cache factory is mandatory")
		return
	}

	// Create and populate the object:
	result = &HandlerWrapper{
		logger:       b.logger,
		cacheFactory: b.cacheFactory,
	}

	return
}

// Wrap creates a new caching handler that wraps the given one.
func (w *HandlerWrapper) Wrap(h http.Handler) http.Handler {
	return &handler{
		logger:       w.logger,
		cacheFactory: w.cacheFactory,
		handler:      h,
	}
}

// ServeHTTP is the implementation of the HTTP handler interface.
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cache := CacheFromContext(ctx)
	if cache == nil {
		var err error
		cache, err = h.cacheFactory(ctx)
		if err != nil {
			h.logger.Error(ctx, "Can't create cache: %v", err)
			errors.SendInternalServerError(w, r)
			return
		}
		ctx = CacheIntoContext(ctx, cache)
		r = r.WithContext(ctx)
	}
	h.handler.ServeHTTP(w, r)
}

// defaultHandlerWrapperCacheFactory is the default function used to create caches.
func defaultHandlerWrapperCacheFactory(ctx context.Context) (cache Cache, err error) {
	cache, err = NewMemoryCache().Build(ctx)
	return
}
