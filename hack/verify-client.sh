#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/init.sh"

METAMODEL="${1:-.}"
TARGET_DIR="${2:-.}"

VERIFY_GEN_TARGET_DIR=$(mktemp -d -t verify-client-XXXX)
"$(dirname "${BASH_SOURCE}")/generate-client.sh" "${METAMODEL}" "${VERIFY_GEN_TARGET_DIR}"

echo "checking tmp content in ${VERIFY_GEN_TARGET_DIR}"
for dir in "${GENERATED_CLIENT_DIRS[@]}"
do
  diff -r "${VERIFY_GEN_TARGET_DIR}/${dir}" "${TARGET_DIR}/${dir}"
done

# clean temporary output
rm -rf "${VERIFY_GEN_TARGET_DIR}"
