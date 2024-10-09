package models

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
	"github.com/google/uuid"
)

type FileError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type FileValidateInputParams struct {
	ReceivedFile       string                 `json:"received_file"`
	Encoding           constants.EncodingType `json:"encoding"`
	Separator          rune                   `json:"separator"`
	HasHeader          bool                   `json:"has_header"`
	ValidationCallback func(FileValidationResult)
	Destination        string `json:"destination"`
	ConfigFile         string `json:"config_file"`
	Debug              bool   `json:"debug"`
	LogToFile          bool   `json:"log-file"`
}

type ConfigFields struct {
	HasHeader bool                   `json:"has_header"`
	Separator string                 `json:"separator"`
	Encoding  constants.EncodingType `json:"encoding"`
}

type FileValidationParams struct {
	FileUUID     uuid.UUID              `json:"file_uuid"`
	ReceivedFile string                 `json:"received_filename"`
	Encoding     constants.EncodingType `json:"detected_encoding"`
	Delimiter    rune                   `json:"detected_delimiter"`
	Header       []string               `json:"header"`
}

type FileValidationResult struct {
	ReceivedFile string                 `json:"received_filename"`
	Encoding     constants.EncodingType `json:"encoding"`
	FileUUID     uuid.UUID              `json:"uuid"`
	SizeInBytes  int64                  `json:"size_bytes"`
	Delimiter    rune                   `json:"delimiter"`
	Error        *FileError             `json:"error"`
	Status       string                 `json:"status"`
	HasHeader    bool                   `json:"has_header"`
	Destination  string                 `json:"dest_folder"`
}
