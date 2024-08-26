package row

import (
	"encoding/json"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const tempDirectory = "dex-csv-row-validation-test-temp"

type EventMetadata struct {
	ReceivedFilename string `json:"received_filename"`
	DataStreamID     string `json:"data_stream_id"`
	SenderID         string `json:"sender_id"`
	DataProducerID   string `json:"data_producer_id"`
	DataStreamRoute  string `json:"data_stream_route"`
	Jurisdiction     string `json:"jurisdiction"`
	Version          string `json:"version"`
}

type MockSendEventsToDestination struct {
	result      interface{}
	destination string
}

func (m *MockSendEventsToDestination) callback(result interface{}, destination string) {
	m.result = result
	m.destination = destination
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected: %s, but got: %s", expected, actual)
	}
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	err := os.Mkdir(tempDirectory, 0755)
	if err != nil {
		tb.Fatalf("%s: %v", constants.DIRECTORY_CREATE_ERROR, err)
	}

	testFilesWithContent := map[string]string{
		"UTF8Encoding.csv":    "Hello, World! This is US-ASCII.\nLine 2: More text.",
		"UTF8BomEncoding.csv": "\xEF\xBB\xBFName,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok,cmallok@example.com",
		//"USASCIIEncoding.csv":     "Chris~Wilson,DevOps engineer, ensures CI/CD pipelines and *NIX server maintenance.",
		//"Windows1252Encoding.csv": "L'éƒté en France,München Äpfel\nJosé DíažŸ,François Dupont",
		//"ISO8859_1Encoding.csv":   "José Dí^az,Software engineer, working on CSV & Golang.",
		"SemicolonDelimiter.csv": "CSV; with; semicolons; as; delimiter",
		"NoDelimiter.csv":        "Lorem ipsum dolor sit amet",
	}

	for file, content := range testFilesWithContent {
		tempFile := filepath.Join(tempDirectory, file)
		err := os.WriteFile(tempFile, []byte(content), 0644)
		if err != nil {
			tb.Fatalf("%s %s: %v", constants.FILE_WRITE_ERROR, file, err)
		}

		event := EventMetadata{
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
		//replace file extension from .csv to .json
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
	params := []models.FileValidationParams{
		models.FileValidationParams{
			ReceivedFile: tempDirectory + "/UTF8Encoding.csv",
			Encoding:     constants.UTF8_BOM,
			Delimiter:    ",", // TODO: Should this Delimiter field be a rune type?
			FileUUID:     uuid.New(),
		},
		models.FileValidationParams{
			ReceivedFile: tempDirectory + "/UTF8BomEncoding.csv",
			Encoding:     constants.UTF8_BOM,
			Delimiter:    ",",
			FileUUID:     uuid.New(),
		},
	}

	expectedResults := []models.RowValidationResult{
		models.RowValidationResult{
			FileUUID:  params[0].FileUUID, // TODO: There is a probably a better way to set up a 2D matrix of structs for creating input and output data; otherwise you deal with this odd hard-coding indices
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash1",
			RowNumber: 1,
		},
		models.RowValidationResult{
			FileUUID:  params[1].FileUUID,
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash2",
			RowNumber: 2,
		},
		// ... add more expected results if needed
	}

	sendEventsToDestination := MockSendEventsToDestination{}

	for i, param := range params {
		Validate(param, sendEventsToDestination.callback)
		assertEqual(t, expectedResults[i], sendEventsToDestination.result)
		sendEventsToDestination = MockSendEventsToDestination{} // reset struct contents
	}
}

// TestValidate_ErrorCreatingReader tests when Validate() fails to create a CSV file reader.
func TestValidate_ErrorCreatingReader(t *testing.T) {
	params := models.FileValidationParams{
		ReceivedFile: tempDirectory + "/NonExistent.csv", // non-existent file
		Encoding:     constants.UTF8_BOM,
		Delimiter:    ",",
		FileUUID:     uuid.New(),
	}

	expectedResult := models.RowValidationResult{
		FileUUID: params.FileUUID,
		Status:   constants.STATUS_FAILED,
		Error: &models.RowError{
			Message:  "CSV reader error",
			Severity: "Failure",
			Line:     -1,
			Column:   -1,
		},
	}

	sendEventsToDestination := MockSendEventsToDestination{}

	Validate(params, sendEventsToDestination.callback)

	assertEqual(t, expectedResult, sendEventsToDestination.result)
}

// TestValidate_ErrorReadingRow tests when Validate() encounters a read error.
func TestValidate_ErrorReadingRow(t *testing.T) {

	params := []models.FileValidationParams{
		models.FileValidationParams{
			ReceivedFile: tempDirectory + "/SemicolonDelimiter.csv",
			Encoding:     constants.UTF8_BOM,
			Delimiter:    ",",
			FileUUID:     uuid.New(),
		},
		models.FileValidationParams{
			ReceivedFile: tempDirectory + "/NoDelimiter.csv",
			Encoding:     constants.UTF8_BOM,
			Delimiter:    ",",
			FileUUID:     uuid.New(),
		},
	}

	expectedResults := []models.RowValidationResult{
		models.RowValidationResult{
			FileUUID:  params[0].FileUUID, // TODO: There is a probably a better way to set up a 2D matrix of structs for creating input and output data; otherwise you deal with this odd hard-coding indices
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash1",
			RowNumber: 1,
		},
		models.RowValidationResult{
			FileUUID: params[1].FileUUID,
			Status:   constants.STATUS_FAILED,
			Error: &models.RowError{
				Message:  "Mismatched field counts",
				Severity: "Failure",
				Line:     -1,
				Column:   -1,
			},
		},
	}

	sendEventsToDestination := MockSendEventsToDestination{}

	for i, param := range params {
		Validate(param, sendEventsToDestination.callback)
		assertEqual(t, expectedResults[i], sendEventsToDestination.result)
		sendEventsToDestination = MockSendEventsToDestination{}
	}
}
