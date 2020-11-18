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

// This file contains the the implementation of the configuration object.

package configuration

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

// Builder contains the data and logic needed to populate a configuration object. Don't create
// instances of this type directly, use the New function instead.
type Builder struct {
	sources []interface{}
}

// Object contains configuration data.
type Object struct {
	tree *yaml.Node
}

// New creates a new builder that can be use to populate a configuration object.
func New() *Builder {
	return &Builder{}
}

// Load adds the given objects as sources where the configuration will be loaded from.
//
// If a source is a string it will be interpreted as the name of a file containing the YAML text.
//
// If a source is an array of bytes it will be interpreted as the actual YAML text.
//
// If a source implements the io.Reader interface, then it will be sued to read in memory the YAML
// text.
//
// If the source can also be a yaml.Node or another configuration Object. In those cases the
// content of the source will be copied.
//
// If the source is any other kind of object then it will be serialized as YAML and then loaded.
func (b *Builder) Load(sources ...interface{}) *Builder {
	b.sources = append(b.sources, sources...)
	return b
}

// Build uses the information stored in the builder to create and populate a configuration
// object.
func (b *Builder) Build() (object *Object, err error) {
	// Merge the sources:
	tree := &yaml.Node{}
	for _, current := range b.sources {
		switch source := current.(type) {
		case string:
			err = b.mergeFile(source, tree)
		case []byte:
			err = b.mergeBytes(source, tree)
		case io.Reader:
			err = b.mergeReader(source, tree)
		case yaml.Node:
			err = b.mergeNode(&source, tree)
		case *yaml.Node:
			err = b.mergeNode(source, tree)
		case *Object:
			err = b.mergeNode(source.tree, tree)
		case Object:
			err = b.mergeNode(source.tree, tree)
		default:
			err = b.mergeAny(source, tree)
		}
		if err != nil {
			return
		}
	}

	// Process the tags:
	err = processTags(tree)
	if err != nil {
		return
	}

	// Create and populate the object:
	object = &Object{
		tree: tree,
	}

	return
}

func (b *Builder) mergeString(src string, dst *yaml.Node) error {
	return b.mergeBytes([]byte(src), dst)
}

func (b *Builder) mergeReader(src io.Reader, dst *yaml.Node) error {
	buffer, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	return b.mergeBytes(buffer, dst)
}

func (b *Builder) mergeFile(src string, dst *yaml.Node) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return b.mergeDir(src, dst)
	}
	buffer, err := ioutil.ReadFile(src) // #nosec G304
	if err != nil {
		return err
	}
	return b.mergeBytes(buffer, dst)
}

func (b *Builder) mergeDir(src string, dst *yaml.Node) error {
	infos, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	files := make([]string, 0, len(infos))
	for _, info := range infos {
		name := info.Name()
		ext := filepath.Ext(name)
		if ext == ".yaml" || ext == ".yml" {
			files = append(files, filepath.Join(src, name))
		}
	}
	sort.Strings(files)
	for _, file := range files {
		err = b.mergeFile(file, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Builder) mergeAny(src interface{}, dst *yaml.Node) error {
	buffer, err := yaml.Marshal(src)
	if err != nil {
		return err
	}
	return b.mergeBytes(buffer, dst)
}

func (b *Builder) mergeBytes(src []byte, dst *yaml.Node) error {
	var tree yaml.Node
	err := yaml.Unmarshal(src, &tree)
	if err != nil {
		return err
	}
	return b.mergeNode(&tree, dst)
}

func (b *Builder) mergeNode(src, dst *yaml.Node) error {
	if src.Kind != dst.Kind {
		b.deepCopy(src, dst)
		return nil
	}
	switch src.Kind {
	case 0:
		return b.mergeEmpty(src, dst)
	case yaml.DocumentNode:
		return b.mergeDocument(src, dst)
	case yaml.SequenceNode:
		return b.mergeSequence(src, dst)
	case yaml.MappingNode:
		return b.mergeMapping(src, dst)
	case yaml.ScalarNode:
		return b.mergeScalar(src, dst)
	case yaml.AliasNode:
		return b.mergeAlias(src, dst)
	default:
		return fmt.Errorf("don't know how to handle YAML node of type %d", src.Kind)
	}
}

func (b *Builder) mergeDocument(src, dst *yaml.Node) error {
	return b.mergeNode(src.Content[0], dst.Content[0])
}

func (b *Builder) mergeEmpty(src, dst *yaml.Node) error {
	return nil
}

func (b *Builder) mergeSequence(src, dst *yaml.Node) error {
	size := len(src.Content)
	nodes := make([]*yaml.Node, size)
	for i := 0; i < size; i++ {
		nodes[i] = &yaml.Node{}
		b.deepCopy(src.Content[i], nodes[i])
	}
	dst.Content = append(dst.Content, nodes...)
	return nil
}

func (b *Builder) mergeMapping(src, dst *yaml.Node) error {
	srcSize := len(src.Content) / 2
	i := 0
	for i < srcSize {
		srcKey := src.Content[2*i]
		srcValue := src.Content[2*i+1]
		dstSize := len(dst.Content) / 2
		j := 0
		for j < dstSize {
			dstKey := dst.Content[2*j]
			dstValue := dst.Content[2*j+1]
			if srcKey.Value == dstKey.Value {
				err := b.mergeNode(srcValue, dstValue)
				if err != nil {
					return err
				}
				break
			}
			j++
		}
		if j == dstSize {
			dstKey := &yaml.Node{}
			b.deepCopy(srcKey, dstKey)
			dstValue := &yaml.Node{}
			b.deepCopy(srcValue, dstValue)
			dst.Content = append(dst.Content, dstKey, dstValue)
		}
		i++
	}
	return nil
}

func (b *Builder) mergeScalar(src, dst *yaml.Node) error {
	b.deepCopy(src, dst)
	return nil
}

func (b *Builder) mergeAlias(src, dst *yaml.Node) error {
	b.deepCopy(src, dst)
	return nil
}

func (b *Builder) deepCopy(src, dst *yaml.Node) {
	*dst = *src
	if src.Content != nil {
		size := len(src.Content)
		dst.Content = make([]*yaml.Node, size)
		for i := 0; i < size; i++ {
			dst.Content[i] = &yaml.Node{}
			b.deepCopy(src.Content[i], dst.Content[i])
		}
	}
}

// Populate populates the given destination object with the information stored in this
// configuration object. The destination object should be a pointer to a variable containing
// the same tags used by the yaml.Unmarshal method of the YAML library.
func (o *Object) Populate(v interface{}) error {
	return o.tree.Decode(v)
}

// MarshalYAML is the implementation of the yaml.Marshaller interface. This is intended to be able
// use the type for fields inside other structs. Refrain from calling this method for any other
// use.
func (o *Object) MarshalYAML() (result interface{}, err error) {
	result = o.tree
	return
}

// UnmarshalYAML is the implementation of the yaml.Unmarshaller interface. This is intended to be
// able to use the type for fields inside structs. Refraim from calling this method for any other
// use.
func (o *Object) UnmarshalYAML(value *yaml.Node) error {
	o.tree = value
	return nil
}
