package cli

import (
	"encoding/json"
	"flag"
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

func TestParseFlagsWithConfigFile(t *testing.T) {
	testCaseNameValidConfigFIle := "Test Case with valid config.json file"
	tests := []struct {
		name           string
		configFields   map[string]interface{}
		expectedResult models.FileValidateInputParams
	}{
		{
			name: testCaseNameValidConfigFIle,
			configFields: map[string]interface{}{
				"encoding":  "UTF-8",
				"separator": ",",
				"hasHeader": true,
			},
			expectedResult: models.FileValidateInputParams{
				ReceivedFile: "testFile.csv",
				Destination:  "C://destination/folder",
				HasHeader:    true,
			},
		},
		{
			name:         "Test case with an empty config.json file",
			configFields: map[string]interface{}{},
			expectedResult: models.FileValidateInputParams{
				ReceivedFile: "testFile.csv",
				Destination:  "C://destination/folder",
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Create config.json file
			configFile, err := os.CreateTemp("", "config.json")
			if err != nil {
				t.Fatalf("Failed to create temp config file: %v", err)
			}
			defer os.Remove(configFile.Name())

			configAsJson, err := json.Marshal(testCase.configFields)
			if err != nil {
				t.Fatalf("%v", err)
			}

			if _, err := configFile.Write(configAsJson); err != nil {
				t.Fatalf("%s %v", constants.FILE_WRITE_ERROR, err)
			}

			if err := configFile.Close(); err != nil {
				t.Fatalf("%s %v", constants.FILE_CLOSE_ERROR, err)
			}

			// Reset flags before test run
			resetFlags()
			os.Args = []string{
				"cmd",
				"-fileURL", "testFile.csv",
				"-destination", "C://destination/folder",
				"-config", configFile.Name(),
			}

			actualResult := ParseFlags()

			if actualResult.Encoding != constants.UTF8 && testCase.name == testCaseNameValidConfigFIle {
				t.Errorf("Expected encoding is UTF-8, but got %s", actualResult.Encoding)
			}
			if actualResult.Separator != constants.COMMA && testCase.name == testCaseNameValidConfigFIle {
				t.Errorf("Expected separator is `,`, but got %s", string(actualResult.Separator))
			}
			if actualResult.HasHeader != testCase.expectedResult.HasHeader {
				t.Errorf("hasHeader is expected to be %v, but got %v", testCase.expectedResult.HasHeader, actualResult.HasHeader)
			}

			if testCase.expectedResult.Destination != actualResult.Destination {
				t.Errorf("Expected destination: %s, got: %s", testCase.expectedResult.Destination, actualResult.Destination)
			}
			if testCase.expectedResult.ReceivedFile != actualResult.ReceivedFile {
				t.Errorf("Expected file: %s, got: %s", testCase.expectedResult.ReceivedFile, actualResult.ReceivedFile)
			}
		})
	}
}

/*
Will come back to the test cases below
I had some trouble redirecting output from stdoud and capturing exit code
*/
func TestParseFlagsWithMissingRequiredFlags(t *testing.T)  {}
func TestParseFlagsWithNonExistentConfigFile(t *testing.T) {}
