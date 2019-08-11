#
# Copyright (c) 2018 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Directory containing the model:
model:=/files/projects/ocm-api-model/model

.PHONY: examples
examples:
	cd examples && \
	for i in *.go; do \
		go build -mod=readonly $${i} || exit 1; \
	done

.PHONY: test
test:
	ginkgo -r pkg \
		$(NULL)

.PHONY: fmt
fmt:
	gofmt -s -l -w \
		pkg \
		examples \
		$(NULL)

.PHONY: lint
lint:
	golangci-lint run \
		--no-config \
		--issues-exit-code=1 \
		--deadline=15m \
		--skip-dirs=pkg/client/accountsmgmt \
		--skip-dirs=pkg/client/clustersmgmt \
		--skip-dirs=pkg/client/errors \
		--skip-dirs=pkg/client/helpers \
		--disable-all \
		--enable=deadcode \
		--enable=gas \
		--enable=goconst \
		--enable=gofmt \
		--enable=golint \
		--enable=ineffassign \
		--enable=interfacer \
		--enable=lll \
		--enable=megacheck \
		--enable=misspell \
		--enable=structcheck \
		--enable=unconvert \
		--enable=varcheck \
		$(NULL)

.PHONY: generate
generate:
	rm -rf \
		pkg/client/accountsmgmt \
		pkg/client/clustersmgmt \
		pkg/client/errors \
		pkg/client/helpers
	ocm-metamodel-tool generate \
		--model=$(model) \
		--base=github.com/openshift-online/uhc-sdk-go/pkg/client \
		--output=pkg/client
