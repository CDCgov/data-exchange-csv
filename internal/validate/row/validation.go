package row

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/google/uuid"
)

type ValidationResult struct {
	FileUUID  uuid.UUID `json:"file_uuid"`
	RowNumber int       `json:"row_number"`
	RowUUID   uuid.UUID `json:"row_uuid"`
	Hash      []byte    `json:"row_hash"`
	Error     error     `json:"error"`
}

func (vr *ValidationResult) Validate(reader *csv.Reader, fileUUID uuid.UUID) (bool, error) {
	// verify the row against rules/requirements
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("error ", err)
		}

		fmt.Println(row)

	}
	return true, nil
}
