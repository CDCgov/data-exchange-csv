package row

import (
	"encoding/csv"
	"io"

	"github.com/google/uuid"
)

type ValidationResult struct {
	FileUUID  uuid.UUID `json:"file_uuid"`
	RowNumber int       `json:"row_number"`
	RowUUID   uuid.UUID `json:"row_uuid"`
	Hash      [32]byte  `json:"row_hash"`
	Error     error     `json:"error"`
}

func (vr *ValidationResult) Validate(reader *csv.Reader, fileUUID uuid.UUID, seperator string) {
	vr.FileUUID = fileUUID

	rowCount := 0

	for {
		row, err := reader.Read()

		vr.Hash = ComputeHash(row, seperator)

		if err == io.EOF {
			break
		}
		vr.RowNumber = rowCount
		rowCount++

		vr.RowUUID = uuid.New()

		if err != nil {
			//send to DLQ
			// PS API
			//TODO-ERROR Object
			vr.Error = err
		}
		//
		// valid row
		// ready to transform to json

	}
}
