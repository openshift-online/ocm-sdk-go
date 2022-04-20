package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/openshift-online/ocm-sdk-go/logging"
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

	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of users:
	collection := connection.ClustersMgmt().V1().
		Clusters().
		Cluster("123").
		IdentityProviders().
		IdentityProvider("456").
		HtpasswdUsers()

	// Read the list of users:
	users := []*cmv1.HTPasswdUser{}
	reader := csv.NewReader(strings.NewReader(usersCSV))
	for {
		var record []string
		record, err = reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't read record: %v\n", err)
			os.Exit(1)
		}
		user, err := cmv1.NewHTPasswdUser().
			Username(record[0]).
			Password(record[1]).
			Build()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't build user: %v\n", err)
			os.Exit(1)
		}
		users = append(users, user)
	}

	// Send the request to import the users:
	_, err = collection.Import().
		Items(users).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't import users: %v\n", err)
		os.Exit(1)
	}
}

const usersCSV = `
joe,joe123
mary,mary123
`
