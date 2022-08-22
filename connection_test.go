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

// This file contains tests for the connection.

package sdk

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2/dsl/core" // nolint
	. "github.com/onsi/gomega"             // nolint
	. "github.com/onsi/gomega/gbytes"      // nolint

	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

var _ = Describe("Connection", func() {
	It("Can be created with access token", func() {
		accessToken := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with refresh token", func() {
		refreshToken := MakeTokenString("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with offline access token", func() {
		offlineToken := MakeTokenString("Offline", 0)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with access and refresh tokens", func() {
		accessToken := MakeTokenString("Bearer", 5*time.Minute)
		refreshToken := MakeTokenString("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken, refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with access and offline tokens", func() {
		accessToken := MakeTokenString("Bearer", 5*time.Minute)
		offlineToken := MakeTokenString("Offline", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken, offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with user name and password", func() {
		connection, err := NewConnectionBuilder().
			Logger(logger).
			User("myuser", "mypassword").
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with client identifier and secret", func() {
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Client("myclientid", "myclientsecret").
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with metrics subsystem", func() {
		accessToken := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken).
			MetricsSubsystem("my_subsystem").
			Build()
		Expect(err).ToNot(HaveOccurred())
		Expect(connection).ToNot(BeNil())
		defer func() {
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())
		}()
		Expect(connection.MetricsSubsystem()).To(Equal("my_subsystem"))
	})

	It("Selects default OpenID server with default access token", func() {
		accessToken := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(accessToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		tokenURL := connection.TokenURL()
		Expect(tokenURL).To(Equal(DefaultTokenURL))
		clientID, clientSecret := connection.Client()
		Expect(clientID).To(Equal(DefaultClientID))
		Expect(clientSecret).To(Equal(DefaultClientSecret))
	})

	It("Selects default OpenID server with default refresh token", func() {
		refreshToken := MakeTokenString("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(refreshToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		tokenURL := connection.TokenURL()
		Expect(tokenURL).To(Equal(DefaultTokenURL))
		clientID, clientSecret := connection.Client()
		Expect(clientID).To(Equal(DefaultClientID))
		Expect(clientSecret).To(Equal(DefaultClientSecret))
	})

	It("Selects default OpenID server with default offline access token", func() {
		offlineToken := MakeTokenString("Offline", 0)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		tokenURL := connection.TokenURL()
		Expect(tokenURL).To(Equal(DefaultTokenURL))
		clientID, clientSecret := connection.Client()
		Expect(clientID).To(Equal(DefaultClientID))
		Expect(clientSecret).To(Equal(DefaultClientSecret))
	})

	It("Honours explicitly provided OpenID server with user name and password", func() {
		connection, err := NewConnectionBuilder().
			Logger(logger).
			User("myuser", "mypassword").
			TokenURL(DefaultTokenURL).
			Client(DefaultClientID, DefaultClientSecret).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()
		tokenURL := connection.TokenURL()
		Expect(tokenURL).To(Equal(DefaultTokenURL))
		clientID, clientSecret := connection.Client()
		Expect(clientID).To(Equal(DefaultClientID))
		Expect(clientSecret).To(Equal(DefaultClientSecret))
	})

	It("Use transport wrapper", func() {
		// Create a connection:
		transport := NewTestTransport()
		connection, err := NewConnectionBuilder().
			Logger(logger).
			User("test", "test").
			TransportWrapper(func(wrapped http.RoundTripper) http.RoundTripper {
				return transport
			}).
			Build()
		Expect(err).ToNot(HaveOccurred())
		defer connection.Close()

		// Try to get the tokens using a explicit and short timeout to make the test run
		// faster (by default it takes up to 15 seconds) but give it enough time to retry
		// a few times:
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, _, err = connection.TokensContext(ctx)

		// Check that the transport was called at least three times:
		Expect(transport.called).To(BeNumerically(">=", 3))
		Expect(err).To(HaveOccurred())
	})

	It("Can be created with one alternative URL", func() {
		// Create the connection:
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(token).
			URL("https://my.server.com").
			AlternativeURL("/api/clusters_mgmt", "https://your.server.com").
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Check that the URLs are set:
		Expect(connection.URL()).To(Equal("https://my.server.com"))
		alternativeURLs := connection.AlternativeURLs()
		Expect(alternativeURLs).To(HaveLen(1))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/clusters_mgmt",
			"https://your.server.com",
		))
	})

	It("Can be created with two alternative URLs", func() {
		// Create the connection:
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(token).
			URL("https://my.server.com").
			AlternativeURL("/api/clusters_mgmt", "https://your.server.com").
			AlternativeURL("/api/accounts_mgmt", "https://her.server.com").
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Check that the URLs are set:
		Expect(connection.URL()).To(Equal("https://my.server.com"))
		alternativeURLs := connection.AlternativeURLs()
		Expect(alternativeURLs).To(HaveLen(2))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/clusters_mgmt",
			"https://your.server.com",
		))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/accounts_mgmt",
			"https://her.server.com",
		))
	})

	It("Can be created with a map of alternative URLs", func() {
		// Create the connection:
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(token).
			URL("https://my.server.com").
			AlternativeURLs(map[string]string{
				"/api/clusters_mgmt": "https://your.server.com",
				"/api/accounts_mgmt": "https://her.server.com",
			}).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Check that the URLs are set:
		Expect(connection.URL()).To(Equal("https://my.server.com"))
		alternativeURLs := connection.AlternativeURLs()
		Expect(alternativeURLs).To(HaveLen(2))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/clusters_mgmt",
			"https://your.server.com",
		))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/accounts_mgmt",
			"https://her.server.com",
		))
	})

	It("Altering returned alternative URLs doesn't affect internal state", func() {
		// Create the connection:
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(token).
			URL("https://my.server.com").
			AlternativeURLs(map[string]string{
				"/api/clusters_mgmt": "https://your.server.com",
				"/api/accounts_mgmt": "https://her.server.com",
			}).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Try to modify the returned map of alternative URLs:
		alternativeURLs := connection.AlternativeURLs()
		alternativeURLs["/api/service_logs"] = "https://his.server.com"

		// Check that map used internall hasn't changed:
		alternativeURLs = connection.AlternativeURLs()
		Expect(alternativeURLs).To(HaveLen(2))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/clusters_mgmt",
			"https://your.server.com",
		))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/accounts_mgmt",
			"https://her.server.com",
		))
	})

	It("Can't be created with invalid alternative URL prefix", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		_, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(token).
			AlternativeURL("junk", "https://api.openshift.com").
			Build()
		Expect(err).To(HaveOccurred())
	})

	It("Can't be created with invalid alternative URL", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		_, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(token).
			AlternativeURL("/api/clusters_mgmt", ":junk").
			Build()
		Expect(err).To(HaveOccurred())
	})

	It("Can be configured with a YAML string", func() {
		// Create temporary files for the trusted CAs:
		tmp, err := os.MkdirTemp("", "*.test.cas")
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = os.RemoveAll(tmp)
			Expect(err).ToNot(HaveOccurred())
		}()
		err = os.WriteFile(filepath.Join(tmp, "myca.pem"), mycaPEM, 0600)
		Expect(err).ToNot(HaveOccurred())
		err = os.WriteFile(filepath.Join(tmp, "yourca.pem"), yourcaPEM, 0600)
		Expect(err).ToNot(HaveOccurred())

		// Create the YAML configuration string:
		fileAccess := MakeTokenString("Bearer", 5*time.Minute)
		fileRefresh := MakeTokenString("Refresh", 10*time.Hour)
		content := EvaluateTemplate(
			`
			url: https://my.server.com
			alternative_urls:
			  /api/clusters_mgmt: https://your.server.com
			  /api/accounts_mgmt: https://her.server.com
			token_url: https://openid.server.com
			user: myuser
			password: mypassword
			client_id: myclient
			client_secret: mysecret
			tokens:
			- {{ .AccessToken }}
			- {{ .RefreshToken }}
			scopes:
			- openid
			- myscope
			insecure: true
			trusted_cas:
			- {{ .Tmp }}/myca.pem
			- {{ .Tmp }}/yourca.pem
			agent: myagent
			retry_limit: 4
			metrics_subsystem: mysubsystem
			`,
			"Tmp", tmp,
			"AccessToken", fileAccess,
			"RefreshToken", fileRefresh,
		)

		// Create the connection and verify it has been created with the configuration
		// stored in the YAML string:
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Load(content).
			Build()
		Expect(err).ToNot(HaveOccurred())
		Expect(connection.URL()).To(Equal("https://my.server.com"))
		alternativeURLs := connection.AlternativeURLs()
		Expect(alternativeURLs).To(HaveLen(2))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/clusters_mgmt", "https://your.server.com",
		))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/accounts_mgmt", "https://her.server.com",
		))
		Expect(connection.TokenURL()).To(Equal("https://openid.server.com"))
		user, password := connection.User()
		Expect(user).To(Equal("myuser"))
		Expect(password).To(Equal("mypassword"))
		client, secret := connection.Client()
		Expect(client).To(Equal("myclient"))
		Expect(secret).To(Equal("mysecret"))
		returnedAccess, returnedRefresh, err := connection.Tokens()
		Expect(err).ToNot(HaveOccurred())
		Expect(returnedAccess).To(Equal(fileAccess))
		Expect(returnedRefresh).To(Equal(fileRefresh))
		defer func() {
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())
		}()
		Expect(connection.Scopes()).To(ConsistOf("openid", "myscope"))
		Expect(connection.Insecure()).To(BeTrue())
		Expect(connection.Agent()).To(Equal("myagent"))
		Expect(connection.RetryLimit()).To(Equal(4))
		Expect(connection.MetricsSubsystem()).To(Equal("mysubsystem"))
	})

	It("Can be configured with a YAML file", func() {
		// Create temporary files for the trusted CAs:
		tmp, err := os.MkdirTemp("", "*.test.cas")
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = os.RemoveAll(tmp)
			Expect(err).ToNot(HaveOccurred())
		}()
		err = os.WriteFile(filepath.Join(tmp, "myca.pem"), mycaPEM, 0600)
		Expect(err).ToNot(HaveOccurred())
		err = os.WriteFile(filepath.Join(tmp, "yourca.pem"), yourcaPEM, 0600)
		Expect(err).ToNot(HaveOccurred())

		// Create a temporary YAML file containing the configuration:
		fileAccess := MakeTokenString("Bearer", 5*time.Minute)
		fileRefresh := MakeTokenString("Refresh", 10*time.Hour)
		content := EvaluateTemplate(
			`
			url: https://my.server.com
			alternative_urls:
			  /api/clusters_mgmt: https://your.server.com
			  /api/accounts_mgmt: https://her.server.com
			token_url: https://openid.server.com
			user: myuser
			password: mypassword
			client_id: myclient
			client_secret: mysecret
			tokens:
			- {{ .AccessToken }}
			- {{ .RefreshToken }}
			scopes:
			- openid
			- myscope
			insecure: true
			trusted_cas:
			- {{ .Tmp }}/myca.pem
			- {{ .Tmp }}/yourca.pem
			agent: myagent
			retry_limit: 4
			metrics_subsystem: mysubsystem
			`,
			"Tmp", tmp,
			"AccessToken", fileAccess,
			"RefreshToken", fileRefresh,
		)
		file, err := os.CreateTemp("", "*.yaml")
		Expect(err).ToNot(HaveOccurred())
		path := file.Name()
		defer func() {
			err = os.Remove(path)
			Expect(err).ToNot(HaveOccurred())
		}()
		_, err = file.WriteString(content)
		Expect(err).ToNot(HaveOccurred())
		err = file.Close()
		Expect(err).ToNot(HaveOccurred())

		// Create the connection and verify it has been created with the configuration
		// stored in the YAML file:
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Load(path).
			Build()
		Expect(err).ToNot(HaveOccurred())
		Expect(connection.URL()).To(Equal("https://my.server.com"))
		alternativeURLs := connection.AlternativeURLs()
		Expect(alternativeURLs).To(HaveLen(2))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/clusters_mgmt", "https://your.server.com",
		))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/accounts_mgmt", "https://her.server.com",
		))
		Expect(connection.TokenURL()).To(Equal("https://openid.server.com"))
		user, password := connection.User()
		Expect(user).To(Equal("myuser"))
		Expect(password).To(Equal("mypassword"))
		client, secret := connection.Client()
		Expect(client).To(Equal("myclient"))
		Expect(secret).To(Equal("mysecret"))
		returnedAccess, returnedRefresh, err := connection.Tokens()
		Expect(err).ToNot(HaveOccurred())
		Expect(returnedAccess).To(Equal(fileAccess))
		Expect(returnedRefresh).To(Equal(fileRefresh))
		defer func() {
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())
		}()
		Expect(connection.Scopes()).To(ConsistOf("openid", "myscope"))
		Expect(connection.Insecure()).To(BeTrue())
		Expect(connection.Agent()).To(Equal("myagent"))
		Expect(connection.RetryLimit()).To(Equal(4))
		Expect(connection.MetricsSubsystem()).To(Equal("mysubsystem"))
	})

	It("Method calls after load override configuration file", func() {
		// Create temporary files for the trusted CAs:
		tmp, err := os.MkdirTemp("", "*.test.cas")
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = os.RemoveAll(tmp)
			Expect(err).ToNot(HaveOccurred())
		}()
		err = os.WriteFile(filepath.Join(tmp, "myca.pem"), mycaPEM, 0600)
		Expect(err).ToNot(HaveOccurred())
		err = os.WriteFile(filepath.Join(tmp, "yourca.pem"), yourcaPEM, 0600)
		Expect(err).ToNot(HaveOccurred())

		// Create a temporary YAML file containing the configuration:
		fileAccess := MakeTokenString("Bearer", 5*time.Minute)
		fileRefresh := MakeTokenString("Refresh", 10*time.Hour)
		content := EvaluateTemplate(
			`
			url: https://my.server.com
			alternative_urls:
			  /api/clusters_mgmt: https://your.server.com
			  /api/accounts_mgmt: https://her.server.com
			token_url: https://openid.server.com
			user: myuser
			password: mypassword
			client_id: myclient
			client_secret: mysecret
			tokens:
			- {{ .AccessToken }}
			- {{ .RefreshToken }}
			scopes:
			- openid
			- myscope
			insecure: true
			trusted_cas:
			- {{ .Tmp }}/myca.pem
			- {{ .Tmp }}/yourca.pem
			agent: myagent
			retry_limit: 5
			metrics_subsystem: mysubsystem
			`,
			"Tmp", tmp,
			"AccessToken", fileAccess,
			"RefreshToken", fileRefresh,
		)
		file, err := os.CreateTemp("", "*.yaml")
		Expect(err).ToNot(HaveOccurred())
		path := file.Name()
		defer func() {
			err = os.Remove(path)
			Expect(err).ToNot(HaveOccurred())
		}()
		_, err = file.WriteString(content)
		Expect(err).ToNot(HaveOccurred())
		err = file.Close()
		Expect(err).ToNot(HaveOccurred())

		// Load the configuration file and then configure the connection with method
		// calls:
		overridenAccess := MakeTokenString("Bearer", 5*time.Minute)
		overridenRefresh := MakeTokenString("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Load(path).
			URL("https://overriden.my.server.com").
			AlternativeURL("/api/clusters_mgmt", "https://overriden.your.server.com").
			AlternativeURL("/api/accounts_mgmt", "https://overriden.her.server.com").
			TokenURL("https://overriden.openid.server.com").
			User("overriden.myuser", "overriden.mypassword").
			Client("overriden.myclient", "overriden.mysecret").
			Tokens(overridenAccess, overridenRefresh).
			Scopes("openid", "overriden.myscope").
			Insecure(false).
			Agent("overriden.myagent").
			RetryLimit(4).
			MetricsSubsystem("overriden_mysubsystem").
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Check that the actual settings are the ones set with the method calls:
		Expect(connection.URL()).To(Equal("https://overriden.my.server.com"))
		alternativeURLs := connection.AlternativeURLs()
		Expect(alternativeURLs).To(HaveLen(2))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/clusters_mgmt", "https://overriden.your.server.com",
		))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/accounts_mgmt", "https://overriden.her.server.com",
		))
		Expect(connection.TokenURL()).To(Equal("https://overriden.openid.server.com"))
		user, password := connection.User()
		Expect(user).To(Equal("overriden.myuser"))
		Expect(password).To(Equal("overriden.mypassword"))
		client, secret := connection.Client()
		Expect(client).To(Equal("overriden.myclient"))
		Expect(secret).To(Equal("overriden.mysecret"))
		returnedAccess, returnedRefresh, err := connection.Tokens()
		Expect(err).ToNot(HaveOccurred())
		Expect(returnedAccess).To(Equal(overridenAccess))
		Expect(returnedRefresh).To(Equal(overridenRefresh))
		defer func() {
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())
		}()
		Expect(connection.Scopes()).To(ConsistOf("openid", "overriden.myscope"))
		Expect(connection.Insecure()).To(BeFalse())
		Expect(connection.Agent()).To(Equal("overriden.myagent"))
		Expect(connection.RetryLimit()).To(Equal(4))
		Expect(connection.MetricsSubsystem()).To(Equal("overriden_mysubsystem"))
	})

	It("Method calls before load don't override configuration file", func() {
		// Create temporary files for the trusted CAs:
		tmp, err := os.MkdirTemp("", "*.test.cas")
		Expect(err).ToNot(HaveOccurred())
		defer func() {
			err = os.RemoveAll(tmp)
			Expect(err).ToNot(HaveOccurred())
		}()
		err = os.WriteFile(filepath.Join(tmp, "myca.pem"), mycaPEM, 0600)
		Expect(err).ToNot(HaveOccurred())
		err = os.WriteFile(filepath.Join(tmp, "yourca.pem"), yourcaPEM, 0600)
		Expect(err).ToNot(HaveOccurred())

		// Create a temporary YAML file containing the configuration:
		fileAccess := MakeTokenString("Bearer", 5*time.Minute)
		fileRefresh := MakeTokenString("Refresh", 10*time.Hour)
		content := EvaluateTemplate(
			`
			url: https://my.server.com
			alternative_urls:
			  /api/clusters_mgmt: https://your.server.com
			  /api/accounts_mgmt: https://her.server.com
			token_url: https://openid.server.com
			user: myuser
			password: mypassword
			client_id: myclient
			client_secret: mysecret
			tokens:
			- {{ .AccessToken }}
			- {{ .RefreshToken }}
			scopes:
			- openid
			- myscope
			insecure: true
			trusted_cas:
			- {{ .Tmp }}/myca.pem
			- {{ .Tmp }}/yourca.pem
			agent: myagent
			retry_limit: 5
			metrics_subsystem: mysubsystem
			`,
			"Tmp", tmp,
			"AccessToken", fileAccess,
			"RefreshToken", fileRefresh,
		)
		file, err := os.CreateTemp("", "*.yaml")
		Expect(err).ToNot(HaveOccurred())
		path := file.Name()
		defer func() {
			err = os.Remove(path)
			Expect(err).ToNot(HaveOccurred())
		}()
		_, err = file.WriteString(content)
		Expect(err).ToNot(HaveOccurred())
		err = file.Close()
		Expect(err).ToNot(HaveOccurred())

		// Configure the connection with methods call and then load the configuration file:
		overridenAccess := MakeTokenString("Bearer", 5*time.Minute)
		overridenRefresh := MakeTokenString("Refresh", 10*time.Hour)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("https://overriden.my.server.com").
			AlternativeURL("/api/clusters_mgmt", "https://overriden.your.server.com").
			AlternativeURL("/api/accounts_mgmt", "https://overriden.her.server.com").
			TokenURL("https://overriden.openid.server.com").
			User("overriden.myuser", "overriden.mypassword").
			Client("overriden.myclient", "overriden.mysecret").
			Tokens(overridenAccess, overridenRefresh).
			Scopes("openid", "overriden.myscope").
			Insecure(false).
			Agent("overriden.myagent").
			RetryLimit(4).
			MetricsSubsystem("overriden_mysubsystem").
			Load(path).
			Build()
		Expect(err).ToNot(HaveOccurred())

		// Check that the actual settings are the ones from the file:
		Expect(connection.URL()).To(Equal("https://my.server.com"))
		alternativeURLs := connection.AlternativeURLs()
		Expect(alternativeURLs).To(HaveLen(2))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/clusters_mgmt", "https://your.server.com",
		))
		Expect(alternativeURLs).To(HaveKeyWithValue(
			"/api/accounts_mgmt", "https://her.server.com",
		))
		Expect(connection.TokenURL()).To(Equal("https://openid.server.com"))
		user, password := connection.User()
		Expect(user).To(Equal("myuser"))
		Expect(password).To(Equal("mypassword"))
		client, secret := connection.Client()
		Expect(client).To(Equal("myclient"))
		Expect(secret).To(Equal("mysecret"))
		returnedAccess, returnedRefresh, err := connection.Tokens()
		Expect(err).ToNot(HaveOccurred())
		Expect(returnedAccess).To(Equal(fileAccess))
		Expect(returnedRefresh).To(Equal(fileRefresh))
		defer func() {
			err = connection.Close()
			Expect(err).ToNot(HaveOccurred())
		}()
		Expect(connection.Scopes()).To(ConsistOf("openid", "myscope"))
		Expect(connection.Insecure()).To(BeTrue())
		Expect(connection.Agent()).To(Equal("myagent"))
		Expect(connection.RetryLimit()).To(Equal(5))
		Expect(connection.MetricsSubsystem()).To(Equal("mysubsystem"))
	})

	It("Returns configured URL when there are no alternative URLs", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("https://my.server.com").
			Tokens(token).
			Build()
		Expect(err).ToNot(HaveOccurred())
		Expect(connection.URL()).To(Equal("https://my.server.com"))
	})

	It("Returns configured URL when there are alternative URLs", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("https://my.server.com").
			AlternativeURL("/api/clusters_mgmt", "https://your.server.com").
			AlternativeURL("/api/accounts_mgmt", "https://her.server.com").
			Tokens(token).
			Build()
		Expect(err).ToNot(HaveOccurred())
		Expect(connection.URL()).To(Equal("https://my.server.com"))
	})

	It("Can't be created with URL without scheme", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("my.server.com").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("my.server.com"))
		Expect(message).To(ContainSubstring("scheme"))
		Expect(message).To(ContainSubstring("http"))
		Expect(message).To(ContainSubstring("https"))
	})

	It("Can't be created with URL with wrong scheme", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("junk://my.server.com").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("junk://my.server.com"))
		Expect(message).To(ContainSubstring("scheme"))
		Expect(message).To(ContainSubstring("http"))
		Expect(message).To(ContainSubstring("https"))
		Expect(message).To(ContainSubstring("junk"))
	})

	It("Can't be created with alternative URL wihout scheme", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			AlternativeURL("/api/clusters_mtmt", "my.server.com").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("my.server.com"))
		Expect(message).To(ContainSubstring("scheme"))
		Expect(message).To(ContainSubstring("http"))
		Expect(message).To(ContainSubstring("https"))
	})

	It("Can't be created with alternative URL with wrong scheme", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			AlternativeURL("/api/clusters_mtmt", "junk://my.server.com").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("junk://my.server.com"))
		Expect(message).To(ContainSubstring("scheme"))
		Expect(message).To(ContainSubstring("http"))
		Expect(message).To(ContainSubstring("https"))
		Expect(message).To(ContainSubstring("junk"))
	})

	It("Can't be created with URL without host name", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		_, err := NewConnectionBuilder().
			Logger(logger).
			URL("http:///mypath").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(ContainSubstring("http:///mypath"))
		Expect(message).To(ContainSubstring("host name"))
	})

	It("Can't be created with alternative URL without host name", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			AlternativeURL("/api/clusters_mgmt", "http:///mypath").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("http:///mypath"))
		Expect(message).To(ContainSubstring("host name"))
	})

	It("Can't be created with token URL without scheme", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("https://my.server.com").
			TokenURL("your.server.com").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("your.server.com"))
		Expect(message).To(ContainSubstring("scheme"))
		Expect(message).To(ContainSubstring("http"))
		Expect(message).To(ContainSubstring("https"))
	})

	It("Can't be created with token URL with wrong scheme", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("https://my.server.com").
			TokenURL("junk://your.server.com").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("junk://your.server.com"))
		Expect(message).To(ContainSubstring("scheme"))
		Expect(message).To(ContainSubstring("http"))
		Expect(message).To(ContainSubstring("https"))
		Expect(message).To(ContainSubstring("junk"))
	})

	It("Can't be created with token URL without host name", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("http://my.server.com").
			TokenURL("http:///yourpath").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("http:///yourpath"))
		Expect(message).To(ContainSubstring("host name"))
	})

	It("Can be created with Unix network and host", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("unix://my.server.com/tmp/api.socket").
			Tokens(token).
			Build()
		Expect(err).ToNot(HaveOccurred())
		Expect(connection).ToNot(BeNil())
	})

	It("Can be created with Unix network and HTTPS", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("unix+https://my.server.com/tmp/api.socket").
			Tokens(token).
			Build()
		Expect(err).ToNot(HaveOccurred())
		Expect(connection).ToNot(BeNil())
	})

	It("Can't be created with Unix network and no host", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("unix:/tmp/api.socket").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("unix:/tmp/api.socket"))
		Expect(message).To(ContainSubstring("host"))
		Expect(message).To(ContainSubstring("mandatory"))
	})

	It("Can't be created with incorrect network", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("junk+https://my.server.com").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("junk"))
		Expect(message).To(ContainSubstring("network"))
		Expect(message).To(ContainSubstring("tcp"))
		Expect(message).To(ContainSubstring("unix"))
	})

	It("Can't be created with Unix network and no socket", func() {
		token := MakeTokenString("Bearer", 5*time.Minute)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			URL("unix://my.server.com").
			Tokens(token).
			Build()
		Expect(err).To(HaveOccurred())
		Expect(connection).To(BeNil())
		message := err.Error()
		Expect(message).To(ContainSubstring("unix"))
		Expect(message).To(ContainSubstring("socket"))
		Expect(message).To(ContainSubstring("path"))
	})

	It("Function Close returns nil when trying to close a closed connection", func() {
		offlineToken := MakeTokenString("Offline", 0)
		connection, err := NewConnectionBuilder().
			Logger(logger).
			Tokens(offlineToken).
			Build()
		Expect(err).ToNot(HaveOccurred())
		err = connection.Close()
		Expect(err).ToNot(HaveOccurred())
		// Try to close the connection again
		err = connection.Close()
		Expect(err).To(BeNil())
	})
})

type TestTransport struct {
	called int
}

func (t *TestTransport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	t.called++
	header := http.Header{}
	header.Add("Content-type", "application/json")
	response = &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     header,
		Body:       BufferWithBytes([]byte("{}")),
	}
	return response, nil
}

func NewTestTransport() *TestTransport {
	return &TestTransport{called: 0}
}

// This certificate is used only to verify that the connection can load a certificate from
// a file. It isn't valid for any other thing. It has been generated with the following
// commmad:
//
//	openssl req \
//	-x509 \
//	-newkey rsa:4096 \
//	-keyout myca.key \
//	-nodes \
//	-out myca.crt \
//	-subj '/CN=myca.com' \
//	-days 3650
var mycaPEM = []byte(`
-----BEGIN CERTIFICATE-----
MIIFCzCCAvOgAwIBAgIUEy+dp9sWPdu59sahER7AxvapYOgwDQYJKoZIhvcNAQEL
BQAwFTETMBEGA1UEAwwKeW91cmNhLmNvbTAeFw0yMDExMTgxNjUwMTNaFw0zMDEx
MTYxNjUwMTNaMBUxEzARBgNVBAMMCnlvdXJjYS5jb20wggIiMA0GCSqGSIb3DQEB
AQUAA4ICDwAwggIKAoICAQC7IXXeem/arTEAvujthKpzxMOaimpIq276rIaaehSf
PKfwwFScz6KzkcRcCjzGlmQIUfe0VunL9xMfWcOBQ8u0LofJpcRE+AuYXdgAuuyH
ijWukGZ4o1QGoSmS90TVOOLGA38gnPbQTAgJN8DzxccoOTVtdqsAZMK5zKGJ0IUa
0ZkPJvs0QYfVYMgYCjiRCeTNze4Cwb/ecj9CZump2IUNm3oE28dUkzRCBysNAvKu
Ms9HktnQft7BxoefdCZK04o8a1BLzZVSPe7qV+bk0+hD1xygoEPATzhuGl7Al9EI
lkOJ8fv3uompnz5bHOeP2dNuwn9efeINtJcLlx8wySkU0oNTqQq9MVI7gRwUUAT0
dLzETULngRhvjGSYEyST3vT27V444fVYkSVIjmji+SmzSejfZq/A1NTQ8M9TUWIs
7dL562GJnsalPnI+m9XR5m3oajY+CYtcd5q1iIus+WrMXups8fQnpssJHioFs86s
NEQ0Evbl0OGxxYivJwKbT6Oo86Uh3nUXXx/xxBI5HRmQap38EK6K0WZR3V7MBnpF
xVv5vUO/zXc7DxAghcXb41XSTWGOM+AqaCIv+zys87/F6x1dmCkA8DCNxYlEeK88
6xHuK5265iu/to9NSvzCAxmrz4fDea3eAxpZus39yN3N2ud2IAlMguicjZgLF3mg
uQIDAQABo1MwUTAdBgNVHQ4EFgQUf37ek+qeiMRhZ0q6whksuPY4ezQwHwYDVR0j
BBgwFoAUf37ek+qeiMRhZ0q6whksuPY4ezQwDwYDVR0TAQH/BAUwAwEB/zANBgkq
hkiG9w0BAQsFAAOCAgEArP/sjXMURWamJFckKwpam+w8WPW3b0wq6GkQky6XDcXK
sym5vJfQtQgzZV/rxb0RcO4ywPKYJK2ViREqksmlD5XLL/6grbe+rcIY55IVcFKe
3ZfE7toa08gWV8kb9VP2KVNt5505jJUVtF8FxRxsu5W0x6b6Kegyotsd5/7tads9
9qMOoiWOyWxdP4FZalNM8PXaF8pspNqeo/cvgFWvDTqtFvgH4vkLehMueAWmDML2
lqYHCaMworpY3e8vfk6jK8b9fRmuXaGMlOTpY7XoF8OOSMI1LdPVSO/lo8DCJbge
RoBHZ97fA8ShB+WRjBuAuh5ST/TEqTha+razhmauVT6CYtw9SSC0SK0ZbNg2oZoG
CzrQhYf9gHXXPnp0qsuPbtHrMm3DPBHBUrfrlvVVmWjKCVipFdPyRXth9FtNguZn
d5NUX3JwGRoez/xHv8nHyjdKmTCu8pmP/9SAoV/HUgpcSaEtXyZdlNd+SgIcgV/A
00eSLesNr9/auzWklxh92oqDnd96IRueamFm6W0BCh64if//BCmAFelPHlWSP1ws
nzNA++GMw2OXJ1cE/9GTU3or0eDGRgB9XIu+T/SLPXW9xBm0hZdt5gyYfEDmppHa
nctMPznTWc+iYCMAwroHzJV40ZrVhllhNYrrOLigA7NfAiXelLmSbLx316TnoZM=
-----END CERTIFICATE-----
`)

// This certificate is used only to verify that the connection can load a certificate from
// a file. It isn't valid for any other thing. It has been generated with the following
// commmad:
//
//	openssl req \
//	-x509 \
//	-newkey rsa:4096 \
//	-keyout yourca.key \
//	-nodes \
//	-out yourca.crt \
//	-subj '/CN=yourca.com' \
//	-days 3650
var yourcaPEM = []byte(`
-----BEGIN CERTIFICATE-----
MIIFCzCCAvOgAwIBAgIUOBKDkme46UAOif9G7fIfN0FTmAswDQYJKoZIhvcNAQEL
BQAwFTETMBEGA1UEAwwKeW91cmNhLmNvbTAeFw0yMDExMTgxNjU0MTBaFw0zMDEx
MTYxNjU0MTBaMBUxEzARBgNVBAMMCnlvdXJjYS5jb20wggIiMA0GCSqGSIb3DQEB
AQUAA4ICDwAwggIKAoICAQCjCUIsLIE0gp7JQZHvwrl12IjYEQjJWEpsEbp/jpux
ztUAVju5Cq8+V5DYIzHi6WHlottpp6obh4TaCHZroZxwXKCoUARJwtPqADhDX+tr
Jy8gA2y5ixGxryyXVsAT3YkEBW05i82aKa+FB05T0eUS52SBS2V6Fd4XU19denx5
rb1eFNor/rmG0gMCAsh/4oWw8DdBEU0qc/9vQ3lsWvGU1noYt/kwfAcaSydrqaIA
EvK/sG1/hxWD+JBOUwrah0zxT+7x6FbzqX83m03HM7ZHJR+qNjtg31loQwocsn20
qOM3vMQkqyqjnXMHtIleyhw7fWNqQcvS/f+A3QUosf3h90c/BVmoJ+QUsm9gOUFT
jGWybFWRnukSY/CMazVQvSF3N6GDz0iQQFQEwtpLhe0UhpFs/UUOXUi2+JjLRkLf
fIeyk+K9EAz8Nd+vOMgr7ud70MykF7X5FzGLwJfz5bj62XVVifA8yMNoxIdlRFZp
H8OXzMwUO0+ktCf1StXCEV8/HoBP8BeKRl/PPqyHlY2rqXrNYq+ZXR+I25HWJcyJ
UUWQBq66yfxEIR2L5tJoVf0P7Z+WcplX7bo8T06n92zV5pU6WPi2+xAob2cmR9M4
thlHl5uUDIYRgXZ2RQnMhea2GcDtFi8zfeWIJTZ54CmUi326sToqvCbOrHMkxus9
xwIDAQABo1MwUTAdBgNVHQ4EFgQU5anvmjhvY6yPqy9qKU6cZAjX88MwHwYDVR0j
BBgwFoAU5anvmjhvY6yPqy9qKU6cZAjX88MwDwYDVR0TAQH/BAUwAwEB/zANBgkq
hkiG9w0BAQsFAAOCAgEAH2x/DTPnfxw06RZmjGCWOJJOUiO9uW4dVtLUzdCmOkuX
zngsqAB6Mqy+GXb4jsNdKdVR74Lb/9gv9WkXTxLPnW8sBmxC7NxjJZbQVlHuQK3U
s8GGo1wwuR/kCcnckdRAuDf8BbllqpPq+zUm1ZzM2NnMtiptkojxpleJugttXGiw
SHRh6hY3RZAKD6s6eyg67O6Dx3bFgyzYt91cG+YJPbSQh9hBhhlmp5GwOMg27u6N
skSNZIK2SNs5Bael+WfiUiEB1cFwUc0TYPUSJEkLXvLcqqVzuj53IEU7UCifYqHQ
xlhdROagJc8fSOQ0yEEwBPqDVRT3fJipAGRB7h8a0pEtfbD8M0df6DGkcRvOf/My
B2Ss8ZrL+tLDKEJji6aZXlkFbs6aKko0cKbZQvquISgEdcZp3hQ4oc+eUmkOfkk6
0D+7ZY3m4JDAr38tVDw3lG2I3THvmT8zdPZzujkZjHUvVqWaoEX1k6IUSZ5DnQ3H
NIf/SfR7aXBeCsoJnCE+nsN/ba2twS8Wx1evuWDdlVGYrE4ujXpCAKspwi0mPirx
bkAQVDU7e/Zr6P8ZI9P4w1MYZvKagPo1+hCCiEAaeNdgMoOwQbLpw6py0x7B+mxI
ppX5DVNV8wJAb9KqeSXwd89Z5unpeS6KZsMcb5qiK60Lj12aLmZ8ip6s7xjAS8Q=
-----END CERTIFICATE-----
`)
