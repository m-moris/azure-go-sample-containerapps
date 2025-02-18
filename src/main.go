package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request on /hello endpoint")
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
