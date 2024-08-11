package transform

import (
	"encoding/json"
	"strconv"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/google/uuid"
)

func RowToJson(row []string, params models.FileValidationParams,
	rowUUID uuid.UUID, sendEventsToDestination func(result interface{}, destination string)) {

	transformationResult := models.RowTransformationResult{
		FileUUID: params.FileUUID,
		RowUUID:  rowUUID,
	}

	parsedRow := make(map[string]string)

	if len(params.Header) > 0 {
		for index, column := range params.Header {
			parsedRow[column] = row[index]
		}
	} else {
		for index, field := range row {
			/*
				Use strconv.Itoa() to convert row slice indices to strings for use as keys in the map.
				This is needed because map keys must be strings.
				Note: `string()` converts integers to Unicode code points (e.g., `string(65)` → "A"),
				whereas `strconv.Itoa()` converts integers to their string representation (e.g., `strconv.Itoa(65)` → "65").
			*/
			parsedRow[strconv.Itoa(index)] = field
		}
	}
	transformedRow, err := json.Marshal(parsedRow)

	if err != nil {
		transformationResult.Error = err
		transformationResult.Status = constants.STATUS_FAILED
		sendEventsToDestination(transformationResult, constants.DEAD_LETTER_QUEUE)
		return
	}

	transformationResult.Status = constants.STATUS_SUCCESS
	transformationResult.JsonRow = transformedRow
	sendEventsToDestination(transformationResult, constants.TRANSFORMED_ROW_REPORTS)
}
