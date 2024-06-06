package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/detect"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
	"github.com/google/uuid"
)

type ValidationResult struct {
	ReceivedFile    string                 `json:"received_filename"`
	Source          string                 `json:"source"`
	Size            int64                  `json:"size"`
	FileUUID        uuid.UUID              `json:"uuid"`
	Error           *Error                 `json:"error"`
	Encoding        constants.EncodingType `json:"encoding"`
	Delimiter       string                 `json:"delimiter"`
	Status          string                 `json:"status"` // or object?
	Jurisdiction    string                 `json:"jurisdiction"`
	DataStreamID    string                 `json:"data_stream_id"`
	DataStreamRoute string                 `json:"data_stream_route"`
	SenderID        string                 `json:"sender_id"`
	DataProducerID  string                 `json:"data_producer_id"`
	Version         string                 `json:"version"`
}
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (vr *ValidationResult) Validate(configFile string) {

	vr.FileUUID = uuid.New()
	vr.processMetadataFields(configFile)

	file, err := os.Open(vr.ReceivedFile)
	if err != nil {
		vr.Error = &Error{Message: err.Error(), Code: 13}
		copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
		return
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			vr.Error = &Error{Message: err.Error(), Code: 13}
			copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
			return
		}
	}(file)

	data, err := utils.ReadFileRandomly(file)
	if err != nil {
		vr.Error = &Error{Message: err.Error(), Code: 13}
		copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
		return
	}

	detectedDelimiter := detect.Delimiter(string(data))
	vr.Delimiter = constants.DelimiterCharacters[detectedDelimiter]

	if vr.Delimiter == constants.DelimiterCharacters[0] {
		vr.Error = &Error{Message: constants.UNSUPPORTED_DELIMITER_ERROR, Code: 13}
		copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
	}

	hasBOM, err := detect.BOM(file)

	if err != nil {
		vr.Error = &Error{Message: err.Error(), Code: 13}
		copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
		return
	}
	if hasBOM {
		vr.Encoding = constants.UTF8_BOM
		copyToDestination(vr, constants.FILE_REPORTS)
		return
	}

	detectedEncoding := detect.Encoding(data)
	vr.Encoding = detectedEncoding
	copyToDestination(vr, constants.FILE_REPORTS)
}

func copyToDestination(result *ValidationResult, destination string) error {
	//This is temporary function that copies result  into destination
	// Serialize the struct to JSON
	jsonContent, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to marshal struct to JSON: %s", err)
	}

	// Create or open destination file in append mode
	destFilePath := filepath.Join(destination, "output.json")
	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to create/open destination file: %s", err)
	}
	defer destFile.Close()

	// Write JSON content to the destination file
	if _, err := destFile.Write(jsonContent); err != nil {
		return fmt.Errorf("failed to write to destination file: %s", err)
	}

	return nil
}
