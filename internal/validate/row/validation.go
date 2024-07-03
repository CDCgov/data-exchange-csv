package row

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/transform"
	"github.com/google/uuid"
)

type RowValidationResult struct {
	FileUUID  uuid.UUID `json:"file_uuid"`
	RowNumber int       `json:"row_number"`
	RowUUID   uuid.UUID `json:"row_uuid"`
	Hash      [32]byte  `json:"row_hash"`
	Error     *Error    `json:"error"`
	Status    string    `json:"status"`
}

type Error struct {
	Message string `json:"message"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
}

func Validate(reader *csv.Reader, fileUUID uuid.UUID, separator string) {
	validationResult := RowValidationResult{}
	validationResult.FileUUID = fileUUID

	rowCount := 0

	for {
		row, err := reader.Read()
		validationResult.Hash = ComputeHash(row, separator)

		if err == io.EOF {
			break
		}
		validationResult.RowNumber = rowCount
		rowCount++

		validationResult.RowUUID = uuid.New()

		if err != nil {
			validationResult.Error = processRowError(err)
			validationResult.Status = constants.STATUS_FAILED
			fmt.Println("Invalid Row Result: ", validationResult) //temp
		}

		validationResult.Status = constants.STATUS_SUCCESS
		fmt.Println("Valid Row Result: ", validationResult) //temp

		// ready to transform to json
		transform.RowToJson(row, validationResult.FileUUID, validationResult.RowUUID)

	}

}
func processRowError(err error) *Error {
	rowError := &Error{}

	if parseErr, ok := err.(*csv.ParseError); ok {
		rowError.Line = parseErr.Line
		rowError.Column = parseErr.Column

		if parseErr.Err == csv.ErrFieldCount {
			rowError.Message = constants.ERR_MISMATCHED_FIELD_COUNTS
			return rowError
		}

		if parseErr.Err == csv.ErrQuote {
			rowError.Message = constants.ERR_UNESCAPED_QUOTES
			return rowError
		}

		if parseErr.Err == csv.ErrBareQuote {
			rowError.Message = constants.ERR_BARE_QUOTE
			return rowError
		}

		rowError.Message = err.Error()

	}

	return rowError

}
