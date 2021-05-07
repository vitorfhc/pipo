package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
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
	}

	logrus.Debug(loginData)
}
