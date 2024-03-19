package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strings"
	"testing"
)

/*
//commenting out for now, revisit

	func TestLineBreakCR(t *testing.T) {
		csvData := "Name,Email\rJane Doe,johndoe@example.com\rJane Smith,janesmith@example.com\rChris Mallok,cmallok@example.com"

		reader := csv.NewReader(strings.NewReader(csvData))
		reader.LazyQuotes = true
		records, err := reader.ReadAll()
		fmt.Println(records)
		if err != nil {
			t.Errorf("Error reading CSV file %e", err)
		}

		expectedNumberOfRecords := 3
		actualNumberOfRecords := len(records) - 1 //subtract header row

		if actualNumberOfRecords != expectedNumberOfRecords {
			t.Errorf("Expected %d records, and got %d ", expectedNumberOfRecords, actualNumberOfRecords)
		}

}
*/
func TestLineBreakCRLF(t *testing.T) {
	//Line break CRLF
	csvData := "Name,Email\r\nJane Doe,johndoe@example.com\r\nJane Smith,janesmith@example.com\r\nChris Mallok,cmallok@example.com"

	reader := csv.NewReader(strings.NewReader(csvData))

	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Error reading CSV file %e", err)
	}

	expectedNumberOfRecords := 3
	actualNumberOfRecords := len(records) - 1 //subtract header row

	if actualNumberOfRecords != expectedNumberOfRecords {
		t.Errorf("Expected %d records, and got %d ", expectedNumberOfRecords, actualNumberOfRecords)
	}

}
func TestLineBreakLF(t *testing.T) {
	//Line break LF
	csvData := "Name,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok, cmallok@example.com"

	reader := csv.NewReader(strings.NewReader(csvData))

	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Error reading CSV file: %e", err)
	}

	expectedNumberOfRecords := 3
	actualNumberOfRecords := len(records) - 1 //subtract header row

	if actualNumberOfRecords != expectedNumberOfRecords {
		t.Errorf("Expected %d records, and got %d ", expectedNumberOfRecords, actualNumberOfRecords)
	}

}

func TestLineBreakAtTheEnd(t *testing.T) {
	//Line breaks on the last record
	csvData := "Name,Email\nJane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok, cmallok@example.com\n"

	reader := csv.NewReader(strings.NewReader(csvData))

	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Error reading CSV file: %e", err)
	}

	expectedNumberOfRecords := 3
	actualNumberOfRecords := len(records) - 1 //subtract header row

	if actualNumberOfRecords != expectedNumberOfRecords {
		t.Errorf("Expected %d records, and got %d ", expectedNumberOfRecords, actualNumberOfRecords)
	}

}

func TestNoHeader(t *testing.T) {
	//test with csv data that does not have header
	csvData := "Jane Doe,johndoe@example.com\nJane Smith,janesmith@example.com\nChris Mallok, cmallok@example.com\n"

	reader := csv.NewReader(strings.NewReader(csvData))

	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Error reading CSV file: %e", err)
	}

	expectedNumberOfRecords := 3
	actualNumberOfRecords := len(records)

	if actualNumberOfRecords != expectedNumberOfRecords {
		t.Errorf("Expected %d records, and got %d ", expectedNumberOfRecords, actualNumberOfRecords)
	}

}

func TestFieldOneOrMoreFields(t *testing.T) {
	//Within the header and within each record, there may be one or more fields, separated by commas.
	csvData := "Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com"

	reader := csv.NewReader(strings.NewReader(csvData))

	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Error reading CSV file: %e", err)
	}

	expectedNumberOfRecords := 10
	actualNumberOfRecords := len(records) - 1 //subtract header row

	if actualNumberOfRecords != expectedNumberOfRecords {
		t.Errorf("Expected %d records, and got %d ", expectedNumberOfRecords, actualNumberOfRecords)
	}

}
func TestRecordWithDifferentNumberOfFields(t *testing.T) {
	//Each record should contain the same number of fields throughout the file
	csvData := "Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com"

	reader := csv.NewReader(strings.NewReader(csvData))

	_, err := reader.ReadAll()

	if err != nil {
		expectedErrorMessage := errors.New("record on line 4: wrong number of fields")
		fmt.Println(expectedErrorMessage)
		fmt.Println(err.Error())
		if err.Error() == expectedErrorMessage.Error() {
			t.Logf("Expected error %e:", expectedErrorMessage)
		} else {
			t.Errorf("Unexpected error reading CSV file: %e", err)
		}

	}

}

func TestRecordFieldsWithSpaces(t *testing.T) {
	//Spaces are considered part of a field and should not be ignored.
	csvData := "Name,Email\nJohn,john@example.com    \nJane    ,jane@example.com\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com       \nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry    ,henry@example.com"

	reader := csv.NewReader(strings.NewReader(csvData))

	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Unexpected error reading CSV file: %e", err)
	}

	expectedOutput := [][]string{
		{"Name", "Email"},
		{"John", "john@example.com    "},
		{"Jane    ", "jane@example.com"},
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
		{"Charlie", "charlie@example.com"},
		{"Diana", "diana@example.com       "},
		{"Eva", "eva@example.com"},
		{"Frank", "frank@example.com"},
		{"Grace", "grace@example.com"},
		{"Henry    ", "henry@example.com"},
	}

	for rowIndex, record := range records {
		for fieldIndex, field := range record {
			if field != expectedOutput[rowIndex][fieldIndex] {
				t.Errorf("Expected field: %s,  and actual field: %s", expectedOutput[rowIndex][fieldIndex], field)
			}
		}
	}
}

func TestLastFieldInRecordFollowedByComma(t *testing.T) {
	//The last field in the record must not be followed by a comma.
	csvData := "Name,Email\nJohn,john@example.com\nJane,jane@example.com\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com,"

	reader := csv.NewReader(strings.NewReader(csvData))

	_, err := reader.ReadAll()

	if err != nil {
		expectedErrorMessage := errors.New("record on line 11: wrong number of fields")
		fmt.Println(expectedErrorMessage)
		fmt.Println(err.Error())
		if err.Error() == expectedErrorMessage.Error() {
			t.Logf("Expected error %e:", expectedErrorMessage)
		} else {
			t.Errorf("Unexpected error reading CSV file: %e", err)
		}
	}
}

func TestQuotesInFieldNotEnclosedWithDoubleQuotes(t *testing.T) {
	//If fields are not enclosed with double quotes, then double quotes may not appear inside the fields
	csvData := "Name,Email\nJohn,john@example.com\nJane,\"jane@example.com\"\nAlice,alice@example.com\n\"Bob\",\"bob@example.com\"\nCharlie,charlie@example.com\n\"Diana\",\"diana@example.com\"\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com"

	reader := csv.NewReader(strings.NewReader(csvData))
	reader.LazyQuotes = false
	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Unexpected error reading CSV file: %e", err)
	}

	expectedOutput := [][]string{
		{"Name", "Email"},
		{"John", "john@example.com"},
		{"Jane", "jane@example.com"},
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
		{"Charlie", "charlie@example.com"},
		{"Diana", "diana@example.com"},
		{"Eva", "eva@example.com"},
		{"Frank", "frank@example.com"},
		{"Grace", "grace@example.com"},
		{"Henry", "henry@example.com"},
	}

	for rowIndex, record := range records {
		for fieldIndex, field := range record {
			if field != expectedOutput[rowIndex][fieldIndex] {
				t.Errorf("Expected field: %s,  and actual field: %s", expectedOutput[rowIndex][fieldIndex], field)
			}
		}
	}

}

func TestDoubleQuotesInsideQuotedField(t *testing.T) {
	//If double-quotes are used to enclose fields, then a double-quote appearing inside a field
	// must be escaped by preceding it with another double quote.
	csvData := "Name,Email\nJohn,john@example.com\nJane,\"jane\"\"@example.com\"\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,henry@example.com"
	reader := csv.NewReader(strings.NewReader(csvData))
	reader.LazyQuotes = true // to treat double quotes as part of the field value
	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Unexpected error reading CSV file: %e", err)
	}
	expectedOutput := [][]string{
		{"Name", "Email"},
		{"John", "john@example.com"},
		{"Jane", "jane\"@example.com"},
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
		{"Charlie", "charlie@example.com"},
		{"Diana", "diana@example.com"},
		{"Eva", "eva@example.com"},
		{"Frank", "frank@example.com"},
		{"Grace", "grace@example.com"},
		{"Henry", "henry@example.com"},
	}

	for rowIndex, record := range records {
		for fieldIndex, field := range record {
			if field != expectedOutput[rowIndex][fieldIndex] {
				t.Errorf("Expected field: %s,  and actual field: %s", expectedOutput[rowIndex][fieldIndex], field)
			}
		}
	}
}

func TestFieldsWithLineBreaksInsideQuotes(t *testing.T) {
	csvData := "Name,Email\nJohn,john@example.com\nJane,\"jane   \n@example.com\"\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,\"henry@example.com       \n\""
	reader := csv.NewReader(strings.NewReader(csvData))
	reader.LazyQuotes = true // to treat double quotes as part of the field value
	records, err := reader.ReadAll()
	fmt.Println(records)

	if err != nil {
		t.Errorf("Unexpected error reading CSV file: %e", err)
	}
	expectedOutput := [][]string{
		{"Name", "Email"},
		{"John", "john@example.com"},
		{"Jane", "jane   \n@example.com"},
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
		{"Charlie", "charlie@example.com"},
		{"Diana", "diana@example.com"},
		{"Eva", "eva@example.com"},
		{"Frank", "frank@example.com"},
		{"Grace", "grace@example.com"},
		{"Henry", "henry@example.com       \n"},
	}

	for rowIndex, record := range records {
		for fieldIndex, field := range record {
			if field != expectedOutput[rowIndex][fieldIndex] {
				t.Errorf("Expected field: %s,  and actual field: %s", expectedOutput[rowIndex][fieldIndex], field)
			}
		}
	}

}
func TestFieldsWithCommasInsideQuotes(t *testing.T) {
	csvData := "Name,Email\nJohn,john@example.com\nJane,\"jane,@example.com\"\nAlice,alice@example.com\nBob,bob@example.com\nCharlie,charlie@example.com\nDiana,diana@example.com\nEva,eva@example.com\nFrank,frank@example.com\nGrace,grace@example.com\nHenry,\"henry,@example.com\""
	reader := csv.NewReader(strings.NewReader(csvData))
	reader.LazyQuotes = true // to treat double quotes as part of the field value
	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Unexpected error reading CSV file: %e", err)
	}
	expectedOutput := [][]string{
		{"Name", "Email"},
		{"John", "john@example.com"},
		{"Jane", "jane,@example.com"},
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
		{"Charlie", "charlie@example.com"},
		{"Diana", "diana@example.com"},
		{"Eva", "eva@example.com"},
		{"Frank", "frank@example.com"},
		{"Grace", "grace@example.com"},
		{"Henry", "henry,@example.com"},
	}

	for rowIndex, record := range records {
		for fieldIndex, field := range record {
			if field != expectedOutput[rowIndex][fieldIndex] {
				t.Errorf("Expected field: %s,  and actual field: %s", expectedOutput[rowIndex][fieldIndex], field)
			}
		}
	}

}

func main() {

	testing.Main(nil, nil, nil, nil)
}
