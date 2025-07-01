#!/bin/bash -e

# This script ensures that all the api-model submodules used across the project are all the same version.

# Group 1: Shared modules to verify across root and examples
COMMON_MODULES=(
  "github.com/openshift-online/ocm-api-model/model"
  "github.com/openshift-online/ocm-api-model/clientapi"
)

COMMON_DIRS=(
  "."
  "examples"
)

# Group 2: metamodel_generator module to only be checked in metamodel_generator
METAMODEL_MODULE="github.com/openshift-online/ocm-api-model/metamodel_generator"
METAMODEL_DIR="metamodel_generator"

# Function to robustly extract module version from go.mod
extract_version() {
  local module=$1
  local mod_file=$2

  awk -v mod="$module" '
    $1 == "require" && $2 == mod { print $3; exit }
    $1 == mod { print $2; exit }
  ' "$mod_file"
}

echo "Verifying shared module versions across ./ and ./examples..."

status=0

# Step 1: Check shared modules in . and examples
for module in "${COMMON_MODULES[@]}"; do
  echo "Checking module: $module"
  ref_version=""
  consistent=true

  for dir in "${COMMON_DIRS[@]}"; do
    mod_file="$dir/go.mod"
    version=$(extract_version "$module" "$mod_file")

    if [ -z "$version" ]; then
      echo "  ❌ $module not found in $mod_file"
      consistent=false
      status=1
      continue
    fi

    echo "  ✔ $dir/go.mod: $version"

    if [ -z "$ref_version" ]; then
      ref_version="$version"
    elif [ "$version" != "$ref_version" ]; then
      echo "  ❌ Version mismatch in $mod_file: $version (expected: $ref_version)"
      consistent=false
      status=1
    fi
  done

  if $consistent; then
    echo "✅ All versions match for $module: $ref_version"
  else
    echo "❌ Version mismatch found for $module"
  fi

  echo
done

# Step 2: Check metamodel_generator module in metamodel_generator/go.mod
echo "Checking metamodel module in $METAMODEL_DIR: $METAMODEL_MODULE"

mod_file="$METAMODEL_DIR/go.mod"
metamodel_version=$(extract_version "$METAMODEL_MODULE" "$mod_file")

if [ -z "$metamodel_version" ]; then
  echo "❌ $METAMODEL_MODULE not found in $mod_file"
  status=1
elif [ "$metamodel_version" != "$ref_version" ]; then
  echo "  ❌ Version mismatch in $mod_file: $metamodel_version (expected: $ref_version)"
  consistent=false
  status=1
fi

if $consistent; then
  echo "✅ All versions match for $METAMODEL_MODULE: $ref_version"
else
  echo "❌ Version mismatch found for $METAMODEL_MODULE"
fi

exit $status
