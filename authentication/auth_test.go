package authentication

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/http/httptest"

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

	It("Requires TLS 1.3 by default", func() {
		client, err := NewAuthCodeConfig().oauth2HTTPClient(ctx)
		Expect(err).ToNot(HaveOccurred())

		transport, ok := client.Transport.(*http.Transport)
		Expect(ok).To(BeTrue())
		Expect(transport.TLSClientConfig.MinVersion).To(Equal(uint16(tls.VersionTLS13)))
	})

	It("Lowers the floor to TLS 1.2 when AllowTLS12 is enabled", func() {
		client, err := NewAuthCodeConfig().
			AllowTLS12(true).
			oauth2HTTPClient(ctx)
		Expect(err).ToNot(HaveOccurred())

		transport, ok := client.Transport.(*http.Transport)
		Expect(ok).To(BeTrue())
		Expect(transport.TLSClientConfig.MinVersion).To(Equal(uint16(tls.VersionTLS12)))
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

	It("Refuses a TLS 1.2-only server by default", func() {
		server := httptest.NewUnstartedServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		))
		server.TLS = &tls.Config{MaxVersion: tls.VersionTLS12}
		server.StartTLS()
		defer server.Close()

		cas := x509.NewCertPool()
		cas.AddCert(server.Certificate())

		client, err := NewAuthCodeConfig().
			Logger(logger).
			TrustedCAs(cas).
			Insecure(true). // bypass 127.0.0.1 hostname check; version floor still enforced
			oauth2HTTPClient(ctx)
		Expect(err).ToNot(HaveOccurred())

		_, err = client.Get(server.URL)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("protocol version"))
	})

	It("Connects to a TLS 1.2-only server when AllowTLS12 is enabled", func() {
		server := httptest.NewUnstartedServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		))
		server.TLS = &tls.Config{MaxVersion: tls.VersionTLS12}
		server.StartTLS()
		defer server.Close()

		cas := x509.NewCertPool()
		cas.AddCert(server.Certificate())

		client, err := NewAuthCodeConfig().
			Logger(logger).
			TrustedCAs(cas).
			Insecure(true).
			AllowTLS12(true).
			oauth2HTTPClient(ctx)
		Expect(err).ToNot(HaveOccurred())

		resp, err := client.Get(server.URL)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(resp.TLS.Version).To(Equal(uint16(tls.VersionTLS12)))
		Expect(resp.Body.Close()).To(Succeed())
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

var _ = Describe("Token request TLS minimum version", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	It("Requires TLS 1.3 for the token request client by default", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Client("myclient", "mysecret").
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() { Expect(wrapper.Close()).To(Succeed()) }()

		Expect(wrapper.clientSelector.AllowTLS12()).To(BeFalse())
	})

	It("Propagates the TLS 1.2 opt-out to the token request client", func() {
		wrapper, err := NewTransportWrapper().
			Logger(logger).
			Client("myclient", "mysecret").
			AllowTLS12(true).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() { Expect(wrapper.Close()).To(Succeed()) }()

		Expect(wrapper.clientSelector.AllowTLS12()).To(BeTrue())
	})
})
