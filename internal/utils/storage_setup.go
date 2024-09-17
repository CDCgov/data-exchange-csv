package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func SetupEnvToStoreResults(root string) (err error) {
	fileDir := filepath.Join(root, constants.FILE)
	rowDir := filepath.Join(root, constants.ROW)

	err = os.MkdirAll(fileDir, os.ModeNamedPipe)

	if err != nil {
		return fmt.Errorf("failed to create directory")
	}
	err = os.MkdirAll(rowDir, os.ModeNamedPipe)

	if err != nil {
		return fmt.Errorf("failed to create directory")
	}

	return nil
}
