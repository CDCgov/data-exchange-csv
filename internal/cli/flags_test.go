package cli

import (
	"flag"
	"os"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
)

func resetFlags() {
	//We need to reinitialize the command-line flags for each test case
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestParseFlagsWithRequiredFlags(t *testing.T) {
	//reset flags before test run
	resetFlags()
	os.Args = []string{
		"cmd",
		"-fileURL", "testFile.csv",
		"-destination", "C://destination/folder",
	}

	expectedResult := models.FileValidateInputParams{
		ReceivedFile: "testFile.csv",
		Destination:  "C://destination/folder",
	}

	actualResult := ParseFlags()

	if actualResult.ReceivedFile != expectedResult.ReceivedFile {
		t.Errorf("Expected file: %s and got: %s", expectedResult.ReceivedFile, actualResult.ReceivedFile)
	}

	if actualResult.Destination != expectedResult.Destination {
		t.Errorf("Expected file: %s and got: %s", expectedResult.Destination, actualResult.Destination)
	}
}

func TestParseFlagsWithMissingRequiredFlags(t *testing.T) {
	//reset flags before test run
	//TO DO
	resetFlags()
	os.Args = []string{
		"cmd",
		"-fileURL", "testFile.csv",
	}

}
func TestParseFlagsWithOptionalFlags(t *testing.T) {
	//reset flags before test run
	resetFlags()
	os.Args = []string{
		"cmd",
		"-fileURL", "testFile.csv",
		"-destination", "C://destination/folder",
		"-debug=true",
		"-log-file=true",
	}

	expectedResult := models.FileValidateInputParams{
		ReceivedFile: "testFile.csv",
		Destination:  "C://destination/folder",
		Debug:        true,
		LogToFile:    true,
	}

	actualResult := ParseFlags()
	if actualResult.ReceivedFile != expectedResult.ReceivedFile {
		t.Errorf("Expected file: %s and got: %s", expectedResult.ReceivedFile, actualResult.ReceivedFile)
	}

	if actualResult.Destination != expectedResult.Destination {
		t.Errorf("Expected file: %s and got: %s", expectedResult.Destination, actualResult.Destination)
	}

	if actualResult.Debug != expectedResult.Debug {
		t.Errorf("Expected file: %s and got: %s", expectedResult.ReceivedFile, actualResult.ReceivedFile)
	}

	if actualResult.LogToFile != expectedResult.LogToFile {
		t.Errorf("Expected file: %s and got: %s", expectedResult.ReceivedFile, actualResult.ReceivedFile)

	}
}

func TestParseFlagsWithOptionalConfigFile(t *testing.T) {
	//reset flags before test run
	resetFlags()
}

func TestParseFlagsWithEmptyConfigFile(t *testing.T) {
	//reset flags before test run
	resetFlags()
}

func TestParseFlagsWithNonExistentConfigFile(t *testing.T) {
	//reset flags before test run
	resetFlags()
}
