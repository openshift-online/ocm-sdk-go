#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/init.sh"

TARGET_DIR="${1:-.}"

# clean existing output
rm -rf "${TARGET_DIR}"