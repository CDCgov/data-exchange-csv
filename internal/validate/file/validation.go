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
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
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
	result := FileValidationResult{}
	return result
}
