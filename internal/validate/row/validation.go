package row

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/transform"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"
	"github.com/google/uuid"
	"golang.org/x/text/encoding/charmap"
)

func createReader(file *os.File, encoding constants.EncodingType, delimiter string) (*csv.Reader, error) {
	var reader *csv.Reader

	switch encoding {
	case constants.WINDOWS1252:
		decoder := charmap.Windows1252.NewDecoder()
		reader = csv.NewReader(decoder.Reader(file))
	case constants.ISO8859_1:
		decoder := charmap.ISO8859_1.NewDecoder()
		reader = csv.NewReader(decoder.Reader(file))
	default:
		if encoding == constants.UTF8_BOM {
			if _, err := file.Seek(constants.BOM_LENGTH, 0); err != nil {
				return nil, err
			}
		}
		reader = csv.NewReader(file)
	}
	//If the file is tab-separated (TSV), update the reader's separator to TAB.
	//This ensures that the reader correctly parses each field based on the tab delimiter.
	if delimiter == string(constants.TAB) {
		reader.Comma = constants.TAB
	}

	return reader, nil
}

func Validate(params models.FileValidationParams, sendEventsToDestination func(result interface{}, destination string)) {
	//initialize logger from sloger package
	logger := sloger.With(constants.PACKAGE, constants.ROW)
	logger.Info(fmt.Sprintf(constants.MSG_ROW_VALIDATION_BEGIN, params.FileUUID))

	file, _ := os.Open(params.ReceivedFile)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(file)

	validationResult := models.RowValidationResult{
		FileUUID: params.FileUUID,
	}

	reader, err := createReader(file, params.Encoding, params.Delimiter)
	if err != nil {
		validationResult.Error = &models.RowError{Message: constants.CSV_READER_ERROR, Severity: constants.Failure}
		logger.Error(fmt.Sprintf(constants.MSG_CSV_READER_FAILURE, err.Error()))
		sendEventsToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
	}

	//If header is present, skip the header to ensure header row is not validated or transformed.
	if len(params.Header) != 0 {
		logger.Info(constants.MSG_HEADER_PRESENT_SKIP_FIRST_ROW)
		reader.Read()
	}

	rowCount := 1

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		validationResult.RowUUID = uuid.New()
		logger.Info(fmt.Sprintf(constants.MSG_ROW_UUID, validationResult.RowUUID))
		validationResult.Hash = ComputeHash(row, params.Delimiter)
		logger.Info(fmt.Sprintf(constants.MSG_ROW_COMPUTED_HASH, validationResult.Hash))
		validationResult.RowNumber = rowCount
		logger.Info(fmt.Sprintf(constants.MSG_ROW_NUMBER, rowCount))
		rowCount++

		if err != nil {
			validationResult.Error = processRowError(err)
			validationResult.Status = constants.STATUS_FAILED
			logger.Error(fmt.Sprintf(constants.MSG_ROW_VALIDATION_FAILURE, validationResult.Error.Message))
			sendEventsToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
			continue
		}

		validationResult.Status = constants.STATUS_SUCCESS
		logger.Info(constants.MSG_ROW_VALIDATION_SUCCESS)
		sendEventsToDestination(validationResult, constants.ROW_REPORTS)
		// valid row, ready to transform to json
		transform.RowToJson(row, params, validationResult.RowUUID, sendEventsToDestination)

	}

}

func processRowError(err error) *models.RowError {
	rowError := &models.RowError{}

	if parseErr, ok := err.(*csv.ParseError); ok {
		rowError.Line = parseErr.Line
		rowError.Column = parseErr.Column

		if parseErr.Err == csv.ErrFieldCount {
			rowError.Message = constants.ERR_MISMATCHED_FIELD_COUNTS
			rowError.Severity = constants.Failure
			return rowError
		}

		if parseErr.Err == csv.ErrQuote {
			rowError.Message = constants.ERR_UNESCAPED_QUOTES
			rowError.Severity = constants.Failure
			return rowError
		}

		if parseErr.Err == csv.ErrBareQuote {
			rowError.Message = constants.ERR_BARE_QUOTE
			rowError.Severity = constants.Failure
			return rowError
		}

		rowError.Message = err.Error()
	}

	return rowError
}
