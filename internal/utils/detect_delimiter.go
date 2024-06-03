package utils

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

type Delimiter struct {
	Character rune
}

type DelimiterDetector struct{}

func (s *DelimiterDetector) DetectDelimiter(sample string) *Delimiter {
	delimiterAsRune := &Delimiter{Character: constants.COMMA}

	// Try to detect most frequently repeated delimiter
	delimiter := s.mostFrequentDelimiter(sample)
	if delimiter != 0 {
		delimiterAsRune.Character = delimiter
	} else {
		delimiterAsRune.Character = 0
	}

	return delimiterAsRune
}

func (s *DelimiterDetector) mostFrequentDelimiter(sample string) rune {

	supportedDelimiters := []rune{constants.COMMA, constants.TAB}

	counts := make(map[rune]int, len(supportedDelimiters))

	for _, delimiter := range supportedDelimiters {
		counts[delimiter] = 0
	}

	for _, char := range sample {
		if _, exists := counts[char]; exists {
			counts[char]++
		}
	}

	var mostFrequent int
	var detectedDelimiter rune
	for r, count := range counts {
		if count > mostFrequent {
			mostFrequent = count
			detectedDelimiter = r
		}
	}
	return detectedDelimiter
}
