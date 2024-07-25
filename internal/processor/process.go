package processor

/*
This is temporary and will be re-written for processing events to DLQ and routing service
*/
import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
)

func ProcessFileValidationResult(validationResult models.FileValidationResult) {
	if validationResult.Status == constants.STATUS_SUCCESS {
		SendEventsToRouting(validationResult, constants.FILE_REPORTS)
	} else {
		SendEventsToDLQ(validationResult, constants.DEAD_LETTER_QUEUE)
	}

}

func SendEventsToDLQ(result interface{}, destination string) {
	//This is temporary function that copies event  into destination
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
