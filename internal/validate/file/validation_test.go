package file

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

type eventConfig struct {
	ReceivedFilename string `json:"received_filename"`
	DataStreamID     string `json:"data_stream_id"`
	SenderID         string `json:"sender_id"`
	DataProducerID   string `json:"data_producer_id"`
	DataStreamRoute  string `json:"data_stream_route"`
	Jurisdiction     string `json:"jurisdiction"`
	Version          string `json:"version"`
}

type expectedValidationResult struct {
	dataStreamID    string
	senderID        string
	dataProducerID  string
	dataStreamRoute string
	jurisdiction    string
	version         string
	delimiter       string
	encoding        string
}

func verifyValidationResult(t *testing.T, source string, expectedResult expectedValidationResult) {
	t.Helper()
	validationResult := &ValidationResult{}
	validationResult.Validate(source)

	if expectedResult.encoding != string(validationResult.Encoding) {
		t.Errorf("Expected encoding %s, got %s", expectedResult.encoding, validationResult.Encoding)
	}

	if expectedResult.dataStreamID != validationResult.DataStreamID {
		t.Errorf("Expected DataStreamID %s, got %s", expectedResult.dataStreamID, validationResult.DataStreamID)
	}

	if expectedResult.senderID != validationResult.SenderID {
		t.Errorf("Expected SenderID %s, got %s", expectedResult.senderID, validationResult.SenderID)
	}

	if expectedResult.dataProducerID != validationResult.DataProducerID {
		t.Errorf("Expected DataProducerID %s, got %s", expectedResult.dataProducerID, validationResult.DataProducerID)
	}

	if expectedResult.jurisdiction != validationResult.Jurisdiction {
		t.Errorf("Expected Jurisdiction %s, got %s", expectedResult.jurisdiction, validationResult.Jurisdiction)
	}

	if expectedResult.version != validationResult.Version {
		t.Errorf("Expected Version %s, got %s", expectedResult.version, validationResult.Version)
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
		"UTF8BomEncoding.csv":     "\xEF\xBB\xBF\nName,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok,cmallok@example.com",
		"USASCIIEncoding.csv":     "Chris~Wilson,DevOps engineer, ensures CI/CD pipelines and *NIX server maintenance.",
		"Windows1252Encoding.csv": "Brontë Sisters,München Äpfel,José Díaz,François Dupont",
		"ISO8859_1Encoding.csv":   "José Díaz,Software engineer, working on CSV & Golang.",
		"TestCSVHeader.csv":       "Index,Name,Description\nJosé Díaz,Software engineer, working on C++ & Python.\nFrançois Dupont,Product manager: oversees marketing & sales.",
		"TestTSVHeader.tsv":       "Index\tName\tDescription\nJosé Díaz\tSoftware engineer\tworking on C++ & Python.\nFrançois Dupont\tProduct manager: oversees marketing & sales.",
	}

	for file, content := range testFilesWithContent {
		tempFile := filepath.Join(tempDirectory, file)

		err := os.WriteFile(tempFile, []byte(content), 0644)
		if err != nil {
			tb.Fatalf("%s %s: %v", constants.FILE_WRITE_ERROR, file, err)
		}

		var receivedFilenameBuilder strings.Builder
		receivedFilenameBuilder.WriteString(tempDirectory + "/")
		receivedFilenameBuilder.WriteString(file)

		event := eventConfig{
			ReceivedFilename: receivedFilenameBuilder.String(),
			DataStreamID:     constants.DATA_STREAM_ID,
			SenderID:         constants.SENDER_ID,
			DataProducerID:   constants.DATA_PRODUCER_ID,
			Jurisdiction:     constants.JURISDICTION,
			Version:          constants.VERSION,
		}

		eventAsJson, err := json.MarshalIndent(event, "", "    ")
		if err != nil {
			tb.Fatalf("%s %s: %v", constants.ERROR_CONVERTING_STRUCT_TO_JSON, file, err)
		}

		eventFileName := file[:len(file)-4] + constants.JSON_EXTENSION // remove .csv from file replace with .json
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
	source := "dex-csv-file-validation-test-temp/UTF8Encoding.json"
	validationResult := expectedValidationResult{}
	validationResult.dataProducerID = constants.DATA_PRODUCER_ID
	validationResult.dataStreamID = constants.DATA_STREAM_ID
	validationResult.dataStreamRoute = constants.DATA_STREAM_ROUTE
	validationResult.delimiter = constants.CSV
	validationResult.encoding = string(constants.UTF8)
	validationResult.jurisdiction = constants.JURISDICTION

	verifyValidationResult(t, source, validationResult)

}
func TestValidateUTF8BomEncodedCSVFile(t *testing.T) {
	source := "dex-csv-file-validation-test-temp/UTF8BomEncoding.json"
	validationResult := expectedValidationResult{}
	validationResult.dataProducerID = constants.DATA_PRODUCER_ID
	validationResult.dataStreamID = constants.DATA_STREAM_ID
	validationResult.dataStreamRoute = constants.DATA_STREAM_ROUTE
	validationResult.delimiter = constants.CSV
	validationResult.encoding = string(constants.UTF8_BOM)
	validationResult.jurisdiction = constants.JURISDICTION

	verifyValidationResult(t, source, validationResult)
}
func TestValidateUSASCIIEncodedCSVFile(t *testing.T) {
	source := "dex-csv-file-validation-test-temp/USASCIIEncoding.json"
	validationResult := expectedValidationResult{}
	validationResult.dataProducerID = constants.DATA_PRODUCER_ID
	validationResult.dataStreamID = constants.DATA_STREAM_ID
	validationResult.dataStreamRoute = constants.DATA_STREAM_ROUTE
	validationResult.delimiter = constants.CSV
	validationResult.encoding = string(constants.UTF8)
	validationResult.jurisdiction = constants.JURISDICTION

	verifyValidationResult(t, source, validationResult)

}
func TestValidateWindows1252EncodedCSVFile(t *testing.T) {
	source := "dex-csv-file-validation-test-temp/Windows1252Encoding.json"
	validationResult := expectedValidationResult{}
	validationResult.dataProducerID = constants.DATA_PRODUCER_ID
	validationResult.dataStreamID = constants.DATA_STREAM_ID
	validationResult.dataStreamRoute = constants.DATA_STREAM_ROUTE
	validationResult.delimiter = constants.CSV
	validationResult.encoding = string(constants.WINDOWS1252)
	validationResult.jurisdiction = constants.JURISDICTION

	verifyValidationResult(t, source, validationResult)
}
func TestValidateISO8859_1EncodedCSVFile(t *testing.T) {
	source := "dex-csv-file-validation-test-temp/ISO8859_1Encoding.json"
	validationResult := expectedValidationResult{}
	validationResult.dataProducerID = constants.DATA_PRODUCER_ID
	validationResult.dataStreamID = constants.DATA_STREAM_ID
	validationResult.dataStreamRoute = constants.DATA_STREAM_ROUTE
	validationResult.delimiter = constants.CSV
	validationResult.encoding = string(constants.ISO8859_1)
	validationResult.jurisdiction = constants.JURISDICTION

	verifyValidationResult(t, source, validationResult)

}
func TestValidateCSVFileHeader(t *testing.T) {
	//TO DO
	_ = "dex-csv-file-validation-test-temp/TestCSVHeader.json"

}
