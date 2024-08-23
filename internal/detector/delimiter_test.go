package detector

import (
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func TestDelimiterComma(t *testing.T) {
	csvData := "José Díaz,Software engineer, working on CSV & Golang."
	csvDataAsBytes := []rune(csvData)

	expectedDelimiter := constants.COMMA
	actualDelimiter := constants.DelimiterCharacters[DetectDelimiter(csvDataAsBytes)]

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %c records, and got %c ", expectedDelimiter, actualDelimiter)
	}

}
func TestDelimiterTab(t *testing.T) {
	tsvData := "Index\tName\tDescription\nJosé Díaz\tSoftware engineer\tworking on C++ & Python.\nFrançois Dupont\tProduct manager: oversees marketing & sales."
	tsvDataAsBytes := []rune(tsvData)

	expectedDelimiter := constants.TAB
	actualDelimiter := constants.DelimiterCharacters[DetectDelimiter(tsvDataAsBytes)]

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %c records, and got %c ", expectedDelimiter, actualDelimiter)
	}
}
func TestUnsupportedDelimiter(t *testing.T) {
	invalidData := "name|age\nOlya|64\nBobby|68"
	invalidDataAsBytes := []rune(invalidData)

	expectedDelimiter := constants.UNSUPPORTED
	actualDelimiter := DetectDelimiter(invalidDataAsBytes)

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %c records, and got %c ", expectedDelimiter, actualDelimiter)
	}

}
