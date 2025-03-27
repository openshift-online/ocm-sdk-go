#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/init.sh"

METAMODEL="${1:-.}"
TARGET_DIR="${2:-./openapi}"

VERIFY_GEN_TARGET_DIR=$(mktemp -d -t verify-openapi-XXXX)
"$(dirname "${BASH_SOURCE}")/generate-openapi.sh" "${METAMODEL}" "${VERIFY_GEN_TARGET_DIR}"

echo "checking tmp content in ${VERIFY_GEN_TARGET_DIR}"
diff -r "${VERIFY_GEN_TARGET_DIR}" "${TARGET_DIR}"

# clean temporary output
rm -rf "${VERIFY_GEN_TARGET_DIR}"
