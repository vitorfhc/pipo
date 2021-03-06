package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/vitorfhc/heimdall/handlers"
	_ "github.com/vitorfhc/heimdall/logger"
)

func main() {
	http.HandleFunc("/auth", handlers.AuthHandler)

	addr, exists := os.LookupEnv("HEIMDALL_ADDR")
	if !exists {
		addr = "0.0.0.0:9000"
	}

	logrus.Info("Starting server at ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logrus.Error("Error starting server: ", err)
	}
}
