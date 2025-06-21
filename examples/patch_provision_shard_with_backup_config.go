package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/openshift-online/ocm-sdk-go/logging"

	v1 "github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := logging.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	url := "http://localhost:9000" // local proxy URL

	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		URL(url).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	managementClusterName := "" // update this with the name of the management cluster
	search := fmt.Sprintf("management_cluster = '%s'", managementClusterName)
	response, err := connection.ClustersMgmt().V1().ProvisionShards().List().Search(search).Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve the list of provision shards: %s\n", err)
		os.Exit(1)
	}

	if response.Size() < 1 {
		fmt.Fprintf(os.Stderr, "Can't find provision shard for management cluster %s", managementClusterName)
		os.Exit(1)
	}

	var currentShard cmv1.ProvisionShard
	response.Items().Each(func(item *cmv1.ProvisionShard) bool {
		currentShard = *item
		return true
	})

	backupConfigBuilder1 := cmv1.NewAWSBackupConfig().
		AccountID("account1").
		RoleArn("arn:aws:iam::123456789011:role/RH-Installer").
		IdentityProviderArn("arn:aws:iam::123456789012:oidc-provider/rh-oidc.s3.us-east-1.amazonaws.com/123").
		S3Bucket("bucket1")

	backupConfigBuilder2 := cmv1.NewAWSBackupConfig().
		AccountID("account2").
		RoleArn("arn:aws:iam::123456789012:role/RH-Installer").
		IdentityProviderArn("arn:aws:iam::123456789012:oidc-provider/rh-oidc.s3.us-east-1.amazonaws.com/456").
		S3Bucket("bucket2")

	backupConfigs := map[string]*cmv1.AWSBackupConfigBuilder{"Management Cluster 1": backupConfigBuilder1, "Management Cluster 2": backupConfigBuilder2}

	newShard, err := v1.NewProvisionShard().
		ID(currentShard.ID()).
		HypershiftConfig(
			v1.NewServerConfig().
				Topology(currentShard.HypershiftConfig().Topology()).
				AWSShard(v1.NewAWSShard().
					ECRRepositoryURLs(currentShard.HypershiftConfig().AWSShard().ECRRepositoryURLs()...).
					BackupConfigs(backupConfigs),
				),
		).Build()

	patchResponse, err := connection.ClustersMgmt().V1().ProvisionShards().ProvisionShard(currentShard.ID()).Update().Body(newShard).SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't update the provision shard: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Updated backup configs of provision shard with ID: %s\n", currentShard.ID())

	for mc, bkpCfg := range patchResponse.Body().HypershiftConfig().AWSShard().BackupConfigs() {
		fmt.Printf("Management Cluster: %s\n", mc)
		fmt.Printf("Backup Configs: %+v\n", bkpCfg)
	}
}
