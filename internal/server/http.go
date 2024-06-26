package server

import (
	"encoding/json"
	"fmt"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"log/slog"
	"net/http"
)

const port = "8080" // TODO: Replace with env variable
const endpoint = ":" + port

// TODO: Do we need to authenticate sender?

// New creates a new HTTP server that serves as REST API
func New() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/v1/api/health", healthCheckHandler) // TODO: Not hard code API version
	mux.HandleFunc("/v1/api/validate/csv", validateCSVHandler)

	run := func() error {
		svr := &http.Server{
			Addr:    endpoint,
			Handler: mux,
		}

		slog.Info(fmt.Sprintf("Server listening on port %s...", port))
		// TODO: Certs can probably go into an env variable
		// TODO: Use HTTPS in prod?
		// Certs are handled on Kubernetes-level
		// log.Error("server.New(): %s", "error", svr.ListenAndServeTLS("server.crt", "server.key"))
		err := svr.ListenAndServe()
		slog.Error("server.New():", "error", err)
		return err
	}

	return run()
}

// defaultHandler is the default handler that writes 404 HTTP status to response header
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	slog.Warn("Connected to default handler. Upstream error?", "method", r.Method, "protocol", r.Proto)
	w.WriteHeader(http.StatusNotFound)
}

// validateCSVHandler processes a URL to CSV file in payload and validates it
func validateCSVHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("Connected to %s", endpoint), "method", r.Method, "protocol", r.Proto)

	switch r.Method {

	case http.MethodPost:
		if !hasValidContentType(r.Header) {
			slog.Warn("Sender sent a non-CSV or non-TSV file")
			http.Error(w, "Invalid Content-Type for validation", http.StatusBadRequest)
			break
		}

		if hasEmptyBody(r) {
			slog.Warn("Sender sent request with empty body")
			http.Error(w, "Expect a non-empty body", http.StatusBadRequest)
			break
		}

		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte("Hello, World!"))

		slog.Info("Calling validation function")
		validationResult := file.Validate("") // TODO: Replace with a file path or an URL to remote resource to be validated
		serializedResult, _ := json.Marshal(validationResult)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(serializedResult)

	default:
		slog.Warn("Sender used an unsupported HTTP method")
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

// hasValidContentType determines if an HTTP header has valid Content-Type
func hasValidContentType(header http.Header) bool {
	contentType := header.Get("Content-Type")
	return contentType == "text/csv" || contentType == "text/tab-separated-values"
}

// hasEmptyBody determines if HTTP request has empty body
func hasEmptyBody(r *http.Request) bool {
	// Note that Content-Length header may be absent so don't get value from r.Header
	return r.ContentLength == 0
}

// healthCheckHandler writes 200 HTTP status to response header
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Connected to health check handler.", "method", r.Method, "protocol", r.Proto)
	switch r.Method {

	case http.MethodGet:
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
