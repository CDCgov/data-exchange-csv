package main

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/cli"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"
)

func main() {
	logger := sloger.With(constants.PACKAGE, constants.MAIN)
	logger.Info(constants.APPLICATION_STARTED)

	validationInputParams := cli.ParseFlags()
	rootDir := validationInputParams.Destination

	err := utils.SetupEnvToStoreResults(rootDir)

	if err == nil {
		file.Validate(validationInputParams)
	} else {
		logger.Error("Unable to setup environment to store validation results")
		return
	}
}
