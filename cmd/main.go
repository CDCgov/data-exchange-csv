package main

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/processor"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/row"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"
)

func main() {
	logger := sloger.With(constants.PACKAGE, constants.MAIN)
	logger.Info(constants.APPLICATION_STARTED)

	event := "data/event_source.json"

	fileValidationResult := file.Validate(event)
	processor.ProcessFileValidationResult(fileValidationResult)

	if fileValidationResult.Status == constants.STATUS_SUCCESS {
		params := models.FileValidationParams{
			FileUUID:     fileValidationResult.FileUUID,
			ReceivedFile: fileValidationResult.ReceivedFile,
			Encoding:     fileValidationResult.Encoding,
			Delimiter:    fileValidationResult.Delimiter,
			Header:       fileValidationResult.Config.HeaderValidationResult.Header,
		}

		logger.Info(constants.MSG_FILE_VALIDATION_SUCCESS)
		row.Validate(params, processor.SendEventsToDestination)
	}

}
