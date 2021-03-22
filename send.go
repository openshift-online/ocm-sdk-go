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

package sdk

import (
	"context"
	"fmt"
	"html"
	"io/ioutil"
	"mime"
	"net/http"
	"path"
	"regexp"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/openshift-online/ocm-sdk-go/internal"
)

var wsRegex = regexp.MustCompile(`\s+`)

// RoundTrip is the implementation of the http.RoundTripper interface.
func (c *Connection) RoundTrip(request *http.Request) (response *http.Response, err error) {
	// Check if the connection is closed:
	err = c.checkClosed()
	if err != nil {
		return
	}

	// Get the context from the request:
	ctx := request.Context()

	// Check the request URL:
	if request.URL.Path == "" {
		err = fmt.Errorf("request path is mandatory")
		return
	}
	if request.URL.Scheme != "" || request.URL.Host != "" || !path.IsAbs(request.URL.Path) {
		err = fmt.Errorf("request URL '%s' isn't absolute", request.URL)
		return
	}

	// Add the base URL to the request URL:
	base, err := c.selectBaseURL(ctx, request)
	if err != nil {
		return
	}
	request.URL = base.URL.ResolveReference(request.URL)

	// Check the request method and body:
	switch request.Method {
	case http.MethodGet, http.MethodDelete:
		if request.Body != nil {
			c.logger.Warn(ctx,
				"Request body is not allowed for the '%s' method",
				request.Method,
			)
		}
	case http.MethodPost, http.MethodPatch, http.MethodPut:
		// POST and PATCH and PUT don't need to have a body. It is up to the server to decide if
		// this is acceptable.
	default:
		err = fmt.Errorf("method '%s' is not allowed", request.Method)
		return
	}

	// Get the access token:
	token, _, err := c.TokensContext(ctx)
	if err != nil {
		err = fmt.Errorf("can't get access token: %w", err)
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
	case http.MethodPost, http.MethodPatch, http.MethodPut:
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("Accept", "application/json")

	// Select the client:
	client, err := c.clientSelector.Select(ctx, base)
	if err != nil {
		return
	}

	// Send the request and get the response:
	response, err = client.Do(request)
	if err != nil {
		err = fmt.Errorf("can't send request: %w", err)
		return
	}

	// Check that the response content type is JSON:
	err = c.checkContentType(response)
	if err != nil {
		return
	}

	return
}

// checkContentType checks that the content type of the given response is JSON. Note that if the
// content type isn't JSON this method will consume the complete body in order to generate an error
// message containing a summary of the content.
func (c *Connection) checkContentType(response *http.Response) error {
	var err error
	var mediaType string
	contentType := response.Header.Get("Content-Type")
	if contentType != "" {
		mediaType, _, err = mime.ParseMediaType(contentType)
		if err != nil {
			return err
		}
	} else {
		mediaType = contentType
	}
	if !strings.EqualFold(mediaType, "application/json") {
		var summary string
		summary, err = c.contentSummary(mediaType, response)
		if err != nil {
			return fmt.Errorf(
				"expected response content type 'application/json' but received "+
					"'%s' and couldn't obtain content summary: %w",
				mediaType, err,
			)
		}
		return fmt.Errorf(
			"expected response content type 'application/json' but received '%s' and "+
				"content '%s'",
			mediaType, summary,
		)
	}
	return nil
}

// contentSummary reads the body of the given response and returns a summary it. The summary will
// be the complete body if it isn't too log. If it is too long then the summary will be the
// beginning of the content followed by ellipsis.
func (c *Connection) contentSummary(mediaType string, response *http.Response) (summary string, err error) {
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	limit := 250
	runes := []rune(string(body))
	if strings.EqualFold(mediaType, "text/html") && len(runes) > limit {
		content := html.UnescapeString(strip.StripTags(string(body)))
		content = wsRegex.ReplaceAllString(strings.TrimSpace(content), " ")
		runes = []rune(content)
	}
	if len(runes) > limit {
		summary = fmt.Sprintf("%s...", string(runes[:limit]))
	} else {
		summary = string(runes)
	}
	return
}

// selectBaseURL selects the base URL that should be used for the given request, according its path
// and the alternative URLs configured when the connection was created.
func (c *Connection) selectBaseURL(ctx context.Context,
	request *http.Request) (base *internal.ServerAddress, err error) {
	// Select the base URL that has the longest matching prefix. Note that it is enough to pick
	// the first match because the entries have already been sorted by descending prefix length
	// when the connection was created.
	for _, entry := range c.urlTable {
		if entry.re.MatchString(request.URL.Path) {
			base = entry.url
			return
		}
	}
	if base == nil {
		err = fmt.Errorf(
			"can't find any matching URL for request path '%s'",
			request.URL.Path,
		)
	}
	return
}
