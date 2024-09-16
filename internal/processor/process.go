package processor

/*
This is temporary and will be re-written for processing events to DLQ and routing service
*/

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
)

func ProcessFileValidationResult(validationResult models.FileValidationResult) {
	jsonData := structToJson(validationResult)
	var filename string

	if validationResult.Status == constants.STATUS_SUCCESS {
		filename = fmt.Sprintf(constants.FILE_VALIDATION_REPORT_NAME, validationResult.FileUUID, constants.STATUS_VALID, constants.JSON_EXTENSION)
	} else {
		filename = fmt.Sprintf(constants.FILE_VALIDATION_REPORT_NAME, validationResult.FileUUID, constants.STATUS_INVALID, constants.JSON_EXTENSION)
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
	destinationPath := filepath.Join(destination, filename)
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
