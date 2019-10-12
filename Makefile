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

# Enable Go modules:
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

# Details of the model to use:
model_version:=HEAD
model_url:=/files/projects/ocm-api-model

# Details of the metamodel to use:
metamodel_version:=v0.0.8
metamodel_url:=https://github.com/openshift-online/ocm-api-metamodel.git

.PHONY: examples
examples:
	cd examples && \
	for i in *.go; do \
		go build -mod=readonly $${i} || exit 1; \
	done

.PHONY: test
test:
	ginkgo -r .

.PHONY: fmt
fmt:
	gofmt -s -l -w .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: generate
generate: model metamodel
	rm -rf \
		accountsmgmt \
		authorizations \
		clustersmgmt \
		errors \
		helpers
	metamodel/ocm-metamodel-tool \
		generate \
		--model=model/model \
		--base=github.com/openshift-online/ocm-sdk-go \
		--output=.

.PHONY: model
model:
	[ -d "$@" ] || git clone "$(model_url)" "$@"
	cd "$@" && git fetch --tags origin
	cd "$@" && git checkout -B build "$(model_version)"

.PHONY: metamodel
metamodel:
	[ -d "$@" ] || git clone "$(metamodel_url)" "$@"
	cd "$@" && git fetch --tags origin
	cd "$@" && git checkout -B build "$(metamodel_version)"
	make -C "$@"

.PHONY: clean
clean:
	rm -rf \
		.gobin \
		metamodel \
		model \
		$(NULL)
