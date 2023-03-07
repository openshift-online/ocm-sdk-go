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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	time "time"
)

// VersionBuilder contains the data and logic needed to build 'version' objects.
//
// Representation of an _OpenShift_ version.
type VersionBuilder struct {
	bitmap_            uint32
	id                 string
	href               string
	availableUpgrades  []string
	channelGroup       string
	endOfLifeTimestamp time.Time
	rawID              string
	releaseImage       string
	rosaEnabled        bool
	default_           bool
	enabled            bool
	hypershiftEnabled  bool
}

// NewVersion creates a new builder of 'version' objects.
func NewVersion() *VersionBuilder {
	return &VersionBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *VersionBuilder) Link(value bool) *VersionBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *VersionBuilder) ID(value string) *VersionBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *VersionBuilder) HREF(value string) *VersionBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *VersionBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// ROSAEnabled sets the value of the 'ROSA_enabled' attribute to the given value.
func (b *VersionBuilder) ROSAEnabled(value bool) *VersionBuilder {
	b.rosaEnabled = value
	b.bitmap_ |= 8
	return b
}

// AvailableUpgrades sets the value of the 'available_upgrades' attribute to the given values.
func (b *VersionBuilder) AvailableUpgrades(values ...string) *VersionBuilder {
	b.availableUpgrades = make([]string, len(values))
	copy(b.availableUpgrades, values)
	b.bitmap_ |= 16
	return b
}

// ChannelGroup sets the value of the 'channel_group' attribute to the given value.
func (b *VersionBuilder) ChannelGroup(value string) *VersionBuilder {
	b.channelGroup = value
	b.bitmap_ |= 32
	return b
}

// Default sets the value of the 'default' attribute to the given value.
func (b *VersionBuilder) Default(value bool) *VersionBuilder {
	b.default_ = value
	b.bitmap_ |= 64
	return b
}

// Enabled sets the value of the 'enabled' attribute to the given value.
func (b *VersionBuilder) Enabled(value bool) *VersionBuilder {
	b.enabled = value
	b.bitmap_ |= 128
	return b
}

// EndOfLifeTimestamp sets the value of the 'end_of_life_timestamp' attribute to the given value.
func (b *VersionBuilder) EndOfLifeTimestamp(value time.Time) *VersionBuilder {
	b.endOfLifeTimestamp = value
	b.bitmap_ |= 256
	return b
}

// HypershiftEnabled sets the value of the 'hypershift_enabled' attribute to the given value.
func (b *VersionBuilder) HypershiftEnabled(value bool) *VersionBuilder {
	b.hypershiftEnabled = value
	b.bitmap_ |= 512
	return b
}

// RawID sets the value of the 'raw_ID' attribute to the given value.
func (b *VersionBuilder) RawID(value string) *VersionBuilder {
	b.rawID = value
	b.bitmap_ |= 1024
	return b
}

// ReleaseImage sets the value of the 'release_image' attribute to the given value.
func (b *VersionBuilder) ReleaseImage(value string) *VersionBuilder {
	b.releaseImage = value
	b.bitmap_ |= 2048
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *VersionBuilder) Copy(object *Version) *VersionBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	b.rosaEnabled = object.rosaEnabled
	if object.availableUpgrades != nil {
		b.availableUpgrades = make([]string, len(object.availableUpgrades))
		copy(b.availableUpgrades, object.availableUpgrades)
	} else {
		b.availableUpgrades = nil
	}
	b.channelGroup = object.channelGroup
	b.default_ = object.default_
	b.enabled = object.enabled
	b.endOfLifeTimestamp = object.endOfLifeTimestamp
	b.hypershiftEnabled = object.hypershiftEnabled
	b.rawID = object.rawID
	b.releaseImage = object.releaseImage
	return b
}

// Build creates a 'version' object using the configuration stored in the builder.
func (b *VersionBuilder) Build() (object *Version, err error) {
	object = new(Version)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	object.rosaEnabled = b.rosaEnabled
	if b.availableUpgrades != nil {
		object.availableUpgrades = make([]string, len(b.availableUpgrades))
		copy(object.availableUpgrades, b.availableUpgrades)
	}
	object.channelGroup = b.channelGroup
	object.default_ = b.default_
	object.enabled = b.enabled
	object.endOfLifeTimestamp = b.endOfLifeTimestamp
	object.hypershiftEnabled = b.hypershiftEnabled
	object.rawID = b.rawID
	object.releaseImage = b.releaseImage
	return
}
