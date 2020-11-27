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

package sdk

import (
	"bytes"
	"crypto/rsa"
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"testing"
	"text/template"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/onsi/gomega/ghttp"

	"github.com/openshift-online/ocm-sdk-go/logging"

	. "github.com/onsi/ginkgo" // nolint
	. "github.com/onsi/gomega" // nolint
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client")
}

// Logger used during the tests:
var logger logging.Logger

var _ = BeforeSuite(func() {
	var err error

	// Load the keys used to sign and verify tokens:
	jwtPublicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(jwtPublicKeyPEM))
	Expect(err).ToNot(HaveOccurred())
	jwtPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtPrivateKeyPEM))
	Expect(err).ToNot(HaveOccurred())

	// Create the logger:
	logger, err = logging.NewStdLoggerBuilder().
		Streams(GinkgoWriter, GinkgoWriter).
		Debug(true).
		Build()
	Expect(err).ToNot(HaveOccurred())
})

// MakeServer creates a test server configured so that it sends log messages to the Ginkgo
// writer. It returns the created server and the name of a temporary file that contains the
// CA certificate that the client should trust in order to connect to the server. It is the
// responsibility of the caller to delete this temporary file when it is no longer needed.
func MakeServer() (server *ghttp.Server, ca string) {
	// Create and configure the server:
	server = ghttp.NewTLSServer()
	server.Writer = GinkgoWriter
	server.HTTPTestServer.Config.ErrorLog = log.New(GinkgoWriter, "", log.LstdFlags)

	// Fetch the CA certificate:
	addr, err := url.Parse(server.URL())
	Expect(err).ToNot(HaveOccurred())
	conn, err := tls.Dial("tcp", addr.Host, &tls.Config{
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
	ca = file.Name()

	return
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

// DefaultToken generates a token issued by the default OpenID server and with the given type and
// with the given life. If the life is zero the token will never expire. If the life is positive the
// token will be valid, and expire after that time.  If the life is negative the token will be
// already expired that time ago.
func DefaultToken(typ string, life time.Duration) string {
	return issueToken("https://sso.redhat.com/auth/realms/redhat-external", typ, life)
}

// issueToken generates a token issued by the given issuer, the given type and the given life time.
func issueToken(issuer string, typ string, life time.Duration) string {
	iat := time.Now()
	exp := iat.Add(life)
	claims := jwt.MapClaims{
		"iss": issuer,
		"iat": iat.Unix(),
		"typ": typ,
		"exp": exp.Unix(),
	}
	plain := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := plain.SignedString(jwtPrivateKey)
	Expect(err).ToNot(HaveOccurred())
	return token
}

// Public and private key that will be used to sign and verify tokens in the tests:
var (
	jwtPublicKey  *rsa.PublicKey
	jwtPrivateKey *rsa.PrivateKey
)

// Public key in PEM format:
const jwtPublicKeyPEM = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7bKPFZi7LJ5Oc/XefBDe
byQ1i38Sc3f7Jq0vh8aZC2W6SyqIlv3uUDWFozw0bdkS4MGN6eFjql0JIMIIoq/C
A3aNDCJXKFyVOepe7kgWQ5WY2HH03D/gzUM773TPIkeLCUDJhWi+KMcoMtyxgwr+
X4UVRz/o73fKMrv1bKq7ajAu2Wq1Cjp7zeoirnVz2uplpEtholrySyuhKFmjlRvg
eaLzlc/krB24+IPdJrklGyuwyr8jHDjYBJIsNuqtOzMibdhKPtAhswgZ/lyCFWt+
xAvLsVAJtfNwuED/Cac2KdY60tZzeWsknSuZKL76OARHxlPOWrMsw4jrqpkXM7Ns
LQIDAQAB
-----END PUBLIC KEY-----
`

// Private key in PEM format:
const jwtPrivateKeyPEM = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA7bKPFZi7LJ5Oc/XefBDebyQ1i38Sc3f7Jq0vh8aZC2W6SyqI
lv3uUDWFozw0bdkS4MGN6eFjql0JIMIIoq/CA3aNDCJXKFyVOepe7kgWQ5WY2HH0
3D/gzUM773TPIkeLCUDJhWi+KMcoMtyxgwr+X4UVRz/o73fKMrv1bKq7ajAu2Wq1
Cjp7zeoirnVz2uplpEtholrySyuhKFmjlRvgeaLzlc/krB24+IPdJrklGyuwyr8j
HDjYBJIsNuqtOzMibdhKPtAhswgZ/lyCFWt+xAvLsVAJtfNwuED/Cac2KdY60tZz
eWsknSuZKL76OARHxlPOWrMsw4jrqpkXM7NsLQIDAQABAoIBAQCqZXenTr7XZIDv
JhGhNOKQIA/2eVi7yAYWGs7Y8aijAAEFg285dr3RaBzuAOnA2X1r+7UFNZsh9OHn
RtGz9nUJ0IGacj+y4nPjeb0l6i1zs5lHiKG1BmHcI9eieEVI2Kq2LmiIp6ayStrp
Y0Ypn8bsqNWxJwKQMHqV1iJBDT+fBZuk2kEouRiWCBazIlUgd110v0veWul9eHgy
kqrrhogq2/RdY9+EiKzPSr0A8zdFJHfRGSuk3rK9sUQ9HbvL6ZG+q+O3IUKfO33O
lYniHG+/FmY2ESHqRBSd93/zKqubIhbj0Ha/JiSm66ranGDtqeBqdUZxIVt9TH8U
qltTsKmtAoGBAP61QMouox/AIgLPZ4GUVYc8EtQz3SNioQ24Bn+c4idL30vwewrR
Togz7WDF1EUiL2ljE7Dxxo0FJbEasETZojoXHt7ZCpaSEK+GcvKmV/NdCrprrYhO
GPwuttaEMlOXNtFsjWGhhg+d18aCZTWpIUmsQS8823nwWiGDuYbxlYzXAoGBAO7n
N5P5McAMRi2NdJJXbVvBbC2xwrvOt/qZZnAvFAX4NI088CBj1EGW1b9Mq+uuOfkM
8DP3EkGZ/YixTAIub5xL55CpyU38+aIoKgfLJuk+DHGBJD24qYYTzsXPEFuDvjyU
+0VnbCMInyga6iW2ux3FCYLnowv7yI8GFCmw90qbAoGBAKPgjn0HIKEvBzLzqs7u
V1EZT6wEaoV30dN30YaNs9xArry3TxMYtARiFJqs7fRFGGgf/O1dwbe91hBq8Xp6
5Cun8I7E6lESTXYBdTe12uSTunFTEmWeiejHTZAboh2yLuzzgMuOFyk5DzmDcAbk
eKxkDdSMvVFpWTQzAk1WZjglAoGAQie8+Dj1Ud0UQeD9+thC7DmvnyeuaRthTv/T
ohUnUk/dHY9WX2HFkTQXlJXCtczVOOYgTgOJBqmBz6xpA+Gf/oP2Z9TcbcAz0HeW
y/mxmL0Z7QR56K2OJBawF46zVOQydcw7mIh/JWRpzk1FsZPcVO4PKDTErbjXXOOu
Ca17jSkCgYBRwxm+l3gCduco5byxzMftqyMBm+JUDtFdkQseSzF2YYHW7cPylmi+
Br3bhh0/sYVONO3a0EGr37J6d8pESpVIHsmVKPNsaLb5vMOwE0hAP5Aj83MkFlo5
fD77PZoNGoJiJ9PCF3f7fZSwcAsA1hbulzR/hl5MuRxhybAYbfx6xg==
-----END RSA PRIVATE KEY-----
`
