package testing

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/onsi/gomega/ghttp"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"

	. "github.com/onsi/ginkgo/v2/dsl/core" // nolint
	. "github.com/onsi/gomega"             // nolint
)

var _ = Describe("RespondWithOCMObjectMarshal", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = MakeTCPServer()
	})

	AfterEach(func() {
		server.Close()
	})

	It("should respond with JSON encoded object using provided marshal function", func() {
		// Create a mock cluster object (we'll use a simple struct for testing)
		cluster, err := v1.NewCluster().
			ID("test-cluster").
			Name("Test Cluster").
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Configure the server to respond with the cluster
		server.AppendHandlers(
			RespondWithOcmObjectMarshal(http.StatusOK, cluster, v1.MarshalCluster),
		)

		// Make a request
		resp, err := http.Get(server.URL())
		Expect(err).ToNot(HaveOccurred())
		DeferCleanup(resp.Body.Close)

		// Verify response
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(resp.Header.Get("Content-Type")).To(Equal("application/json"))

		// Read and verify the response body contains JSON
		body, err := io.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())
		Expect(string(body)).To(ContainSubstring("test-cluster"))
		Expect(string(body)).To(ContainSubstring("Test Cluster"))
	})

	It("should handle nil objects gracefully", func() {
		server.AppendHandlers(
			RespondWithOcmObjectMarshal(http.StatusOK, nil, v1.MarshalCluster),
		)

		resp, err := http.Get(server.URL())
		Expect(err).ToNot(HaveOccurred())
		DeferCleanup(resp.Body.Close)

		Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
	})

	It("should handle invalid marshal function", func() {
		cluster := &v1.Cluster{}
		server.AppendHandlers(
			RespondWithOcmObjectMarshal(http.StatusOK, cluster, "not-a-function"),
		)

		resp, err := http.Get(server.URL())
		Expect(err).ToNot(HaveOccurred())
		DeferCleanup(resp.Body.Close)

		Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
	})

	It("should work as a direct HTTP handler", func() {
		// Create a test cluster
		cluster, err := v1.NewCluster().
			ID("test-cluster-123").
			Name("Test Cluster for Unit Test").
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Create a test request
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		reqRecorder := httptest.NewRecorder()

		// Create the handler
		handler := RespondWithOcmObjectMarshal(http.StatusOK, cluster, v1.MarshalCluster)

		// Call the handler
		handler(reqRecorder, req)

		// Check the status code
		Expect(reqRecorder.Code).To(Equal(http.StatusOK))

		// Check the content type
		Expect(reqRecorder.Header().Get("Content-Type")).To(Equal("application/json"))

		// Check that the response body contains expected data
		body := reqRecorder.Body.String()
		Expect(body).To(ContainSubstring("test-cluster-123"))
		Expect(body).To(ContainSubstring("Test Cluster for Unit Test"))
	})

	It("should allow direct marshal function testing", func() {
		// Create a test cluster
		cluster, err := v1.NewCluster().
			ID("direct-test").
			Name("Direct Marshal Test").
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Test marshaling directly
		var buf bytes.Buffer
		err = v1.MarshalCluster(cluster, &buf)
		Expect(err).ToNot(HaveOccurred())

		// Verify the output
		result := buf.String()
		Expect(result).To(ContainSubstring("direct-test"))
		Expect(result).To(ContainSubstring("Direct Marshal Test"))
	})
})
