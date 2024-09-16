package main

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/processor"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"

	//"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"
)

func main() {
	logger := sloger.With(constants.PACKAGE, constants.MAIN)
	logger.Info(constants.APPLICATION_STARTED)

	inputParams := models.FileValidateInputParams{
		FileURL:            "data/file-with-headers-100-rows.csv",
		HasHeader:          false,
		Destination:        "storage",
		ValidationCallback: processor.ProcessFileValidationResult,
	}

	file.Validate(inputParams)
}
