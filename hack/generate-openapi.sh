#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/init.sh"

METAMODEL="${1:-.}"
TARGET_DIR="${2:-./openapi}"

# clean existing output
rm -rf "${TARGET_DIR}"

${METAMODEL} generate openapi \
  --model=vendor/github.com/openshift-online/ocm-api-model/model \
  --output="${TARGET_DIR}"