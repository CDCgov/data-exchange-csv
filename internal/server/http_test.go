package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type unitTest struct {
	method      string
	contentType string // Using RFC 6838 standards (MIME types)
	body        io.Reader
	httpStatus  int // Expected status
}

type httpTests struct {
	endpoint string
	queryStr string
	tests    []unitTest
}

// TestDefaultHandler tests if root returns a 403 HTTP status
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

// TestValidateCSVHandler tests non-POST methods to /validate/csv endpoint.
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

// TestValidateCSVHandlerPOST tests POST method to /validate/csv endpoint.
// POST methods allow CSV and TSV files in payload.
// Note: This test doesn't validate business logic of CSV validation.
func TestValidateCSVHandlerPOST(t *testing.T) {
	csvTestData := "first_name,last_name\nJohn,Doe\nMary,Jane"
	tsvTestData := "first_name\tlast_name\nJohn\tDoe\nMary\tJane"
	jsonTestData := `{
{
	first_name: John,
	last_name: Doe
},
{
	first_name: Mary,
	last_name: Jane
}
}`

	csvReader := strings.NewReader(csvTestData)
	tsvReader := strings.NewReader(tsvTestData)
	jsonReader := strings.NewReader(jsonTestData)

	httpTests := httpTests{
		endpoint: "/v1/api/validate/csv",
		tests: []unitTest{
			// Missing file in body
			{method: "POST", httpStatus: http.StatusBadRequest,
				contentType: "text/csv"},
			// Non-CSV file in body
			{method: "POST", httpStatus: http.StatusBadRequest,
				contentType: "application/json", body: jsonReader},
			// CSV file in body but wrong content type (JSON)
			{method: "POST", httpStatus: http.StatusBadRequest,
				contentType: "application/json", body: csvReader},
			// CSV file in body but wrong content type (Excel, CSV adjacent format)
			{method: "POST", httpStatus: http.StatusBadRequest,
				contentType: "application/vnd.ms-excel", body: csvReader},
			// CSV file in body but accepted file format for processing (any format with delimiter-separated values)
			{method: "POST", httpStatus: http.StatusAccepted,
				contentType: "text/tab-separated-values", body: csvReader},
			// TSV file in body
			{method: "POST", httpStatus: http.StatusAccepted,
				contentType: "text/tab-separated-values", body: tsvReader},
			// Base success case
			{method: "POST", httpStatus: http.StatusAccepted,
				contentType: "text/csv", body: csvReader},
		},
	}

	for _, test := range httpTests.tests {
		req, _ := http.NewRequest(test.method, httpTests.endpoint, test.body)
		req.Header.Set("Content-Type", test.contentType)

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
	// TODO: Implement this
}
