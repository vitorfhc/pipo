package gql

import (
	"os"

	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

var GQLClient *graphql.Client

func init() {
	endpoint, exists := os.LookupEnv("HEIMDALL_GRAPHQL")
	if !exists {
		endpoint = "http://localhost:8080/v1/graphql"
	}

	GQLClient = graphql.NewClient(endpoint)
	GQLClient.Log = func(s string) {
		logrus.WithField("package", "graphql").Debug(s)
	}
}
