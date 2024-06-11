package detect

import (
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func TestDelimiterComma(t *testing.T) {
	csvData := "José Díaz,Software engineer, working on CSV & Golang."
	csvDataAsBytes := []byte(csvData)

	expectedDelimiter := constants.CSV
	actualDelimiter := constants.DelimiterCharacters[Delimiter(csvDataAsBytes)]

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %s records, and got %s ", expectedDelimiter, actualDelimiter)
	}

}
func TestDelimiterTab(t *testing.T) {
	tsvData := "Index\tName\tDescription\nJosé Díaz\tSoftware engineer\tworking on C++ & Python.\nFrançois Dupont\tProduct manager: oversees marketing & sales."
	tsvDataAsBytes := []byte(tsvData)

	expectedDelimiter := constants.TSV
	actualDelimiter := constants.DelimiterCharacters[Delimiter(tsvDataAsBytes)]

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %s records, and got %s ", expectedDelimiter, actualDelimiter)
	}
}
func TestUnsupportedDelimiter(t *testing.T) {
	invalidData := "name|age\nOlya|64\nBobby|68"
	invalidDataAsBytes := []byte(invalidData)

	expectedDelimiter := constants.UNSUPPORTED
	actualDelimiter := constants.DelimiterCharacters[Delimiter(invalidDataAsBytes)]

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %s records, and got %s ", expectedDelimiter, actualDelimiter)
	}

}
