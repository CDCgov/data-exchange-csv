package utils

import (
	"strings"

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
	// Count occurrences of common delimiters
	supportedDelimiters := []rune{constants.COMMA, constants.TAB}
	counts := make(map[rune]int)
	for _, candidate := range supportedDelimiters {
		counts[candidate] = strings.Count(sample, string(candidate))
	}

	// Return the most common delimiter
	var maxCount int
	var detectedDelimiter rune
	for r, count := range counts {
		if count > maxCount {
			maxCount = count
			detectedDelimiter = r
		}
	}

	return detectedDelimiter
}
