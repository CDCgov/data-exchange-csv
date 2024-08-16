package file

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
)

type EventMetadata struct {
	ReceivedFilename string `json:"received_filename"`
	DataStreamID     string `json:"data_stream_id"`
	SenderID         string `json:"sender_id"`
	DataProducerID   string `json:"data_producer_id"`
	DataStreamRoute  string `json:"data_stream_route"`
	Jurisdiction     string `json:"jurisdiction"`
	Version          string `json:"version"`
}

type ExpectedValidationResult struct {
	Jurisdiction string
	Delimiter    string
	Encoding     string
	Metadata     *models.MetadataValidationResult
	Config       *models.ConfigValidationResult
}

func verifyValidationResult(t *testing.T, source string, expectedResult ExpectedValidationResult) {
	t.Helper()

	validationResult := Validate(source)
	assertEqual(t, "encoding", expectedResult.Encoding, string(validationResult.Encoding))
	assertEqual(t, "data_stream_id", expectedResult.Metadata.DataStreamID, validationResult.Metadata.DataStreamID)
	assertEqual(t, "sender_id", expectedResult.Metadata.SenderID, validationResult.Metadata.SenderID)
	assertEqual(t, "data_producer_id", expectedResult.Metadata.DataProducerID, validationResult.Metadata.DataProducerID)
	assertEqual(t, "jurisdiction", expectedResult.Jurisdiction, validationResult.Metadata.Jurisdiction)
	assertEqual(t, "version", expectedResult.Metadata.Version, validationResult.Metadata.Version)
}

func assertEqual(t *testing.T, field string, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected %s: %s, but got: %s", field, expected, actual)
	}
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	tempDirectory := "dex-csv-file-validation-test-temp"
	err := os.Mkdir(tempDirectory, 0755)
	if err != nil {
		tb.Fatalf("%s: %v", constants.DIRECTORY_CREATE_ERROR, err)
	}

	testFilesWithContent := map[string]string{
		"UTF8Encoding.csv":        "Hello, World! This is US-ASCII.\nLine 2: More text.",
		"UTF8BomEncoding.csv":     "\xEF\xBB\xBFName,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok,cmallok@example.com",
		"USASCIIEncoding.csv":     "Chris~Wilson,DevOps engineer, ensures CI/CD pipelines and *NIX server maintenance.",
		"Windows1252Encoding.csv": "L'éƒté en France,München Äpfel\nJosé DíažŸ,François Dupont",
		"ISO8859_1Encoding.csv":   "José Dí^az,Software engineer, working on CSV & Golang.",
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

func TestValidateUTF8EncodedCSVFile(t *testing.T) {
	validationResult := ExpectedValidationResult{
		Delimiter:    constants.CSV,
		Encoding:     string(constants.UTF8),
		Jurisdiction: constants.CSV_JURISDICTION,
		Metadata: &models.MetadataValidationResult{
			DataProducerID:  constants.CSV_DATA_PRODUCER_ID,
			DataStreamID:    constants.CSV_DATA_STREAM_ID,
			DataStreamRoute: constants.CSV_DATA_STREAM_ROUTE,
			SenderID:        constants.CSV_SENDER_ID,
			Version:         constants.VERSION,
		},
	}
	verifyValidationResult(t, "dex-csv-file-validation-test-temp/UTF8Encoding.json", validationResult)
}

func TestValidateUTF8BomEncodedCSVFile(t *testing.T) {
	validationResult := ExpectedValidationResult{
		Delimiter:    constants.CSV,
		Encoding:     string(constants.UTF8_BOM),
		Jurisdiction: constants.CSV_JURISDICTION,
		Metadata: &models.MetadataValidationResult{
			DataProducerID:  constants.CSV_DATA_PRODUCER_ID,
			DataStreamID:    constants.CSV_DATA_STREAM_ID,
			DataStreamRoute: constants.CSV_DATA_STREAM_ROUTE,
			SenderID:        constants.CSV_SENDER_ID,
			Version:         constants.VERSION,
		},
	}
	verifyValidationResult(t, "dex-csv-file-validation-test-temp/UTF8BomEncoding.json", validationResult)
}

func TestValidateUSASCIIEncodedCSVFile(t *testing.T) {
	validationResult := ExpectedValidationResult{
		Delimiter:    constants.CSV,
		Encoding:     string(constants.UTF8),
		Jurisdiction: constants.CSV_JURISDICTION,
		Metadata: &models.MetadataValidationResult{
			DataProducerID:  constants.CSV_DATA_PRODUCER_ID,
			DataStreamID:    constants.CSV_DATA_STREAM_ID,
			DataStreamRoute: constants.CSV_DATA_STREAM_ROUTE,
			SenderID:        constants.CSV_SENDER_ID,
			Version:         constants.VERSION,
		},
	}
	verifyValidationResult(t, "dex-csv-file-validation-test-temp/USASCIIEncoding.json", validationResult)
}

func TestValidateWindows1252EncodedCSVFile(t *testing.T) {
	validationResult := ExpectedValidationResult{
		Delimiter:    constants.CSV,
		Encoding:     string(constants.WINDOWS1252),
		Jurisdiction: constants.CSV_JURISDICTION,
		Metadata: &models.MetadataValidationResult{
			DataProducerID:  constants.CSV_DATA_PRODUCER_ID,
			DataStreamID:    constants.CSV_DATA_STREAM_ID,
			DataStreamRoute: constants.CSV_DATA_STREAM_ROUTE,
			SenderID:        constants.CSV_SENDER_ID,
			Version:         constants.VERSION,
		},
	}
	verifyValidationResult(t, "dex-csv-file-validation-test-temp/Windows1252Encoding.json", validationResult)
}

func TestValidateISO8859_1EncodedCSVFile(t *testing.T) {
	validationResult := ExpectedValidationResult{
		Delimiter:    constants.CSV,
		Encoding:     string(constants.ISO8859_1),
		Jurisdiction: constants.CSV_JURISDICTION,
		Metadata: &models.MetadataValidationResult{
			DataProducerID:  constants.CSV_DATA_PRODUCER_ID,
			DataStreamID:    constants.CSV_DATA_STREAM_ID,
			DataStreamRoute: constants.CSV_DATA_STREAM_ROUTE,
			SenderID:        constants.CSV_SENDER_ID,
			Version:         constants.VERSION,
		},
	}
	verifyValidationResult(t, "dex-csv-file-validation-test-temp/ISO8859_1Encoding.json", validationResult)
}
