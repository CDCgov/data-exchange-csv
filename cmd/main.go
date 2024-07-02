package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/row"
	"golang.org/x/text/encoding/charmap"
)

func main() {
	event := "data/event_config.json"

	validationResult := file.Validate(event)

	file, _ := os.Open(validationResult.ReceivedFile)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	detectedEncoding := validationResult.Encoding

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
		validationResult.Encoding = constants.UNDEF
		return
	}

	// if detected delimiter is TSV, change the seperator for a csv.Reader to a tab rune
	if validationResult.Delimiter == constants.TSV {
		reader.Comma = constants.TAB
	}
	row.Validate(reader, validationResult.FileUUID, validationResult.Delimiter)
}
