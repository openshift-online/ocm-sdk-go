#!/bin/bash -e

# ==============================================================================
# Script: update-model.sh
#
# This script ensures that all the `ocm-api-model` submodules used across the
# project are updated to the latest version or pinned to a specific commit SHA.
#
# The OCM SDK can be generated simply by running the following after all changes have been made:
#
# ./hack/update-model.sh
# make update
#
# USAGE:
#   ./update-model.sh [COMMIT_SHA]
#
# ARGUMENTS:
#   COMMIT_SHA (optional) - If provided, all listed modules will be pinned to
#                           this specific commit instead of updating to latest.
#
# EXAMPLES:
#   ./update-model.sh
#     → Updates all modules to their latest versions.
#
#   ./update-model.sh f67fb59980981bdc81d95d1379a82da5bcec57bf
#     → Pins all modules to the specified commit SHA.
#
# ==============================================================================

# Optional commit SHA
COMMIT_SHA="$1"

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
  "github.com/openshift-online/ocm-api-model/metamodel_generator"
)

echo "Updating Go modules..."
if [ -n "$COMMIT_SHA" ]; then
  echo "Using specific commit SHA: $COMMIT_SHA"
fi

for dir in "${MODULE_DIRS[@]}"; do
  echo "Updating in directory: $dir"
  pushd "$dir" > /dev/null || {
    echo "Failed to enter $dir"
    continue
  }

  for mod in "${MODULES[@]}"; do
    if [ -n "$COMMIT_SHA" ]; then
      echo "  - Pinning $mod to $COMMIT_SHA"
      go get "${mod}@${COMMIT_SHA}"
    else
      echo "  - Updating $mod to latest"
      go get -u "$mod"
    fi
  done

  echo "  - Tidying up module"
  go mod tidy

  popd > /dev/null
done

echo "Done."
