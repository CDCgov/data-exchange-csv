package utils

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"golang.org/x/net/html/charset"
)

func DetectEncoding(randomBytes []byte) constants.EncodingType {
	fmt.Println("THIS IS TEST")

	// Detect encoding
	var encoding constants.EncodingType

	if utf8.Valid(randomBytes) {

		fmt.Println("THIS IS VALID UTF-8")
		encoding = constants.UTF8
		//still need to call functions to check character sets if
	}
	enc, name, certain := charset.DetermineEncoding(randomBytes, "")
	fmt.Println(enc, certain)
	fmt.Println("Detected encoding is ", name)
	fmt.Println(isValidUSASCII(randomBytes))
	fmt.Println(isValidISO_8859_1(randomBytes))
	fmt.Println(isValidWindows1252(randomBytes))

	return encoding

}

func isValidUSASCII(bytes []byte) bool {
	for _, b := range bytes {
		if b > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func isValidISO_8859_1(bytes []byte) bool {
	for _, b := range bytes {
		if b > 127 && b < 160 {
			return false
		}
	}
	return true
}

func isValidWindows1252(bytes []byte) bool {
	for _, b := range bytes {
		if (b < 32 && b != 9 && b != 10 && b != 13) || (b > 127 && b < 160) {
			fmt.Println("its not windows1252 due to ", b, string(b))
			return false
		}

	}
	return true
}
