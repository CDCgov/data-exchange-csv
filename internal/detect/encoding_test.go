package detect

import (
	"testing"
)

func TestIsValidUSASCII(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{"Empty Input Test", []byte{}, false},
		{"Valid Single Byte Sequence", []byte{45, 65, 56, 87, 122}, true},
		{"Invalid Single Byte Sequence", []byte{129, 133, 140, 150, 177}, false},
		{"Mix single Byte and invalid >single Byte Sequence", []byte{122, 234, 23, 112}, false},
		{"All Valid bytes", []byte{110, 111, 112, 124, 126, 98, 78}, true},
		{"All invalid bytes", []byte{129, 130, 131, 132, 133, 134}, false},
		{"Start with valid byte, followed by invalid bytes", []byte{125, 200, 201, 222, 223}, false},
		{"Start with invalid byte, followed by valid bytes", []byte{200, 125, 110, 112}, false},
		{"Byte above upper range", []byte{128, 130, 131}, false},
		{"Byte below lower range", []byte{0, 11, 112}, true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isValidUSASCII(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %v, got %v for input %v", testCase.expected, result, testCase.input)
			}
		})
	}

}

func TestIsValidUTF8(t *testing.T) {

}
func TestIsValidWindows1552(t *testing.T) {

}
func TestIsValidISO8859_1(t *testing.T) {

}
