package main

import (
	"log/slog"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/processor"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/row"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"
)

func init() {
	logFile, err := os.OpenFile("logs/csv-validation.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		panic("Failed to Open log file:" + logFile.Name())
	}

	//create a new logger with
	logger := slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo}))

	//set the logger as default logger
	sloger.SetDefaultLogger(logger)

}
func main() {

	event := "data/event_source.json"
	logger := sloger.With("event", "main")

	logger.Info("Application Started")

	fileValidationResult := file.Validate(event)
	processor.ProcessFileValidationResult(fileValidationResult)

	if fileValidationResult.Status == constants.STATUS_SUCCESS {

		//valid file, proceed with row validation
		params := models.FileValidationParams{
			FileUUID:     fileValidationResult.FileUUID,
			ReceivedFile: fileValidationResult.ReceivedFile,
			Encoding:     fileValidationResult.Encoding,
			Delimiter:    fileValidationResult.Delimiter,
			Header:       fileValidationResult.Config.HeaderValidationResult.Header,
		}

		row.Validate(params, processor.SendEventsToDestination)
	}

}
