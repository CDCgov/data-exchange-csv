package file

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"reflect"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/detector"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
	"github.com/google/uuid"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type configIdentifier struct {
	DataStreamID    string   `json:"data_stream_id"`
	DataStreamRoute string   `json:"data_stream_route"`
	Header          []string `json:"header"`
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
type ConfigValidationResult struct {
	Error  *Error                 `json:"error"`
	Status string                 `json:"status"`
	Header HeaderValidationResult `json:"header_validation_result"`
}
type HeaderValidationResult struct {
	Status string   `json:"status"`
	Error  *Error   `json:"error"`
	Header []string `json:"header"`
}
type FileValidationParams struct {
	FileUUID     uuid.UUID              `json:"file_uuid"`
	ReceivedFile string                 `json:"received_filename"`
	Encoding     constants.EncodingType `json:"detected_encoding"`
	Delimiter    string                 `json:"detected_delimiter"`
	Header       []string               `json:"header"`
}

type FileValidationResult struct {
	ReceivedFile string                   `json:"received_filename"`
	Encoding     constants.EncodingType   `json:"encoding"`
	FileUUID     uuid.UUID                `json:"uuid"`
	Size         int64                    `json:"size"`
	Delimiter    string                   `json:"delimiter"`
	Error        *Error                   `json:"error"`
	Status       string                   `json:"status"`
	Metadata     MetadataValidationResult `json:"metadata_validation_result"`
	Config       ConfigValidationResult   `json:"config_validation_result"`
}

func Validate(eventMetadataFileURL string) FileValidationResult {
	metadataValidationResult := validateMetadataFile(eventMetadataFileURL)

	fileValidationResult := validateFile(metadataValidationResult.ReceivedFile)
	fileValidationResult.Metadata = metadataValidationResult

	configValidationResult := validateConfigFile(metadataValidationResult.DataStreamID, metadataValidationResult.ReceivedFile)
	fileValidationResult.Config = configValidationResult

	if metadataValidationResult.Status != constants.STATUS_SUCCESS {
		fileValidationResult.Status = constants.STATUS_FAILED
		return fileValidationResult
	}
	fileValidationResult.Status = constants.STATUS_SUCCESS

	return fileValidationResult
}

func validateMetadataFile(fileMetadata string) MetadataValidationResult {
	validationResult := MetadataValidationResult{}

	file, err := os.Open(fileMetadata)
	if err != nil {
		validationResult.Error = &Error{Message: constants.FILE_OPEN_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}
	defer file.Close()

	fields, err := io.ReadAll(file)
	if err != nil {
		validationResult.Error = &Error{Message: constants.FILE_READ_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	var metadataMap map[string]string
	err = json.Unmarshal(fields, &metadataMap)
	if err != nil {
		validationResult.Error = &Error{Message: constants.ERROR_UNMARSHALING_JSON, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}
	validationResult.Jurisdiction = metadataMap[constants.JURISDICTION]

	if filename, ok := metadataMap[constants.RECEIVED_FILENAME]; ok {
		validationResult.ReceivedFile = filename
	} else {
		validationResult.Error = &Error{Message: constants.FILE_MISSING_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	validationResult.DataStreamID = metadataMap[constants.DATA_STREAM_ID]
	validationResult.DataProducerID = metadataMap[constants.DATA_PRODUCER_ID]
	validationResult.Version = metadataMap[constants.VERSION]
	validationResult.DataStreamRoute = metadataMap[constants.DATA_STREAM_ROUTE]
	validationResult.Status = constants.STATUS_SUCCESS
	validationResult.SenderID = metadataMap[constants.SENDER_ID]

	return validationResult
}

func validateConfigFile(dataStreamId string, received_file string) ConfigValidationResult {

	validationResult := ConfigValidationResult{}

	headerValidationResult := HeaderValidationResult{}

	var identifiers []configIdentifier

	fields, err := os.ReadFile(constants.CONFIG_FILE)
	if err != nil {
		validationResult.Error = &Error{Message: constants.FILE_READ_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	err = json.Unmarshal(fields, &identifiers)
	if err != nil {
		validationResult.Error = &Error{Message: constants.ERROR_UNMARSHALING_JSON, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	var expectedHeader []string
	for _, identifier := range identifiers {
		if identifier.DataStreamID == dataStreamId && len(identifier.Header) > 0 {
			expectedHeader = identifier.Header
		}
	}

	//open file to get actual header to compare with expected one
	file, err := os.Open(received_file)
	if err != nil {
		validationResult.Error = &Error{Message: constants.FILE_OPEN_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			validationResult.Error = &Error{Message: constants.FILE_CLOSE_ERROR, Code: 13}
			validationResult.Status = constants.STATUS_FAILED
		}
	}(file)

	//if byte order mark is detected, move pointer 3 bytes to exclude it
	hasBom, err := detector.DetectBOM(file)
	if err != nil {
		validationResult.Status = constants.BOM_NOT_DETECTED_ERROR
	}

	if hasBom {
		file.Seek(3, 0)
	}

	//get the actual header from the file, compare with expected and update status
	reader := csv.NewReader(file)
	actualHeader, err := reader.Read()
	if err != nil {
		validationResult.Status = constants.CSV_READER_ERROR
	}
	if reflect.DeepEqual(expectedHeader, actualHeader) {
		headerValidationResult.Status = constants.STATUS_SUCCESS
		headerValidationResult.Header = actualHeader
	} else {
		headerValidationResult.Status = constants.STATUS_FAILED
		validationResult.Error = &Error{Message: constants.ERR_HEADER_VALIDATION, Code: 13}
	}

	validationResult.Header = headerValidationResult

	return validationResult
}

func validateFile(receivedFile string) FileValidationResult {
	validationResult := FileValidationResult{}

	fileUUID := uuid.New()
	validationResult.FileUUID = fileUUID
	validationResult.ReceivedFile = receivedFile

	file, err := os.Open(receivedFile)
	if err != nil {
		validationResult.Error = &Error{Message: constants.FILE_OPEN_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			validationResult.Error = &Error{Message: constants.FILE_CLOSE_ERROR, Code: 13}
			validationResult.Status = constants.STATUS_FAILED
			return
		}
	}(file)
	hasBOM, err := detector.DetectBOM(file)

	if err != nil {
		validationResult.Error = &Error{Message: constants.BOM_NOT_DETECTED_ERROR, Code: 13}
	}

	data, err := utils.ReadFileRandomly(file)
	if err != nil {
		validationResult.Error = &Error{Message: err.Error(), Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult

	}

	detectedDelimiter := detector.DetectDelimiter(data)
	validationResult.Delimiter = constants.DelimiterCharacters[detectedDelimiter]

	if validationResult.Delimiter == constants.DelimiterCharacters[0] {
		validationResult.Error = &Error{Message: constants.UNSUPPORTED_DELIMITER_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	if hasBOM {
		validationResult.Encoding = constants.UTF8_BOM
	} else {
		detectedEncoding := detector.DetectEncoding(data)
		if detectedEncoding == constants.UNDEF {
			validationResult.Error = &Error{Message: constants.UNSUPPORTED_ENCODING_ERROR, Code: 13}
			validationResult.Status = constants.STATUS_FAILED
			return validationResult
		}
		validationResult.Encoding = detectedEncoding
	}

	validationResult.Status = constants.STATUS_SUCCESS

	return validationResult
}
