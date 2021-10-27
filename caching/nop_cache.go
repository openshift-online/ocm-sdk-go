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

// This file contains a cache implementation that doesn't store key/value pairs. This is intended
// for tests only.

package caching

import (
	"context"
)

// NopCacheBuilder contains the data and logic needed to create a cache that doesn't store key/value
// pairs. Don't create instances of this type directly, use the NewNopCache function instead.
type NopCacheBuilder struct {
}

// NopCache is an implementation of the cache interface that doesn't store key/value pairs.
type NopCache struct {
}

// NewNopCache creates a builder that can then be used to configure and create a nop cache.
func NewNopCache() *NopCacheBuilder {
	return &NopCacheBuilder{}
}

// Build uses the data stored in the builder to create a new nop cache.
func (b *NopCacheBuilder) Build(ctx context.Context) (result *NopCache, err error) {
	// Create and populate the object:
	result = &NopCache{}

	return
}

// Get is part of the implementation of the Cache interface.
func (c *NopCache) Get(ctx context.Context, key interface{}) (value interface{}, ok bool) {
	// This does nothing on purpose.
	return
}

// Put is part of the implementation of the Cache interface.
func (c *NopCache) Put(ctx context.Context, key interface{}, value interface{}) {
	// This does nothing on purpose.
}
