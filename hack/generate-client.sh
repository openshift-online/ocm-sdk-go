#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/init.sh"

METAMODEL="${1:-metamodel_generator/metamodel}"
TARGET_DIR="${2:-.}"

# clean existing output
$(dirname "${BASH_SOURCE}")/clean-client.sh "${TARGET_DIR}"

${METAMODEL} generate go \
  --model=vendor/github.com/openshift-online/ocm-api-model/model \
  --base=github.com/openshift-online/ocm-sdk-go \
  --apiBase=github.com/openshift-online/ocm-api-model/clientapi \
  --generators=builders-alias,clients,errors,helpers,json-alias,request-json,metrics,openapi,types-alias \
  --output="${TARGET_DIR}"