package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/utils"
	"golang.org/x/text/encoding/charmap"
)

func IsValid(source string) constants.FileValidationResult {
	file, err := os.Open(source)

	if err != nil {
		slog.Error(constants.FILE_OPEN_ERROR)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error(constants.FILE_CLOSE_ERROR)
		}
	}(file)

	//initialize file object
	var result constants.FileValidationResult
	result.Name = "csv file"
	result.Source = source

	//check for bom
	hasBOM, err := utils.DetectBOM(file)
	if err != nil {
		fmt.Println("Error Detecting Bom", err)
	}

	//get sample data at random places to detect delimiter, and encoding
	randomSampleData, err := utils.GetRandomSample(file)
	if err != nil {
		fmt.Println("Error getting randomSampleData ", randomSampleData, err)
	}

	detector := &utils.DelimiterDetector{}
	delimiter := detector.DetectDelimiter(string(randomSampleData))
	result.Delimiter = constants.DelimiterCharacters[delimiter.Character]

	enc := constants.DelimiterCharacters[delimiter.Character]
	fmt.Println(enc)
	result.Delimiter = enc

	if hasBOM {
		result.Encoding = constants.UTF8_BOM
	} else {
		//detect encoding with random sample data
		//enc := utils.DetectEncoding(randomSampleData)
		decoder := charmap.Windows1252.NewDecoder()

		//tempcode-> check if windows1252 decoder can correctly parse the csv file
		_, err = file.Seek(0, 0)
		if err != nil {
			result.Error = err
			return result
		}
		reader := csv.NewReader(decoder.Reader(file))

		fmt.Println(reader.ReadAll())
		for {
			record, err := reader.Read()
			if err == io.EOF {
				fmt.Println(err)
				break
			}
			fmt.Println(record)
		}
		result.Encoding = utils.DetectEncoding(randomSampleData)
	}

	return result
}
