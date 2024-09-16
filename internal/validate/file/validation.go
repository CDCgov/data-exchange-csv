package file

import (
	"fmt"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/detector"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/row"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
	"github.com/google/uuid"
)

func Validate(fileInputParams models.FileValidateInputParams) {
	//initialize logger from sloger package
	logger := sloger.With(constants.PACKAGE, constants.FILE)
	logger.Info(constants.MSG_FILE_VALIDATION_BEGIN)

	fileValidationResult := validateFile(fileInputParams)
	fileInputParams.ValidationCallback(fileValidationResult)
	if fileValidationResult.Status == constants.STATUS_SUCCESS {
		row.Validate(fileValidationResult)
	}
}

func validateFile(params models.FileValidateInputParams) models.FileValidationResult {
	validationResult := models.FileValidationResult{
		FileUUID:     uuid.New(),
		ReceivedFile: params.FileURL,
	}

	//update destination where results will be stored
	validationResult.Destination = fmt.Sprintf("%s%s", params.Destination, constants.FILE_REPORTS)
	fmt.Println(validationResult.Destination)

	file, err := os.Open(params.FileURL)
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

	data, err := utils.ReadFileRandomly(file)
	if err != nil {
		validationResult.Error = &models.FileError{Message: err.Error(), Code: 13}
		validationResult.Status = constants.STATUS_FAILED
		return validationResult
	}

	if params.Encoding != "" {
		validationResult.Encoding = params.Encoding
	} else {
		//Encoding is not supplied, we need to detect it
		hasBOM, err := detector.DetectBOM(file)

		if err != nil {
			validationResult.Error = &models.FileError{Message: constants.BOM_NOT_DETECTED_ERROR, Code: 13}
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
	}

	if params.Separator == constants.TAB || params.Separator == constants.COMMA {
		validationResult.Delimiter = params.Separator
	} else {
		//delimiter is not supplied, we need to detect it
		detectedDelimiter := detector.DetectDelimiter(data)
		if detectedDelimiter == constants.DelimiterCharacters[0] {
			validationResult.Error = &models.FileError{Message: constants.UNSUPPORTED_DELIMITER_ERROR, Code: 13}
			validationResult.Status = constants.STATUS_FAILED
			return validationResult
		}
		validationResult.Delimiter = detectedDelimiter
	}

	if params.HasHeader {
		validationResult.HasHeader = true
	}

	//get size of the file
	fileInfo, err := file.Stat()
	if err != nil {
		validationResult.Status = "Error computing the size of the file"
	}

	fileSize := fileInfo.Size()
	validationResult.Size = fileSize

	validationResult.Status = constants.STATUS_SUCCESS
	return validationResult
}
