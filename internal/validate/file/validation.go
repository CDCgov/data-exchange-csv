/*
Package file provides functionalities for validating and working with files in the context of file-level validation.
This package is designed to handle CSV and TSV files, within the bounded context of file-level validation.
Key components contributing to the overall validation of a file include:
  - HeaderValidation - if provided in the config file
  - Detecting Delimiter
  - Detecting Encoding type
*/
package file

import (
	"fmt"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
)

type FileValidationResult struct {
	Name      string                 `json:"name"`
	Source    string                 `json:"source"`
	Size      int64                  `json:"size"`
	FileUUID  constants.UUID         `json:"uuid"`
	Error     error                  `json:"error"`
	Encoding  constants.EncodingType `json:"encoding"`
	Delimiter string                 `json:"delimiter"`
	Status    string                 `json:"status"` // or object
}

func Validate(filePath string) FileValidationResult {
	file, err := os.Open(filePath)
	if err != nil {
		return FileValidationResult{
			Source: filePath,
			Error:  err,
		}
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(constants.FILE_CLOSE_ERROR)
			return
		}
	}()

	result := FileValidationResult{
		Name:   file.Name(),
		Source: filePath,
	}

	// Detect BOM
	hasBOM, err := utils.DetectBOM(file)
	if err != nil {
		result.Error = err
		return result
	}

	// Get a random sample of the file data
	sampleData, err := utils.GetRandomSample(file)
	if err != nil {
		result.Error = err
		return result
	}

	// Detect delimiter
	delimiterDetector := &utils.DelimiterDetector{}
	detectedDelimiter := delimiterDetector.DetectDelimiter(string(sampleData))
	result.Delimiter = constants.DelimiterCharacters[detectedDelimiter.Character]

	// Detect encoding
	if hasBOM {
		result.Encoding = constants.UTF8_BOM
	} else {
		result.Encoding = utils.DetectEncoding(sampleData)
	}

	return result
}
