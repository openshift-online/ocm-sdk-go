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

// This file contains the cache interface.

package caching

import (
	"context"
)

// Cache is an object that knows how to store key/value pairs.
type Cache interface {
	// Get returns the value that is associated with the given key and a boolean indicating if
	// there is such an object.
	Get(ctx context.Context, key interface{}) (value interface{}, ok bool)

	// Put puts in the cache the given given value associated to the given key.
	Put(ctx context.Context, key, value interface{})
}

// CacheFactory is a function that knows how to create caches.
type CacheFactory func(ctx context.Context) (cache Cache, err error)
