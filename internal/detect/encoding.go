package detect

import (
	"unicode/utf8"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func Encoding(data []byte) constants.EncodingType {

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
	// Initialize most significant bit
	const MSBMask = 0x80
	for _, byteVal := range data {
		if byteVal&MSBMask != 0 {
			return false
		}
	}
	return true
}

func isValidISO_8859_1(data []byte) bool {
	const (
		invalidStart = 0x80 // as decimal 128
		invalidEnd   = 0x9F // as decimal 159
	)
	for _, byteVal := range data {
		if byteVal >= invalidStart && byteVal <= invalidEnd {
			return false
		}
	}
	return true
}

func isValidWindows1252(data []byte) bool {
	const (
		validStart = 0x80 // as decimal 128
		validEnd   = 0x9F // as decimal 159
	)
	for _, byteVal := range data {
		if byteVal >= validStart && byteVal <= validEnd {
			return true
		}
	}
	return false
}
