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

// ClusterOperatorsConditions represents the values of the 'cluster_operators_conditions' type.
//
// Counts of operators conditions inside a cluster.
type ClusterOperatorsConditions struct {
	available *int
	degraded  *int
	failing   *int
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterOperatorsConditions) Empty() bool {
	return o == nil || (o.available == nil &&
		o.degraded == nil &&
		o.failing == nil &&
		true)
}

// Available returns the value of the 'available' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of "available" cluster operators.
func (o *ClusterOperatorsConditions) Available() int {
	if o != nil && o.available != nil {
		return *o.available
	}
	return 0
}

// GetAvailable returns the value of the 'available' attribute and
// a flag indicating if the attribute has a value.
//
// Number of "available" cluster operators.
func (o *ClusterOperatorsConditions) GetAvailable() (value int, ok bool) {
	ok = o != nil && o.available != nil
	if ok {
		value = *o.available
	}
	return
}

// Degraded returns the value of the 'degraded' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of "degraded" cluster operators.
func (o *ClusterOperatorsConditions) Degraded() int {
	if o != nil && o.degraded != nil {
		return *o.degraded
	}
	return 0
}

// GetDegraded returns the value of the 'degraded' attribute and
// a flag indicating if the attribute has a value.
//
// Number of "degraded" cluster operators.
func (o *ClusterOperatorsConditions) GetDegraded() (value int, ok bool) {
	ok = o != nil && o.degraded != nil
	if ok {
		value = *o.degraded
	}
	return
}

// Failing returns the value of the 'failing' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of "failing" cluster operators.
func (o *ClusterOperatorsConditions) Failing() int {
	if o != nil && o.failing != nil {
		return *o.failing
	}
	return 0
}

// GetFailing returns the value of the 'failing' attribute and
// a flag indicating if the attribute has a value.
//
// Number of "failing" cluster operators.
func (o *ClusterOperatorsConditions) GetFailing() (value int, ok bool) {
	ok = o != nil && o.failing != nil
	if ok {
		value = *o.failing
	}
	return
}

// ClusterOperatorsConditionsListKind is the name of the type used to represent list of objects of
// type 'cluster_operators_conditions'.
const ClusterOperatorsConditionsListKind = "ClusterOperatorsConditionsList"

// ClusterOperatorsConditionsListLinkKind is the name of the type used to represent links to list
// of objects of type 'cluster_operators_conditions'.
const ClusterOperatorsConditionsListLinkKind = "ClusterOperatorsConditionsListLink"

// ClusterOperatorsConditionsNilKind is the name of the type used to nil lists of objects of
// type 'cluster_operators_conditions'.
const ClusterOperatorsConditionsListNilKind = "ClusterOperatorsConditionsListNil"

// ClusterOperatorsConditionsList is a list of values of the 'cluster_operators_conditions' type.
type ClusterOperatorsConditionsList struct {
	href  *string
	link  bool
	items []*ClusterOperatorsConditions
}

// Len returns the length of the list.
func (l *ClusterOperatorsConditionsList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterOperatorsConditionsList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterOperatorsConditionsList) Get(i int) *ClusterOperatorsConditions {
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
func (l *ClusterOperatorsConditionsList) Slice() []*ClusterOperatorsConditions {
	var slice []*ClusterOperatorsConditions
	if l == nil {
		slice = make([]*ClusterOperatorsConditions, 0)
	} else {
		slice = make([]*ClusterOperatorsConditions, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterOperatorsConditionsList) Each(f func(item *ClusterOperatorsConditions) bool) {
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
func (l *ClusterOperatorsConditionsList) Range(f func(index int, item *ClusterOperatorsConditions) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
