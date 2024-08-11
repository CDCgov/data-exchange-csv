package detector

import (
	"unicode/utf8"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func DetectEncoding(data []rune) constants.EncodingType {

	if isValidUSASCII(data) {
		return constants.UTF8
	}

	if isValidISO_8859_1(data) {
		return constants.ISO8859_1
	}

	if isValidWindows1252(data) {
		return constants.WINDOWS1252
	}

	if utf8.Valid([]byte(string(data))) {
		return constants.UTF8
	}
	return constants.UNDEF

}

func isValidUSASCII(data []rune) bool {
	if len(data) == 0 {
		return false
	}
	for _, runeVal := range data {
		if runeVal >= constants.MSBMask {
			return false
		}
	}
	return true
}
func isValidISO_8859_1(data []rune) bool {
	if len(data) == 0 {
		return false
	}

	for _, runeVal := range data {
		if (runeVal >= constants.MSBMask && runeVal <= constants.Windows1252RunThreshold) || runeVal > constants.SingleByteSequenceEnd {
			return false
		}
	}
	return true
}
func isValidWindows1252(data []rune) bool {
	if len(data) == 0 {
		return false
	}

	for _, runeVal := range data {
		if runeVal > constants.SingleByteSequenceEnd {
			return false
		}
	}
	return true
}
