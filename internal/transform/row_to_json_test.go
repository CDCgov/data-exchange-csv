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

// callback will store the transformation result from RowToJson
func (m *MockTransformedResult) callback(params models.RowCallbackParams) error {
	m.transformedResult = params.TransformationResult.(string)

	return nil
}

func TestRowToJsonSuccess(t *testing.T) {
	mockSender := &MockTransformedResult{}

	//initialize input parameters for `RowToJson``
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
	//call the function being tested
	RowToJson(rowToBeTransformed, params, rowUUID, false, header, mockSender.callback)

	//deserialize the transformed result to a map
	var actualTransformedResult map[string]interface{}
	//convert transformed result to slice of bytes before Unmarshalling
	convertTransformedResultToByteSlice := []byte(mockSender.transformedResult)
	err := json.Unmarshal(convertTransformedResultToByteSlice, &actualTransformedResult)
	if err != nil {
		t.Errorf("Unexpected error occurred during JSON Unmarshalling %v", err.Error())
	}

	//validate File UUID
	fileUUIDAsString := expectedResult.FileUUID.String()
	if fileUUIDAsString != actualTransformedResult["file_uuid"] {
		t.Errorf("Expected File UUID  %v and actual  %v do not match", fileUUIDAsString, actualTransformedResult["file_uuid"])
	}
	//validate  Row UUID
	rowUUIDAsString := expectedResult.RowUUID.String()
	if rowUUIDAsString != actualTransformedResult["row_uuid"] {
		t.Errorf("Expected Row UUID  %v and actual  %v do not match", rowUUIDAsString, actualTransformedResult["row_uuid"])
	}
	//Validate Status
	if expectedResult.Status != actualTransformedResult["status"] {
		t.Errorf("Expected status  %v and actual  %v do not match", expectedResult.Status, actualTransformedResult["status"])
	}

	//Unmarshall expected JSON row for comparison
	var jsonMap map[string]interface{}
	err = json.Unmarshal(expectedResult.JsonRow, &jsonMap)
	if err != nil {
		t.Errorf("Error unmarshalling expectedResult.JsonRow %s", err)
	}

	//validate transformed JSON row
	if !reflect.DeepEqual(jsonMap, actualTransformedResult["json_row"]) {
		t.Errorf("Expected jsonRow  %v and actual  %v do not match", jsonMap, actualTransformedResult["json_row"])
	}

}

/*
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
*/
