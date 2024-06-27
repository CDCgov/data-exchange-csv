package utils

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func ReadFileRandomly(file *os.File) ([]byte, error) {

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()

	if fileSize < int64(constants.MAX_READ_THRESHOLD) {
		buffer := make([]byte, fileSize)
		_, err := file.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}

	randomBytes := make([]byte, 0, constants.MAX_READ_THRESHOLD)

	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano()))

	reader := bufio.NewReader(file)

	startTime := time.Now()

	for len(randomBytes) < constants.MAX_READ_THRESHOLD && time.Since(startTime) < constants.MAX_EXECUTION_TIME {
		offset := randomNumber.Int63n(fileSize)
		_, err := file.Seek(offset, 0)

		if err != nil {
			return nil, err
		}

		buffer := make([]byte, constants.MAX_READ_THRESHOLD-len(randomBytes))
		n, err := reader.Read(buffer)

		if err != nil {
			return nil, err
		}
		randomBytes = append(randomBytes, buffer[:n]...)
	}
	// reset the pointer in the begining of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil

}
