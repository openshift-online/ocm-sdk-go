#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

ARTIFACT_DIR=${ARTIFACT_DIR:-$(mktemp -d)}
TEMP_DIR="${ARTIFACT_DIR}"/gofmt

mkdir "${TEMP_DIR}"

DIFF_FILE="${TEMP_DIR}"/gofmt.diff
gofmt -d -s -l $(find . -maxdepth 1 -type d  ! -name 'vendor' ! -name '.') > "${DIFF_FILE}"

if [ -s "${DIFF_FILE}" ]; then
  cat "${DIFF_FILE}"
  exit 1
else
  exit 0
fi