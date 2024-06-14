package row

import (
	"crypto/sha256"
	"strings"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func ComputeHash(row []string, delimiter string) [32]byte {
	separator := string(constants.CSV)

	if delimiter == constants.TSV {
		separator = string(constants.TSV)
	}

	contatenatedRow := strings.Join(row, separator)
	hash := sha256.Sum256([]byte(contatenatedRow))

	return hash
}
