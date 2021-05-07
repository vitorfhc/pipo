package gql

import (
	"context"

	"github.com/sirupsen/logrus"
)

type EmployerObject struct {
	CNPJ     string
	UUID     string
	Password string
}

type employerResponse struct {
	Employer []EmployerObject
}

func GetEmployer(cnpj string) EmployerObject {
	req := newAuthenticatedRequest(getEmployerQuery)
	req.Var("cnpj", cnpj)
	var res employerResponse

	ctx := context.Background()
	err := GQLClient.Run(ctx, req, &res)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"query_name":   "GetEmployer",
			"query_string": getEmployerQuery,
		}).Error("Error making query: ", err)
	}

	if len(res.Employer) != 0 {
		return res.Employer[0]
	}

	return EmployerObject{}
}
