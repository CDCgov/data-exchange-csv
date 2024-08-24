package transform

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/CDCgov/data-exchange-csv/cmd/pkg/sloger"
	"github.com/google/uuid"
)

func RowToJson(row []string, params models.FileValidationParams,
	rowUUID uuid.UUID, sendEventsToDestination func(result interface{}, destination string)) {

	//initialize logger using sloger package
	logger := sloger.With(constants.PACKAGE, constants.TRANSFORM)
	logger.Info(fmt.Sprintf(constants.MSG_ROW_TRANSFORMATION_BEGIN, rowUUID))

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
		logger.Error(fmt.Sprintf(constants.MSG_ROW_TRANSFORM_ERROR, err.Error()))
		sendEventsToDestination(transformationResult, constants.DEAD_LETTER_QUEUE)
		return
	}

	transformationResult.Status = constants.STATUS_SUCCESS
	transformationResult.JsonRow = transformedRow
	logger.Info(constants.MSG_ROW_TRANSFORM_SUCCESS)
	sendEventsToDestination(transformationResult, constants.TRANSFORMED_ROW_REPORTS)
}
