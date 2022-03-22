/*
Copyright (c) 2018 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file is a main program for the examples. It takes the name of the example from the first
// command line parameter and then runs that example passing the rest of the command line
// parameters.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

// Global logger
var logger logr.Logger

// Examples indexed by name:
var examples = map[string]func(ctx context.Context, args []string) error{
	"client_credentials_grant":     clientCredentialsGrant,
	"create_cluster":               createCluster,
	"create_product":               createProduct,
	"create_syncset":               createSyncset,
	"delete_cluster":               deleteCluster,
	"delete_product":               deleteProduct,
	"delete_subscription":          deleteSubscription,
	"dump_config":                  dumpConfig,
	"existing_token":               existingToken,
	"export_control_review":        exportControlReview,
	"get_cluster_credentials":      getClusterCredentials,
	"get_cluster":                  getCluster,
	"get_cluster_logs":             getClusterLogs,
	"get_metadata":                 getMetadata,
	"get_service":                  getService,
	"list_applications":            listApplications,
	"list_cloud_providers":         listCloudProviders,
	"list_cluster_creators":        listClusterCreators,
	"list_clusters":                listClusters,
	"list_products":                listProducts,
	"list_quota_cost":              listQuotaCost,
	"list_status_updates":          listStatusUpdates,
	"list_versions":                listVersions,
	"load_config":                  loadConfig,
	"prometheus_metrics":           prometheusMetrics,
	"pushpop_job_queue":            pushpopJobQueue,
	"resource_review":              resourceReview,
	"run_cluster_operator_metrics": runClusterOperatorMetrics,
	"run_metric_query":             runMetricQuery,
	"sync_addons":                  syncAddons,
	"transport_wrapper":            transportWrapper,
	"update_cluster":               updateCluster,
	"update_product":               updateProduct,
}

func main() {
	// Create the context:
	ctx := context.Background()

	// Create the logger:
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.OutputPaths = []string{
		"stdout",
	}
	zapper, err := config.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create logger: %v\n", err)
		os.Exit(1)
	}
	logger = zapr.NewLogger(zapper)

	// Get the name of the example:
	argc := len(os.Args)
	if argc < 2 {
		logger.Error(
			nil,
			"First command line parameter must be the example name",
			"argc", argc,
		)
		os.Exit(1)
	}
	name := os.Args[1]
	example, ok := examples[name]
	if !ok {
		logger.Error(
			nil,
			"Failed to find example",
			"name", name,
		)
		os.Exit(1)
	}

	// Run the example:
	err = example(ctx, os.Args[2:])
	if err != nil {
		logger.Error(
			err,
			"Example failed",
			"name", name,
		)
		os.Exit(1)
	}
}
