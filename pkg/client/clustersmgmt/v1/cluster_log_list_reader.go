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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

import (
	"fmt"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// clusterLogListData is type used internally to marshal and unmarshal lists of objects
// of type 'cluster_log'.
type clusterLogListData []*clusterLogData

// UnmarshalClusterLogList reads a list of values of the 'cluster_log'
// from the given source, which can be a slice of bytes, a string, an io.Reader or a
// json.Decoder.
func UnmarshalClusterLogList(source interface{}) (list *ClusterLogList, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	var data clusterLogListData
	err = decoder.Decode(&data)
	if err != nil {
		return
	}
	list, err = data.unwrap()
	return
}

// wrap is the method used internally to convert a list of values of the
// 'cluster_log' value to a JSON document.
func (l *ClusterLogList) wrap() (data clusterLogListData, err error) {
	if l == nil {
		return
	}
	data = make(clusterLogListData, len(l.items))
	for i, item := range l.items {
		data[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// list of values of the 'cluster_log' type.
func (d clusterLogListData) unwrap() (list *ClusterLogList, err error) {
	if d == nil {
		return
	}
	items := make([]*ClusterLog, len(d))
	for i, item := range d {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(ClusterLogList)
	list.items = items
	return
}

// clusterLogListLinkData is type used internally to marshal and unmarshal links
// to lists of objects of type 'cluster_log'.
type clusterLogListLinkData struct {
	Kind  *string           "json:\"kind,omitempty\""
	HREF  *string           "json:\"href,omitempty\""
	Items []*clusterLogData "json:\"items,omitempty\""
}

// wrapLink is the method used internally to convert a list of values of the
// 'cluster_log' value to a link.
func (l *ClusterLogList) wrapLink() (data *clusterLogListLinkData, err error) {
	if l == nil {
		return
	}
	items := make([]*clusterLogData, len(l.items))
	for i, item := range l.items {
		items[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	data = new(clusterLogListLinkData)
	data.Items = items
	data.HREF = l.href
	data.Kind = new(string)
	if l.link {
		*data.Kind = ClusterLogListLinkKind
	} else {
		*data.Kind = ClusterLogListKind
	}
	return
}

// unwrapLink is the function used internally to convert a JSON link to a list
// of values of the 'cluster_log' type to a list.
func (d *clusterLogListLinkData) unwrapLink() (list *ClusterLogList, err error) {
	if d == nil {
		return
	}
	items := make([]*ClusterLog, len(d.Items))
	for i, item := range d.Items {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(ClusterLogList)
	list.items = items
	list.href = d.HREF
	if d.Kind != nil {
		switch *d.Kind {
		case ClusterLogListKind:
			list.link = false
		case ClusterLogListLinkKind:
			list.link = true
		default:
			err = fmt.Errorf(
				"expected kind '%s' or '%s' but got '%s'",
				ClusterLogListKind,
				ClusterLogListLinkKind,
				*d.Kind,
			)
			return
		}
	}
	return
}
