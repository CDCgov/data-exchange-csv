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

// TestDefaultHandler tests if / returns a 200 HTTP status
func TestDefaultHandler(t *testing.T) {

}

// TestValidateCSVHandler tests if /validate/csv returns a 200 HTTP status
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

// TestHealthCheckHandler tests if /health returns a 200 HTTP status
func TestHealthCheckHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/api/health", nil)
	w := httptest.NewRecorder()

	healthCheckHandler(w, req)

	assert.Equal(t, 200, w.Code)
}

// TestDuplicateServer confirms if only one HTTP server can be active at any time
func TestDuplicateServer(t *testing.T) {

}
