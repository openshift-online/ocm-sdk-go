//go:build !windows
// +build !windows

/*
Copyright (c) 2024 Red Hat, Inc.

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

package securestore

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2" // nolint
	. "github.com/onsi/gomega"    // nolint

	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
)

// This test requires `pass` to be installed.
// macOS: `brew install pass`
// linux: `sudo apt-get install pass` or `sudo yum install pass`

const keyring_dir = "keyring-pass-test-*"

func runCmd(cmds ...string) {
	cmd := exec.Command(cmds[0], cmds[1:]...) //nolint:gosec
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(cmd)
		fmt.Println(string(out))
		Fail(err.Error())
	}
}

var _ = Describe("Pass Keyring", Ordered, func() {
	const backend = "pass"

	BeforeAll(func() {
		pwd, err := os.Getwd()
		if err != nil {
			Fail(err.Error())
		}
		pwdParent := filepath.Dir(pwd)

		// the default temp directory can't be used because gpg-agent complains with "socket name too long"
		tmpdir, err := os.MkdirTemp("/tmp", keyring_dir)
		if err != nil {
			Fail(err.Error())

		}
		tmpdirPass, err := os.MkdirTemp("/tmp", ".password-store-*")
		if err != nil {
			Fail(err.Error())
		}

		// Initialise a blank GPG homedir; import & trust the test key
		gnupghome := filepath.Join(tmpdir, ".gnupg")
		err = os.Mkdir(gnupghome, os.FileMode(int(0700)))
		if err != nil {
			Fail(err.Error())
		}
		os.Setenv("GNUPGHOME", gnupghome)
		os.Setenv("PASSWORD_STORE_DIR", tmpdirPass)
		os.Unsetenv("GPG_AGENT_INFO")
		os.Unsetenv("GPG_TTY")

		runCmd("gpg", "--batch", "--import", filepath.Join(pwdParent, "testdata", "test-gpg.key"))
		runCmd("gpg", "--batch", "--import-ownertrust", filepath.Join(pwdParent, "testdata", "test-ownertrust-gpg.txt"))
		runCmd("pass", "init", "ocm-devel@redhat.com")

		DeferCleanup(func() {
			os.Unsetenv("GNUPGHOME")
			os.Unsetenv("PASSWORD_STORE_DIR")
			os.RemoveAll(filepath.Join("/tmp", keyring_dir))
		})
	})

	BeforeEach(func() {
		err := RemoveConfigFromKeyring(backend)
		Expect(err).To(BeNil())
	})

	When("Listing Keyrings", func() {
		It("Lists pass as a valid keyring", func() {
			backends := AvailableBackends()
			Expect(backends).To(ContainElement(backend))
		})
	})

	When("Using Pass", func() {
		It("Stores/Removes configuration in Pass", func() {
			// Create the token
			accessToken := MakeTokenString("Bearer", 15*time.Minute)

			// Run insert
			err := UpsertConfigToKeyring(backend, []byte(accessToken))

			Expect(err).To(BeNil())

			// Check the content of the keyring
			result, err := GetConfigFromKeyring(backend)
			Expect(result).To(Equal([]byte(accessToken)))
			Expect(err).To(BeNil())

			// Remove the configuration from the keyring
			err = RemoveConfigFromKeyring(backend)
			Expect(err).To(BeNil())

			// Ensure the keyring is empty
			result, err = GetConfigFromKeyring(backend)
			Expect(result).To(Equal([]byte("")))
			Expect(err).To(BeNil())
		})
	})
})
