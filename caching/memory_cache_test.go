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

var _ = Describe("Memory cache", func() {
	var ctx context.Context
	var cache Cache

	BeforeEach(func() {
		var err error

		// Create a context:
		ctx = context.Background()

		// Create a cache:
		cache, err = NewMemoryCache().Build(ctx)
		Expect(err).ToNot(HaveOccurred())
	})

	It("Get after put", func() {
		cache.Put(ctx, "mykey", "myvalue")
		value, ok := cache.Get(ctx, "mykey")
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal("myvalue"))
	})

	It("Get without put", func() {
		value, ok := cache.Get(ctx, "mykey")
		Expect(ok).To(BeFalse())
		Expect(value).To(BeNil())
	})
})
