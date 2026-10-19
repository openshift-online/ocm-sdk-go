package internal

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2/dsl/core" // nolint
	. "github.com/onsi/gomega"             // nolint
)

var _ = Describe("TLS version negotiation (behavioral)", func() {
	var (
		ctx    context.Context
		server *httptest.Server
		cas    *x509.CertPool
	)

	BeforeEach(func() {
		ctx = context.Background()

		// A server that will NOT speak anything newer than TLS 1.2:
		server = httptest.NewUnstartedServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		))
		server.TLS = &tls.Config{MaxVersion: tls.VersionTLS12}
		server.StartTLS()

		cas = x509.NewCertPool()
		cas.AddCert(server.Certificate())
	})

	AfterEach(func() {
		server.Close()
	})

	It("Default client (TLS 1.3 floor) REFUSES a TLS 1.2-only server", func() {
		selector, err := NewClientSelector().
			Logger(logger).
			TrustedCAs(cas).
			Insecure(true). // bypass 127.0.0.1 hostname check; version floor still enforced
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() { Expect(selector.Close()).To(Succeed()) }()

		address, err := ParseServerAddress(ctx, server.URL)
		Expect(err).ToNot(HaveOccurred())
		client, err := selector.Select(ctx, address)
		Expect(err).ToNot(HaveOccurred())

		_, err = client.Get(server.URL)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("protocol version"))
	})

	It("Opt-out client (AllowTLS12) CONNECTS to a TLS 1.2-only server", func() {
		selector, err := NewClientSelector().
			Logger(logger).
			TrustedCAs(cas).
			Insecure(true).
			AllowTLS12(true).
			Build(ctx)
		Expect(err).ToNot(HaveOccurred())
		defer func() { Expect(selector.Close()).To(Succeed()) }()

		address, err := ParseServerAddress(ctx, server.URL)
		Expect(err).ToNot(HaveOccurred())
		client, err := selector.Select(ctx, address)
		Expect(err).ToNot(HaveOccurred())

		resp, err := client.Get(server.URL)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(resp.TLS.Version).To(Equal(uint16(tls.VersionTLS12)))
		Expect(resp.Body.Close()).To(Succeed())
	})
})
