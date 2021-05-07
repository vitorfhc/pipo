package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/vitorfhc/heimdall/auth"
	"github.com/vitorfhc/heimdall/gql"
)

type loginJSON struct {
	Username string
	Password string
}

func AuthHandler(w http.ResponseWriter, req *http.Request) {
	logrus.WithFields(logrus.Fields{
		"handler": "AuthHandler",
		"method":  req.Method,
		"uri":     req.RequestURI,
	}).Info("Request received")

	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("The authentication must be done through a POST"))
		if err != nil {
			logrus.WithField("handler", "AuthHandler").Error("Error writing response: ", err)
		}
		return
	}

	jsonDecoder := json.NewDecoder(req.Body)
	loginData := loginJSON{}
	err := jsonDecoder.Decode(&loginData)
	if err != nil {
		logrus.WithField("handler", "AuthHandler").Error("Error trying to decode JSON from request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	employer := gql.GetEmployer(loginData.Username)
	hashBytes := md5.Sum([]byte(loginData.Password))
	hash := hex.EncodeToString(hashBytes[:])

	if employer.Password != hash {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := auth.GenerateToken(employer)
	if err != nil {
		logrus.WithField("handler", "AuthHandler").Error("Error trying to sign token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "pipo-token",
		Value:    tokenString,
		HttpOnly: true,
		// Uncomment below on production
		// Secure: true,
		// Domain: "domain.com",
	})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
