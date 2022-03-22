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

// This file contains the implementations of the methods of the connection that are used to dump to
// the log the details of HTTP requests and responses.

package sdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/go-logr/logr"
	jsoniter "github.com/json-iterator/go"

	"github.com/openshift-online/ocm-sdk-go/v2/helpers"
)

// dumpTransportWrapper is a transport wrapper that creates round trippers that dump the details of
// the request and the responses to the log.
type dumpTransportWrapper struct {
	logger logr.Logger
}

// Wrap creates a round tripper on top of the given one that sends to the log the details of
// requests and responses.
func (w *dumpTransportWrapper) Wrap(transport http.RoundTripper) http.RoundTripper {
	return &dumpRoundTripper{
		logger: w.logger,
		next:   transport,
	}
}

// dumpRoundTripper is a round tripper that dumps the details of the requests and the responses to
// the log.
type dumpRoundTripper struct {
	logger logr.Logger
	next   http.RoundTripper
}

// Make sure that we implement the http.RoundTripper interface:
var _ http.RoundTripper = &dumpRoundTripper{}

// RoundTrip is he implementation of the http.RoundTripper interface.
func (d *dumpRoundTripper) RoundTrip(request *http.Request) (response *http.Response, err error) {
	// Get the context:
	ctx := request.Context()

	// Read the complete body in memory, in order to send it to the log, and replace it with a
	// reader that reads it from memory:
	if request.Body != nil {
		var body []byte
		body, err = ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}
		err = request.Body.Close()
		if err != nil {
			return
		}
		d.dumpRequest(ctx, request, body)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	} else {
		d.dumpRequest(ctx, request, nil)
	}

	// Call the next round tripper:
	response, err = d.next.RoundTrip(request)
	if err != nil {
		return
	}

	// Read the complete response body in memory, in order to send it the log, and replace it
	// with a reader that reads it from memory:
	if response.Body != nil {
		var body []byte
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
		err = response.Body.Close()
		if err != nil {
			return
		}
		d.dumpResponse(ctx, response, body)
		response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	} else {
		d.dumpResponse(ctx, response, nil)
	}

	return
}

const (
	// redactionStr replaces sensitive values in output.
	redactionStr = "***"
)

// redactFields are removed from log output when dumped.
var redactFields = map[string]bool{
	"access_token":  true,
	"admin":         true,
	"id_token":      true,
	"refresh_token": true,
	"password":      true,
	"client_secret": true,
	"kubeconfig":    true,
	"ssh":           true,
}

// dumpRequest dumps to the log, in debug level, the details of the given HTTP request.
func (d *dumpRoundTripper) dumpRequest(ctx context.Context, request *http.Request, body []byte) {
	var pairs []interface{}
	pairs = append(
		pairs,
		"method", request.Method,
		"url", request.URL,
	)
	header := http.Header{}
	if request.Host != "" {
		header.Add("Host", request.Host)
	}
	for name, values := range request.Header {
		if strings.ToLower(name) == "authorization" {
			header.Add(name, "***")
		} else {
			for _, value := range values {
				header.Add(name, value)
			}
		}
	}
	pairs = append(pairs, "header", header)
	if body != nil {
		pairs = append(pairs, "body", d.dumpBody(header, body))
	}
	d.logger.Info("Response", pairs...)
}

// dumpResponse dumps to the log, in debug level, the details of the given HTTP response.
func (d *dumpRoundTripper) dumpResponse(ctx context.Context, response *http.Response, body []byte) {
	var pairs []interface{}
	pairs = append(
		pairs,
		"protocol", response.Proto,
		"status", response.Status,
		"header", response.Header,
	)
	if body != nil {
		pairs = append(pairs, "body", d.dumpBody(response.Header, body))
	}
	d.logger.Info("Response", pairs...)
}

// dumpBody checks the content type used in the given header and then it converts the given body
// into an object that can be added as a field of a log message.
func (d *dumpRoundTripper) dumpBody(header http.Header, body []byte) interface{} {
	// Try to parse the content type:
	var mediaType string
	contentType := header.Get("Content-Type")
	if contentType != "" {
		var err error
		mediaType, _, err = mime.ParseMediaType(contentType)
		if err != nil {
			d.logger.Error(
				err,
				"Can't parse content type",
				"content_type", contentType,
			)
		}
	} else {
		mediaType = contentType
	}

	// Dump the body according to the content type:
	switch mediaType {
	case "application/x-www-form-urlencoded":
		return d.dumpForm(body)
	case "application/json", "":
		return d.dumpJSON(body)
	default:
		return body
	}
}

// dumpForm prepares the given form data for use as a field in the log, excluding security sensitive
// parts.
func (d *dumpRoundTripper) dumpForm(data []byte) interface{} {
	// Parse the form:
	form, err := url.ParseQuery(string(data))
	if err != nil {
		return data
	}

	// Redact values corresponding to security sensitive fields:
	for name, values := range form {
		if redactFields[name] {
			for i := range values {
				values[i] = redactionStr
			}
		}
	}

	// Remove values of sensitive fields:
	redactedForm := url.Values{}
	for name, values := range form {
		redactedValues := make([]string, len(values))
		for i, value := range values {
			if redactFields[name] {
				redactedValues[i] = "***"
			} else {
				redactedValues[i] = value
			}
		}
		redactedForm[name] = redactedValues
	}

	// Get and sort the names of the fields of the form, so that the generated output will be
	// predictable:
	names := make([]string, len(form))
	i := 0
	for name := range form {
		names[i] = name
		i++
	}
	sort.Strings(names)

	// Send the redacted data to the log:
	result := map[string]interface{}{}
	for _, name := range names {
		var value interface{}
		values := redactedForm[name]
		switch len(values) {
		case 1:
			value = values[0]
		default:
			value = values
		}
		result[name] = value

	}
	return result
}

// dumpJSON tries to parse the given data as a JSON document. If that works, then it returns the
// result, otherwise it returns it as is.
func (d *dumpRoundTripper) dumpJSON(data []byte) interface{} {
	iterator, err := helpers.NewIterator(data)
	if err != nil {
		return data
	}
	var buffer bytes.Buffer
	stream := helpers.NewStream(&buffer)
	d.redactSensitive(iterator, stream)
	err = stream.Flush()
	if err != nil {
		return data
	}
	data = buffer.Bytes()
	var result interface{}
	err = jsoniter.Unmarshal(data, &result)
	if err != nil {
		return data
	}
	return result
}

// redactSensitive replaces sensitive fields within a response with redactionStr.
func (d *dumpRoundTripper) redactSensitive(it *jsoniter.Iterator, str *jsoniter.Stream) {
	switch it.WhatIsNext() {
	case jsoniter.ObjectValue:
		str.WriteObjectStart()
		first := true
		for field := it.ReadObject(); field != ""; field = it.ReadObject() {
			if !first {
				str.WriteMore()
			}
			first = false
			str.WriteObjectField(field)
			if v, ok := redactFields[field]; ok && v {
				str.WriteString(redactionStr)
				it.Skip()
				continue
			}
			d.redactSensitive(it, str)
		}
		str.WriteObjectEnd()
	case jsoniter.ArrayValue:
		str.WriteArrayStart()
		first := true
		for it.ReadArray() {
			if !first {
				str.WriteMore()
			}
			first = false
			d.redactSensitive(it, str)
		}
		str.WriteArrayEnd()
	case jsoniter.StringValue:
		str.WriteString(it.ReadString())
	case jsoniter.NumberValue:
		v := it.ReadNumber()
		i, err := v.Int64()
		if err == nil {
			str.WriteInt64(i)
			break
		}
		f, err := v.Float64()
		if err == nil {
			str.WriteFloat64(f)
		}
	case jsoniter.BoolValue:
		str.WriteBool(it.ReadBool())
	case jsoniter.NilValue:
		str.WriteNil()
		// Skip because no reading from it is involved
		it.Skip()
	}
}
