package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

const port = "8080" // TODO: Replace with env variable
const endpoint = ":" + port

// New creates a new HTTP server
func New() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/v1/api/health", healthCheckHandler) // TODO: Not hard code API version
	mux.HandleFunc("/v1/api/validate/csv", validateCSVHandler)

	svr := &http.Server{
		Addr:    endpoint,
		Handler: mux,
	}

	slog.Info(fmt.Sprintf("Server listening on port %s...", port))
	// TODO: Certs can probably go into an env variable
	// TODO: Use HTTPS in prod?
	// Certs are handled on Kubernetes-level
	// log.Error("server.New(): %s", "error", svr.ListenAndServeTLS("server.crt", "server.key"))
	slog.Error("server.New():", "error", svr.ListenAndServe())
}

// defaultHandler is the default handler that writes 404 HTTP status to response header
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	slog.Warn("Connected to default handler. Upstream error?", "method", r.Method, "protocol", r.Proto)
}

// validateCSVHandler processes a URL to CSV file in payload and validates it
func validateCSVHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("Connected to %s", endpoint), "method", r.Method, "protocol", r.Proto)
	_, _ = w.Write([]byte("Hello, World!"))
	// TODO: Call CSV validation business logic
}

// healthCheckHandler writes 200 HTTP status to response header
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	slog.Info("Connected to health check handler.", "method", r.Method, "protocol", r.Proto)
}
