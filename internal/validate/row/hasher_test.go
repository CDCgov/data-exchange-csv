package row

import (
	"crypto/sha256"
	"encoding/base64"
	"testing"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func computeExpectedHash(concatenatedRow string) string {
	hash := sha256.Sum256([]byte(concatenatedRow))
	hashSlice := hash[:]

	return base64.StdEncoding.EncodeToString(hashSlice)
}

func TestComputeHash(t *testing.T) {
	tests := []struct {
		name      string
		row       []string
		delimiter rune
		expected  string
	}{
		{
			name:      "CSV file with the valid input",
			row:       []string{"George", "30", "Engineer"},
			delimiter: constants.COMMA,
			expected:  computeExpectedHash("George,30,Engineer"),
		},
		{
			name:      "TSV file with the valid input",
			row:       []string{"Bob", "25", "Doctor"},
			delimiter: constants.TAB,
			expected:  computeExpectedHash("Bob\t25\tDoctor"),
		},
	}
	for _, test := range tests {
		result := ComputeHash(test.row, string(test.delimiter))
		if result != test.expected {
			t.Errorf("ActualHash(%v, %q) = %q; ExpectedHash %q", test.row, test.delimiter, result, test.expected)
		}
	}

}
