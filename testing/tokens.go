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
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	. "github.com/onsi/gomega" // nolint
)

// MakeTokenObject generates a token with the claims resulting from merging the default claims and
// the claims explicitly given.
func MakeTokenObject(claims jwt.MapClaims) *jwt.Token {
	merged := jwt.MapClaims{}
	for name, value := range MakeClaims() {
		merged[name] = value
	}
	for name, value := range claims {
		if value == nil {
			delete(merged, name)
		} else {
			merged[name] = value
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, merged)
	token.Header["kid"] = "123"
	var err error
	token.Raw, err = token.SignedString(jwtPrivateKey)
	Expect(err).ToNot(HaveOccurred())
	return token
}

// MakeClaims generates a default set of claims to be used to issue a token.
func MakeClaims() jwt.MapClaims {
	iat := time.Now()
	exp := iat.Add(1 * time.Minute)
	return jwt.MapClaims{
		"iss": "https://sso.redhat.com/auth/realms/redhat-external",
		"iat": iat.Unix(),
		"typ": "Bearer",
		"exp": exp.Unix(),
	}
}

// MakeTokenString generates a token issued by the default OpenID server and with the given type and
// with the given life. If the life is zero the token will never expire. If the life is positive the
// token will be valid, and expire after that time. If the life is negative the token will be
// already expired that time ago.
func MakeTokenString(typ string, life time.Duration) string {
	token := MakeTokenObject(jwt.MapClaims{
		"typ": typ,
		"exp": time.Now().Add(life).Unix(),
	})
	return token.Raw
}

func RespondWithTokens(accessToken, refreshToken string) http.HandlerFunc {
	return RespondWithJSONTemplate(
		http.StatusOK,
		`{
			"access_token": "{{ .AccessToken }}",
			"refresh_token": "{{ .RefreshToken }}"
		}`,
		"AccessToken", accessToken,
		"RefreshToken", refreshToken,
	)
}

func RespondWithTokenError(err, description string) http.HandlerFunc {
	return RespondWithJSONTemplate(
		http.StatusUnauthorized,
		`{
			"error": "{{ .Error }}",
			"error_description": "{{ .Description }}"
		}`,
		"Error", err,
		"Description", description,
	)
}

// DefaultJWKS generates the JSON web key set used for tests.
func DefaultJWKS() []byte {
	// Create a temporary file containing the JSON web key set:
	bigE := big.NewInt(int64(jwtPublicKey.E))
	bigN := jwtPublicKey.N
	return []byte(fmt.Sprintf(
		`{
			"keys": [{
				"kid": "123",
				"kty": "RSA",
				"alg": "RS256",
				"e": "%s",
				"n": "%s"
			}]
		}`,
		base64.RawURLEncoding.EncodeToString(bigE.Bytes()),
		base64.RawURLEncoding.EncodeToString(bigN.Bytes()),
	))
}

// Public and private key that will be used to sign and verify tokens in the tests:
var (
	jwtPublicKey  *rsa.PublicKey
	jwtPrivateKey *rsa.PrivateKey
)

func init() {
	var err error

	// Load the keys used to sign and verify tokens:
	jwtPublicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(jwtPublicKeyPEM))
	if err != nil {
		panic(err)
	}
	jwtPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtPrivateKeyPEM))
	if err != nil {
		panic(err)
	}
}

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
