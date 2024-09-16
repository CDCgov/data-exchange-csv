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
	FileURL            string                 `json:"file_url"`
	Encoding           constants.EncodingType `json:"encoding"`
	Separator          rune                   `json:"separator"`
	HasHeader          bool                   `json:"has_header"`
	ValidationCallback func(FileValidationResult)
	Destination        string `json:"destination"`
}

type ConfigIdentifier struct {
	DataStreamID    string   `json:"data_stream_id"`
	DataStreamRoute string   `json:"data_stream_route"`
	Header          []string `json:"header"`
}

type MetadataValidationResult struct {
	ReceivedFile    string     `json:"received_filename"`
	Error           *FileError `json:"error"`
	Status          string     `json:"status"`
	Jurisdiction    string     `json:"jurisdiction"`
	DataStreamID    string     `json:"data_stream_id"`
	DataStreamRoute string     `json:"data_stream_route"`
	SenderID        string     `json:"sender_id"`
	DataProducerID  string     `json:"data_producer_id"`
	Version         string     `json:"version"`
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
