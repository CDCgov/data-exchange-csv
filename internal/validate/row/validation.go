package row

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/transform"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/google/uuid"
	"golang.org/x/text/encoding/charmap"
)

type RowValidationResult struct {
	FileUUID  uuid.UUID `json:"file_uuid"`
	RowNumber int       `json:"row_number"`
	RowUUID   uuid.UUID `json:"row_uuid"`
	Hash      string    `json:"row_hash"`
	Error     *Error    `json:"error"`
	Status    string    `json:"status"`
}

type Error struct {
	Message  string             `json:"message"`
	Line     int                `json:"line"`
	Column   int                `json:"column"`
	Severity constants.Severity `json:"severity"`
}

func Validate(params file.FileValidationParams,
	dlqCallback, routingCallback func(result interface{}, destination string)) {

	file, _ := os.Open(params.ReceivedFile)
	//handle error
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	detectedEncoding := params.Encoding

	var reader *csv.Reader

	if detectedEncoding == constants.UTF8 {
		reader = csv.NewReader(file)
	} else if detectedEncoding == constants.UTF8_BOM {
		file.Seek(3, 0)
		reader = csv.NewReader(file)
	} else if detectedEncoding == constants.ISO8859_1 {
		decoder := charmap.ISO8859_1.NewDecoder()
		reader = csv.NewReader(decoder.Reader(file))
	} else {
		decoder := charmap.Windows1252.NewDecoder()
		reader = csv.NewReader(decoder.Reader(file))
	}

	// if detected delimiter is TSV, change the seperator for a csv.Reader to a tab rune
	if params.Delimiter == constants.TSV {
		reader.Comma = constants.TAB
	}

	//initialize row validation result
	validationResult := RowValidationResult{
		FileUUID: params.FileUUID,
	}

	//if header  is present or failed skip the first row
	if len(params.Header) != 0 {
		reader.Read()
	}

	rowCount := 1
	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}
		validationResult.RowUUID = uuid.New()
		validationResult.Hash = ComputeHash(row, params.Delimiter)
		validationResult.RowNumber = rowCount

		rowCount++

		if err != nil {
			validationResult.Error = processRowError(err)
			validationResult.Status = constants.STATUS_FAILED
			dlqCallback(validationResult, constants.DEAD_LETTER_QUEUE)
			continue
		}

		validationResult.Status = constants.STATUS_SUCCESS
		routingCallback(validationResult, constants.ROW_REPORTS)
		// valid row, ready to transform to json
		transform.RowToJson(row, params, validationResult.RowUUID, dlqCallback, routingCallback)

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
