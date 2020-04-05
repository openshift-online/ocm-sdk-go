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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterAlertsFiring represents the values of the 'cluster_alerts_firing' type.
//
// Counts of different alerts firing inside a cluster.
type ClusterAlertsFiring struct {
	critical *int
	high     *int
	none     *int
	warning  *int
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterAlertsFiring) Empty() bool {
	return o == nil || (o.critical == nil &&
		o.high == nil &&
		o.none == nil &&
		o.warning == nil &&
		true)
}

// Critical returns the value of the 'critical' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of "critical" alerts firing.
func (o *ClusterAlertsFiring) Critical() int {
	if o != nil && o.critical != nil {
		return *o.critical
	}
	return 0
}

// GetCritical returns the value of the 'critical' attribute and
// a flag indicating if the attribute has a value.
//
// Number of "critical" alerts firing.
func (o *ClusterAlertsFiring) GetCritical() (value int, ok bool) {
	ok = o != nil && o.critical != nil
	if ok {
		value = *o.critical
	}
	return
}

// High returns the value of the 'high' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of "high" alerts firing.
func (o *ClusterAlertsFiring) High() int {
	if o != nil && o.high != nil {
		return *o.high
	}
	return 0
}

// GetHigh returns the value of the 'high' attribute and
// a flag indicating if the attribute has a value.
//
// Number of "high" alerts firing.
func (o *ClusterAlertsFiring) GetHigh() (value int, ok bool) {
	ok = o != nil && o.high != nil
	if ok {
		value = *o.high
	}
	return
}

// None returns the value of the 'none' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of "none" alerts firing.
func (o *ClusterAlertsFiring) None() int {
	if o != nil && o.none != nil {
		return *o.none
	}
	return 0
}

// GetNone returns the value of the 'none' attribute and
// a flag indicating if the attribute has a value.
//
// Number of "none" alerts firing.
func (o *ClusterAlertsFiring) GetNone() (value int, ok bool) {
	ok = o != nil && o.none != nil
	if ok {
		value = *o.none
	}
	return
}

// Warning returns the value of the 'warning' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of "warning" alerts firing.
func (o *ClusterAlertsFiring) Warning() int {
	if o != nil && o.warning != nil {
		return *o.warning
	}
	return 0
}

// GetWarning returns the value of the 'warning' attribute and
// a flag indicating if the attribute has a value.
//
// Number of "warning" alerts firing.
func (o *ClusterAlertsFiring) GetWarning() (value int, ok bool) {
	ok = o != nil && o.warning != nil
	if ok {
		value = *o.warning
	}
	return
}

// ClusterAlertsFiringListKind is the name of the type used to represent list of objects of
// type 'cluster_alerts_firing'.
const ClusterAlertsFiringListKind = "ClusterAlertsFiringList"

// ClusterAlertsFiringListLinkKind is the name of the type used to represent links to list
// of objects of type 'cluster_alerts_firing'.
const ClusterAlertsFiringListLinkKind = "ClusterAlertsFiringListLink"

// ClusterAlertsFiringNilKind is the name of the type used to nil lists of objects of
// type 'cluster_alerts_firing'.
const ClusterAlertsFiringListNilKind = "ClusterAlertsFiringListNil"

// ClusterAlertsFiringList is a list of values of the 'cluster_alerts_firing' type.
type ClusterAlertsFiringList struct {
	href  *string
	link  bool
	items []*ClusterAlertsFiring
}

// Len returns the length of the list.
func (l *ClusterAlertsFiringList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterAlertsFiringList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterAlertsFiringList) Get(i int) *ClusterAlertsFiring {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterAlertsFiringList) Slice() []*ClusterAlertsFiring {
	var slice []*ClusterAlertsFiring
	if l == nil {
		slice = make([]*ClusterAlertsFiring, 0)
	} else {
		slice = make([]*ClusterAlertsFiring, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterAlertsFiringList) Each(f func(item *ClusterAlertsFiring) bool) {
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
func (l *ClusterAlertsFiringList) Range(f func(index int, item *ClusterAlertsFiring) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
