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

// This file contains tests for the support for response body compression.

package sdk

import (
	"compress/gzip"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2/dsl/core" // nolint
	. "github.com/onsi/gomega"             // nolint

	"github.com/onsi/gomega/ghttp"

	. "github.com/openshift-online/ocm-sdk-go/v2/testing" // nolint
)

var _ = Describe("Compression", func() {
	var (
		server     *ghttp.Server
		connection *Connection
	)

	BeforeEach(func() {
		var err error

		// Create the tokens:
		token := MakeTokenString("Bearer", 5*time.Minute)

		// Create the server:
		server = MakeTCPServer()

		// Create the connection:
		connection, err = NewConnection().
			Logger(logger).
			URL(server.URL()).
			Tokens(token).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		// Close the connection:
		err := connection.Close()
		Expect(err).ToNot(HaveOccurred())

		// Stop the server:
		server.Close()
	})

	It("Decompresses response body", func() {
		// Prepare the server:
		server.AppendHandlers(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				encoding := r.Header.Get("Accept-Encoding")
				Expect(encoding).To(Equal("gzip"))
				body := []byte(`{
					"kind": "Cluster",
					"id": "123",
					"href": "/api/clusters_mgmt/v1/clusters/123",
					"name": "mycluster"
				}`)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Content-Encoding", "gzip")
				w.WriteHeader(http.StatusOK)
				compressor := gzip.NewWriter(w)
				_, err := compressor.Write(body)
				Expect(err).ToNot(HaveOccurred())
				err = compressor.Close()
				Expect(err).ToNot(HaveOccurred())
			}),
		)

		// Send the request:
		response, err := connection.ClustersMgmt().V1().Clusters().Cluster("123").Get().
			Send()
		Expect(err).ToNot(HaveOccurred())
		Expect(response).ToNot(BeNil())
		result := response.Body()
		Expect(result.Kind()).To(Equal("Cluster"))
		Expect(result.ID()).To(Equal("123"))
		Expect(result.HREF()).To(Equal("/api/clusters_mgmt/v1/clusters/123"))
		Expect(result.Name()).To(Equal("mycluster"))
	})
})
