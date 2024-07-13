package transform

import (
	"encoding/json"
	"strconv"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/google/uuid"
)

type RowTransformationResult struct {
	FileUUID uuid.UUID `json:"file_uuid"`
	RowUUID  uuid.UUID `json:"row_uuid"`
	JsonRow  string    `json:"json_row"`
	Error    error     `json:"error"`
	Status   string    `json:"status"`
}

func RowToJson(row []string, fileUUID uuid.UUID, rowUUID uuid.UUID, header []string) {
	transformaionResult := RowTransformationResult{
		FileUUID: fileUUID,
		RowUUID:  rowUUID,
	}

	parsedRow := make(map[string]string)
	if len(header) > 0 {
		for index, column := range header {
			parsedRow[column] = row[index]
		}
	} else {
		for index, field := range row {
			/*
				using strconv.Itoa() function here to ensure each index of the
				row slice can be used as a string key. This is essential because
				JsonRow map keys need to be strings.
				Note that string() can not be used as it converts integer to a
				Unicode code point(rune).
				Example string()-> string(65)-> "A"
				Example strconv.Itoa() ->strconv.Itoa(65) -> "65"
			*/
			parsedRow[strconv.Itoa(index)] = field
		}
	}
	transformedRow, err := json.Marshal(parsedRow)

	if err != nil {
		transformaionResult.Error = err
		file.CopyToDestination(transformaionResult, constants.DEAD_LETTER_QUEUE)
	}

	transformaionResult.Status = constants.STATUS_SUCCESS
	transformaionResult.JsonRow = string(transformedRow)
	file.CopyToDestination(transformaionResult, constants.TRANSFORMED_ROW_REPORTS)

}
