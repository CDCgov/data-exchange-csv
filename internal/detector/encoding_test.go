package detector

import (
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func TestDetectEncoding(t *testing.T) {
	testCases := []struct {
		name     string
		input    []rune
		expected constants.EncodingType
	}{
		// Note: The Go UTF-8 package will return true for an empty rune slice, indicating that it is valid UTF-8.
		{constants.EMPTY_INPUT, []rune{}, constants.UTF8},
		/*
			Ensure that encoding is UTF-8 for valid US-ASCII characters.
			UTF-8 uses a single-byte sequence for ASCII characters, this means any valid ASCII character
			is also a valid UTF-8.
		*/
		{constants.VALID_US_ASCII, []rune("Hello World"), constants.UTF8},
		{constants.INVALID_US_ASCII, []rune("Smörgåsbord"), constants.ISO8859_1},
		{constants.VALID_WINDOWS_1252, []rune{128, 142, 149, 151, 153, 154}, constants.WINDOWS1252},
		{constants.INVALID_WINDOWS_1252, []rune{159, 160, 172, 200, 256, 98, 78}, constants.UTF8},
		{constants.VALID_ISO_8859_1, []rune("Surströmming"), constants.ISO8859_1},
		{constants.INVALID_ISO_8859_1, []rune{126, 200, 201, 256, 300, 408}, constants.UTF8},
		{constants.VALID_UTF_8, []rune{256, 255, 300, 340, 50}, constants.UTF8}}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := DetectEncoding(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %v, got %v for input %v", testCase.expected, result, testCase.input)
			}
		})
	}

}
