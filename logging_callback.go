/*
Copyright (c) 2020 Red Hat, Inc.

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

import "net/http"

type LoggingCallback struct {
	logger Logger
}

func NewLoggingCallback(logger Logger) *LoggingCallback {
	return &LoggingCallback{logger: logger}
}

func (c *LoggingCallback) Pre(request *http.Request) *http.Request {
	c.logger.Info(request.Context(), "Sending request %s '%s'", request.Method, request.URL.String())
	return request
}

func (c *LoggingCallback) Post(response *http.Response, err error) (*http.Response, error){
	if err != nil {
		c.logger.Error(response.Request.Context(), "Got error sending request")
	} else {
		c.logger.Info(response.Request.Context(), "Got response status code %d", response.StatusCode)
	}
	return response, err
}
