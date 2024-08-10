package models

import (
	"encoding/json"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/google/uuid"
)

type RowValidationResult struct {
	FileUUID  uuid.UUID `json:"file_uuid"`
	RowNumber int       `json:"row_number"`
	RowUUID   uuid.UUID `json:"row_uuid"`
	Hash      string    `json:"row_hash"`
	Error     *RowError `json:"error"`
	Status    string    `json:"status"`
}

type RowTransformationResult struct {
	FileUUID uuid.UUID       `json:"file_uuid"`
	RowUUID  uuid.UUID       `json:"row_uuid"`
	JsonRow  json.RawMessage `json:"json_row"`
	Error    error           `json:"error"`
	Status   string          `json:"status"`
}

type RowError struct {
	Message  string             `json:"message"`
	Line     int                `json:"line"`
	Column   int                `json:"column"`
	Severity constants.Severity `json:"severity"`
}
