package gql

import (
	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

var GQLClient *graphql.Client

func init() {
	endpoint := "http://localhost:8080/v1/graphql"
	GQLClient = graphql.NewClient(endpoint)
	GQLClient.Log = func(s string) {
		logrus.WithField("package", "graphql").Debug(s)
	}
}
