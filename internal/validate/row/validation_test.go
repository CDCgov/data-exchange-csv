package row

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockSendEventsToDestination struct {
	result      interface{}
	destination string
}

func (m *MockSendEventsToDestination) callback(result interface{}, destination string) {
	m.result = result
	m.destination = destination
}

// TestValidate_Success	tests positive path in Validate().
func TestValidate_Success(t *testing.T) {
	params := []models.FileValidationParams{
		models.FileValidationParams{
			ReceivedFile: "testdata/test.csv", // TODO: Replace with actual path to test csv
			Encoding:     constants.UTF8_BOM,
			Delimiter:    ",", // TODO: Should this Delimiter field be a rune type?
			FileUUID:     uuid.New(),
		},
		models.FileValidationParams{
			ReceivedFile: "testdata/test2.csv",
			Encoding:     constants.UTF8_BOM,
			Delimiter:    ",",
			FileUUID:     uuid.New(),
		},
	}

	expectedResults := []models.RowValidationResult{
		models.RowValidationResult{
			FileUUID:  params[1].FileUUID, // TODO: There is a probably a better way to set up a 2D matrix of structs for creating input and output data; otherwise you deal with this odd hard-coding indices
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash1",
			RowNumber: 1,
		},
		models.RowValidationResult{
			FileUUID:  params[2].FileUUID,
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash2",
			RowNumber: 2,
		},
		// ... add more expected results if needed
	}

	sendEventsToDestination := MockSendEventsToDestination{}

	for i, param := range params {
		Validate(param, sendEventsToDestination.callback)
		assert.Equal(t, expectedResults[i], sendEventsToDestination.result)
		sendEventsToDestination = MockSendEventsToDestination{} // reset struct contents
	}
}

// TestValidate_ErrorCreatingReader tests when Validate() fails to create a CSV file reader.
func TestValidate_ErrorCreatingReader(t *testing.T) {
	params := models.FileValidationParams{
		ReceivedFile: "testdata/nonexistent.csv", // non-existent file
		Encoding:     constants.UTF8_BOM,
		Delimiter:    ",",
		FileUUID:     uuid.New(),
	}

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

	sendEventsToDestination := MockSendEventsToDestination{}

	Validate(params, sendEventsToDestination.callback)

	assert.Equal(t, expectedResult, sendEventsToDestination.result)
}

// TestValidate_ErrorReadingRow tests when Validate() encounters a read error.
func TestValidate_ErrorReadingRow(t *testing.T) {

	params := []models.FileValidationParams{
		models.FileValidationParams{
			ReceivedFile: "testdata/invalid.csv",
			Encoding:     constants.UTF8_BOM,
			Delimiter:    ",",
			FileUUID:     uuid.New(),
		},
		models.FileValidationParams{
			ReceivedFile: "testdata/invalid2.csv",
			Encoding:     constants.UTF8_BOM,
			Delimiter:    ",",
			FileUUID:     uuid.New(),
		},
	}

	expectedResults := []models.RowValidationResult{
		models.RowValidationResult{
			FileUUID:  params[1].FileUUID,
			Status:    constants.STATUS_SUCCESS,
			Error:     nil,
			RowUUID:   uuid.New(),
			Hash:      "hash1",
			RowNumber: 1,
		},
		models.RowValidationResult{
			FileUUID: params[2].FileUUID,
			Status:   constants.STATUS_FAILED,
			Error: &models.RowError{
				Message:  "Mismatched field counts",
				Severity: "Failure",
				Line:     -1,
				Column:   -1,
			},
		},
	}

	sendEventsToDestination := MockSendEventsToDestination{}

	for i, param := range params {
		Validate(param, sendEventsToDestination.callback)
		assert.Equal(t, expectedResults[i], sendEventsToDestination.result)
		sendEventsToDestination = MockSendEventsToDestination{}
	}
}
