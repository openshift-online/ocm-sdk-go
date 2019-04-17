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

// This file contains the implementation of the methods of the connection that are used to send HTTP
// requests and receive HTTP responses.

package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"
)

func (c *Connection) RoundTrip(request *http.Request) (response *http.Response, err error) {
	// Check if the connection is closed:
	err = c.checkClosed()
	if err != nil {
		return
	}

	// Measure the time that it takes to send the request and receive the resposne, and report
	// it in the log:
	var before time.Time
	var after time.Time
	var elapsed time.Duration
	if c.logger.DebugEnabled() {
		before = time.Now()
	}
	response, err = c.send(request)
	if c.logger.DebugEnabled() {
		after = time.Now()
		elapsed = after.Sub(before)
		c.logger.Debug("Response received in %s", elapsed)
	}
	return
}

func (c *Connection) send(request *http.Request) (response *http.Response, err error) {
	// Check that the request URL:
	if request.URL.Path == "" {
		err = fmt.Errorf("request path is mandatory")
		return
	}
	if request.URL.Scheme != "" || request.URL.Host != "" || !path.IsAbs(request.URL.Path) {
		err = fmt.Errorf("request URL '%s' isn't absolute", request.URL)
		return
	}

	// Add the API URL to the request URL:
	request.URL = c.apiURL.ResolveReference(request.URL)

	// Check the request method and body:
	switch request.Method {
	case http.MethodGet, http.MethodDelete:
		if request.Body != nil {
			err = fmt.Errorf(
				"request body is not allowed for the '%s' method",
				request.Method,
			)
			return
		}
	case http.MethodPost, http.MethodPatch:
		if request.Body == nil {
			err = fmt.Errorf(
				"request body is mandatory for the '%s' method",
				request.Method,
			)
			return
		}
	default:
		err = fmt.Errorf("method '%s' is not allowed", request.Method)
		return
	}

	// Get the access token:
	token, _, err := c.Tokens(request.Context())
	if err != nil {
		err = fmt.Errorf("can't get access token: %v", err)
		return
	}

	// Add the default headers:
	if request.Header == nil {
		request.Header = make(http.Header)
	}
	if c.agent != "" {
		request.Header.Set("User-Agent", c.agent)
	}
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}
	switch request.Method {
	case http.MethodPost, http.MethodPatch:
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("Accept", "application/json")

	// If debug is enabled then we need to read the complete body in memory, in order to send it
	// to the log, and we need to replace the original with a reader that reads it from memory:
	if c.logger.DebugEnabled() {
		if request.Body != nil {
			var body []byte
			body, err = ioutil.ReadAll(request.Body)
			if err != nil {
				err = fmt.Errorf("can't read request body: %v", err)
				return
			}
			err = request.Body.Close()
			if err != nil {
				err = fmt.Errorf("can't close request body: %v", err)
				return
			}
			c.dumpRequest(request, body)
			request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		} else {
			c.dumpRequest(request, nil)
		}
	}

	// Send the request and get the response:
	response, err = c.client.Do(request)
	if err != nil {
		err = fmt.Errorf("can't send request: %v", err)
		return
	}

	// If debug is enabled then we need to read the complete response body in memory, in order
	// to send it the log, and we need to replace the original with a reader that reads it from
	// memory:
	if c.logger.DebugEnabled() {
		if response.Body != nil {
			var body []byte
			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				err = fmt.Errorf("can't read response body: %v", err)
				return
			}
			err = response.Body.Close()
			if err != nil {
				err = fmt.Errorf("can't close response body: %v", err)
				return
			}
			c.dumpResponse(response, body)
			response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		} else {
			c.dumpResponse(response, nil)
		}
	}

	return
}
