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
	time "time"

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// logEntryData is the data structure used internally to marshal and unmarshal
// objects of type 'log_entry'.
type logEntryData struct {
	Kind         *string    "json:\"kind,omitempty\""
	ID           *string    "json:\"id,omitempty\""
	HREF         *string    "json:\"href,omitempty\""
	ClusterUUID  *string    "json:\"cluster_uuid,omitempty\""
	Description  *string    "json:\"description,omitempty\""
	InternalOnly *bool      "json:\"internal_only,omitempty\""
	ServiceName  *string    "json:\"service_name,omitempty\""
	Severity     *Severity  "json:\"severity,omitempty\""
	Summary      *string    "json:\"summary,omitempty\""
	Timestamp    *time.Time "json:\"timestamp,omitempty\""
}

// MarshalLogEntry writes a value of the 'log_entry' to the given target,
// which can be a writer or a JSON encoder.
func MarshalLogEntry(object *LogEntry, target interface{}) error {
	encoder, err := helpers.NewEncoder(target)
	if err != nil {
		return err
	}
	data, err := object.wrap()
	if err != nil {
		return err
	}
	return encoder.Encode(data)
}

// wrap is the method used internally to convert a value of the 'log_entry'
// value to a JSON document.
func (o *LogEntry) wrap() (data *logEntryData, err error) {
	if o == nil {
		return
	}
	data = new(logEntryData)
	data.ID = o.id
	data.HREF = o.href
	data.Kind = new(string)
	if o.link {
		*data.Kind = LogEntryLinkKind
	} else {
		*data.Kind = LogEntryKind
	}
	data.ClusterUUID = o.clusterUUID
	data.Description = o.description
	data.InternalOnly = o.internalOnly
	data.ServiceName = o.serviceName
	data.Severity = o.severity
	data.Summary = o.summary
	data.Timestamp = o.timestamp
	return
}

// UnmarshalLogEntry reads a value of the 'log_entry' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalLogEntry(source interface{}) (object *LogEntry, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(logEntryData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'log_entry' type.
func (d *logEntryData) unwrap() (object *LogEntry, err error) {
	if d == nil {
		return
	}
	object = new(LogEntry)
	object.id = d.ID
	object.href = d.HREF
	if d.Kind != nil {
		object.link = *d.Kind == LogEntryLinkKind
	}
	object.clusterUUID = d.ClusterUUID
	object.description = d.Description
	object.internalOnly = d.InternalOnly
	object.serviceName = d.ServiceName
	object.severity = d.Severity
	object.summary = d.Summary
	object.timestamp = d.Timestamp
	return
}
