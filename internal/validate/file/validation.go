package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/detector"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
	"github.com/google/uuid"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
type MetadataValidationResult struct {
	ReceivedFile    string `json:"received_filename"`
	Error           *Error `json:"error"`
	Status          string `json:"status"`
	Jurisdiction    string `json:"jurisdiction"`
	DataStreamID    string `json:"data_stream_id"`
	DataStreamRoute string `json:"data_stream_route"`
	SenderID        string `json:"sender_id"`
	DataProducerID  string `json:"data_producer_id"`
	Version         string `json:"version"`
}
type configValidationResult struct {
	//TODO
	Status   string `json:"status"`
	FileName string `json:"file_name"`
}

type fileValidationResult struct {
	ReceivedFile string                   `json:"received_filename"`
	Encoding     constants.EncodingType   `json:"encoding"`
	FileUUID     uuid.UUID                `json:"uuid"`
	Size         int64                    `json:"size"`
	Delimiter    string                   `json:"delimiter"`
	Error        *Error                   `json:"error"`
	Status       string                   `json:"status"` // or object?
	Metadata     MetadataValidationResult `json:"metadataValidationResult"`
	Config       configValidationResult   `json:"configValidationResult"`
}

func Validate(configFile string) fileValidationResult {
	metadataValidationResult := validateMetadataFile(configFile)
	if metadataValidationResult.Status != constants.STATUS_SUCCESS {
		CopyToDestination(metadataValidationResult, constants.DEAD_LETTER_QUEUE)
	}
	configValidationResult := validateConfig(configFile)
	if configValidationResult.Status != constants.STATUS_SUCCESS {
		CopyToDestination(configValidationResult, constants.DEAD_LETTER_QUEUE)
	}

	fileValidationResult := validateFile(metadataValidationResult.ReceivedFile)

	fileValidationResult.Config = configValidationResult
	fileValidationResult.Metadata = metadataValidationResult

	if fileValidationResult.Status != constants.STATUS_SUCCESS {
		CopyToDestination(fileValidationResult, constants.DEAD_LETTER_QUEUE)
	}

	CopyToDestination(fileValidationResult, constants.FILE_REPORTS)
	return fileValidationResult
}

func validateMetadataFile(fileMetadata string) MetadataValidationResult {
	validationResult := MetadataValidationResult{}

	file, err := os.Open(fileMetadata)
	if err != nil {
		validationResult.Error = &Error{Message: constants.FILE_OPEN_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
	}
	defer file.Close()

	fields, err := io.ReadAll(file)
	if err != nil {
		validationResult.Error = &Error{Message: err.Error(), Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
	}

	var metadataMap map[string]string
	err = json.Unmarshal(fields, &metadataMap)
	if err != nil {
		validationResult.Error = &Error{Message: err.Error(), Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)

	}
	validationResult.Jurisdiction = metadataMap[constants.JURISDICTION]
	if filename, ok := metadataMap[constants.RECEIVED_FILENAME]; ok {
		validationResult.ReceivedFile = filename
	} else {
		validationResult.Error = &Error{Message: constants.RECEIVED_FILENAME, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
	}

	validationResult.DataStreamID = metadataMap[constants.DATA_STREAM_ID]
	validationResult.DataProducerID = metadataMap[constants.DATA_PRODUCER_ID]
	validationResult.Version = metadataMap[constants.VERSION]
	validationResult.DataStreamRoute = metadataMap[constants.DATA_STREAM_ROUTE]
	validationResult.Status = constants.STATUS_SUCCESS
	validationResult.SenderID = metadataMap[constants.CSV_SENDER_ID]

	return validationResult
}

func validateConfig(configFile string) configValidationResult {
	//TODO
	validationResult := configValidationResult{}
	validationResult.Status = constants.STATUS_SUCCESS
	validationResult.FileName = configFile

	return validationResult
}

func validateFile(fileURI string) fileValidationResult {
	validationResult := fileValidationResult{}

	fileUUID := uuid.New()
	validationResult.FileUUID = fileUUID
	validationResult.ReceivedFile = fileURI

	file, err := os.Open(fileURI)
	if err != nil {
		validationResult.Error = &Error{Message: err.Error(), Code: 13}
		CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
		return validationResult
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			validationResult.Error = &Error{Message: err.Error(), Code: 13}
			CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
			return
		}
	}(file)
	hasBOM, err := detector.DetectBOM(file)

	if err != nil {
		validationResult.Error = &Error{Message: err.Error(), Code: 13}
		CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
		return validationResult
	}

	data, err := utils.ReadFileRandomly(file)
	if err != nil {
		validationResult.Error = &Error{Message: err.Error(), Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
		return validationResult

	}

	detectedDelimiter := detector.DetectDelimiter(data)
	validationResult.Delimiter = constants.DelimiterCharacters[detectedDelimiter]

	if validationResult.Delimiter == constants.DelimiterCharacters[0] {
		validationResult.Error = &Error{Message: constants.UNSUPPORTED_DELIMITER_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
		return validationResult
	}

	if hasBOM {
		validationResult.Encoding = constants.UTF8_BOM
	} else {
		validationResult.Encoding = detector.DetectEncoding(data)
	}

	validationResult.Status = constants.STATUS_SUCCESS

	return validationResult
}
func CopyToDestination(result interface{}, destination string) error {
	//This is temporary function that copies result  into destination
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
