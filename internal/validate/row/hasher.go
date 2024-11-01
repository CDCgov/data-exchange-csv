package row

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func ComputeHash(row []string, delimiter rune) string {
	separator := constants.COMMA

	if delimiter == constants.TAB {
		separator = constants.TAB
	}
	concatenatedRow := strings.Join(row, string(separator))
	hash := sha256.Sum256([]byte(concatenatedRow))

	//we need to convert [32]byte to []byte before using base64.StdEncoding.EncodeToString()
	hashSlice := hash[:]

	return base64.StdEncoding.EncodeToString(hashSlice)
}
