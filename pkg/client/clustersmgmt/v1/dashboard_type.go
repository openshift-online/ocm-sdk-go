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

// DashboardKind is the name of the type used to represent objects
// of type 'dashboard'.
const DashboardKind = "Dashboard"

// DashboardLinkKind is the name of the type used to represent links
// to objects of type 'dashboard'.
const DashboardLinkKind = "DashboardLink"

// DashboardNilKind is the name of the type used to nil references
// to objects of type 'dashboard'.
const DashboardNilKind = "DashboardNil"

// Dashboard represents the values of the 'dashboard' type.
//
// Collection of metrics intended to render a graphical dashboard.
type Dashboard struct {
	id      *string
	href    *string
	link    bool
	metrics *MetricList
}

// Kind returns the name of the type of the object.
func (o *Dashboard) Kind() string {
	if o == nil {
		return DashboardNilKind
	}
	if o.link {
		return DashboardLinkKind
	}
	return DashboardKind
}

// ID returns the identifier of the object.
func (o *Dashboard) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Dashboard) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Dashboard) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Dashboard) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Dashboard) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Metrics returns the value of the 'metrics' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Metrics included in the dashboard.
func (o *Dashboard) Metrics() *MetricList {
	if o == nil {
		return nil
	}
	return o.metrics
}

// GetMetrics returns the value of the 'metrics' attribute and
// a flag indicating if the attribute has a value.
//
// Metrics included in the dashboard.
func (o *Dashboard) GetMetrics() (value *MetricList, ok bool) {
	ok = o != nil && o.metrics != nil
	if ok {
		value = o.metrics
	}
	return
}

// DashboardList is a list of values of the 'dashboard' type.
type DashboardList struct {
	items []*Dashboard
}

// Len returns the length of the list.
func (l *DashboardList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *DashboardList) Slice() []*Dashboard {
	var slice []*Dashboard
	if l == nil {
		slice = make([]*Dashboard, 0)
	} else {
		slice = make([]*Dashboard, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *DashboardList) Each(f func(item *Dashboard) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *DashboardList) Range(f func(index int, item *Dashboard) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
