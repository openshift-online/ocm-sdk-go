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

package sdk

import (
	"bytes"
	"context"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint
)

var _ = Describe("DumpRoundTripper", func() {
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	var stdLogger *StdLogger
	var d *dumpRoundTripper

	BeforeEach(func() {
		var err error

		// Clear buffers
		stdOut.Reset()
		stdErr.Reset()

		// Create logger which writes to buffers
		stdLogger, err = NewStdLoggerBuilder().
			Streams(&stdOut, &stdErr).
			Debug(true).
			Build()
		Expect(err).ToNot(HaveOccurred())

		d = &dumpRoundTripper{stdLogger, nil}
	})

	It("dumpJson ordering test", func() {
		d := dumpRoundTripper{stdLogger, nil}
		json := `{ "z": 0, "y": null, "a": { "a": 5 }, "b": [ 1, { "a": 5 }, 3], "c": true }`
		d.dumpJSON(context.Background(), []byte(json))
		expectedJSON := `{
  "z": 0,
  "y": null,
  "a": {
    "a": 5
  },
  "b": [
    1,
    {
      "a": 5
    },
    3
  ],
  "c": true
}
`
		Expect(stdOut.String()).To(Equal(expectedJSON))
	})

	It("dumpJson redacting test", func() {
		json := `{ "z": 0, "access_token": null, "a": { "a": 5 }, "b": [ 1, { "password": 5 }, 3], "ssh": true }`
		d.dumpJSON(context.Background(), []byte(json))
		expectedJSON := `{
  "z": 0,
  "access_token": "***",
  "a": {
    "a": 5
  },
  "b": [
    1,
    {
      "password": "***"
    },
    3
  ],
  "ssh": "***"
}
`
		Expect(stdOut.String()).To(Equal(expectedJSON))
	})

	It("dumpJson empty test", func() {
		d := dumpRoundTripper{stdLogger, nil}
		json := ``
		d.dumpJSON(context.Background(), []byte(json))
		expectedJSON := `
`
		Expect(stdOut.String()).To(Equal(expectedJSON))
	})

	It("dumpJson empty object test", func() {
		d := dumpRoundTripper{stdLogger, nil}
		json := `{}`
		d.dumpJSON(context.Background(), []byte(json))
		expectedJSON := "{\n  \n}\n"
		Expect(stdOut.String()).To(Equal(expectedJSON))
	})
})
