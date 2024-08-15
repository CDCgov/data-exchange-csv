package row

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_Success(t *testing.T) {
	params := models.FileValidationParams{
		ReceivedFile: "testdata/test.csv", // TODO: Replace with actual path to test csv
		Encoding:     constants.UTF8_BOM,
		Delimiter:    ",", // TODO: Should this Delimiter field be a rune type?
		FileUUID:     uuid.New(),
	}

	var eventsSent []interface{} // Stores events that are sent inside Validate() call

	// TODO: Figure out how to implement/abstract a function that is passed in as a parameter for testing
	sendEventsToDestination := func(result interface{}, destination string) {
		eventsSent = append(eventsSent, result)
	}

	Validate(params, sendEventsToDestination)

	assert.Len(t, eventsSent, 3)

	expectedResults := []models.RowValidationResult{
		models.RowValidationResult{
			FileUUID:  params.FileUUID,
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash1",
			RowNumber: 1,
		},
		models.RowValidationResult{
			FileUUID:  params.FileUUID,
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash2",
			RowNumber: 2,
		},
		// ... add more expected results if needed
	}

	for i, result := range expectedResults {
		assert.Equal(t, result, eventsSent[i])
	}
}

func TestValidate_ErrorCreatingReader(t *testing.T) {

	params := models.FileValidationParams{
		ReceivedFile: "testdata/nonexistent.csv", // non-existent file
		Encoding:     constants.UTF8_BOM,
		Delimiter:    ",",
		FileUUID:     uuid.New(),
	}

	var eventsSent []interface{}
	sendEventsToDestination := func(result interface{}, destination string) {
		eventsSent = append(eventsSent, result)
	}

	Validate(params, sendEventsToDestination)

	assert.Len(t, eventsSent, 1)

	expectedResult := models.RowValidationResult{
		FileUUID: params.FileUUID,
		Status:   constants.STATUS_FAILED,
		Error: &models.RowError{
			Message:  "CSV reader error",
			Severity: "Failure",
			Line:     -1,
			Column:   -1,
		},
	}

	assert.Equal(t, expectedResult, eventsSent[0])
}

func TestValidate_ErrorReadingRow(t *testing.T) {

	params := models.FileValidationParams{
		ReceivedFile: "testdata/invalid.csv",
		Encoding:     constants.UTF8_BOM,
		Delimiter:    ",",
		FileUUID:     uuid.New(),
	}

	var eventsSent []interface{}
	sendEventsToDestination := func(result interface{}, destination string) {
		eventsSent = append(eventsSent, result)
	}

	Validate(params, sendEventsToDestination)

	assert.Len(t, eventsSent, 2)

	expectedResults := []models.RowValidationResult{
		models.RowValidationResult{
			FileUUID:  params.FileUUID,
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash1",
			RowNumber: 1,
		},
		models.RowValidationResult{
			FileUUID: params.FileUUID,
			Status:   constants.STATUS_FAILED,
			Error: &models.RowError{
				Message:  "Mismatched field counts",
				Severity: "Failure",
				Line:     -1,
				Column:   -1,
			},
		},
	}

	for i, result := range expectedResults {
		assert.Equal(t, result, eventsSent[i])
	}
}
