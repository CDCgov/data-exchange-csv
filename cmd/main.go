package main

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/cli"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/processor"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"
)

func main() {

	validationInputParams := cli.ParseFlags()

	logToFile := validationInputParams.LogToFile
	debug := validationInputParams.Debug

	sloger.InitLogger(logToFile, debug)

	logger := sloger.With(constants.PACKAGE, constants.MAIN)
	logger.Info(constants.APPLICATION_STARTED)

	rootDir := validationInputParams.Destination
	err := utils.SetupEnvironment(rootDir)

	if err == nil {
		fileValidationResult := file.Validate(validationInputParams)
		processor.StoreFileValidationResult(fileValidationResult)
	} else {
		logger.Error("Unable to setup environment to store validation results")
		return
	}
}
