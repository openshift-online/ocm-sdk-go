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
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"text/template"

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
		return
	}
}

// VerifyCookie checks that the request contains a cookie with the given name and value.
func VerifyCookie(name, value string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(name)
		Expect(err).ToNot(HaveOccurred())
		Expect(cookie).ToNot(BeNil())
		Expect(cookie.Value).To(Equal(value))
		return
	}
}

// EvaluateTemplate generates a string from the given templlate source and name value pairs. For
// example the following code:
//
//	EvaluateTemplate(
//		`{
//			"access_token": "{{ .AccessToken }}"
//			"refresh_token": "{{ .RefreshToken }}"
//		}`,
//		"AccessToken", "myaccesstoken",
//		"RefreshToken", "myrefreshtoken",
//	)
//
// Will generate the following string:
//
//	{
//		"access_token": "myaccesstoken"
//		"access_token": "myrefreshtoken"
//	}
//
// To simplify embeding of the templates in Go source the function also removes the leading tabs
// from the generated text.
func EvaluateTemplate(source string, args ...interface{}) string {
	// Check that there is an even number of args, and that the first of each pair is a string:
	count := len(args)
	Expect(count%2).To(
		Equal(0),
		"Template '%s' should have an even number of arguments, but it has %d",
		source, count,
	)
	for i := 0; i < count; i = i + 2 {
		name := args[i]
		_, ok := name.(string)
		Expect(ok).To(
			BeTrue(),
			"Argument %d of template '%s' is a key, so it should be a string, "+
				"but its type is %T",
			i, source, name,
		)
	}

	// Put the variables in the map that will be passed as the data object for the execution of
	// the template:
	data := make(map[string]interface{})
	for i := 0; i < count; i = i + 2 {
		name := args[i].(string)
		value := args[i+1]
		data[name] = value
	}

	// Parse the template:
	tmpl, err := template.New("").Parse(source)
	Expect(err).ToNot(
		HaveOccurred(),
		"Can't parse template '%s': %v",
		source, err,
	)

	// Execute the template:
	buffer := new(bytes.Buffer)
	err = tmpl.Execute(buffer, data)
	Expect(err).ToNot(
		HaveOccurred(),
		"Can't execute template '%s': %v",
		source, err,
	)
	result := buffer.String()

	// Remove the leading tabs:
	result = RemoveLeadingTabs(result)

	return result
}

// RemoveLeadingTabs removes the leading tabs from the lines of the given string.
func RemoveLeadingTabs(s string) string {
	return leadingTabsRE.ReplaceAllString(s, "")
}

// leadingTabsRE is the regular expression used to remove leading tabs from strings generated with
// the EvaluateTemplate function.
var leadingTabsRE = regexp.MustCompile(`(?m)^\t*`)
