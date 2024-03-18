package main

import (
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

func TestLineBreakCR(t *testing.T) {
	//This test fails as csv package in Go  does not support carriage returns as line breaker
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
func TestLineBreakCRLF(t *testing.T) {
	//sample csv data
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
	//sample csv data
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

func main() {

	testing.Main(nil, nil, nil, nil)
}
