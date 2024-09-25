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

type ConfigIdentifier struct {
	DataStreamID    string   `json:"data_stream_id"`
	DataStreamRoute string   `json:"data_stream_route"`
	Header          []string `json:"header"`
}

type ConfigValidationResult struct {
	Error                  *FileError             `json:"error"`
	Status                 string                 `json:"status"`
	HeaderValidationResult HeaderValidationResult `json:"header_validation_result"`
}

type HeaderValidationResult struct {
	Status string     `json:"status"`
	Error  *FileError `json:"error"`
	Header []string   `json:"header"`
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
	Size         int64                  `json:"size"`
	Delimiter    rune                   `json:"delimiter"`
	Error        *FileError             `json:"error"`
	Status       string                 `json:"status"`
	HasHeader    bool                   `json:"has_header"`
	Destination  string                 `json:"dest_folder"`
}

type RowCallbackParams struct {
	Destination          string      `json:"destination"`
	ValidationResult     interface{} `json:"validation_result"`
	TransformationResult interface{} `json:"transformation_result"`
	IsFirst              bool        `json:"is_first"`
	IsLast               bool        `json:"is_last"`
}
