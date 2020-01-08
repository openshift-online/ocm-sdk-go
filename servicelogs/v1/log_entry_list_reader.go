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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/servicelogs/v1

import (
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// logEntryListData is type used internally to marshal and unmarshal lists of objects
// of type 'log_entry'.
type logEntryListData []*logEntryData

// UnmarshalLogEntryList reads a list of values of the 'log_entry'
// from the given source, which can be a slice of bytes, a string, an io.Reader or a
// json.Decoder.
func UnmarshalLogEntryList(source interface{}) (list *LogEntryList, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	var data logEntryListData
	err = decoder.Decode(&data)
	if err != nil {
		return
	}
	list, err = data.unwrap()
	return
}

// wrap is the method used internally to convert a list of values of the
// 'log_entry' value to a JSON document.
func (l *LogEntryList) wrap() (data logEntryListData, err error) {
	if l == nil {
		return
	}
	data = make(logEntryListData, len(l.items))
	for i, item := range l.items {
		data[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// list of values of the 'log_entry' type.
func (d logEntryListData) unwrap() (list *LogEntryList, err error) {
	if d == nil {
		return
	}
	items := make([]*LogEntry, len(d))
	for i, item := range d {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(LogEntryList)
	list.items = items
	return
}

// logEntryListLinkData is type used internally to marshal and unmarshal links
// to lists of objects of type 'log_entry'.
type logEntryListLinkData struct {
	Kind  *string         "json:\"kind,omitempty\""
	HREF  *string         "json:\"href,omitempty\""
	Items []*logEntryData "json:\"items,omitempty\""
}

// wrapLink is the method used internally to convert a list of values of the
// 'log_entry' value to a link.
func (l *LogEntryList) wrapLink() (data *logEntryListLinkData, err error) {
	if l == nil {
		return
	}
	items := make([]*logEntryData, len(l.items))
	for i, item := range l.items {
		items[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	data = new(logEntryListLinkData)
	data.Items = items
	data.HREF = l.href
	data.Kind = new(string)
	if l.link {
		*data.Kind = LogEntryListLinkKind
	} else {
		*data.Kind = LogEntryListKind
	}
	return
}

// unwrapLink is the function used internally to convert a JSON link to a list
// of values of the 'log_entry' type to a list.
func (d *logEntryListLinkData) unwrapLink() (list *LogEntryList, err error) {
	if d == nil {
		return
	}
	items := make([]*LogEntry, len(d.Items))
	for i, item := range d.Items {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(LogEntryList)
	list.items = items
	list.href = d.HREF
	if d.Kind != nil {
		list.link = *d.Kind == LogEntryListLinkKind
	}
	return
}
