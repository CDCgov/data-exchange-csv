package detector

import "github.com/CDCgov/data-exchange-csv/cmd/internal/constants"

func DetectDelimiter(data []byte) byte {

	delimiters := []byte{constants.COMMA, constants.TAB}
	delimiterFrequency := make(map[byte]int, len(delimiters))

	for _, delimiter := range delimiters {
		delimiterFrequency[delimiter] = 0
	}

	for _, character := range data {
		if _, exists := delimiterFrequency[character]; exists {
			delimiterFrequency[character]++
		}
	}

	var mostFrequentDelimiter int
	var detectedDelimiter byte

	for delimiter, frequency := range delimiterFrequency {
		if frequency > mostFrequentDelimiter {
			mostFrequentDelimiter = frequency
			detectedDelimiter = delimiter
		}
	}

	return detectedDelimiter
}
