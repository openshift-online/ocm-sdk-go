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

package testing

import (
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/onsi/gomega/ghttp"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint
)

// MakeTCPServer creates a test server that listens in a TCP socket and configured so that it
// sends log messages to the Ginkgo writer.
func MakeTCPServer() *ghttp.Server {
	server := ghttp.NewUnstartedServer()
	server.Writer = GinkgoWriter
	server.HTTPTestServer.Config.ErrorLog = log.New(GinkgoWriter, "", log.LstdFlags)
	server.HTTPTestServer.EnableHTTP2 = true
	server.HTTPTestServer.Start()
	return server
}

// MakeUnixServer creates a test server that listens in a Unix socket and configured so that it
// sends log messages to the Ginkgo writer. It returns the created server and name of a temporary
// file containing the Unix socket. This file will be in a temporary directory, and the caller is
// resposible for removing the directory once it is no longer needed.
func MakeUnixServer() (server *ghttp.Server, socket string) {
	// Create a temporary directory for the Unix sockets:
	sockets, err := ioutil.TempDir("", "sockets")
	Expect(err).ToNot(HaveOccurred())
	socket = filepath.Join(sockets, "server.socket")

	// Create the listener:
	listener, err := net.Listen("unix", socket)
	Expect(err).ToNot(HaveOccurred())

	// Create and configure the server:
	server = ghttp.NewUnstartedServer()
	server.Writer = GinkgoWriter
	server.HTTPTestServer.Config.ErrorLog = log.New(GinkgoWriter, "", log.LstdFlags)
	server.HTTPTestServer.EnableHTTP2 = true
	server.HTTPTestServer.Listener = listener
	server.HTTPTestServer.Start()

	return
}

// MakeTCPTLSServer creates a test server configured so that it sends log messages to the Ginkgo
// writer. It returns the created server and the name of a temporary file that contains the CA
// certificate that the client should trust in order to connect to the server. It is the
// responsibility of the caller to delete this temporary file when it is no longer needed.
func MakeTCPTLSServer() (server *ghttp.Server, ca string) {
	// Create and configure the server:
	server = ghttp.NewUnstartedServer()
	server.Writer = GinkgoWriter
	server.HTTPTestServer.Config.ErrorLog = log.New(GinkgoWriter, "", log.LstdFlags)
	server.HTTPTestServer.EnableHTTP2 = true
	server.HTTPTestServer.StartTLS()

	// Fetch the CA certificate:
	address, err := url.Parse(server.URL())
	Expect(err).ToNot(HaveOccurred())
	ca = fetchCACertificate("tcp", address.Host)

	return
}

// MakeUnixTLSServer creates a test server that listens in a Unix socket and configured so that it
// sends log messages to the Ginkgo writer. It returns the created server, the name of a temporary
// file that contains the CA certificate that the client should trust in order to connect to the
// server and the name of a directory containing the Unix sockets. This file will be in a temporary
// directory. It is the responsibility of the caller to remove these temporary directories and
// files.
func MakeUnixTLSServer() (server *ghttp.Server, ca, socket string) {
	// Create a temporary directory for the Unix sockets:
	sockets, err := ioutil.TempDir("", "sockets")
	Expect(err).ToNot(HaveOccurred())
	socket = filepath.Join(sockets, "server.socket")

	// Create the listener:
	listener, err := net.Listen("unix", socket)
	Expect(err).ToNot(HaveOccurred())

	// Create and configure the server:
	server = ghttp.NewUnstartedServer()
	server.Writer = GinkgoWriter
	server.HTTPTestServer.Config.ErrorLog = log.New(GinkgoWriter, "", log.LstdFlags)
	server.HTTPTestServer.EnableHTTP2 = true
	server.HTTPTestServer.Listener = listener
	server.HTTPTestServer.StartTLS()

	// Fetch the CA certificate:
	ca = fetchCACertificate("unix", socket)

	return
}

// MakeTCPH2CServer creates a test server that supports HTTP/2 without TLS, configured so that it
// sends log messages to the Ginkgo writer.
func MakeTCPH2CServer() *ghttp.Server {
	// Create the server that supports HTTP/2 without TLS:
	h2s := &http2.Server{}

	// Create the regular server:
	server := ghttp.NewUnstartedServer()
	server.Writer = GinkgoWriter
	server.HTTPTestServer.Config.ErrorLog = log.New(GinkgoWriter, "", log.LstdFlags)
	server.HTTPTestServer.EnableHTTP2 = true

	// Wrap the handler of the regular server with the handler that detects HTTP/2 requests
	// without TLS and delegates them to the HTTP/2 server that supports that:
	server.HTTPTestServer.Config.Handler = h2c.NewHandler(server.HTTPTestServer.Config.Handler, h2s)

	// Start the server:
	server.Start()

	return server
}

// MakeUnixH2cServer creates a test server that listens in a Unix socket and supports HTTP/2 without
// TLS, configured so that it sends log messages to the Ginkgo writer. It returns the created server
// and name of a temporary file containing the Unix socket. This file will be in a temporary
// directory, and the caller is resposible for removing the directory once it is no longer needed.
func MakeUnixH2CServer() (server *ghttp.Server, socket string) {
	// Create a temporary directory for the Unix sockets:
	sockets, err := ioutil.TempDir("", "sockets")
	Expect(err).ToNot(HaveOccurred())
	socket = filepath.Join(sockets, "server.socket")

	// Create the listener:
	listener, err := net.Listen("unix", socket)
	Expect(err).ToNot(HaveOccurred())

	// Create the server that supports HTTP/2 without TLS:
	h2s := &http2.Server{}

	// Create the regular server:
	server = ghttp.NewUnstartedServer()
	server.Writer = GinkgoWriter
	server.HTTPTestServer.Config.ErrorLog = log.New(GinkgoWriter, "", log.LstdFlags)
	server.HTTPTestServer.EnableHTTP2 = true
	server.HTTPTestServer.Listener = listener

	// Wrap the handler of the regular server with the handler that detects HTTP/2 requests
	// without TLS and delegates them to the HTTP/2 server that supports that:
	server.HTTPTestServer.Config.Handler = h2c.NewHandler(server.HTTPTestServer.Config.Handler, h2s)

	// Start the server:
	server.Start()

	return
}

// fetchCACertificates connects to the given network address and extracts the CA certificate from
// the TLS handshake. It returns the path of a temporary file containing that CA certificate encoded
// in PEM format. It is the responsibility of the caller to delete that file when it is no longer
// needed.
func fetchCACertificate(network, address string) string {
	// Connect to the server and do the TLS handshake to obtain the certificate chain:
	conn, err := tls.Dial(network, address, &tls.Config{
		InsecureSkipVerify: true, // nolint
	})
	Expect(err).ToNot(HaveOccurred())
	defer func() {
		err = conn.Close()
		Expect(err).ToNot(HaveOccurred())
	}()
	err = conn.Handshake()
	Expect(err).ToNot(HaveOccurred())
	certs := conn.ConnectionState().PeerCertificates
	Expect(certs).ToNot(BeNil())
	Expect(len(certs)).To(BeNumerically(">=", 1))
	cert := certs[len(certs)-1]
	Expect(cert).ToNot(BeNil())

	// Serialize the CA certificate:
	Expect(cert.Raw).ToNot(BeNil())
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}
	buffer := pem.EncodeToMemory(block)
	Expect(buffer).ToNot(BeNil())

	// Store the CA certificate in a temporary file:
	file, err := ioutil.TempFile("", "*.test.ca")
	Expect(err).ToNot(HaveOccurred())
	_, err = file.Write(buffer)
	Expect(err).ToNot(HaveOccurred())
	err = file.Close()
	Expect(err).ToNot(HaveOccurred())

	// Return the path of the temporary file:
	return file.Name()
}

// RespondeWithContent responds with the given status code, content type and body.
func RespondWithContent(status int, contentType, body string) http.HandlerFunc {
	return ghttp.RespondWith(
		status,
		body,
		http.Header{
			"Content-Type": []string{
				contentType,
			},
		},
	)
}

// RespondWithJSON responds with the given status code and JSON body.
func RespondWithJSON(status int, body string) http.HandlerFunc {
	return RespondWithContent(status, "application/json", body)
}

// RespondWithJSONTemplate responds with the given status code and with a JSON body that is
// generated from the given template and arguments. See the EvaluateTemplate function for details
// on how the template and the arguments are combined.
func RespondWithJSONTemplate(status int, source string, args ...interface{}) http.HandlerFunc {
	return RespondWithJSON(status, EvaluateTemplate(source, args...))
}

// RespondWithCookie responds to the request adding a cookie with the given name and value.
func RespondWithCookie(name, value string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:  name,
			Value: value,
		})
	}
}

// VerifyCookie checks that the request contains a cookie with the given name and value.
func VerifyCookie(name, value string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(name)
		Expect(err).ToNot(HaveOccurred())
		Expect(cookie).ToNot(BeNil())
		Expect(cookie.Value).To(Equal(value))
	}
}

// LocalhostKeyPair returns a TLS key pair valid for the name `localhost` DNS
// name, for the `127.0.0.1` IPv4 address and for the `::1` IPv6 address. The
// key pair has been generated with the following command:
//
//	openssl req \
//	-x509 \
//	-newkey rsa:4096 \
//	-nodes \
//	-keyout tls.key \
//	-out tls.crt \
//	-subj '/CN=localhost' \
//	-addext 'subjectAltName=DNS:localhost,IP:127.0.0.1,IP:::1' \
//	-days 3650
func LocalhostKeyPair() tls.Certificate {
	pair, err := tls.X509KeyPair(localhostCrt, localhostKey)
	Expect(err).ToNot(HaveOccurred())
	return pair
}

var localhostCrt = []byte(`
-----BEGIN CERTIFICATE-----
MIIFODCCAyCgAwIBAgIUdhdcmZ5JmTWpGQiuLnTMT6UxZb0wDQYJKoZIhvcNAQEL
BQAwFDESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTIxMDgxMjEwMjQyM1oXDTMxMDgx
MDEwMjQyM1owFDESMBAGA1UEAwwJbG9jYWxob3N0MIICIjANBgkqhkiG9w0BAQEF
AAOCAg8AMIICCgKCAgEAs4I2qpFnkss60FIxKTboLv6QpBt8QnR4l2xT1egCBr3x
kqOkobZjrfZ61JbDyz9ZPcrwzksB+6/Xp8VNmbCvYx8hKAv0EMBTl3Lczv7jIrg8
wXXncjSguRolXkeuQAHqk37CkVhEjfZuvd4ZCoumQXzmw61AZlxDOHDMDM3Um95m
ZYNtVwDn/eZRFIPQbmCcDKC2v2/nlpuviUOuMwnPU/ev4eLaDFUqAO0llTVpJKdb
E0KIyWX7fGFTtibfw5azEGxPEb/7FAUtaa3Rms70I7fwZU/FBy8iXuSzW/g7Ms3a
zZo+rTQFCbNHdU22w4bhZtBvZkXyyDEtxypx77UZZIrKOVflu/VU2lNQS98LnaXF
GtV0qVQ7hFdDRw/XmxDZPsG8xDORgx7PADPTqj0Hylg2oawPZdidqkqP5hgERuWk
vBTMeJ6RjoKLV70FSnIzlOh0fFnln6iZB4dDXE2bLtHtB3oRDrN/VA+N1bCPZmWY
7/s9OqlQ5xRsT9sz+iqyzaOk0XCRV3z+QgohmkOd0GCoZ+QDpogMaaXy0hQlFpaX
hTBO/0FPXzg2ipop7ItgDZYQgl5wuva40R6j1KRqBZLNMuZhosIQ4qf9geeVbZfu
A+QNUSJPrtXXNaCCstiIGJgHWGl1KWBHnmhLSb7yajEFfmQM+yVb/T/TEacjOYMC
AwEAAaOBgTB/MB0GA1UdDgQWBBRK6ctms4WQmh8Qy+VKSHLgc74z3jAfBgNVHSME
GDAWgBRK6ctms4WQmh8Qy+VKSHLgc74z3jAPBgNVHRMBAf8EBTADAQH/MCwGA1Ud
EQQlMCOCCWxvY2FsaG9zdIcEfwAAAYcQAAAAAAAAAAAAAAAAAAAAATANBgkqhkiG
9w0BAQsFAAOCAgEAQhcYqXuWIK0/4oJfJBabbtUUEN1As98gaomcQKMTk30K+aZD
KejPeBZnOJSWpZRl4ypm4Z3u/jIHptLHKSSoSIyEIikV1MsZy9c8E/aetRbrJ4G/
Y+TM/jCKZ9Rz6zxZznN52fME7Gh9de75QsUrmDl8wmgZ9Y09UYX6hG4+rCR0R+lX
GQv3LAVCHXGtCqHS8zXXyymRQghqE8Tz7dzWtpo0AIVA0k8GL4bGxlAAuQL7TXHE
jfLNh35RjULxM6JONYj/uQnaret5kE5ZuzQp+L4HHs/5qIE2wvTyFKyAtJP0E6AP
hI754OEvOVMXfBOYWDalOcyka9kKUCTkDztW5VY/bpapqfTLN6McRqzEijLatg2c
ADHabdIbhKhKKfH1ZmZFXthrPiPfc+Y7oTsf7gmIzR3actjfI6ZWaUhhaipFil3c
7sedG68X7W7I4F5cdPTbZb3fnUb7Xk2mReivDgY2Cz1rMVBYI7wT+m82oPTRwsRU
YZrbnOIXoWmTo6hVtvi7XXNVFj2vDgVQTSqJBUWl2DKqUQ+EALj9AVnDD4TS9Af1
v7215I4iIaeVbmilS7GCoayiPczJfKDpMUf4OEV5r4EomJwWzgqxx5qioRpRvU9k
9kFc5TTcHQENjK3Og/Ii5jpQegAo01sduDnCN05YfAZ2En6M/rqjWG/WiyY=
-----END CERTIFICATE-----
`)

var localhostKey = []byte(`
-----BEGIN PRIVATE KEY-----
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
kQa+wVDL7TzyYOgTVJ1YtorASxSqJBg=
-----END PRIVATE KEY-----
`)
