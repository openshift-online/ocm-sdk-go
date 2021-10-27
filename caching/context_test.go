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

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint
)

var _ = Describe("Context", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	It("Adds cache to existing context", func() {
		// Create a cache:
		cache, err := NewMemoryCache().Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Add the cache to the context:
		ctx = CacheIntoContext(ctx, cache)
		Expect(ctx).ToNot(BeNil())

		// Check the result:
		result := CacheFromContext(ctx)
		Expect(result).To(Equal(cache))
	})

	It("Creates new context if needed", func() {
		// Create a cache:
		cache, err := NewMemoryCache().Build(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Add the cache to the context:
		ctx = CacheIntoContext(nil, cache) // nolint
		Expect(ctx).ToNot(BeNil())

		// Check the result:
		result := CacheFromContext(ctx)
		Expect(result).To(Equal(cache))
	})

	It("Returns nil if there is no cache in the context", func() {
		result := CacheFromContext(ctx)
		Expect(result).To(BeNil())
	})
})
