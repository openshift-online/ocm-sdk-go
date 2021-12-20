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
	"io"
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalIdentityProvider writes a value of the 'identity_provider' type to the given writer.
func MarshalIdentityProvider(object *IdentityProvider, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeIdentityProvider(object, stream)
	stream.Flush()
	return stream.Error
}

// writeIdentityProvider writes a value of the 'identity_provider' type to the given stream.
func writeIdentityProvider(object *IdentityProvider, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if object.bitmap_&1 != 0 {
		stream.WriteString(IdentityProviderLinkKind)
	} else {
		stream.WriteString(IdentityProviderKind)
	}
	count++
	if object.bitmap_&2 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if object.bitmap_&4 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = object.bitmap_&8 != 0 && object.ldap != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("ldap")
		writeLDAPIdentityProvider(object.ldap, stream)
		count++
	}
	present_ = object.bitmap_&16 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("challenge")
		stream.WriteBool(object.challenge)
		count++
	}
	present_ = object.bitmap_&32 != 0 && object.github != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("github")
		writeGithubIdentityProvider(object.github, stream)
		count++
	}
	present_ = object.bitmap_&64 != 0 && object.gitlab != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("gitlab")
		writeGitlabIdentityProvider(object.gitlab, stream)
		count++
	}
	present_ = object.bitmap_&128 != 0 && object.google != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("google")
		writeGoogleIdentityProvider(object.google, stream)
		count++
	}
	present_ = object.bitmap_&256 != 0 && object.htpasswd != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("htpasswd")
		writeHTPasswdIdentityProvider(object.htpasswd, stream)
		count++
	}
	present_ = object.bitmap_&512 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("login")
		stream.WriteBool(object.login)
		count++
	}
	present_ = object.bitmap_&1024 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("mapping_method")
		stream.WriteString(string(object.mappingMethod))
		count++
	}
	present_ = object.bitmap_&2048 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("name")
		stream.WriteString(object.name)
		count++
	}
	present_ = object.bitmap_&4096 != 0 && object.openID != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("open_id")
		writeOpenIDIdentityProvider(object.openID, stream)
		count++
	}
	present_ = object.bitmap_&8192 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("type")
		stream.WriteString(string(object.type_))
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalIdentityProvider reads a value of the 'identity_provider' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalIdentityProvider(source interface{}) (object *IdentityProvider, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readIdentityProvider(iterator)
	err = iterator.Error
	return
}

// readIdentityProvider reads a value of the 'identity_provider' type from the given iterator.
func readIdentityProvider(iterator *jsoniter.Iterator) *IdentityProvider {
	object := &IdentityProvider{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == IdentityProviderLinkKind {
				object.bitmap_ |= 1
			}
		case "id":
			object.id = iterator.ReadString()
			object.bitmap_ |= 2
		case "href":
			object.href = iterator.ReadString()
			object.bitmap_ |= 4
		case "ldap":
			value := readLDAPIdentityProvider(iterator)
			object.ldap = value
			object.bitmap_ |= 8
		case "challenge":
			value := iterator.ReadBool()
			object.challenge = value
			object.bitmap_ |= 16
		case "github":
			value := readGithubIdentityProvider(iterator)
			object.github = value
			object.bitmap_ |= 32
		case "gitlab":
			value := readGitlabIdentityProvider(iterator)
			object.gitlab = value
			object.bitmap_ |= 64
		case "google":
			value := readGoogleIdentityProvider(iterator)
			object.google = value
			object.bitmap_ |= 128
		case "htpasswd":
			value := readHTPasswdIdentityProvider(iterator)
			object.htpasswd = value
			object.bitmap_ |= 256
		case "login":
			value := iterator.ReadBool()
			object.login = value
			object.bitmap_ |= 512
		case "mapping_method":
			text := iterator.ReadString()
			value := IdentityProviderMappingMethod(text)
			object.mappingMethod = value
			object.bitmap_ |= 1024
		case "name":
			value := iterator.ReadString()
			object.name = value
			object.bitmap_ |= 2048
		case "open_id":
			value := readOpenIDIdentityProvider(iterator)
			object.openID = value
			object.bitmap_ |= 4096
		case "type":
			text := iterator.ReadString()
			value := IdentityProviderType(text)
			object.type_ = value
			object.bitmap_ |= 8192
		default:
			iterator.ReadAny()
		}
	}
	return object
}
