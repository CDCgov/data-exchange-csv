package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/validate/file"
	"golang.org/x/text/encoding/charmap"
)

func main() {
	//source := "data/file-with-headers-100-rows.csv"
	//source := "data/file-with-headers-100-rows_with_BOM.csv"
	//source := "data/file-with-headers-100-rows_US_ASCII.csv"
	//source := "data/file-with-headers-rows_iso8859-1.csv"
	source := "data/file-with-headers-windows1252.csv"

	fileValidationResult := file.Validate(source)

	fmt.Println("file validation result: ", fileValidationResult)

	//detect encoding with random sample data
	//enc := utils.DetectEncoding(randomSampleData)
	decoder := charmap.Windows1252.NewDecoder()

	//tempcode-> check if windows1252 decoder can correctly parse the csv file
	file, err := os.Open(source)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(decoder.Reader(file))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
