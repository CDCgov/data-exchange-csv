package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test suite functions follow this flow: prepare data -> define unit tests -> execute tests -> validate results

// MockEvent represents a sample event from a message broker
// We will be using Azure so mimicking Azure Service Bus' event message structure, but this should be platform agnostic
type MockEvent struct {
	ID              string    `json:"id"`
	EventType       string    `json:"eventType"`
	Subject         string    `json:"subject"`
	EventTime       time.Time `json:"eventTime"`
	Data            fileData  `json:"data"`
	DataVersion     string    `json:"dataVersion,omitempty"`
	MetadataVersion string    `json:"metadataVersion,omitempty"`
}

type fileData struct {
	FileURL string `json:"fileUrl"`
}

type httpUnitTest struct {
	name        string    // Optional name for unit test
	method      string    // HTTP method
	contentType string    // Content-Type using RFC 6838 standards (MIME types)
	body        io.Reader // Request body
	status      int       // Expected HTTP status
}

type httpTests struct {
	endpoint string         // Relative endpoint
	queryStr string         // Optional query string
	tests    []httpUnitTest // List of unit tests cases
}

// runTest executes a list of defined subtests against passed in HTTP handler function.
func runTest(t *testing.T, httpTests httpTests, handler func(w http.ResponseWriter, r *http.Request)) {
	for _, test := range httpTests.tests {
		testName := test.name
		if testName == "" {
			testName = test.method
		}

		t.Run(testName, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, httpTests.endpoint, test.body)
			req.Header.Set("Content-Type", test.contentType)

			w := httptest.NewRecorder()

			handler(w, req)

			assert.Equal(t, test.status, w.Code)
		})
	}
}

// TestDefaultHandler tests if root returns a 403 HTTP status for unsupported HTTP methods.
func TestDefaultHandler(t *testing.T) {
	// TODO: Is there a way to feed in a table of arguments and expected outputs in Go?
	// Similar to Java JUnit 5 parameterized tests?
	// https://junit.org/junit5/docs/current/user-guide/#writing-tests-parameterized-tests
	httpTests := httpTests{
		endpoint: "/", // TODO: Use a const instead of hard-coding endpoint
		tests: []httpUnitTest{
			{method: "CONNECT", status: http.StatusNotFound},
			{method: "DELETE", status: http.StatusNotFound},
			{method: "GET", status: http.StatusNotFound},
			{method: "HEAD", status: http.StatusNotFound},
			{method: "OPTIONS", status: http.StatusNotFound},
			{method: "PATCH", status: http.StatusNotFound},
			{method: "POST", status: http.StatusNotFound},
			{method: "PUT", status: http.StatusNotFound},
			{method: "TRACE", status: http.StatusNotFound},
		},
	}

	runTest(t, httpTests, defaultHandler)
}

// TestDefaultHandlerWithNonRoot tests if root returns a 403 HTTP status for unsupported HTTP methods.
func TestDefaultHandlerWithNonRoot(t *testing.T) {
	httpTests := httpTests{
		endpoint: "/test",
		tests: []httpUnitTest{
			{method: "CONNECT", status: http.StatusNotFound},
			{method: "DELETE", status: http.StatusNotFound},
			{method: "GET", status: http.StatusNotFound},
			{method: "HEAD", status: http.StatusNotFound},
			{method: "OPTIONS", status: http.StatusNotFound},
			{method: "PATCH", status: http.StatusNotFound},
			{method: "POST", status: http.StatusNotFound},
			{method: "PUT", status: http.StatusNotFound},
			{method: "TRACE", status: http.StatusNotFound},
		},
	}

	runTest(t, httpTests, defaultHandler)
}

// TestValidateCSVHandlerNonPOST tests non-POST methods to /validate/csv endpoint.
func TestValidateCSVHandlerNonPOST(t *testing.T) {
	httpTests := httpTests{
		endpoint: "/v1/api/validate/csv",
		tests: []httpUnitTest{
			{method: "CONNECT", status: http.StatusMethodNotAllowed},
			{method: "DELETE", status: http.StatusMethodNotAllowed},
			{method: "GET", status: http.StatusMethodNotAllowed},
			{method: "HEAD", status: http.StatusMethodNotAllowed},
			{method: "OPTIONS", status: http.StatusMethodNotAllowed},
			{method: "PATCH", status: http.StatusMethodNotAllowed},
			{method: "PUT", status: http.StatusMethodNotAllowed},
			{method: "TRACE", status: http.StatusMethodNotAllowed},
		},
	}

	runTest(t, httpTests, validateCSVHandler)
}

// TODO: Reevaluate this model; I don't think we will get flat files in initial request body because this service
// will read events off a service bus and use details in those event messages to get the endpoint where file is stored
// so we can stream it; 2024-06-28 update: API is just for users to test their code against before committing to
// the actual production pathway which WILL use service bus behind the curtains
// TestValidateCSVHandlerPOST tests POST method to /validate/csv endpoint with body.
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
		tests: []httpUnitTest{
			{name: "Missing file in body",
				method:      "POST",
				status:      http.StatusBadRequest,
				contentType: "text/csv",
				body:        nil},
			{name: "Non-CSV file in body",
				method:      "POST",
				status:      http.StatusBadRequest,
				contentType: "application/json",
				body:        jsonReader},
			{name: "CSV file in body but wrong content type (JSON)",
				method:      "POST",
				status:      http.StatusBadRequest,
				contentType: "application/json",
				body:        csvReader},
			{name: "CSV file in body but wrong content type (Excel, CSV adjacent format)",
				method:      "POST",
				status:      http.StatusBadRequest,
				contentType: "application/vnd.ms-excel",
				body:        csvReader},
			{name: "CSV file in body but accepted file format for processing (any format with delimiter-separated values)",
				method:      "POST",
				status:      http.StatusAccepted,
				contentType: "text/tab-separated-values",
				body:        csvReader},
			{name: "TSV file in body",
				method:      "POST",
				status:      http.StatusAccepted,
				contentType: "text/tab-separated-values",
				body:        tsvReader},
			{name: "CSV file in body",
				method:      "POST",
				status:      http.StatusAccepted,
				contentType: "text/csv",
				body:        csvReader},
		},
	}

	runTest(t, httpTests, validateCSVHandler)
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
		Data: fileData{
			FileURL: "https://example.blob.core.windows.net/container/file.csv",
		},
		DataVersion:     "1.0",
		MetadataVersion: "1.0",
	}

	jsonBytes, _ := json.Marshal(jsonTestData)
	jsonReader := bytes.NewReader(jsonBytes)

	httpTests := httpTests{
		endpoint: "/v1/api/validate/csv",
		tests: []httpUnitTest{
			{method: "POST",
				status:      http.StatusBadRequest,
				contentType: "application/json",
				body:        jsonReader},
		},
	}

	runTest(t, httpTests, validateCSVHandler)
}

// TestHealthCheckHandler tests if GET /health returns a 200 HTTP status and JSON payload
func TestHealthCheckHandler(t *testing.T) {
	httpTests := httpTests{
		endpoint: "/v1/api/health",
		tests: []httpUnitTest{
			{method: "CONNECT", status: http.StatusMethodNotAllowed},
			{method: "DELETE", status: http.StatusMethodNotAllowed},
			{method: "GET", status: http.StatusOK},
			{method: "HEAD", status: http.StatusMethodNotAllowed},
			{method: "OPTIONS", status: http.StatusMethodNotAllowed},
			{method: "PATCH", status: http.StatusMethodNotAllowed},
			{method: "POST", status: http.StatusMethodNotAllowed},
			{method: "PUT", status: http.StatusMethodNotAllowed},
			{method: "TRACE", status: http.StatusMethodNotAllowed},
		},
	}

	runTest(t, httpTests, healthCheckHandler)
}

// TestDuplicateServer confirms if only one HTTP server can be active at any time
func TestDuplicateServer(t *testing.T) {
	// TODO: Implement this
	// TODO: This test may need to be in its own separate package for black-box testing
	// TODO: This unit tests works if it is implemented outside of server package (for example in main_test.go under main package); unsure why this is the case
	err := New()
	fmt.Println("Hello World")
	assert.NotNil(t, err)
}
