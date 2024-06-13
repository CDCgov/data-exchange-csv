package row

import (
	"crypto/sha256"
	"strings"
)

func ComputeHash(row []string) [32]byte {
	contatenatedRow := strings.Join(row, ",")
	hash := sha256.Sum256([]byte(contatenatedRow))
	return hash
}
