package processor

/*
This is temporary and currently being re-designed
*/
import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

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
	storeResult(jsonData, filename)

}
func structToJson(result models.FileValidationResult) []byte {
	jsonData, err := json.Marshal(result)
	if err != nil {
		slog.Error(constants.ERROR_CONVERTING_STRUCT_TO_JSON)
	}
	return jsonData

}
func storeResult(jsonString []byte, filename string) {
	destinationPath := filepath.Join(constants.FILE_REPORTS, filename)
	destFile, err := os.OpenFile(destinationPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		slog.Error(constants.FILE_OPEN_ERROR)
	}
	defer destFile.Close()

	if _, err := destFile.Write(jsonString); err != nil {
		slog.Error(constants.FILE_WRITE_ERROR)
	}
}

func SendEventsToDestination(result interface{}, destination string) {
	jsonContent, err := json.Marshal(result)
	if err != nil {
		slog.Error(constants.ERROR_CONVERTING_STRUCT_TO_JSON)
	}
	// Create or open destination file in append mode
	destFilePath := filepath.Join(destination, "output.json")
	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		slog.Error("failed to create/open destination file")
	}
	defer destFile.Close()

	// Write JSON content to the destination file
	if _, err := destFile.Write(jsonContent); err != nil {
		slog.Error("failed to write to destination file")
	}
}
