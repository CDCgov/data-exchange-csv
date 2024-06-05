package detect

import (
	"bytes"
	"io"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func BOM(file *os.File) (bool, error) {
	bom := make([]byte, 3)

	_, err := file.Read(bom)
	if err != nil && err != io.EOF {
		return false, err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return false, err
	}

	if bytes.Equal(bom, constants.UTF8Bom) {
		return true, nil
	}

	return false, nil
}
