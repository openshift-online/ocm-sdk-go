# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Building and Testing
- `make examples` - Build all example programs (default target)
- `make test` or `make tests` - Run all tests using Ginkgo
- `make lint` - Run golangci-lint for code quality checks
- `make verify` - Run comprehensive verification (lint, examples, model generation, format checks)

### Code Generation and Model Updates
- `make update` or `make generate` - Generate client code from API model
- `./hack/update-model.sh [version]` - Update ocm-api-model dependency to latest or specific version
- `make model` - Run go mod tidy and vendor
- `make metamodel-install` - Build the metamodel generator tool

### Development Tools
- `make fmt` - Format Go code using gofmt
- `make install-hooks` - Install git pre-push hooks
- `make clean` - Clean generated files and build artifacts

### Individual Test Execution
Tests use Ginkgo framework. Run specific tests with:
- `./bin/ginkgo run -r` - Run all tests recursively
- `./bin/ginkgo run ./path/to/test` - Run tests in specific directory

## Project Architecture

### Core Structure
This is the **OCM (OpenShift Cluster Manager) SDK for Go**, a client library for the OCM API at `api.openshift.com`. The SDK uses a **code generation architecture** where client types and API bindings are automatically generated from an API model.

### Key Components

**Connection Management**
- `connection.go` - Core `Connection` type that manages HTTP connections, authentication, and access to service clients
- Connection acts as the entry point to all OCM services and manages expensive resources like HTTP connection pools and auth tokens

**Service Packages** (Auto-generated from ocm-api-model)
- `accountsmgmt/v1` - Account management service client
- `clustersmgmt/v1` - Cluster management service client
- `authorizations/v1` - Authorization service client
- `addonsmgmt/v1` - Add-on management service client
- `accesstransparency/v1` - Access transparency service client
- Plus additional services: `arohcp`, `jobqueue`, `osdfleetmgmt`, `servicelogs`, `servicemgmt`, `statusboard`, `webrca`

**Core Infrastructure**
- `authentication/` - OAuth2 and token-based authentication
- `logging/` - SDK logging framework
- `metrics/` - Prometheus metrics integration
- `retry/` - Request retry logic with backoff
- `errors/` - OCM-specific error types
- `configuration/` - Configuration management utilities
- `internal/` - Private implementation details

**Code Generation**
- `metamodel_generator/` - Tool for generating client code from API model
- `hack/` - Scripts for generation, verification, and updates
- All service packages under `*/v1/` are **generated code** - do not edit manually

### Development Workflow

**Model Updates**: Changes to the API require updating the `ocm-api-model` dependency first, then regenerating client code with `make update`. The metamodel generator reads API specifications and generates Go types, builders, and client interfaces.

**Versioning**: Service APIs use semantic versioning (v1, v2, etc.) with separate packages for each version to maintain backwards compatibility.

**Generated vs Manual Code**:
- Generated: All `*/v1/` service packages, OpenAPI specs
- Manual: Core SDK (connection, auth, logging, metrics), examples, tests

**Testing**: Uses Ginkgo/Gomega testing framework. Tests cover both unit functionality and integration scenarios against the OCM API.

The package name is `sdk` despite the repository being `ocm-sdk-go`. Import as `sdk "github.com/openshift-online/ocm-sdk-go"`.