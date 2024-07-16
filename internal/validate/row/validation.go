package row

import (
	"encoding/csv"
	"io"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/transform"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/google/uuid"
)

type RowValidationResult struct {
	FileUUID   uuid.UUID `json:"file_uuid"`
	RowNumber  int       `json:"row_number"`
	RowUUID    uuid.UUID `json:"row_uuid"`
	RowContent []string  `json:"row_content"`
	Hash       string    `json:"row_hash"`
	Error      *Error    `json:"error"`
	Status     string    `json:"status"`
}

type Error struct {
	Message  string             `json:"message"`
	Line     int                `json:"line"`
	Column   int                `json:"column"`
	Severity constants.Severity `json:"severity"`
}

func Validate(reader *csv.Reader, fileUUID uuid.UUID, separator string, header file.HeaderValidationResult) {
	validationResult := RowValidationResult{}
	validationResult.FileUUID = fileUUID

	//if header row is present or failed to validate skip it.
	if header.Status != constants.EMPTY_FIELD {
		reader.Read()
	}

	rowCount := 0

	for {
		row, err := reader.Read()
		validationResult.Hash = ComputeHash(row, separator)
		validationResult.RowContent = row
		if err == io.EOF {
			break
		}
		validationResult.RowNumber = rowCount
		rowCount++

		validationResult.RowUUID = uuid.New()

		if err != nil {
			validationResult.Error = processRowError(err)
			validationResult.Status = constants.STATUS_FAILED
			file.CopyToDestination(validationResult, constants.DEAD_LETTER_QUEUE)
		}

		validationResult.Status = constants.STATUS_SUCCESS
		file.CopyToDestination(validationResult, constants.ROW_REPORTS)

		// ready to transform to json
		transform.RowToJson(row, validationResult.FileUUID, validationResult.RowUUID, header.Actual)

	}

}
func processRowError(err error) *Error {
	rowError := &Error{}

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
