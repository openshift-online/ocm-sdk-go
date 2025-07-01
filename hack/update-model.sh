#!/bin/bash -e

# This script ensures that all the api-model submodules used across the project are updated to the latest verison.

# List of go.mod directories
MODULE_DIRS=(
  "./metamodel_generator"
  "./examples"
  "."
)

# The modules to update
MODULES=(
  "github.com/openshift-online/ocm-api-model/model"
  "github.com/openshift-online/ocm-api-model/clientapi"
)

echo "Updating Go modules..."

for dir in "${MODULE_DIRS[@]}"; do
  echo "Updating in directory: $dir"
  pushd "$dir" > /dev/null || {
    echo "Failed to enter $dir"
    continue
  }

  for mod in "${MODULES[@]}"; do
    echo "  - Updating $mod to latest"
    go get -u "$mod"
  done

  echo "  - Tidying up module"
  go mod tidy

  popd > /dev/null
done

echo "Done."
