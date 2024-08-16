package transform

import (
	"encoding/json"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/google/uuid"
)

type MockSendEventsToDestination struct {
	result      interface{}
	destination string
}

func (m *MockSendEventsToDestination) callback(result interface{}, destination string) {
	m.result = result
	m.destination = destination
}

func compareTransformationResult(expected, actual models.RowTransformationResult) bool {
	if actual.FileUUID != expected.FileUUID || actual.RowUUID != expected.RowUUID || actual.Status != expected.Status {
		return false
	}
	if (actual.JsonRow == nil && expected.JsonRow != nil) || (actual.JsonRow != nil && expected.JsonRow == nil) {
		return false
	}
	if string(actual.JsonRow) != string(expected.JsonRow) {
		return false
	}
	if (actual.Error == nil && expected.Error != nil) || (actual.Error != nil && expected.Error == nil) {
		return false
	}

	return true
}

func TestRowToJsonSuccess(t *testing.T) {
	mockSender := &MockSendEventsToDestination{}

	rowToBeTransformed := []string{"Doppio", "Cortado", "Galao", "Lungo"}
	params := models.FileValidationParams{
		FileUUID: uuid.New(),
		Header:   []string{"Espresso", "Balanced Espresso", "Latte", "Long-pull Espresso"},
	}

	rowUUID := uuid.New()

	expectedTransformedRow, _ := json.Marshal(map[string]string{"Espresso": "Doppio", "Balanced Espresso": "Cortado", "Latte": "Galao", "Long-pull Espresso": "Lungo"})

	expectedResult := models.RowTransformationResult{
		FileUUID: params.FileUUID,
		RowUUID:  rowUUID,
		Status:   constants.STATUS_SUCCESS,
		JsonRow:  expectedTransformedRow,
	}

	RowToJson(rowToBeTransformed, params, rowUUID, mockSender.callback)

	if mockSender.destination != constants.TRANSFORMED_ROW_REPORTS {
		t.Errorf("Expected destination  %v and actual  %v do not match", constants.TRANSFORMED_ROW_REPORTS, mockSender.destination)
	}
	if !compareTransformationResult(expectedResult, mockSender.result.(models.RowTransformationResult)) {
		t.Errorf("Expected transformation result %v and actual  %v do not match", expectedResult, mockSender.result)
	}
}

func TestRowToJsonNoHeader(t *testing.T) {
	mockSender := &MockSendEventsToDestination{}

	rowToBeTransformed := []string{"Doppio", "Cortado", "Galao", "Lungo"}
	params := models.FileValidationParams{
		FileUUID: uuid.New(),
		Header:   []string{},
	}

	rowUUID := uuid.New()

	expectedTransformedRow, _ := json.Marshal(map[string]string{"0": "Doppio", "1": "Cortado", "2": "Galao", "3": "Lungo"})

	expectedResult := models.RowTransformationResult{
		FileUUID: params.FileUUID,
		RowUUID:  rowUUID,
		Status:   constants.STATUS_SUCCESS,
		JsonRow:  expectedTransformedRow,
	}

	RowToJson(rowToBeTransformed, params, rowUUID, mockSender.callback)

	if mockSender.destination != constants.TRANSFORMED_ROW_REPORTS {
		t.Errorf("Expected destination  %v and actual  %v do not match", constants.TRANSFORMED_ROW_REPORTS, mockSender.destination)
	}

	if !compareTransformationResult(expectedResult, mockSender.result.(models.RowTransformationResult)) {
		t.Errorf("Expected transformation result %v and actual  %v do not match", expectedResult, mockSender.result)
	}
}
