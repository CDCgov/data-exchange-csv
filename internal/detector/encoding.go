package detector

import (
	"unicode/utf8"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func DetectEncoding(data []byte) constants.EncodingType {

	if isValidUSASCII(data) {
		return constants.UTF8
	}

	if isValidISO_8859_1(data) {
		return constants.ISO8859_1
	}

	if isValidWindows1252(data) {
		return constants.WINDOWS1252
	}

	if utf8.Valid(data) {
		return constants.UTF8
	}
	return constants.UNDEF

}

func isValidUSASCII(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	for _, byteVal := range data {
		if byteVal&constants.MSBMask != 0 {
			return false
		}
	}
	return true
}

func isValidISO_8859_1(data []byte) bool {
	if len(data) == 0 {
		return false
	}

	for _, byteVal := range data {
		if byteVal >= constants.InvalidStartISO88591 && byteVal <= constants.InvalidEndISO88591 {
			return false
		}
	}
	return true
}

func isValidWindows1252(data []byte) bool {
	if len(data) == 0 {
		return false
	}

	for _, byteVal := range data {
		if byteVal > constants.ValidEndWindows1252 {
			return false
		}
	}
	return true
}
