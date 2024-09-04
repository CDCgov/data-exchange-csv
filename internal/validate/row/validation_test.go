package row

import (
	"encoding/json"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const tempDirectory = "dex-csv-row-validation-test-temp"

// TODO: Should eventMetadata be a constant type? If so move definition to /models
type eventMetadata struct {
	ReceivedFilename string `json:"received_filename"`
	DataStreamID     string `json:"data_stream_id"`
	SenderID         string `json:"sender_id"`
	DataProducerID   string `json:"data_producer_id"`
	DataStreamRoute  string `json:"data_stream_route"`
	Jurisdiction     string `json:"jurisdiction"`
	Version          string `json:"version"`
}

type rowValidationTest struct {
	name     string // Test name
	input    models.FileValidationParams
	expected []expectedRowValidationResult // A file can have multiple rows, thus one to many row validation results
}

// expectedRowValidationResult contains a subset of the fields of models.RowValidationResult that we want to test against
// (ignoring fields with randomized values like FileUUID and RowUUID)
type expectedRowValidationResult struct {
	RowNumber int
	Error     *models.RowError
	Status    string
}

type mockSendEventsToDestination struct {
	result      interface{}
	destination string // a path where result is routed to
}

func (m *mockSendEventsToDestination) callback(result interface{}, destination string) {
	m.result = result
	m.destination = destination
}

func verifyValidationResult(t *testing.T, expected expectedRowValidationResult, actual models.RowValidationResult) {
	// We want to compare only a subset of the RowValidationResult fields
	// TODO: Do we want to explicitly downcast RowValidationResult to a expectedRowValidationResult? And if so should this be done
	// in this function or before this function call? It feels off that expected and actual args are not the same type.
	t.Helper()

	expectedFields := reflect.TypeOf(expected)
	expectedValues := reflect.ValueOf(expected)

	actualValues := reflect.ValueOf(actual)

	// Looping over each struct field and comparing field values programmatically instead of hard-coding asserts
	for i := 0; i < expectedFields.NumField(); i++ {
		field := expectedFields.Field(i).Name
		expectedValue := expectedValues.Field(i)

		actualValue := actualValues.Field(i)

		assertEqual(t, field, expectedValue, actualValue)
	}
}

func assertEqual(t *testing.T, field string, expected interface{}, actual interface{}) {
	// Need to ensure both types are equal before comparison
	isEqual := reflect.TypeOf(expected) == reflect.TypeOf(actual) && expected == actual

	if !isEqual {
		t.Errorf("Expected %s: %s, but got: %s", field, expected, actual)
	}
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	// TODO: Use T.TempDir?
	_ = os.RemoveAll(tempDirectory) // removing directory created from previously failed runs
	err := os.Mkdir(tempDirectory, 0755)

	if err != nil {
		tb.Fatalf("%s: %v", constants.DIRECTORY_CREATE_ERROR, err)
	}

	testFilesWithContent := map[string]string{
		"UTF8Encoding.csv":       "Hello, World! This is US-ASCII.\nLine 2: More text.",
		"UTF8BomEncoding.csv":    "\xEF\xBB\xBFName,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok,cmallok@example.com",
		"SemicolonDelimiter.csv": "CSV;with;semicolons;as;delimiter",
		"TabDelimiter.csv":       "CSV\twith\ttabs\tas\tdelimiter",
		"NoDelimiter.csv":        "Lorem ipsum dolor sit amet",
		"EmptyFile.csv":          "",
	}

	for file, content := range testFilesWithContent {
		tempFile := filepath.Join(tempDirectory, file)
		err := os.WriteFile(tempFile, []byte(content), 0644)
		if err != nil {
			tb.Fatalf("%s %s: %v", constants.FILE_WRITE_ERROR, file, err)
		}

		// TODO: Remove data stream content as per pivot to standalone Go package
		event := eventMetadata{
			ReceivedFilename: filepath.Join(tempDirectory, file),
			DataStreamID:     constants.CSV_DATA_STREAM_ID,
			SenderID:         constants.CSV_SENDER_ID,
			DataStreamRoute:  constants.CSV_DATA_STREAM_ROUTE,
			DataProducerID:   constants.CSV_DATA_PRODUCER_ID,
			Jurisdiction:     constants.CSV_JURISDICTION,
			Version:          constants.VERSION,
		}

		eventAsJson, err := json.MarshalIndent(event, "", "    ")
		if err != nil {
			tb.Fatalf("%s %s: %v", constants.ERROR_CONVERTING_STRUCT_TO_JSON, file, err)
		}
		// Replace file extension from .csv to .json
		eventFileName := strings.TrimSuffix(file, filepath.Ext(file)) + constants.JSON_EXTENSION
		eventFilePath := filepath.Join(tempDirectory, eventFileName)

		err = os.WriteFile(eventFilePath, eventAsJson, 0644)
		if err != nil {
			tb.Fatalf("%s %s: %v", constants.FILE_WRITE_ERROR, eventFileName, err)
		}
	}

	return func(tb testing.TB) {
		err := os.RemoveAll(tempDirectory)
		if err != nil {
			tb.Errorf("%s %s: %v", constants.DIRECTORY_REMOVE_ERROR, tempDirectory, err)
		}
	}
}

func TestMain(m *testing.M) {
	teardownTest := setupTest(nil)
	executeTestCases := m.Run()

	if teardownTest != nil {
		teardownTest(nil)
	}
	os.Exit(executeTestCases)
}

// TestValidate_Success	tests positive path in Validate().
func TestValidate_Success(t *testing.T) {
	tests := []rowValidationTest{
		{
			name: "Valid UTF8 Encoded CSV Test - 2 Rows",
			input: models.FileValidationParams{
				ReceivedFile: tempDirectory + "/UTF8Encoding.csv",
				Encoding:     constants.UTF8_BOM,
				Delimiter:    ",", // TODO: Should this Delimiter field be a rune type?
				FileUUID:     uuid.New(),
			},
			expected: []expectedRowValidationResult{
				{
					Status:    constants.STATUS_SUCCESS,
					Error:     nil,
					RowNumber: 1,
				},
				{
					Status:    constants.STATUS_SUCCESS,
					Error:     nil,
					RowNumber: 2,
				},
			},
		},
		{
			name: "Valid UTF8 Bom Encoded CSV Test - 1 Row",
			input: models.FileValidationParams{
				ReceivedFile: tempDirectory + "/UTF8BomEncoding.csv",
				Encoding:     constants.UTF8_BOM,
				Delimiter:    ",",
				FileUUID:     uuid.New(),
			},
			expected: []expectedRowValidationResult{
				{
					Status:    constants.STATUS_SUCCESS,
					Error:     nil,
					RowNumber: 1,
				},
			},
		},
	}

	sendEventsToDestination := mockSendEventsToDestination{}

	for _, test := range tests {
		Validate(test.input, sendEventsToDestination.callback)
		// TODO: panic: interface conversion failed. interface{} is models.RowTransformationResult, not models.RowValidationResult
		// This means on a row validation success, result struct is a RowTransformationResult (meaning row is transformed into json format on success)
		// Is RowValidationResult not an emitted/returned?
		// TODO: I think this logic in comparing expected to actual output is wrong because the struct that stores the row errors has a one-to-many relationship
		// with the RowValidationResult. That means we have to loop both the array of expected results and I guess an array that stores the actual RowValidationResults no?
		actualResult := sendEventsToDestination.result.(models.RowValidationResult)
		t.Run(test.name, func(t *testing.T) {
			for _, expectedResult := range test.expected {
				verifyValidationResult(t, expectedResult, actualResult)
			}
		})
		sendEventsToDestination = mockSendEventsToDestination{} // reset struct contents
	}
}

// TestValidate_ErrorCreatingReader tests when Validate() fails to create a CSV file reader.
func TestValidate_ErrorCreatingReader(t *testing.T) {
	tests := []rowValidationTest{
		{
			name: "Non-Existent File",
			input: models.FileValidationParams{
				ReceivedFile: tempDirectory + "/NonExistent.csv", // non-existent file
				Encoding:     constants.UTF8_BOM,
				Delimiter:    ",",
				FileUUID:     uuid.New(),
			},
			expected: []expectedRowValidationResult{
				{
					Status: constants.STATUS_FAILED,
					Error: &models.RowError{
						Message:  constants.CSV_READER_ERROR,
						Severity: constants.Failure,
						Line:     -1,
						Column:   -1,
					},
					RowNumber: 1,
				},
			},
		},
	}

	sendEventsToDestination := mockSendEventsToDestination{}

	for _, test := range tests {
		Validate(test.input, sendEventsToDestination.callback)
		// TODO: panic: interface conversion failed. interface{} is models.RowTransformationResult, not models.RowValidationResult
		// This means on a row validation success, result struct is a RowTransformationResult (meaning row is transformed into json format on success)
		// Is RowValidationResult not an emitted/returned?
		actualResult := sendEventsToDestination.result.(models.RowValidationResult)
		t.Run(test.name, func(t *testing.T) {
			for _, expectedResult := range test.expected {
				verifyValidationResult(t, expectedResult, actualResult)
			}
		})
		sendEventsToDestination = mockSendEventsToDestination{} // reset struct contents
	}
}

// TestValidate_ErrorReadingRow tests when Validate() encounters a read error.
func TestValidate_ErrorReadingRow(t *testing.T) {
	tests := []rowValidationTest{
		{
			name: "Wrong Delimiter - File With Semicolon Delimiter",
			input: models.FileValidationParams{
				ReceivedFile: tempDirectory + "/SemicolonDelimiter.csv",
				Encoding:     constants.UTF8_BOM,
				Delimiter:    ",",
				FileUUID:     uuid.New(),
			},
			expected: []expectedRowValidationResult{
				{
					Status:    constants.STATUS_SUCCESS,
					Error:     nil,
					RowNumber: 1,
				},
			},
		},
		{
			name: "Wrong Delimiter - File With No Delimiter",
			input: models.FileValidationParams{
				ReceivedFile: tempDirectory + "/NoDelimiter.csv",
				Encoding:     constants.UTF8_BOM,
				Delimiter:    ",",
				FileUUID:     uuid.New(),
			},
			expected: []expectedRowValidationResult{
				{
					Status:    constants.STATUS_SUCCESS,
					Error:     nil,
					RowNumber: 1,
				},
				{
					Status: constants.STATUS_FAILED,
					Error: &models.RowError{
						Message:  "Mismatched field counts",
						Severity: "Failure",
						Line:     -1,
						Column:   -1,
					},
					RowNumber: 1,
				},
			},
		},
	}

	sendEventsToDestination := mockSendEventsToDestination{}

	for _, test := range tests {
		Validate(test.input, sendEventsToDestination.callback)
		actualResult := sendEventsToDestination.result.(models.RowValidationResult)
		for _, expectedResult := range test.expected {
			verifyValidationResult(t, expectedResult, actualResult)
		}
		sendEventsToDestination = mockSendEventsToDestination{}
	}
}
