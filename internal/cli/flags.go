package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
)

func ParseFlags() models.FileValidateInputParams {
	var receivedFile string
	var destination string
	var debug bool
	var logToFile bool
	var configFile string

	flag.StringVar(&receivedFile, "fileURL", "", "[Required] The path to the file that will be validated")
	flag.BoolVar(&debug, "debug", false, "[Optional] enable debug mode to generate debug-level logs")
	flag.BoolVar(&logToFile, "log-file", false, "[Optional] If true, logs will be written to logs/validation.json, default is stdout")
	flag.StringVar(&destination, "destination", "", "[Required] The URL to the folder where validation/transformation results will be stored")
	flag.StringVar(&configFile, "config", "", "[Optional] The URL to the config.json. If provided overrides auto-detection of encoding, and separator")

	flag.Parse()

	if receivedFile == "" || destination == "" {
		flag.Usage()
		os.Exit(1)
	}

	fileInputParams := models.FileValidateInputParams{
		ReceivedFile: receivedFile,
		Destination:  destination,
	}

	if debug {
		fileInputParams.HasHeader = true
	}

	if logToFile {
		fileInputParams.LogToFile = true
	}

	//add logic for config.json file
	if configFile != "" {
		configFields := readConfigFile(configFile)

		if value, exists := configFields["encoding"]; exists {
			fileInputParams.Encoding = constants.EncodingType(value.(string))
		}

		if value, exists := configFields["separator"]; exists {
			valueString := value.(string)
			fileInputParams.Separator = []rune(valueString)[0]
		}

		if value, exists := configFields["hasHeader"]; exists && value.(bool) {
			fileInputParams.HasHeader = true
		}
	}
	return fileInputParams
}

func readConfigFile(configFile string) map[string]interface{} {
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println("Failed to open config file")
	}
	defer file.Close()

	fields, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read config file")
	}

	var configFields map[string]interface{}

	err = json.Unmarshal(fields, &configFields)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON from config file")
	}

	return configFields
}
