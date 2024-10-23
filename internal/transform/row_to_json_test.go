package transform

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/models"
	"github.com/google/uuid"
)

type MockTransformedResult struct {
	transformedResult string
}

func (m *MockTransformedResult) callback(params models.RowCallbackParams) error {
	m.transformedResult = params.TransformationResult.(string)
	return nil
}

func validate(t *testing.T, expectedResult models.RowTransformationResult, actualTransformedResult string) {
	var actualResultMap map[string]interface{}
	convertTransformedResultToByteSlice := []byte(actualTransformedResult)
	err := json.Unmarshal(convertTransformedResultToByteSlice, &actualResultMap)
	if err != nil {
		t.Errorf("Unexpected error occurred during JSON Unmarshalling: %v", err.Error())
	}

	// Validate File UUID
	if expectedResult.FileUUID.String() != actualResultMap["file_uuid"] {
		t.Errorf("Expected File UUID %v and actual %v do not match", expectedResult.FileUUID.String(), actualResultMap["file_uuid"])
	}
	// Validate Row UUID
	if expectedResult.RowUUID.String() != actualResultMap["row_uuid"] {
		t.Errorf("Expected Row UUID %v and actual %v do not match", expectedResult.RowUUID.String(), actualResultMap["row_uuid"])
	}
	// Validate Status
	if expectedResult.Status != actualResultMap["status"] {
		t.Errorf("Expected status %v and actual %v do not match", expectedResult.Status, actualResultMap["status"])
	}

	// Unmarshal expected JSON row for comparison
	var expectedJsonRowMap map[string]interface{}
	err = json.Unmarshal(expectedResult.JsonRow, &expectedJsonRowMap)
	if err != nil {
		t.Errorf("Error unmarshalling expectedResult.JsonRow: %s", err)
	}

	// Validate transformed JSON row
	if !reflect.DeepEqual(expectedJsonRowMap, actualResultMap["json_row"]) {
		t.Errorf("Expected jsonRow %v and actual %v do not match", expectedJsonRowMap, actualResultMap["json_row"])
	}
}

func TestRowToJsonSuccess(t *testing.T) {
	mockSender := &MockTransformedResult{}

	rowToBeTransformed := []string{"Doppio", "Cortado", "Galao", "Lungo"}
	header := []string{"Espresso", "Balanced Espresso", "Latte", "Long-pull Espresso"}
	params := models.FileValidationResult{
		FileUUID:  uuid.New(),
		HasHeader: true,
	}
	rowUUID := uuid.New()

	expectedTransformedRow, _ := json.Marshal(map[string]string{"Espresso": "Doppio", "Balanced Espresso": "Cortado", "Latte": "Galao", "Long-pull Espresso": "Lungo"})
	expectedResult := models.RowTransformationResult{
		FileUUID: params.FileUUID,
		RowUUID:  rowUUID,
		Status:   constants.STATUS_SUCCESS,
		JsonRow:  expectedTransformedRow,
	}

	RowToJson(rowToBeTransformed, params, rowUUID, false, header, mockSender.callback)
	validate(t, expectedResult, mockSender.transformedResult)
}

func TestRowToJsonNoHeader(t *testing.T) {
	mockSender := &MockTransformedResult{}

	rowToBeTransformed := []string{"Doppio", "Cortado", "Galao", "Lungo"}
	params := models.FileValidationResult{
		FileUUID:  uuid.New(),
		HasHeader: false,
	}
	rowUUID := uuid.New()

	expectedTransformedRow, _ := json.Marshal(map[string]string{"0": "Doppio", "1": "Cortado", "2": "Galao", "3": "Lungo"})
	expectedResult := models.RowTransformationResult{
		FileUUID: params.FileUUID,
		RowUUID:  rowUUID,
		Status:   constants.STATUS_SUCCESS,
		JsonRow:  expectedTransformedRow,
	}

	RowToJson(rowToBeTransformed, params, rowUUID, false, nil, mockSender.callback)
	validate(t, expectedResult, mockSender.transformedResult)
}
