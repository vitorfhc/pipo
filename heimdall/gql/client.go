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

func newAuthenticatedRequest(query string) *graphql.Request {
	adminSecret, ok := os.LookupEnv("HEIMDALL_HASURA_ADMIN_SECRET")

	if !ok {
		return graphql.NewRequest(query)
	}

	req := graphql.NewRequest(query)
	req.Header.Set("x-hasura-admin-secret", adminSecret)
	return req
}
