package constants

import "time"

type EncodingType string

const (
	MAX_READ_THRESHOLD                           = 1024
	MAX_EXECUTION_TIME                           = 500 * time.Millisecond
	UTF8                            EncodingType = "UTF-8"
	UTF8_BOM                        EncodingType = "UTF-8 WITH BOM"
	USASCII                         EncodingType = "US-ASCII"
	ISO8859_1                       EncodingType = "ISO-8859-1"
	WINDOWS1252                     EncodingType = "windows-1252"
	UNDEF                           EncodingType = "UNDEFINED"
	CSV                                          = "CSV"
	TSV                                          = "TSV"
	UNSUPPORTED                                  = "UNSUPPORTED"
	TAB                                          = '\t'
	COMMA                                        = ','
	NO_DELIMITERS_DETECTED                       = "No delimiters were detected in the file. Please ensure the file has the correct format."
	FILE_READ_ERROR                              = "Error reading the file. Check if the file is accessible and not corrupted."
	FILE_OPEN_ERROR                              = "Error opening the file. Verify the file path and permissions."
	FILE_WRITE_ERROR                             = "Error writing to the file. Ensure you have the necessary write permissions and the file is not locked."
	FILE_CLOSE_ERROR                             = "Error closing the file. This may indicate an issue with file system resources."
	FILE_CREATE_ERROR                            = "Error creating the temp file."
	CSV_READER_ERROR                             = "Error creating CSV reader. Please check the CSV format and ensure it is correctly formatted."
	DIRECTORY_CREATE_ERROR                       = "Failed to create temporary directory."
	DIRECTORY_REMOVE_ERROR                       = "Failed to remove the test directory. Verify that the directory exists and you have the necessary permissions."
	ERROR_CONVERTING_STRUCT_TO_JSON              = "Error converting the struct to JSON. Check the struct definition for compatibility with JSON marshalling."
	JSON_EXTENSION                               = ".json"
	CSV_FILENAME_WITH_BOM                        = "HasBOM.csv"
	CSV_FILENAME_WITHOUT_BOM                     = "NoBOM.csv"
	UNSUPPORTED_DELIMITER_ERROR                  = "Unsupported delimiter found in the file. Please use a supported delimiter and try again."
	UNSUPPORTED_ENCODING_ERROR                   = "Unsupported encoding detected. Ensure the file is encoded in a supported format."
	BOM_NOT_DETECTED_ERROR                       = "Byte Order Mark was not detected."
	FILE_MISSING_ERROR                           = "received_filename is a required metadata field."
)

const (
	DATA_STREAM_ID    = "dex-csv-test"
	SENDER_ID         = "filevalidation-csv-test"
	RECEIVED_FILENAME = "received_filename"
	DATA_PRODUCER_ID  = "dex-csv"
	DATA_STREAM_ROUTE = "dex-route"
	JURISDICTION      = "NJ"
	VERSION           = "1.0.0"
)

const (
	DEAD_LETTER_QUEUE = "results/DLQ"
	FILE_REPORTS      = "results/filereports"
)

const (
	MSBMask               byte = 0x80 // most significant bit, binary:10000000 decimal: 128
	InvalidStartISO88591  byte = 0x80 // as decimal 128
	InvalidEndISO88591    byte = 0x9F // as decimal 159
	ValidStartWindows1252 byte = 0x80 // as decimal 128
	ValidEndWindows1252   byte = 0x9F // as decimal 159

)

var DelimiterCharacters = map[rune]string{
	0:  UNSUPPORTED,
	9:  TSV,
	44: CSV,
}

var (
	UTF8Bom   = []byte{0xEF, 0xBB, 0xBF}
	UTF8NoBom = []byte("Name, Role, Age")
)

type RowValidationResult struct{}

type JSONTransformerResult struct{}
