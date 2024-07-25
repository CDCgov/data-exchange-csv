package main

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/processor"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/row"
)

func main() {
	event := "data/event_source.json"

	fileValidationResult := file.Validate(event)
	processor.ProcessFileValidationResult(fileValidationResult)

	if fileValidationResult.Status == constants.STATUS_SUCCESS {

		//valid file, proceed with row validation
		params := models.FileValidationParams{
			FileUUID:     fileValidationResult.FileUUID,
			ReceivedFile: fileValidationResult.ReceivedFile,
			Encoding:     fileValidationResult.Encoding,
			Delimiter:    fileValidationResult.Delimiter,
			Header:       fileValidationResult.Config.Header.Header,
		}

		row.Validate(params, processor.SendEventsToDLQ, processor.SendEventsToRouting)
	}

}
