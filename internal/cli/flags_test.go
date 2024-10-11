package cli

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
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

	// redirect stdout to buffer
	var buf bytes.Buffer
	flag.CommandLine.SetOutput(&buf)

	//reset flags before test run
	resetFlags()

	//destination that is required is missing
	os.Args = []string{
		"cmd",
		"-fileURL", "testFile.csv",
	}

	ParseFlags()

	output := buf.String()
	fmt.Println(output)
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
	//Create config.json file
	configFile, err := os.CreateTemp("", "config.json")

	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}
	defer os.Remove(configFile.Name())

	configFields := map[string]interface{}{
		"encoding":  "UTF-8",
		"separator": ",",
		"hasHeader": true,
	}

	configAsJson, err := json.Marshal(configFields)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if _, err := configFile.Write(configAsJson); err != nil {
		t.Fatalf("%s %v", constants.FILE_WRITE_ERROR, err)
	}

	if err := configFile.Close(); err != nil {
		t.Fatalf("%s %v", constants.FILE_CLOSE_ERROR, err)
	}

	//reset flags before test run
	resetFlags()
	os.Args = []string{
		"cmd",
		"-fileURL", "testFile.csv",
		"-destination", "C://destination/folder",
		"-config", configFile.Name(),
	}
	actualResult := ParseFlags()

	if actualResult.Encoding != constants.UTF8 {
		t.Errorf("Expected encoding is UTF-8, but got %s", actualResult.Encoding)
	}
	if actualResult.Separator != constants.COMMA {
		t.Errorf("Expected separator is `,` , but got %s", string(actualResult.Separator))
	}
	if !actualResult.HasHeader {
		t.Errorf("hasHeader is expected to be true")
	}
}

func TestParseFlagsWithEmptyConfigFile(t *testing.T) {
	//reset flags before test run
	resetFlags()
}

func TestParseFlagsWithNonExistentConfigFile(t *testing.T) {
	//reset flags before test run
	resetFlags()
}
