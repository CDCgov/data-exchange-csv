package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type unitTest struct {
	method     string
	queryStr   string
	body       io.Reader
	httpStatus int // Expected status
}

type httpTests struct {
	endpoint string
	tests    []unitTest
}

// TestDefaultHandler tests if / returns a 403 HTTP status
func TestDefaultHandler(t *testing.T) {
	httpTests := httpTests{
		endpoint: "/",
		tests: []unitTest{
			{method: "CONNECT", httpStatus: http.StatusNotFound},
			{method: "DELETE", httpStatus: http.StatusNotFound},
			{method: "GET", httpStatus: http.StatusNotFound},
			{method: "HEAD", httpStatus: http.StatusNotFound},
			{method: "OPTIONS", httpStatus: http.StatusNotFound},
			{method: "PATCH", httpStatus: http.StatusNotFound},
			{method: "POST", httpStatus: http.StatusNotFound},
			{method: "PUT", httpStatus: http.StatusNotFound},
			{method: "TRACE", httpStatus: http.StatusNotFound},
		},
	}

	for _, test := range httpTests.tests {
		req, _ := http.NewRequest(test.method, httpTests.endpoint, test.body)
		w := httptest.NewRecorder()

		defaultHandler(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	}
}

// TestValidateCSVHandler tests /validate/csv endpoint, allowing only GET, HEAD, and POST methods
func TestValidateCSVHandler(t *testing.T) {
	httpTests := httpTests{
		endpoint: "/v1/api/validate/csv",
		tests: []unitTest{
			{method: "CONNECT", httpStatus: http.StatusMethodNotAllowed},
			{method: "DELETE", httpStatus: http.StatusMethodNotAllowed},
			{method: "GET", httpStatus: http.StatusOK},
			{method: "HEAD", httpStatus: http.StatusOK},
			{method: "OPTIONS", httpStatus: http.StatusMethodNotAllowed},
			{method: "PATCH", httpStatus: http.StatusMethodNotAllowed},
			{method: "POST", httpStatus: http.StatusOK},
			{method: "PUT", httpStatus: http.StatusMethodNotAllowed},
			{method: "TRACE", httpStatus: http.StatusMethodNotAllowed},
		},
	}

	for _, test := range httpTests.tests {
		req, _ := http.NewRequest(test.method, httpTests.endpoint, test.body)
		w := httptest.NewRecorder()

		validateCSVHandler(w, req)

		assert.Equal(t, test.httpStatus, w.Code)
	}
}

// TestHealthCheckHandler tests if GET /health returns a 200 HTTP status
func TestHealthCheckHandler(t *testing.T) {
	httpTests := httpTests{
		endpoint: "/v1/api/health",
		tests: []unitTest{
			{method: "CONNECT", httpStatus: http.StatusMethodNotAllowed},
			{method: "DELETE", httpStatus: http.StatusMethodNotAllowed},
			{method: "GET", httpStatus: http.StatusOK},
			{method: "HEAD", httpStatus: http.StatusOK},
			{method: "OPTIONS", httpStatus: http.StatusMethodNotAllowed},
			{method: "PATCH", httpStatus: http.StatusMethodNotAllowed},
			{method: "POST", httpStatus: http.StatusMethodNotAllowed},
			{method: "PUT", httpStatus: http.StatusMethodNotAllowed},
			{method: "TRACE", httpStatus: http.StatusMethodNotAllowed},
		},
	}

	for _, test := range httpTests.tests {
		req, _ := http.NewRequest(test.method, httpTests.endpoint, test.body)
		w := httptest.NewRecorder()

		healthCheckHandler(w, req)

		assert.Equal(t, test.httpStatus, w.Code)
	}
}

// TestDuplicateServer confirms if only one HTTP server can be active at any time
func TestDuplicateServer(t *testing.T) {

}
