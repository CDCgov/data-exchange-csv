package utils

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func ReadFileRandomly(file *os.File) ([]rune, error) {

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

		//We need to convert []byte to []rune and return
		return []rune(string(buffer)), nil
	}

	randomRunes := make([]rune, 0, constants.MAX_READ_THRESHOLD)

	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano()))

	reader := bufio.NewReader(file)

	startTime := time.Now()

	for len(randomRunes) < constants.MAX_READ_THRESHOLD && time.Since(startTime) < constants.MAX_EXECUTION_TIME {
		offset := randomNumber.Int63n(fileSize)
		_, err := file.Seek(offset, 0)

		if err != nil {
			return nil, err
		}

		r, _, err := reader.ReadRune()

		if err != nil {
			return nil, err
		}

		randomRunes = append(randomRunes, r)
	}

	// Reset the file pointer to the beginning of the file so that next operation
	// starts reading from the begining.
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return randomRunes, nil

}
