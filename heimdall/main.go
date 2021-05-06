package main

import (
	"fmt"
	"net/http"

	"github.com/vitorfhc/heimdall/handlers"
)

func main() {
	http.HandleFunc("/auth", handlers.AuthHandler)

	fmt.Println("Starting server...")
	http.ListenAndServe(":9000", nil)
}
