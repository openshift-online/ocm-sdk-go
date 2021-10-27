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

// This file contains a cache implementation that stores key/value pairs in memory.

package caching

import (
	"context"
	"sync"
)

// MemoryCacheBuilder contains the data and logic needed to create a memory cache. Don't create
// instances of this type directly, use the NewMemoryCache function instead.
type MemoryCacheBuilder struct {
}

// MemoryCache is an implementation of the cache iterface that stores the keys and values in memory.
type MemoryCache struct {
	values *sync.Map
}

// NewMemoryCache creates a builder that can then be used to configure and create a memory cache.
func NewMemoryCache() *MemoryCacheBuilder {
	return &MemoryCacheBuilder{}
}

// Build uses the data stored in the builder to create a new memory cache.
func (b *MemoryCacheBuilder) Build(ctx context.Context) (result *MemoryCache, err error) {
	// Create and populate the object:
	result = &MemoryCache{
		values: &sync.Map{},
	}

	return
}

// Get is part of the implementation of the Cache interface.
func (c *MemoryCache) Get(ctx context.Context, key interface{}) (value interface{}, ok bool) {
	value, ok = c.values.Load(key)
	return
}

// Put is part of the implementation of the Cache interface.
func (c *MemoryCache) Put(ctx context.Context, key interface{}, value interface{}) {
	c.values.Store(key, value)
}
