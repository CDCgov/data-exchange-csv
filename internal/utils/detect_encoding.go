package utils

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"golang.org/x/net/html/charset"
)

func DetectEncoding(randomBytes []byte) constants.EncodingType {

	var encoding constants.EncodingType

	if utf8.Valid(randomBytes) {
		encoding = constants.UTF8
		//TBD- should unparseable characters be checked?
	} else {
		_, name, _ := charset.DetermineEncoding(randomBytes, "")
		encoding = constants.EncodingType(name)
	}
	//TBD-How to utilize the functions below
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
		if b < 32 && b != 9 && b != 10 && b != 13 {
			return false
		}
	}
	return true
}
