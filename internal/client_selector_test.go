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

package internal

import (
	"context"

	. "github.com/onsi/ginkgo/v2/dsl/core" // nolint
	. "github.com/onsi/gomega"             // nolint
)

var _ = Describe("Create client selector", func() {
	It("Can't be created without a logger", func() {
		selector, err := NewClientSelector().Build(context.Background())
		Expect(err).To(HaveOccurred())
		Expect(selector).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("logger"))
		Expect(message).To(ContainSubstring("mandatory"))
	})
})

var _ = Describe("Select client", func() {
	var (
		ctx      context.Context
		selector *ClientSelector
	)

	BeforeEach(func() {
		var err error

		// Create a context:
		ctx = context.Background()

		// Create the selector:
		selector, err = NewClientSelector().
			Logger(logger).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(selector).ToNot(BeNil())
	})

	AfterEach(func() {
		// Close the selector:
		err := selector.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Reuses client for same TCP address", func() {
		address, err := ParseServerAddress(ctx, "tcp://my.server.com")
		Expect(err).ToNot(HaveOccurred())
		firstClient, err := selector.Select(ctx, address)
		Expect(err).ToNot(HaveOccurred())
		secondClient, err := selector.Select(ctx, address)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondClient).To(BeIdenticalTo(firstClient))
	})

	It("Doesn't reuse client for different TCP addresses", func() {
		firstAddress, err := ParseServerAddress(ctx, "tcp://my.server.com")
		Expect(err).ToNot(HaveOccurred())
		secondAddress, err := ParseServerAddress(ctx, "tcp://your.server.com")
		Expect(err).ToNot(HaveOccurred())
		firstClient, err := selector.Select(ctx, firstAddress)
		Expect(err).ToNot(HaveOccurred())
		secondClient, err := selector.Select(ctx, secondAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondClient == firstClient).To(BeFalse())
	})

	It("Reuses client for different TCP protocols", func() {
		firstAddress, err := ParseServerAddress(ctx, "http://my.server.com")
		Expect(err).ToNot(HaveOccurred())
		secondAddress, err := ParseServerAddress(ctx, "https://my.server.com")
		Expect(err).ToNot(HaveOccurred())
		firstClient, err := selector.Select(ctx, firstAddress)
		Expect(err).ToNot(HaveOccurred())
		secondClient, err := selector.Select(ctx, secondAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondClient == firstClient).To(BeTrue())
	})

	It("Doesn't resuse client for different Unix sockets", func() {
		firstAddress, err := ParseServerAddress(ctx, "unix://my.server.com/my.socket")
		Expect(err).ToNot(HaveOccurred())
		secondAddress, err := ParseServerAddress(ctx, "unix://my.server.com/your.socket")
		Expect(err).ToNot(HaveOccurred())
		firstClient, err := selector.Select(ctx, firstAddress)
		Expect(err).ToNot(HaveOccurred())
		secondClient, err := selector.Select(ctx, secondAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondClient == firstClient).To(BeFalse())
	})
})
