package main

import (
	"net/http"

	"golang.org/x/exp/slog"
)

// this is a simple web server
func main() {
	logger := slog.Default()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	logger.Info("Starting server on port 8080")
	http.ListenAndServe(":8080", mux)
}
