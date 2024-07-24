package processor

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
)

func ProcessFileValidationResult(validationResult file.FileValidationResult) {
	//if file is unreadable send it to DLQ
	if validationResult.Status == constants.STATUS_SUCCESS {
		SendEventsToRouting(validationResult, constants.FILE_REPORTS)
	} else {
		SendEventsToDLQ(validationResult, constants.DEAD_LETTER_QUEUE)
	}

}

func SendEventsToDLQ(result interface{}, destination string) {
	//This is temporary function that copies result  into destination
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

func SendEventsToRouting(result interface{}, destination string) {
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
