package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockEvent represents a sample event from a message broker
// We will be using Azure so mimicking Azure Service Bus' event message structure, but this should be agnostic
type MockEvent struct {
	ID              string      `json:"id"`
	EventType       string      `json:"eventType"`
	Subject         string      `json:"subject"`
	EventTime       time.Time   `json:"eventTime"`
	Data            newFileData `json:"data"`
	DataVersion     string      `json:"dataVersion,omitempty"`
	MetadataVersion string      `json:"metadataVersion,omitempty"`
}

type newFileData struct {
	FileURL string `json:"fileUrl"`
}

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

// TestDefaultHandler tests if root returns a 403 HTTP status for unsupported HTTP methods.
func TestDefaultHandler(t *testing.T) {
	// TODO: Refactor this unit test struct b/c if one specific test fails, it is harder to backtrace which one of these elements
	// gave error. Is there a way to feed in a table of arguments and expected outputs in Go?
	// Similar to Java JUnit 5 parameterized tests?
	// https://junit.org/junit5/docs/current/user-guide/#writing-tests-parameterized-tests
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

// TestDefaultHandlerWithNonRoot tests if root returns a 403 HTTP status for unsupported HTTP methods.
func TestDefaultHandlerWithNonRoot(t *testing.T) {
	httpTests := httpTests{
		endpoint: "/test",
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
			{method: "GET", httpStatus: http.StatusMethodNotAllowed},
			{method: "HEAD", httpStatus: http.StatusMethodNotAllowed},
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

// TODO: Reevaluate this model; I don't think we will get flat files in initial request body because this service
// will read events off a service bus and use details in those event messages to get the endpoint where file is stored
// so we can stream it; 2024-06-28 update: API is just for users to test their code against before committing to
// the actual production pathway which WILL use service bus behind the curtains
// TestValidateCSVHandlerWithBody tests POST method to /validate/csv endpoint with body.
// Note: This test doesn't validate business logic of CSV validation.
func TestValidateCSVHandlerWithBody(t *testing.T) {
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
			// CSV file in body
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

// TestValidateCSVHandlerEventMessage tests POST method to /validate/csv endpoint with an event message in body.
// Note: This test doesn't validate business logic of CSV validation.
func TestValidateCSVHandlerEventMessage(t *testing.T) {
	// TODO: Clarify whether this API endpoint is just for end-users to test their CSV file or is this API exclusive
	// for internal use; that affects whether or not users can send a simple POST request with a CSV file in body vs.
	// going through whole pipeline
	jsonTestData := MockEvent{
		ID:        "42",
		EventType: "NewFile",
		Subject:   "csv",
		EventTime: time.Now(),
		Data: newFileData{
			FileURL: "https://example.blob.core.windows.net/container/file.csv",
		},
		DataVersion:     "1.0",
		MetadataVersion: "1.0",
	}

	jsonBytes, _ := json.Marshal(jsonTestData)
	jsonReader := bytes.NewReader(jsonBytes)

	httpTests := httpTests{
		endpoint: "/v1/api/validate/csv",
		tests: []unitTest{
			// Base success case
			{method: "POST", httpStatus: http.StatusBadRequest,
				contentType: "application/json", body: jsonReader},
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

// TestHealthCheckHandler tests if GET /health returns a 200 HTTP status and JSON payload
func TestHealthCheckHandler(t *testing.T) {
	httpTests := httpTests{
		endpoint: "/v1/api/health",
		tests: []unitTest{
			{method: "CONNECT", httpStatus: http.StatusMethodNotAllowed},
			{method: "DELETE", httpStatus: http.StatusMethodNotAllowed},
			{method: "GET", httpStatus: http.StatusOK},
			{method: "HEAD", httpStatus: http.StatusMethodNotAllowed},
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
	_ = New()
	assert.NotNil(t, New())
}
