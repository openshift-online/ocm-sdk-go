#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/init.sh"

METAMODEL="${1:-.}"
TARGET_DIR="${2:-.}"

# clean existing output
for dir in "${GENERATED_CLIENT_DIRS[@]}"
do
  rm -rf "${TARGET_DIR}/${dir}"
done

${METAMODEL} generate go \
  --model=vendor/github.com/openshift-online/ocm-api-model/model \
  --base=github.com/openshift-online/ocm-sdk-go \
  --apiBase=github.com/openshift-online/ocm-api-model/clientapi \
  --generators=builders-alias,clients,errors,helpers,json-alias,request-json,metrics,openapi,types-alias \
  --output="${TARGET_DIR}"