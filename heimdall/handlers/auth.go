package handlers

import (
	"net/http"
)

func AuthHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("The authentication must be done through a POST"))
		return
	}
}
