package detect

import (
	"os"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func TestDetectBOM(t *testing.T) {

	csvFileWithBOM, err := os.CreateTemp("", constants.CSV_FILENAME_WITH_BOM)
	if err != nil {
		t.Errorf(constants.FILE_CREATE_ERROR)
	}

	defer os.Remove(csvFileWithBOM.Name())

	bom := constants.UTF8Bom
	_, err = csvFileWithBOM.Write(bom)
	if err != nil {
		t.Errorf(constants.FILE_WRITE_ERROR)
	}

	file, _ := os.Open(csvFileWithBOM.Name())
	expectedBom := true
	bomFound, err := BOM(file)

	if err != nil {
		t.Errorf(constants.FILE_OPEN_ERROR)
	}

	if expectedBom != bomFound {
		t.Errorf(constants.BOM_NOT_DETECTED_ERROR)

	}
}

func TestNotDetectBOM(t *testing.T) {

	csvFileWithBOM, err := os.CreateTemp("", constants.CSV_FILENAME_WITHOUT_BOM)
	if err != nil {
		t.Errorf(constants.FILE_CREATE_ERROR)
	}
	defer os.Remove(csvFileWithBOM.Name())

	_, err = csvFileWithBOM.Write(constants.UTF8NoBom)
	if err != nil {
		t.Errorf(constants.FILE_WRITE_ERROR)
	}

	file, _ := os.Open(csvFileWithBOM.Name())
	expectedBom := false
	bomFound, err := BOM(file)

	if err != nil {
		t.Errorf(constants.FILE_OPEN_ERROR)
	}

	if expectedBom != bomFound {
		t.Errorf(constants.BOM_NOT_DETECTED_ERROR)

	}
}
