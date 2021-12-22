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

package logging

import (
	. "github.com/onsi/ginkgo/v2" // nolint
	. "github.com/onsi/gomega"    // nolint
)

var _ = Describe("List", func() {
	Describe("All", func() {
		DescribeTable("Examples",
			func(items []string, list string) {
				Expect(All(items)).To(Equal(list))
			},
			Entry("Nil", nil, ""),
			Entry("Empty", []string{}, ""),
			Entry("One", []string{"one"}, "'one'"),
			Entry("Two", []string{"one", "two"}, "'one' and 'two'"),
			Entry("Three", []string{"one", "two", "three"}, "'one', 'two' and 'three'"),
		)

		It("Doesn't modify the input", func() {
			items := []string{"a", "b", "c"}
			All(items)
			Expect(items).To(Equal([]string{"a", "b", "c"}))
		})
	})

	Describe("Any", func() {
		DescribeTable("Examples",
			func(items []string, list string) {
				Expect(Any(items)).To(Equal(list))
			},
			Entry("Nil", nil, ""),
			Entry("Empty", []string{}, ""),
			Entry("One", []string{"one"}, "'one'"),
			Entry("Two", []string{"one", "two"}, "'one' or 'two'"),
			Entry("Three", []string{"one", "two", "three"}, "'one', 'two' or 'three'"),
		)

		It("Doesn't modify the input", func() {
			items := []string{"a", "b", "c"}
			Any(items)
			Expect(items).To(Equal([]string{"a", "b", "c"}))
		})
	})
})
