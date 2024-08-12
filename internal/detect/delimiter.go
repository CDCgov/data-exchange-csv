package detect

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func Delimiter(data string) rune {

	delimiters := []rune{constants.COMMA, constants.TAB}
	delimiterFrequency := make(map[rune]int, len(delimiters))

	for _, delimiter := range delimiters {
		delimiterFrequency[delimiter] = 0
	}

	for _, character := range data {
		if _, exists := delimiterFrequency[character]; exists {
			delimiterFrequency[character]++
		}
	}

	var mostFrequentDelimiter int
	var detectedDelimiter rune

	for delimiter, frequency := range delimiterFrequency {
		if frequency > mostFrequentDelimiter {
			mostFrequentDelimiter = frequency
			detectedDelimiter = delimiter
		}
	}

	return detectedDelimiter
}
