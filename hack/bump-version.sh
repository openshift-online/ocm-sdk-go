#!/bin/bash
# Script to bump the patch version and suggest git commands for tagging

set -euo pipefail

# Repository URLs
OFFICIAL_REPO_SSH="git@github.com:openshift-online/ocm-sdk-go.git"
OFFICIAL_REPO_HTTPS="https://github.com/openshift-online/ocm-sdk-go.git"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}INFO:${NC} $1"
}

print_success() {
    echo -e "${GREEN}SUCCESS:${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}WARNING:${NC} $1"
}

print_error() {
    echo -e "${RED}ERROR:${NC} $1"
}

# Check if we're in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    print_error "Not in a git repository"
    exit 1
fi

# Check if we're on the main branch
CURRENT_BRANCH=$(git branch --show-current 2>/dev/null || echo "")
if [ "$CURRENT_BRANCH" != "main" ]; then
    print_error "You must be on the 'main' branch to create a release tag"
    print_error "Current branch: $CURRENT_BRANCH"
    print_error "Switch to main branch with: git checkout main"
    exit 1
fi

# Get the latest tag
print_info "Finding the latest release tag..."
LATEST_TAG=$(git tag --sort=-version:refname | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | head -1)

if [ -z "$LATEST_TAG" ]; then
    print_error "No semantic version tags found (expected format: vX.Y.Z)"
    exit 1
fi

print_success "Latest tag found: $LATEST_TAG"

# Parse the version components
if [[ $LATEST_TAG =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
    MAJOR=${BASH_REMATCH[1]}
    MINOR=${BASH_REMATCH[2]}
    PATCH=${BASH_REMATCH[3]}
else
    print_error "Invalid tag format: $LATEST_TAG (expected format: vX.Y.Z)"
    exit 1
fi

# Bump patch version
NEW_PATCH=$((PATCH + 1))
NEW_TAG="v${MAJOR}.${MINOR}.${NEW_PATCH}"

print_info "Version bump analysis:"
echo "  Current version: $LATEST_TAG"
echo "  New version:     $NEW_TAG"
echo "  Change:          Patch version bumped from $PATCH to $NEW_PATCH"

# Get git remote information
# Find the remote that points to the official openshift-online repository
REMOTE_NAME=""
REMOTE_URL=""

# Look for the official repository (SSH or HTTPS) with (push)
while read -r line; do
    if [[ "$line" == *"$OFFICIAL_REPO_SSH"*"(push)" ]] || \
       [[ "$line" == *"$OFFICIAL_REPO_HTTPS"*"(push)" ]]; then
        REMOTE_NAME=$(echo "$line" | awk '{print $1}')
        REMOTE_URL=$(echo "$line" | awk '{print $2}')
        break
    fi
done < <(git remote -v)

# If we didn't find the official remote, error out
if [ -z "$REMOTE_NAME" ]; then
    print_error "Could not find a remote pointing to the official repository"
    print_error "Looking for: $OFFICIAL_REPO_SSH or $OFFICIAL_REPO_HTTPS"
    print_error "Available remotes:"
    git remote -v | while read line; do echo "  $line"; done
    print_error "Please add the official repository as a remote, for example:"
    print_error "  git remote add upstream $OFFICIAL_REPO_SSH"
    exit 1
fi

# Current branch was already determined above during branch check

echo ""
print_info "Git repository information:"
echo "  Remote:  $REMOTE_NAME ($REMOTE_URL)"
echo "  Branch:  $CURRENT_BRANCH"

# Check if tag already exists
if git tag | grep -q "^${NEW_TAG}$"; then
    print_error "Tag $NEW_TAG already exists!"
    exit 1
fi

# Check for uncommitted changes
if ! git diff --quiet --exit-code || ! git diff --cached --quiet --exit-code; then
    print_warning "There are uncommitted changes in the repository"
    echo "  Consider committing or stashing changes before creating a tag"
fi

# Display suggested commands
echo ""
print_success "Suggested commands to create and push the new tag:"
echo ""
echo -e "${GREEN}# Create the new tag:${NC}"
echo "git tag -a -m 'Release ${NEW_TAG#v}' $NEW_TAG"
echo ""
echo -e "${GREEN}# Push the tag to remote:${NC}"
echo "git push $REMOTE_NAME $NEW_TAG"
echo ""

