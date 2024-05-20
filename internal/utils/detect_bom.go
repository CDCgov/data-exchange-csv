package utils

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"io"
	"os"
)

func DetectBOM(file *os.File) (bool, error) {
	bom := make([]byte, 3)
	_, err := file.Read(bom)
	if err != nil && err != io.EOF {
		return false, err
	}

	// check if BOM is present and return true
	if string(bom) == constants.UTF8BOMHex || string(bom) == constants.UTF8BOMUnicode {
		return true, nil
	}
	return false, nil
}
