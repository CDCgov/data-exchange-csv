package detect

import (
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func TestIsValidUSASCII(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{constants.EMPTY_INPUT, []byte{}, false},
		{constants.VALID_SINGLE_BYTE_SEQUENCE, []byte{45, 65, 56, 87, 122}, true},
		{constants.INVALID_SINGLE_BYTE_SEQUENCE, []byte{129, 133, 140, 150, 177}, false},
		{constants.MIXED_BYTE_SEQUENCES, []byte{122, 234, 23, 112}, false},
		{constants.ALL_VALID_BYTES, []byte{110, 111, 112, 124, 126, 98, 78}, true},
		{constants.ALL_INVALID_BYTES, []byte{129, 130, 131, 132, 133, 134}, false},
		{constants.VALID_FIRST_BYTE_SEQUENCE, []byte{125, 200, 201, 222, 223}, false},
		{constants.INVALID_FIRST_BYTE_SEQUENCE, []byte{200, 125, 110, 112}, false},
		{constants.BYTE_ABOVE_UPPER_RANGE, []byte{128, 130, 131}, false},
		{constants.BYTE_BELOW_LOWER_RANGE, []byte{0, 11, 112}, true},
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

func TestIsValidISO8859_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{constants.EMPTY_INPUT, []byte{}, false},
		{constants.VALID_SINGLE_BYTE_SEQUENCE, []byte{45, 65, 56, 87, 122}, true},
		{constants.INVALID_SINGLE_BYTE_SEQUENCE, []byte{129, 133, 140, 150, 159}, false},
		{constants.MIXED_BYTE_SEQUENCES, []byte{122, 130, 234, 23, 112}, false},
		{constants.ALL_VALID_BYTES, []byte{160, 165, 166, 200, 180, 189, 177}, true},
		{constants.ALL_INVALID_BYTES, []byte{129, 130, 131, 132, 133, 134}, false},
		{constants.VALID_FIRST_BYTE_SEQUENCE, []byte{125, 130, 134, 140, 150}, false},
		{constants.INVALID_FIRST_BYTE_SEQUENCE, []byte{128, 125, 110, 112}, false},
		{constants.BYTE_ABOVE_UPPER_RANGE, []byte{160, 161, 162}, true},
		{constants.BYTE_BELOW_LOWER_RANGE, []byte{127, 11, 112}, true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isValidISO_8859_1(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %v, got %v for input %v", testCase.expected, result, testCase.input)
			}
		})
	}

}

func TestIsValidWindows1552(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{constants.EMPTY_INPUT, []byte{}, false},
		{constants.VALID_SINGLE_BYTE_SEQUENCE, []byte{129, 133, 140, 150, 160}, false},
		{constants.ALL_VALID_BYTES, []byte{129, 159, 158, 152, 151, 140, 141}, true},
		{constants.ALL_INVALID_BYTES, []byte{1, 126}, false},
		{constants.VALID_FIRST_BYTE_SEQUENCE, []byte{128, 120, 110, 111, 99}, false},
		{constants.INVALID_FIRST_BYTE_SEQUENCE, []byte{127, 128, 129, 130}, false},
		{constants.BYTE_ABOVE_UPPER_RANGE, []byte{129, 130, 131}, true},
		{constants.BYTE_BELOW_LOWER_RANGE, []byte{127, 11, 112}, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isValidWindows1252(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %v, got %v for input %v", testCase.expected, result, testCase.input)
			}
		})
	}
}
