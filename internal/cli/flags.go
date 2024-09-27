package cli

import (
	"flag"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/processor"
)

func ParseFlags() models.FileValidateInputParams {
	inputFile := flag.String("fileURL", "", "[Required] The path to the file that will be validated")
	debug := flag.Bool("debug", false, "[Optional] enable debug mode to generate debug-level logs")
	logToTheFile := flag.Bool("log-file", false, "[Optional] If true, logs will be written to logs/validation.json, default is stdout")
	destination := flag.String("destination", "", "[Required] The URL to the folder where validation/transformation results will be stored")
	configFile := flag.String("config", "", "[Optional] The URL to the config.json. If provided overrides auto-detection of encoding, and separator")
	flag.Parse()

	if *inputFile == "" || *destination == "" {
		flag.Usage()
		os.Exit(1)
	}
	//add logic for config.json file

	return models.FileValidateInputParams{
		ReceivedFile:       *inputFile,
		Destination:        *destination,
		ConfigFile:         *configFile,
		Debug:              *debug,
		LogToFile:          *logToTheFile,
		ValidationCallback: processor.ProcessFileValidationResult,
	}

}
