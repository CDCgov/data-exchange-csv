package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"golang.org/x/text/encoding/charmap"
)

func main() {
	source := "data/event_config.json"

	vaidationResult := &file.ValidationResult{}
	vaidationResult.Validate(source)

	file, _ := os.Open(vaidationResult.ReceivedFile)
	detectedEncoding := vaidationResult.Encoding

	var reader *csv.Reader

	if detectedEncoding == constants.UTF8 || detectedEncoding == constants.UTF8_BOM {
		reader = csv.NewReader(file)
	} else if detectedEncoding == constants.ISO8859_1 {
		decoder := charmap.ISO8859_1.NewDecoder()
		reader = csv.NewReader(decoder.Reader(file))
	} else if detectedEncoding == constants.WINDOWS1252 {
		decoder := charmap.Windows1252.NewDecoder()
		reader = csv.NewReader(decoder.Reader(file))
	} else {
		fmt.Println("TODO: handle unsupported encoding")
		return
	}

	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		}
		// TODO - row validate, file uuid.

	}

}
