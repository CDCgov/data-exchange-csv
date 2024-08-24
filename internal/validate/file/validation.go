package file

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"reflect"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/detector"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
	"github.com/google/uuid"
)

func Validate(eventMetadataFileURL string) models.FileValidationResult {
	//initialize logger from sloger package
	logger := sloger.With(constants.PACKAGE, constants.FILE)
	logger.Info(constants.MSG_FILE_VALIDATION_BEGIN)

	metadataValidationResult := validateMetadataFile(eventMetadataFileURL)

	fileValidationResult := validateFile(metadataValidationResult.ReceivedFile)
	fileValidationResult.Metadata = metadataValidationResult

	configValidationResult := validateConfigFile(metadataValidationResult.DataStreamID, metadataValidationResult.ReceivedFile)
	fileValidationResult.Config = configValidationResult

	if metadataValidationResult.Status != constants.STATUS_SUCCESS {
		fileValidationResult.Status = constants.STATUS_FAILED
		logger.Error(constants.MSG_FILE_VALIDATION_FAIL)
		return fileValidationResult
	}
	fileValidationResult.Status = constants.STATUS_SUCCESS

	return fileValidationResult
}

func validateMetadataFile(fileMetadata string) models.MetadataValidationResult {
	validationResult := models.MetadataValidationResult{}

	file, err := os.Open(fileMetadata)
	if err != nil {
		validationResult.Error = &models.FileError{Message: constants.FILE_OPEN_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}
	defer file.Close()

	fields, err := io.ReadAll(file)
	if err != nil {
		validationResult.Error = &models.FileError{Message: constants.FILE_READ_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	var metadataMap map[string]string
	err = json.Unmarshal(fields, &metadataMap)
	if err != nil {
		validationResult.Error = &models.FileError{Message: constants.ERROR_UNMARSHALING_JSON, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}
	validationResult.Jurisdiction = metadataMap[constants.JURISDICTION]

	if filename, ok := metadataMap[constants.RECEIVED_FILENAME]; ok {
		validationResult.ReceivedFile = filename
	} else {
		validationResult.Error = &models.FileError{Message: constants.FILE_MISSING_ERROR, Code: 13}
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

func validateConfigFile(dataStreamId string, received_file string) models.ConfigValidationResult {

	validationResult := models.ConfigValidationResult{}

	headerValidationResult := models.HeaderValidationResult{}

	var identifiers []models.ConfigIdentifier

	fields, err := os.ReadFile(constants.CONFIG_FILE)
	if err != nil {
		validationResult.Error = &models.FileError{Message: constants.FILE_READ_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	err = json.Unmarshal(fields, &identifiers)
	if err != nil {
		validationResult.Error = &models.FileError{Message: constants.ERROR_UNMARSHALING_JSON, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	var expectedHeader []string
	for _, identifier := range identifiers {
		if identifier.DataStreamID == dataStreamId && len(identifier.Header) > 0 {
			expectedHeader = identifier.Header

		}
	}
	if len(expectedHeader) > 0 {
		//open file to get actual header to compare with expected one
		file, err := os.Open(received_file)
		if err != nil {
			validationResult.Error = &models.FileError{Message: constants.FILE_OPEN_ERROR, Code: 13}
			validationResult.Status = constants.STATUS_FAILED
			return validationResult
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				validationResult.Error = &models.FileError{Message: constants.FILE_CLOSE_ERROR, Code: 13}
				validationResult.Status = constants.STATUS_FAILED
			}
		}(file)

		//we need to move file pointer 3 bytes to exclude byte order mark, if detected
		hasBom, err := detector.DetectBOM(file)
		if err != nil {
			validationResult.Status = constants.BOM_NOT_DETECTED_ERROR
		}

		if hasBom {
			file.Seek(3, 0)
		}

		//we need to get actual header from the file to compare with expected one
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
			validationResult.Error = &models.FileError{Message: constants.ERR_HEADER_VALIDATION, Code: 13}
		}

		validationResult.HeaderValidationResult = headerValidationResult

	}

	return validationResult
}

func validateFile(receivedFile string) models.FileValidationResult {
	validationResult := models.FileValidationResult{}

	fileUUID := uuid.New()
	validationResult.FileUUID = fileUUID
	validationResult.ReceivedFile = receivedFile

	file, err := os.Open(receivedFile)
	if err != nil {
		validationResult.Error = &models.FileError{Message: constants.FILE_OPEN_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			validationResult.Error = &models.FileError{Message: constants.FILE_CLOSE_ERROR, Code: 13}
			validationResult.Status = constants.STATUS_FAILED
			return
		}
	}(file)
	hasBOM, err := detector.DetectBOM(file)

	if err != nil {
		validationResult.Error = &models.FileError{Message: constants.BOM_NOT_DETECTED_ERROR, Code: 13}
	}

	data, err := utils.ReadFileRandomly(file)
	if err != nil {
		validationResult.Error = &models.FileError{Message: err.Error(), Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult

	}

	detectedDelimiter := detector.DetectDelimiter(data)
	validationResult.Delimiter = string(constants.DelimiterCharacters[detectedDelimiter])

	if validationResult.Delimiter == string(constants.DelimiterCharacters[0]) {
		validationResult.Error = &models.FileError{Message: constants.UNSUPPORTED_DELIMITER_ERROR, Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	if hasBOM {
		validationResult.Encoding = constants.UTF8_BOM
	} else {
		detectedEncoding := detector.DetectEncoding(data)
		if detectedEncoding == constants.UNDEF {
			validationResult.Error = &models.FileError{Message: constants.UNSUPPORTED_ENCODING_ERROR, Code: 13}
			validationResult.Status = constants.STATUS_FAILED
			return validationResult
		}
		validationResult.Encoding = detectedEncoding
	}

	validationResult.Status = constants.STATUS_SUCCESS

	return validationResult
}
