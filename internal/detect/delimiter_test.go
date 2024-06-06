package detect

import (
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func TestDelimiterComma(t *testing.T) {
	csvData := "José Díaz,Software engineer, working on CSV & Golang."

	expectedDelimiter := constants.CSV
	actualDelimiter := constants.DelimiterCharacters[Delimiter(csvData)]

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %s records, and got %s ", expectedDelimiter, actualDelimiter)
	}

}
func TestDelimiterTab(t *testing.T) {
	tsvData := "Index\tName\tDescription\nJosé Díaz\tSoftware engineer\tworking on C++ & Python.\nFrançois Dupont\tProduct manager: oversees marketing & sales."

	expectedDelimiter := constants.TSV
	actualDelimiter := constants.DelimiterCharacters[Delimiter(tsvData)]

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %s records, and got %s ", expectedDelimiter, actualDelimiter)
	}
}
func TestUnsupportedDelimiter(t *testing.T) {
	invaliData := "name|age\nOlya|64\nBobby|68"

	expectedDelimiter := constants.UNSUPPORTED
	actualDelimiter := constants.DelimiterCharacters[Delimiter(invaliData)]

	if expectedDelimiter != actualDelimiter {
		t.Errorf("Expected %s records, and got %s ", expectedDelimiter, actualDelimiter)
	}

}
