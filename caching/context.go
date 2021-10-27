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

// This file contains functions to get from the context and add to the context a cache.

package caching

import (
	"context"
)

// contextKeyType is the type of the keys used to store things in the context.
type contextKeyType int

// contextCacheKey is the key used to store the cache in the context.
const contextCacheKey contextKeyType = iota

// CacheFromContext returns the cache associated to the given context, or nil if no cache is
// associated to the context.
func CacheFromContext(ctx context.Context) Cache {
	value := ctx.Value(contextCacheKey)
	if value != nil {
		return value.(Cache)
	}
	return nil
}

// CacheIntoContext adds the given cache to the context. If the given context is nil a new one will
// be created using the context.Background function.
func CacheIntoContext(ctx context.Context, cache Cache) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, contextCacheKey, cache)
}
