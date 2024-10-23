package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
)

const tempDirectory = "dex-csv-file-validation-test-temp"

func verifyValidationResult(t *testing.T, fileValidationInputParams models.FileValidateInputParams, expectedResult models.FileValidationResult) {
	t.Helper()

	validationResult := Validate(fileValidationInputParams)
	assertEqual(t, "encoding", string(expectedResult.Encoding), string(validationResult.Encoding))
	assertEqual(t, "delimiter", string(expectedResult.Delimiter), string(validationResult.Delimiter))
	assertEqual(t, "size", expectedResult.SizeInBytes, validationResult.SizeInBytes)
	assertEqual(t, "header", expectedResult.HasHeader, validationResult.HasHeader)
	assertEqual(t, "status", expectedResult.Status, validationResult.Status)

}

func assertEqual(t *testing.T, field string, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %s: %s, but got: %s", field, expected, actual)
	}
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	err := os.Mkdir(tempDirectory, 0755)
	if err != nil {
		tb.Fatalf("%s: %v", constants.DIRECTORY_CREATE_ERROR, err)
	}

	testFilesWithContent := map[string]string{
		"UTF8Encoding.csv":        "Hello, World! This is US-ASCII.\nLine 2, More text.",
		"UTF8BomEncoding.csv":     "\xEF\xBB\xBFName,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok,cmallok@example.com",
		"USASCIIEncoding.csv":     "Chris~Wilson,DevOps engineer, ensures CI/CD pipelines and *NIX server maintenance.",
		"Windows1252Encoding.csv": "L'éƒté en France,München Äpfel\nJosé DíažŸ,François Dupont",
		"ISO8859_1Encoding.csv":   "José Dí^az,Software engineer, working on CSV & Golang.",
		"UTF8BomEncoding.tsv":     "\xEF\xBB\xBFName\tEmail\nJane Doe\tjohndoe@example.com\nJane Smith\tjanesmith@example.com\nChris Mallok\tcmallok@example.com",
		"emptyFileTest.csv":       "",
	}

	for file, content := range testFilesWithContent {
		tempFile := filepath.Join(tempDirectory, file)
		err := os.WriteFile(tempFile, []byte(content), 0644)
		if err != nil {
			tb.Fatalf("%s %s: %v", constants.FILE_WRITE_ERROR, file, err)
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
	expectedValidationResult := models.FileValidationResult{
		Delimiter:   constants.COMMA,
		Encoding:    constants.UTF8,
		SizeInBytes: 50,
		HasHeader:   false,
		Status:      constants.STATUS_SUCCESS,
	}
	fileValidationInputParams := models.FileValidateInputParams{
		ReceivedFile: filepath.Join(tempDirectory, "UTF8Encoding.csv"),
		Separator:    constants.COMMA,
		HasHeader:    false,
	}
	verifyValidationResult(t, fileValidationInputParams, expectedValidationResult)
}

func TestValidateUTF8BomEncodedCSVFile(t *testing.T) {
	expectedValidationResult := models.FileValidationResult{
		Delimiter:   constants.COMMA,
		Encoding:    constants.UTF8_BOM,
		SizeInBytes: 108,
		HasHeader:   false,
		Status:      constants.STATUS_SUCCESS,
	}
	fileValidationInputParams := models.FileValidateInputParams{
		ReceivedFile: filepath.Join(tempDirectory, "UTF8BomEncoding.csv"),
		Encoding:     constants.UTF8_BOM,
		HasHeader:    false,
	}
	verifyValidationResult(t, fileValidationInputParams, expectedValidationResult)
}

func TestValidateUSASCIIEncodedCSVFile(t *testing.T) {
	expectedValidationResult := models.FileValidationResult{
		Delimiter:   constants.COMMA,
		Encoding:    constants.UTF8,
		SizeInBytes: 82,
		HasHeader:   false,
		Status:      constants.STATUS_SUCCESS,
	}
	fileValidationInputParams := models.FileValidateInputParams{
		ReceivedFile: filepath.Join(tempDirectory, "USASCIIEncoding.csv"),
		Separator:    constants.COMMA,
		HasHeader:    false,
	}
	verifyValidationResult(t, fileValidationInputParams, expectedValidationResult)
}

func TestValidateWindows1252EncodedCSVFile(t *testing.T) {
	expectedValidationResult := models.FileValidationResult{
		Delimiter:   constants.COMMA,
		Encoding:    constants.WINDOWS1252,
		SizeInBytes: 67,
		HasHeader:   false,
		Status:      constants.STATUS_SUCCESS,
	}
	fileValidationInputParams := models.FileValidateInputParams{
		ReceivedFile: filepath.Join(tempDirectory, "Windows1252Encoding.csv"),
		Separator:    constants.COMMA,
		HasHeader:    false,
	}
	verifyValidationResult(t, fileValidationInputParams, expectedValidationResult)
}

func TestValidateISO8859_1EncodedCSVFile(t *testing.T) {
	expectedValidationResult := models.FileValidationResult{
		Delimiter:   constants.COMMA,
		Encoding:    constants.ISO8859_1,
		SizeInBytes: 56,
		HasHeader:   false,
		Status:      constants.STATUS_SUCCESS,
	}
	fileValidationInputParams := models.FileValidateInputParams{
		ReceivedFile: filepath.Join(tempDirectory, "ISO8859_1Encoding.csv"),
		Separator:    constants.COMMA,
		HasHeader:    false,
	}

	verifyValidationResult(t, fileValidationInputParams, expectedValidationResult)
}

func TestValidateUTF8BomEncodedTSVFile(t *testing.T) {
	expectedValidationResult := models.FileValidationResult{
		Delimiter:   constants.TAB,
		Encoding:    constants.UTF8_BOM,
		SizeInBytes: 108,
		HasHeader:   false,
		Status:      constants.STATUS_SUCCESS,
	}
	fileValidationInputParams := models.FileValidateInputParams{
		ReceivedFile: filepath.Join(tempDirectory, "UTF8BomEncoding.tsv"),
		HasHeader:    false,
	}
	verifyValidationResult(t, fileValidationInputParams, expectedValidationResult)
}

func TestEmptyFile(t *testing.T) {
	expectedValidationResult := models.FileValidationResult{
		SizeInBytes: 0,
		HasHeader:   false,
		Status:      constants.STATUS_FAILED,
	}
	fileValidationInputParams := models.FileValidateInputParams{
		ReceivedFile: filepath.Join(tempDirectory, "emptyFileTest.csv"),
		HasHeader:    false,
	}
	verifyValidationResult(t, fileValidationInputParams, expectedValidationResult)
}
