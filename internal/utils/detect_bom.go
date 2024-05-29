package utils

import (
	"bytes"
	"io"
	"os"
)

var (
	UTF8_BOM = []byte{0xEF, 0xBB, 0xBF}
)

func DetectBOM(file *os.File) (bool, error) {
	bom := make([]byte, 3)
	_, err := file.Read(bom)
	if err != nil && err != io.EOF {
		return false, err
	}

	//reset the file pointer to the begining of the file
	if _, err := file.Seek(0, 0); err != nil {
		return false, err
	}
	//compare the first three bytes of the file with the UTF-8 BOM sequence
	if bytes.Equal(bom, UTF8_BOM) {
		return true, nil
	}

	return false, nil
}
