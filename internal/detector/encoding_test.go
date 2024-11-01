package detector

import (
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

const (
	EMPTY_INPUT          = "Empty Input Test"
	VALID_US_ASCII       = "US-ASCII Valid Input"
	INVALID_US_ASCII     = "US-ASCII Invalid Input"
	VALID_WINDOWS_1252   = "Windows 1252 Valid Input"
	INVALID_WINDOWS_1252 = "Windows 1252 Invalid Input"
	VALID_ISO_8859_1     = "ISO 8859-1 Valid Input"
	INVALID_ISO_8859_1   = "ISO 8859-1 Invalid Input"
	VALID_UTF_8          = "UTF-8 Valid Input"
	INVALID_UTF_8        = "UTF-8 Invalid Input"
)

func TestDetectEncoding(t *testing.T) {
	testCases := []struct {
		name     string
		input    []rune
		expected constants.EncodingType
	}{
		// Note: The Go UTF-8 package will return true for an empty rune slice, indicating that it is valid UTF-8.
		{EMPTY_INPUT, []rune{}, constants.UTF8},
		/*
			Ensure that encoding is UTF-8 for valid US-ASCII characters.
			UTF-8 uses a single-byte sequence for ASCII characters, this means any valid ASCII character
			is also a valid UTF-8.
		*/
		{VALID_US_ASCII, []rune("Hello World"), constants.UTF8},
		{INVALID_US_ASCII, []rune("Smörgåsbord"), constants.ISO8859_1},
		{VALID_WINDOWS_1252, []rune("Börš"), constants.WINDOWS1252},
		{INVALID_WINDOWS_1252, []rune("田中太郎"), constants.UTF8},
		{VALID_ISO_8859_1, []rune("Surströmming"), constants.ISO8859_1},
		{INVALID_ISO_8859_1, []rune("山田花子"), constants.UTF8},
		{VALID_UTF_8, []rune("山田花子"), constants.UTF8}}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := DetectEncoding(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %v, got %v for input %v", testCase.expected, result, testCase.input)
			}
		})
	}

}
