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

// This file contains tests for the functions that extract authentication and authorization
// information from contexts.

package authentication

import (
	"context"

	. "github.com/onsi/ginkgo/v2"                      // nolint
	. "github.com/onsi/gomega"                         // nolint
	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

var _ = Describe("Add token to context", func() {
	It("Adds the token", func() {
		token := MakeTokenObject(nil)
		ctx := ContextWithToken(context.TODO(), token)
		extracted := ctx.Value(tokenKeyValue)
		Expect(extracted).To(BeIdenticalTo(token))
	})
})

var _ = Describe("Get token from context", func() {
	It("Succeeds if there is a token", func() {
		token := MakeTokenObject(nil)
		ctx := context.WithValue(context.TODO(), tokenKeyValue, token)
		extracted, err := TokenFromContext(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(extracted).ToNot(BeNil())
		Expect(extracted.Raw).To(Equal(token.Raw))
	})

	It("Succeeds if there is no token", func() {
		ctx := context.TODO()
		extracted, err := TokenFromContext(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(extracted).To(BeNil())
	})
})

var _ = Describe("Get bearer from context", func() {
	It("Succeeds if there is a token", func() {
		token := MakeTokenObject(nil)
		ctx := context.WithValue(context.TODO(), tokenKeyValue, token)
		extracted, err := BearerFromContext(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(extracted).To(Equal(token.Raw))
	})

	It("Succeeds if there is no token", func() {
		ctx := context.TODO()
		extracted, err := BearerFromContext(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(extracted).To(BeEmpty())
	})
})
