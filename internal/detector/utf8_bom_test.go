package detector

import (
	"os"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func createTempFile(content []byte, fileName string) (*os.File, error) {
	file, err := os.CreateTemp("", fileName)
	if err != nil {
		return nil, err
	}

	_, err = file.Write(content)
	if err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}

func checkForBom(filePath *os.File, t *testing.T) bool {
	file, _ := os.Open(filePath.Name())
	bomFound, err := DetectBOM(file)
	defer os.Remove(file.Name())

	if err != nil {
		t.Errorf(constants.FILE_OPEN_ERROR)
	}

	return bomFound

}
func TestDetectBOM(t *testing.T) {

	csvFileWithBOM, err := createTempFile(constants.UTF8Bom, constants.CSV_FILENAME_WITH_BOM)
	if err != nil {
		t.Errorf(constants.FILE_CREATE_ERROR)
	}
	expectedBom := true
	actualBom := checkForBom(csvFileWithBOM, t)

	if expectedBom != actualBom {
		t.Errorf(constants.BOM_NOT_DETECTED_ERROR)

	}
}

func TestNotDetectBOM(t *testing.T) {

	csvFileWithNoBOM, err := createTempFile(constants.UTF8NoBom, constants.CSV_FILENAME_WITHOUT_BOM)
	if err != nil {
		t.Errorf(constants.FILE_CREATE_ERROR)
	}
	expectedBom := false
	actualBom := checkForBom(csvFileWithNoBOM, t)

	if expectedBom != actualBom {
		t.Errorf(constants.BOM_NOT_DETECTED_ERROR)

	}
}
