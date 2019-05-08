/*
Copyright (c) 2018 Red Hat, Inc.

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

// This file contains the implementations of the methods of the connection that handle OpenID
// authentication tokens.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/internal"
)

// Tokens returns the access and refresh tokens that is currently in use by the connection. If it is
// necessary to request a new token because it wasn't requested yet, or because it is expired, this
// method will do it and will return an error if it fails.
//
// This operation is potentially lengthy, as it may require network communication. Consider using a
// context and the TokensContext method.
func (c *Connection) Tokens() (access, refresh string, err error) {
	return c.TokensContext(context.Background())
}

// TokensContext returns the access and refresh tokens that is currently in use by the connection.
// If it is necessary to request a new token because it wasn't requested yet, or because it is
// expired, this method will do it and will return an error if it fails.
func (c *Connection) TokensContext(ctx context.Context) (access, refresh string, err error) {
	// We need to make sure that this method isn't execute concurrently, as we will be updating
	// multiple attributes of the connection:
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	// If the access token is expired, then check the refresh token and either refresh it or
	// request a new one:
	now := time.Now()
	var accessExpires bool
	var accessLeft time.Duration
	if c.accessToken != nil {
		accessExpires, accessLeft, err = c.tokenExpiry(ctx, c.accessToken, now)
		if err != nil {
			return
		}
	}
	var refreshExpires bool
	var refreshLeft time.Duration
	if c.refreshToken != nil {
		refreshExpires, refreshLeft, err = c.tokenExpiry(ctx, c.refreshToken, now)
		if err != nil {
			return
		}
	}
	if c.logger.DebugEnabled() {
		c.debugExpiry(ctx, "Bearer", c.accessToken, accessExpires, accessLeft)
		c.debugExpiry(ctx, "Refresh", c.refreshToken, refreshExpires, refreshLeft)
	}
	if c.accessToken == nil || (accessExpires && accessLeft < 5*time.Second) {
		if c.refreshToken == nil || (refreshExpires && refreshLeft < 10*time.Second) {
			c.logger.Debug(ctx, "Requesting new token")
			err = c.sendRequestTokenForm(ctx)
			if err != nil {
				return
			}
		} else {
			c.logger.Debug(ctx, "Refreshing token")
			err = c.sendRefreshTokenForm(ctx)
			if err != nil {
				return
			}
		}
	}
	if c.accessToken != nil {
		access = c.accessToken.Raw
	}
	if c.refreshToken != nil {
		refresh = c.refreshToken.Raw
	}

	return
}

func (c *Connection) sendRequestTokenForm(ctx context.Context) error {
	form := url.Values{}
	if c.user != "" && c.password != "" {
		form.Set("grant_type", "password")
		form.Set("client_id", c.clientID)
		form.Set("username", c.user)
		form.Set("password", c.password)
	} else if c.clientID != "" && c.clientSecret != "" {
		form.Set("grant_type", "client_credentials")
		form.Set("client_id", c.clientID)
		form.Set("client_secret", c.clientSecret)
	} else {
		return fmt.Errorf(
			"either user name and password or client identifier and secret must " +
				"be provided",
		)
	}
	form.Set("scope", strings.Join(c.scopes, " "))
	return c.sendTokenForm(ctx, form)
}

func (c *Connection) sendRefreshTokenForm(ctx context.Context) error {
	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("client_id", c.clientID)
	form.Set("client_secret", c.clientSecret)
	form.Set("refresh_token", c.refreshToken.Raw)
	return c.sendTokenForm(ctx, form)
}

func (c *Connection) sendTokenForm(ctx context.Context, form url.Values) error {
	// Create the HTTP request:
	body := []byte(form.Encode())
	request, err := http.NewRequest(http.MethodPost, c.tokenURL.String(), bytes.NewReader(body))
	request.Close = true
	header := request.Header
	if c.agent != "" {
		header.Set("User-Agent", c.agent)
	}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	header.Set("Accept", "application/json")
	if err != nil {
		return fmt.Errorf("can't create request: %v", err)
	}

	// Set the context:
	if ctx != nil {
		request = request.WithContext(ctx)
	}

	// Send the HTTP request:
	if c.logger.DebugEnabled() {
		var censoredBody bytes.Buffer
		// Unlike real url.Values.Encode(), this doesn't sort keys.
		for name, values := range form {
			for _, value := range values {
				// Buffer.Write*() don't require error checking but golangci-lint v1.10.2
				// on Jenkins flags them (maybe https://github.com/securego/gosec/issues/267).
				if censoredBody.Len() > 0 {
					censoredBody.WriteByte('&') // #nosec G104
				}
				censoredBody.WriteString(url.QueryEscape(name) + "=") // #nosec G104

				if name == "password" || name == "client_secret" {
					censoredBody.WriteString("***") // #nosec G104

				} else {
					censoredBody.WriteString(url.QueryEscape(value)) // #nosec G104
				}
			}
		}
		c.dumpRequest(ctx, request, censoredBody.Bytes())
	}
	response, err := c.client.Do(request)
	if err != nil {
		return fmt.Errorf("can't send request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body:
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("can't read response: %v", err)
	}
	if c.logger.DebugEnabled() {
		c.dumpResponse(ctx, response, body)
	}

	// Check the response status and content type:
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("token response status is: %s", response.Status)
	}
	header = response.Header
	content := header.Get("Content-Type")
	if content != "application/json" {
		return fmt.Errorf("expected 'application/json' but got '%s'", content)
	}

	// Parse the response body:
	var msg internal.TokenResponse
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return fmt.Errorf("can't parse JSON response: %v", err)
	}
	if msg.Error != nil {
		if msg.ErrorDescription != nil {
			return fmt.Errorf("%s: %s", *msg.Error, *msg.ErrorDescription)
		}
		return fmt.Errorf("%s", *msg.Error)
	}
	if msg.TokenType != nil && *msg.TokenType != "bearer" {
		return fmt.Errorf("expected 'bearer' token type but got '%s", *msg.TokenType)
	}
	if msg.AccessToken == nil {
		return fmt.Errorf("no access token was received")
	}
	accessToken, _, err := c.tokenParser.ParseUnverified(*msg.AccessToken, jwt.MapClaims{})
	if err != nil {
		return err
	}
	if msg.RefreshToken == nil {
		return fmt.Errorf("no refresh token was received")
	}
	refreshToken, _, err := c.tokenParser.ParseUnverified(*msg.RefreshToken, jwt.MapClaims{})
	if err != nil {
		return err
	}

	// Save the new tokens:
	c.accessToken = accessToken
	c.refreshToken = refreshToken

	return nil
}

// tokenExpiry determines if the given token expires, and the time that remains till it expires.
func (c *Connection) tokenExpiry(ctx context.Context, token *jwt.Token,
	now time.Time) (expires bool, left time.Duration, err error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("expected map claims bug got %T", claims)
		return
	}
	claim, ok := claims["exp"]
	if !ok {
		err = fmt.Errorf("token doesn't contain the 'exp' claim")
		return
	}
	exp, ok := claim.(float64)
	if !ok {
		err = fmt.Errorf("expected floating point 'exp' but got %T", claim)
		return
	}
	if exp == 0 {
		expires = false
		left = 0
	} else {
		expires = true
		left = time.Unix(int64(exp), 0).Sub(now)
	}
	return
}

// debugExpiry sends to the log information about the expiration of the given token.
func (c *Connection) debugExpiry(ctx context.Context, typ string, token *jwt.Token, expires bool,
	left time.Duration) {
	if token != nil {
		if expires {
			if left < 0 {
				c.logger.Debug(ctx, "%s token expired %s ago", typ, -left)
			} else if left > 0 {
				c.logger.Debug(ctx, "%s token expires in %s", typ, left)
			} else {
				c.logger.Debug(ctx, "%s token expired just now", typ)
			}
		}
	} else {
		c.logger.Debug(ctx, "%s token isn't available", typ)
	}
}
