package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func SetupEnvironment(root string) (err error) {
	//create file folder
	fileDir := filepath.Join(root, constants.FILE)
	err = os.MkdirAll(fileDir, os.ModeNamedPipe)

	if err != nil {
		return fmt.Errorf("failed to create directory")
	}
	//create row folder
	rowDir := filepath.Join(root, constants.ROW)
	err = os.MkdirAll(rowDir, os.ModeNamedPipe)

	if err != nil {
		return fmt.Errorf("failed to create directory")
	}

	return nil
}
