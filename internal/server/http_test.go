package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type httpTests struct {
	endpoint string
	// tests	[]struct {
	// method: string
	// httpStatus: int
	// body: string
	// }
	method     string
	httpStatus int // Expected status
	body       io.Reader
}

// TestDefaultHandler tests if / returns a 403 HTTP status
func TestDefaultHandler(t *testing.T) {
	tests := []httpTests{
		{method: "GET", endpoint: "/", httpStatus: http.StatusNotFound},
		{method: "HEAD", endpoint: "/", httpStatus: http.StatusNotFound},
		{method: "POST", endpoint: "/", httpStatus: http.StatusNotFound},
		{method: "PUT", endpoint: "/", httpStatus: http.StatusNotFound},
		{method: "DELETE", endpoint: "/", httpStatus: http.StatusNotFound},
		{method: "CONNECT", endpoint: "/", httpStatus: http.StatusNotFound},
		{method: "OPTIONS", endpoint: "/", httpStatus: http.StatusNotFound},
		{method: "TRACE", endpoint: "/", httpStatus: http.StatusNotFound},
		{method: "PATCH", endpoint: "/", httpStatus: http.StatusNotFound},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(test.method, test.endpoint, test.body)
		w := httptest.NewRecorder()

		defaultHandler(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	}
}

// TestValidateCSVHandler tests /validate/csv endpoint, allowing only GET, HEAD, and POST methods
func TestValidateCSVHandler(t *testing.T) {
	tests := []httpTests{
		{method: "GET", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusOK},
		{method: "HEAD", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusOK},
		{method: "POST", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusOK},
		{method: "PUT", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusMethodNotAllowed},
		{method: "DELETE", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusMethodNotAllowed},
		{method: "CONNECT", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusMethodNotAllowed},
		{method: "OPTIONS", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusMethodNotAllowed},
		{method: "TRACE", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusMethodNotAllowed},
		{method: "PATCH", endpoint: "/v1/api/validate/csv", httpStatus: http.StatusMethodNotAllowed},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(test.method, test.endpoint, test.body)
		w := httptest.NewRecorder()

		validateCSVHandler(w, req)

		assert.Equal(t, test.httpStatus, w.Code)
	}
}

// TestHealthCheckHandler tests if GET /health returns a 200 HTTP status
func TestHealthCheckHandler(t *testing.T) {
	tests := []httpTests{
		{method: "GET", endpoint: "/v1/api/health", httpStatus: http.StatusOK},
		{method: "HEAD", endpoint: "/v1/api/health", httpStatus: http.StatusOK},
		{method: "POST", endpoint: "/v1/api/health", httpStatus: http.StatusMethodNotAllowed},
		{method: "PUT", endpoint: "/v1/api/health", httpStatus: http.StatusMethodNotAllowed},
		{method: "DELETE", endpoint: "/v1/api/health", httpStatus: http.StatusMethodNotAllowed},
		{method: "CONNECT", endpoint: "/v1/api/health", httpStatus: http.StatusMethodNotAllowed},
		{method: "OPTIONS", endpoint: "/v1/api/health", httpStatus: http.StatusMethodNotAllowed},
		{method: "TRACE", endpoint: "/v1/api/health", httpStatus: http.StatusMethodNotAllowed},
		{method: "PATCH", endpoint: "/v1/api/health", httpStatus: http.StatusMethodNotAllowed},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(test.method, test.endpoint, test.body)
		w := httptest.NewRecorder()

		healthCheckHandler(w, req)

		assert.Equal(t, test.httpStatus, w.Code)
	}
}

// TestDuplicateServer confirms if only one HTTP server can be active at any time
func TestDuplicateServer(t *testing.T) {

}
