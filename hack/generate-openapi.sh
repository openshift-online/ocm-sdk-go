#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/init.sh"

METAMODEL="${1:-metamodel_generator/metamodel}"
TARGET_DIR="${2:-./openapi}"

# clean existing output
$(dirname "${BASH_SOURCE}")/clean-openapi.sh "${TARGET_DIR}"

${METAMODEL} generate openapi \
  --model=vendor/github.com/openshift-online/ocm-api-model/model \
  --output="${TARGET_DIR}"