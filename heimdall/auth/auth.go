package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/vitorfhc/heimdall/gql"
)

var pKey *rsa.PrivateKey

func init() {
	privateKeyPath, ok := os.LookupEnv("HEIMDALL_PRIVATE_KEY")
	if !ok {
		logrus.Fatal("Environemnt variable HEIMDALL_PRIVATE_KEY is not set")
	}

	privateKeyPath, err := filepath.Abs(privateKeyPath)
	if err != nil {
		logrus.Fatal("Error private key file path: ", err)
	}

	content, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		logrus.Fatal("Error opening private key file: ", err)
	}

	pKey, err = jwt.ParseRSAPrivateKeyFromPEM(content)
	if err != nil {
		logrus.Fatal("Error parsing private key: ", err)
	}
}

func GenerateToken(employer gql.EmployerObject) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"hasura": jwt.MapClaims{
			"x-hasura-user-id":       employer.UUID,
			"x-hasura-default-role":  "employer",
			"x-hasura-allowed-roles": []string{"employer"},
		},
	})

	signed, err := token.SignedString(pKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}
