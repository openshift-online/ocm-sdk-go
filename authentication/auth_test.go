package authentication

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo/v2/dsl/core"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"                   // nolint
	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

var _ = Describe("Auth code OAuth2 HTTP client", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	It("Verifies TLS certificates by default", func() {
		client, err := NewAuthCodeConfig().oauth2HTTPClient(ctx)
		Expect(err).ToNot(HaveOccurred())

		transport, ok := client.Transport.(*http.Transport)
		Expect(ok).To(BeTrue())
		Expect(transport.TLSClientConfig.InsecureSkipVerify).To(BeFalse())
		Expect(transport.TLSClientConfig.RootCAs).To(BeNil())
	})

	It("Rejects server certificates that aren't trusted", func() {
		server, _ := MakeTCPTLSServer()
		defer server.Close()

		client, err := NewAuthCodeConfig().oauth2HTTPClient(ctx)
		Expect(err).ToNot(HaveOccurred())

		response, err := client.Get(server.URL())
		Expect(err).To(HaveOccurred())
		Expect(response).To(BeNil())
	})

	It("Accepts untrusted server certificates when Insecure is enabled", func() {
		server, _ := MakeTCPTLSServer()
		defer server.Close()
		server.AppendHandlers(RespondWith(http.StatusOK, nil))

		client, err := NewAuthCodeConfig().
			Insecure(true).
			oauth2HTTPClient(ctx)
		Expect(err).ToNot(HaveOccurred())

		response, err := client.Get(server.URL())
		Expect(err).ToNot(HaveOccurred())
		Expect(response.StatusCode).To(Equal(http.StatusOK))
		Expect(response.Body.Close()).To(Succeed())
	})

	It("Trusts custom certificate authorities", func() {
		server, ca := MakeTCPTLSServer()
		defer server.Close()
		server.AppendHandlers(RespondWith(http.StatusOK, nil))

		client, err := NewAuthCodeConfig().
			Logger(logger).
			TrustedCAs(ca).
			oauth2HTTPClient(ctx)
		Expect(err).ToNot(HaveOccurred())

		response, err := client.Get(server.URL())
		Expect(err).ToNot(HaveOccurred())
		Expect(response.StatusCode).To(Equal(http.StatusOK))
		Expect(response.Body.Close()).To(Succeed())

		transport := client.Transport.(*http.Transport)
		Expect(transport.TLSClientConfig.InsecureSkipVerify).To(BeFalse())
	})
})
