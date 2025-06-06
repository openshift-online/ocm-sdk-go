#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/init.sh"

TARGET_DIR="${1:-.}"

# clean existing output
for dir in "${GENERATED_CLIENT_DIRS[@]}"
do
  rm -rf "${TARGET_DIR}/${dir}"
done