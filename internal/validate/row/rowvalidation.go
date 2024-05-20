package row

import "github.com/CDCgov/data-exchange-csv/cmd/internal/constants"

type Row struct {
	FileUUID  constants.UUID
	RowNumber int
	RowUUID   constants.UUID
	Location  string
	Hash      []byte
}

func IsValid(row string) (bool, error) {
	// verify the row agains rules/requirements

	return true, nil
}
