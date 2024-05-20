package utils

import (
	"math/rand"
	"os"
	"time"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func GetRandomSample(file *os.File) ([]byte, error) {
	threshold := constants.MAX_READ
	var randomBytes []byte

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	if fileSize < int64(threshold) {
		//load the entire file
		buffer := make([]byte, fileSize)
		_, err := file.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}
	// initialize random seed to generate different sequences of random numbers
	randomNumberGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(randomBytes) < threshold && fileSize > 0 {
		//generate random offset
		offset := randomNumberGenerator.Int63n(fileSize)
		//seek to the random offset
		_, err := file.Seek(offset, 0)
		if err != nil {
			return nil, err
		}

		buffer := make([]byte, threshold)

		n, err := file.Read(buffer)

		if err != nil {
			return nil, err
		}
		randomBytes = append(randomBytes, buffer[:n]...)
	}

	return randomBytes, nil

}
