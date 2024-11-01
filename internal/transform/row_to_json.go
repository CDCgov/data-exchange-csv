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

func RowToJson(row []string, params models.FileValidationResult,
	rowUUID uuid.UUID, isFirst bool, header []string, callback func(params models.RowCallbackParams) error) {

	//initialize logger using sloger package
	logger := sloger.With(constants.PACKAGE, constants.TRANSFORM)
	logger.Debug(fmt.Sprintf(constants.MSG_ROW_TRANSFORMATION_BEGIN, rowUUID))

	transformationResult := models.RowTransformationResult{
		FileUUID: params.FileUUID,
		RowUUID:  rowUUID,
	}

	parsedRow := make(map[string]string)

	if params.HasHeader {
		for index, column := range header {
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

		jsonContent, err := json.Marshal(transformationResult)
		if err != nil {
			logger.Error(constants.ERROR_CONVERTING_STRUCT_TO_JSON)
		}

		callback(models.RowCallbackParams{
			IsFirst:              isFirst,
			FileUUID:             params.FileUUID,
			TransformationResult: string(jsonContent),
			Destination:          params.Destination,
			Transform:            true,
		})

		return
	}

	transformationResult.Status = constants.STATUS_SUCCESS
	transformationResult.JsonRow = transformedRow

	logger.Debug(constants.MSG_ROW_TRANSFORM_SUCCESS)

	jsonContent, err := json.Marshal(transformationResult)
	if err != nil {
		logger.Error(constants.ERROR_CONVERTING_STRUCT_TO_JSON)
	}

	callback(models.RowCallbackParams{
		IsFirst:              isFirst,
		FileUUID:             params.FileUUID,
		TransformationResult: string(jsonContent),
		Destination:          params.Destination,
		Transform:            true,
	})
}
