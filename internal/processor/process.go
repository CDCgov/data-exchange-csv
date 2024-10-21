package processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
)

const (
	JSON_EXTENSION              = ".json"
	FILE_VALIDATION_REPORT_NAME = "%s_%s%s"
)

func StoreFileValidationResult(validationResult models.FileValidationResult) {
	jsonData := structToJson(validationResult)
	var filename string

	if validationResult.Status == constants.STATUS_SUCCESS {
		filename = fmt.Sprintf(FILE_VALIDATION_REPORT_NAME, validationResult.FileUUID, constants.STATUS_VALID, JSON_EXTENSION)
	} else {
		filename = fmt.Sprintf(FILE_VALIDATION_REPORT_NAME, validationResult.FileUUID, constants.STATUS_INVALID, JSON_EXTENSION)
	}
	StoreResult(jsonData, validationResult.Destination, filename)
}

func structToJson(result models.FileValidationResult) []byte {
	jsonData, err := json.Marshal(result)
	if err != nil {
		slog.Error(constants.ERROR_CONVERTING_STRUCT_TO_JSON)
	}
	return jsonData

}
func StoreResult(jsonString interface{}, destination, filename string) {

	destinationPath := filepath.Join(destination+"/file/", filename)
	destFile, err := os.OpenFile(destinationPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		slog.Error(constants.FILE_OPEN_ERROR)
	}
	defer destFile.Close()

	switch data := jsonString.(type) {
	case []byte:
		if _, err := destFile.Write(data); err != nil {
			slog.Error(constants.FILE_WRITE_ERROR)
		}
	case []string:
		strData := strings.Join(data, "\n")
		if _, err := destFile.WriteString(strData); err != nil {
			slog.Error(constants.FILE_WRITE_ERROR)
		}
	}
}

func OnValidateAndTransformRow(params models.RowCallbackParams) error {

	validationPath := fmt.Sprintf("%s/row/validation_result_%s.json", params.Destination, params.FileUUID)
	transformationPath := fmt.Sprintf("%s/row/transformation_result_%s.json", params.Destination, params.FileUUID)

	// Open `validation_result.json` file in append mode to write row validation results
	fileValidation, err := os.OpenFile(validationPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open or create validation_result.json %s: %w", validationPath, err)
	}

	validationWriter := bufio.NewWriter(fileValidation)
	// Write opening bracket if it's the first row
	if params.IsFirst && params.ValidationResult != nil {
		validationWriter.WriteString("[")
	}

	// Append the validation result if not nil
	if params.ValidationResult != nil {
		if !params.IsFirst {
			validationWriter.WriteString(",")
		}
		validationWriter.WriteString(params.ValidationResult.(string))
	}

	var fileTransformation *os.File
	var transformationWriter *bufio.Writer
	if params.Transform {
		// Open `transformation_result.json` file in append mode to write row transformation results
		fileTransformation, err = os.OpenFile(transformationPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open or create transformation_result.json %s: %w", transformationPath, err)
		}
		transformationWriter = bufio.NewWriter(fileTransformation)

		if params.IsFirst && params.TransformationResult != nil {
			transformationWriter.WriteString("[")
		}
		// Append the transformation result if available
		if params.TransformationResult != nil {
			if !params.IsFirst {
				transformationWriter.WriteString(",")
			}
			transformationWriter.WriteString(params.TransformationResult.(string))
		}
	}

	// Write closing bracket and flush the buffers if it's the last row
	if params.IsLast {
		validationWriter.WriteString("]")
		validationWriter.Flush()
		fileValidation.Close()
		if params.Transform {
			transformationWriter.WriteString("]")

			transformationWriter.Flush()

			fileTransformation.Close()
		}
		return nil
	}
	// Flush buffers to ensure the data is written
	validationWriter.Flush()
	if params.Transform {
		transformationWriter.Flush()
	}

	return nil
}
