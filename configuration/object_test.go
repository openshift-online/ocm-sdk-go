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

// This file contains tests for the configuration object.

package configuration

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	. "github.com/onsi/ginkgo"                  // nolint
	. "github.com/onsi/ginkgo/extensions/table" // nolint
	. "github.com/onsi/gomega"                  // nolint
)

var _ = Describe("Object", func() {
	Describe("Load", func() {
		It("Can be loaded from bytes", func() {
			// Load the configuration:
			object, err := New().
				Load([]byte(`mykey: myvalue`)).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
		})

		It("Can be loaded from reader", func() {
			// Load the configuration:
			object, err := New().
				Load(strings.NewReader(`mykey: myvalue`)).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
		})

		It("Can be loaded from file", func() {
			// Create a temporary file containing the configuration:
			tmp, err := ioutil.TempFile("", "*.test.yaml")
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = os.Remove(tmp.Name())
				Expect(err).ToNot(HaveOccurred())
			}()
			_, err = tmp.Write([]byte(`mykey: myvalue`))
			Expect(err).ToNot(HaveOccurred())
			err = tmp.Close()
			Expect(err).ToNot(HaveOccurred())

			// Load the configuration:
			object, err := New().
				Load(tmp.Name()).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
		})

		It("Can be loaded from struct", func() {
			// Load the configuration:
			type Source struct {
				MyKey string `yaml:"mykey"`
			}
			object, err := New().
				Load(&Source{
					MyKey: "myvalue",
				}).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
		})

		It("Can be loaded from map", func() {
			// Load the configuration:
			object, err := New().
				Load(map[string]string{
					"mykey": "myvalue",
				}).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
		})

		It("Can be loaded from YAML node", func() {
			// Create the YAML node:
			var node yaml.Node
			err := yaml.Unmarshal([]byte(`mykey: myvalue`), &node)
			Expect(err).ToNot(HaveOccurred())

			// Load the configuration:
			object, err := New().
				Load(node).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
		})

		It("Can be loaded from configuration object", func() {
			// Load the source configuration:
			source, err := New().
				Load([]byte(`mykey: myvalue`)).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(source).ToNot(BeNil())

			// Load the configuration:
			object, err := New().
				Load(source).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
		})

		It("Can be loaded from directory", func() {
			// Create a temporary directory containing two configuration files:
			tmp, err := ioutil.TempDir("", "*.test.d")
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = os.RemoveAll(tmp)
				Expect(err).ToNot(HaveOccurred())
			}()
			first := filepath.Join(tmp, "my.yaml")
			err = ioutil.WriteFile(first, []byte("mykey: myvalue"), 0600)
			Expect(err).ToNot(HaveOccurred())
			second := filepath.Join(tmp, "your.yaml")
			err = ioutil.WriteFile(second, []byte("yourkey: yourvalue"), 0600)

			// Load the configuration:
			object, err := New().
				Load(tmp).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey   string `yaml:"mykey"`
				YourKey string `yaml:"yourkey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
			Expect(config.YourKey).To(Equal("yourvalue"))
		})

		It("Honours order of files in directory", func() {
			// Create a temporary directory containing two configuration files:
			tmp, err := ioutil.TempDir("", "*.test.d")
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err = os.RemoveAll(tmp)
				Expect(err).ToNot(HaveOccurred())
			}()
			first := filepath.Join(tmp, "0.yaml")
			err = ioutil.WriteFile(first, []byte("mykey: firstvalue"), 0600)
			Expect(err).ToNot(HaveOccurred())
			second := filepath.Join(tmp, "1.yaml")
			err = ioutil.WriteFile(second, []byte("mykey: secondvalue"), 0600)

			// Load the configuration:
			object, err := New().
				Load(tmp).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("secondvalue"))
		})

		It("Decodes binary data", func() {
			// Load the configuration:
			object, err := New().
				Load([]byte(`mykey: !!binary bXl2YWx1ZQ==`)).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("myvalue"))
		})
	})

	Describe("Merge", func() {
		It("Honours order of sources", func() {
			// Load the configuration:
			object, err := New().
				Load([]byte(`mykey: myvalue`)).
				Load([]byte(`mykey: yourvalue`)).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyKey string `yaml:"mykey"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyKey).To(Equal("yourvalue"))
		})

		It("Adds item to slice", func() {
			// Load the configuration:
			object, err := New().
				Load([]byte(`"myslice": [ "firstvalue" ]`)).
				Load([]byte(`"myslice": [ "secondvalue" ]`)).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MySlice []string `yaml:"myslice"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MySlice).To(ConsistOf("firstvalue", "secondvalue"))
		})

		It("Adds entry to map", func() {
			// Load the configuration:
			object, err := New().
				Load([]byte(`"mymap": { firstkey: firstvalue }`)).
				Load([]byte(`"mymap": { secondkey: secondvalue }`)).
				Build()
			Expect(err).ToNot(HaveOccurred())
			Expect(object).ToNot(BeNil())

			// Populate the configuration:
			var config struct {
				MyMap map[string]string `yaml:"mymap"`
			}
			err = object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.MyMap).To(HaveLen(2))
			Expect(config.MyMap).To(HaveKeyWithValue("firstkey", "firstvalue"))
			Expect(config.MyMap).To(HaveKeyWithValue("secondkey", "secondvalue"))
		})
	})

	It("Can be used as a struct field", func() {
		var top, sub func(*Object)

		// This simulates a top level component that uses its own struct to decode the
		// configuration but has no visibility of the struct used by the sub-component.
		top = func(object *Object) {
			// Populate and verify the configuration:
			var config struct {
				TopKey string  `yaml:"topkey"`
				Sub    *Object `yaml:"sub"`
			}
			err := object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.TopKey).To(Equal("topvalue"))
			Expect(config.Sub).ToNot(BeNil())

			// Call the sub-component:
			sub(config.Sub)
		}

		// This simulates a sub-component that uses its own private struct to decode the
		// configuration.
		sub = func(object *Object) {
			// Populate and verify the configuration:
			var config struct {
				SubKey string `yaml:"subkey"`
			}
			err := object.Populate(&config)
			Expect(err).ToNot(HaveOccurred())
			Expect(config.SubKey).To(Equal("subvalue"))
		}

		// Load the configuration:
		source := RemoveLeadingTabs(`
				topkey: topvalue
				sub:
				  subkey: subvalue
			`)
		object, err := New().
			Load([]byte(source)).
			Build()
		Expect(err).ToNot(HaveOccurred())
		Expect(object).ToNot(BeNil())

		// Call the top level components:
		top(object)
	})

	Describe("Tags", func() {
		// This type will be used to check values of fields pupulated using the tags in the
		// tests below:
		type Config struct {
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			ID       int    `yaml:"id"`
			Enabled  bool   `yaml:"enabled"`
		}

		DescribeTable(
			"Successful processing",
			func(source string, variables, files map[string]string, expected Config) {
				var err error

				// Set the environment variables:
				for name, value := range variables {
					err = os.Setenv(name, value)
					Expect(err).ToNot(HaveOccurred())
					defer func() {
						err = os.Unsetenv(name)
						Expect(err).ToNot(HaveOccurred())
					}()
				}

				// Create a temporary directory to contain the temporary files:
				var tmp string
				tmp, err = ioutil.TempDir("", "*.test")
				Expect(err).ToNot(HaveOccurred())
				defer func() {
					err = os.RemoveAll(tmp)
					Expect(err).ToNot(HaveOccurred())
				}()

				// Create the files into the temporary directory:
				for name, content := range files {
					path := filepath.Join(tmp, name)
					err = ioutil.WriteFile(path, []byte(content), 0600)
					Expect(err).ToNot(HaveOccurred())
				}

				// Change into the temporary directory:
				var wd string
				wd, err = os.Getwd()
				Expect(err).ToNot(HaveOccurred())
				err = os.Chdir(tmp)
				Expect(err).ToNot(HaveOccurred())
				defer func() {
					err = os.Chdir(wd)
					Expect(err).ToNot(HaveOccurred())
				}()

				// Remove leading tabs from the source:
				source = RemoveLeadingTabs(source)

				// Do the check:
				object, err := New().Load([]byte(source)).Build()
				Expect(err).ToNot(HaveOccurred())
				var config Config
				err = object.Populate(&config)
				Expect(err).ToNot(HaveOccurred())
				Expect(config).To(Equal(expected))
			},
			Entry(
				"Empty",
				"",
				nil,
				nil,
				Config{},
			),
			Entry(
				"One environment variable",
				`
				user: !variable MYUSER
				`,
				map[string]string{
					"MYUSER": "myuser",
				},
				nil,
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Two environment variables",
				`
				user: !variable MYUSER
				password: !variable MYPASSWORD
				`,
				map[string]string{
					"MYUSER":     "myuser",
					"MYPASSWORD": "mypassword",
				},
				nil,
				Config{
					User:     "myuser",
					Password: "mypassword",
				},
			),
			Entry(
				"Environment variable with `var`",
				`
				user: !var MYUSER
				`,
				map[string]string{
					"MYUSER": "myuser",
				},
				nil,
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Environment variable with `v`",
				`
				user: !v MYUSER
				`,
				map[string]string{
					"MYUSER": "myuser",
				},
				nil,
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Environment variable with leading space",
				`
				password: !variable MYPASSWORD
				`,
				map[string]string{
					"MYPASSWORD": " mypassword",
				},
				nil,
				Config{
					Password: " mypassword",
				},
			),
			Entry(
				"Environment variable with trailing space",
				`
				password: !variable MYPASSWORD
				`,
				map[string]string{
					"MYPASSWORD": "mypassword ",
				},
				nil,
				Config{
					Password: "mypassword ",
				},
			),
			Entry(
				"Environment variable with quotes",
				`
				password: !variable MYPASSWORD
				`,
				map[string]string{
					"MYPASSWORD": "my\"pass\"word",
				},
				nil,
				Config{
					Password: "my\"pass\"word",
				},
			),
			Entry(
				"Numeric environment variable",
				`
				id: !variable/integer MYID
				`,
				map[string]string{
					"MYID": "123",
				},
				nil,
				Config{
					ID: 123,
				},
			),
			Entry(
				"Boolean environment variable",
				`
				enabled: !variable/boolean MYENABLED
				`,
				map[string]string{
					"MYENABLED": "true",
				},
				nil,
				Config{
					Enabled: true,
				},
			),
			Entry(
				"One file",
				`
				user: !file myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": "myuser",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Two files",
				`
				user: !file myuser.txt
				password: !file mypassword.txt
				`,
				nil,
				map[string]string{
					"myuser.txt":     "myuser",
					"mypassword.txt": "mypassword",
				},
				Config{
					User:     "myuser",
					Password: "mypassword",
				},
			),
			Entry(
				"File with `f`",
				`
				user: !f myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": "myuser",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"File with leading space preserved",
				`
				user: !file myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": " myuser",
				},
				Config{
					User: " myuser",
				},
			),
			Entry(
				"File with leading space trimmed",
				`
				user: !file/trim myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": "myuser",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"File with trailing space preserved",
				`
				user: !file myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": "myuser ",
				},
				Config{
					User: "myuser ",
				},
			),
			Entry(
				"File with trailing space trimmed",
				`
				user: !file/trim myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": "myuser ",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"File with trailing line break preserved",
				`
				user: !file myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": "myuser\n",
				},
				Config{
					User: "myuser\n",
				},
			),
			Entry(
				"File with trailing line break trimmed",
				`
				user: !file/trim myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": "myuser\n",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Variable and file",
				`
				user: !variable MYUSER
				password: !file mypassword.txt
				`,
				map[string]string{
					"MYUSER": "myuser",
				},
				map[string]string{
					"mypassword.txt": "mypassword",
				},
				Config{
					User:     "myuser",
					Password: "mypassword",
				},
			),
			Entry(
				"File with quotes",
				`
				user: myuser
				password: !file mypassword.txt
				`,
				nil,
				map[string]string{
					"mypassword.txt": "my\"pass\"word",
				},
				Config{
					User:     "myuser",
					Password: "my\"pass\"word",
				},
			),
			Entry(
				"Numeric file",
				`
				id: !file/integer myid.txt
				`,
				nil,
				map[string]string{
					"myid.txt": "123",
				},
				Config{
					ID: 123,
				},
			),
			Entry(
				"Boolean file",
				`
				enabled: !file/boolean myenabled.txt
				`,
				nil,
				map[string]string{
					"myenabled.txt": "true",
				},
				Config{
					Enabled: true,
				},
			),
			Entry(
				"Script echo variable",
				`
				user: !script echo -n ${MYUSER}
				`,
				map[string]string{
					"MYUSER": "myuser",
				},
				nil,
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Script cat file",
				`
				user: !script cat myuser.txt
				`,
				nil,
				map[string]string{
					"myuser.txt": "myuser",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Include string",
				`
				user: !file/yaml myuser.yaml
				`,
				nil,
				map[string]string{
					"myuser.yaml": "myuser",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Include boolean",
				`
				enabled: !file/yaml myenabled.yaml
				`,
				nil,
				map[string]string{
					"myenabled.yaml": "true",
				},
				Config{
					Enabled: true,
				},
			),
			Entry(
				"Include int",
				`
				id: !file/yaml myid.yaml
				`,
				nil,
				map[string]string{
					"myid.yaml": "123",
				},
				Config{
					ID: 123,
				},
			),
			Entry(
				"Include map",
				`
				!file/yaml mymap.yaml
				`,
				nil,
				map[string]string{
					"mymap.yaml": "{ user: myuser, id: 123 }",
				},
				Config{
					User: "myuser",
					ID:   123,
				},
			),
			Entry(
				"Include chain",
				`
				user: !file/yaml first.yaml
				`,
				nil,
				map[string]string{
					"first.yaml":  "!file/yaml second.yaml",
					"second.yaml": "myuser",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Include chain and variable",
				`
				user: !file/yaml first.yaml
				`,
				map[string]string{
					"MYUSER": "myuser",
				},
				map[string]string{
					"first.yaml":  "!file/yaml second.yaml",
					"second.yaml": "!variable MYUSER",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Include chain and file",
				`
				user: !file/yaml first.yaml
				`,
				nil,
				map[string]string{
					"first.yaml":  "!file/yaml second.yaml",
					"second.yaml": "!file myuser.txt",
					"myuser.txt":  "myuser",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Include chain and script",
				`
				user: !file/yaml first.yaml
				`,
				map[string]string{
					"MYUSER": "myuser",
				},
				map[string]string{
					"first.yaml":  "!file/yaml second.yaml",
					"second.yaml": "!script echo -n ${MYUSER}",
				},
				Config{
					User: "myuser",
				},
			),
			Entry(
				"Parse results of running script",
				`!script/yaml echo user: myuser`,
				nil,
				nil,
				Config{
					User: "myuser",
				},
			),
		)

		It("Fails if first tag isn't supported", func() {
			_, err := New().
				Load([]byte(`mykey: !wrong junk`)).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("wrong"))
		})

		It("Fails if second tag isn't supported", func() {
			_, err := New().
				Load([]byte(`mykey: !script/wrong echo -n junk`)).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("wrong"))
		})

		It("Fails if environment variable doesn't exist", func() {
			_, err := New().
				Load([]byte(`mykey: !variable DOESNOTEXIST`)).
				Build()
			Expect(err).To(HaveOccurred())
		})

		It("Reports location of error for known source name", func() {
			// Create a temporary file containing the configuration:
			tmp, err := ioutil.TempFile("", "*.test.yaml")
			Expect(err).ToNot(HaveOccurred())
			name := tmp.Name()
			defer func() {
				err = os.Remove(name)
				Expect(err).ToNot(HaveOccurred())
			}()
			_, err = tmp.Write([]byte(`mykey: !file /doesnotexist`))
			Expect(err).ToNot(HaveOccurred())
			err = tmp.Close()
			Expect(err).ToNot(HaveOccurred())

			// Check that loading fails and that the error message contains the name of
			// the source file:
			_, err = New().
				Load(name).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring(name + ":1:8"))
		})

		It("Reports location of error for included file", func() {
			// Create the a temporary file containing the text to be included:
			tmp, err := ioutil.TempFile("", "*.test.yaml")
			Expect(err).ToNot(HaveOccurred())
			name := tmp.Name()
			defer func() {
				err = os.Remove(name)
				Expect(err).ToNot(HaveOccurred())
			}()
			_, err = tmp.Write([]byte(`yourkey: !file /doesnotexist`))
			Expect(err).ToNot(HaveOccurred())
			err = tmp.Close()
			Expect(err).ToNot(HaveOccurred())

			// Check that loading fails and that the error message contains the name of
			// the included file:
			_, err = New().
				Load([]byte(`mykey: !file/yaml ` + name)).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring(name + ":1:10:"))
		})

		It("Reports location of error for unknown source name", func() {
			_, err := New().
				Load([]byte(`mykey: !file /doesnotexist.txt`)).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("unknown:1:8:"))
		})

		It("Fails if script writes to stderr", func() {
			_, err := New().
				Load([]byte(`mykey: !script echo myerror 1>&2`)).
				Build()
			Expect(err).To(HaveOccurred())
		})

		It("Script error contains stdout and stderr", func() {
			_, err := New().
				Load([]byte(
					`mykey: !script echo myoutput; echo myerror 1>&2; exit 1`,
				)).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("myoutput"))
			Expect(message).To(ContainSubstring("myerror"))
		})

		It("Fails if script command doesn't exist", func() {
			_, err := New().
				Load([]byte(
					`mykey: !script doesnotexist`,
				)).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("doesnotexist"))
		})

		It("Fails if finds sequence when expecting scalar", func() {
			_, err := New().
				Load([]byte(
					`mykey: !variable []`,
				)).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("variable"))
			Expect(message).To(ContainSubstring("scalar"))
			Expect(message).To(ContainSubstring("sequence"))
		})

		It("Fails if finds mapping when expecting scalar", func() {
			_, err := New().
				Load([]byte(
					`mykey: !file {}`,
				)).
				Build()
			Expect(err).To(HaveOccurred())
			message := err.Error()
			Expect(message).To(ContainSubstring("file"))
			Expect(message).To(ContainSubstring("scalar"))
			Expect(message).To(ContainSubstring("mapping"))
		})
	})
})
