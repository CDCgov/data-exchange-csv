package constants

type EncodingType string
type UUID [16]byte

const (
	MAX_READ                            = 1024
	UTF8BOMHex                          = "\xEF\xBB\xBF"
	UTF8BOMUnicode                      = "\uFEFF"
	UTF8                   EncodingType = "UTF-8"
	UTF8_BOM               EncodingType = "UTF-8 WITH BOM"
	USASCII                EncodingType = "US-ASCII"
	ISO89591               EncodingType = "ISO-8859-1"
	WINDOWS1252            EncodingType = "windows-1252"
	UNDEF                  EncodingType = "UNDEFINED"
	CSV                                 = "CSV"
	TSV                                 = "TSV"
	UNSUPPORTED                         = "unsupported"
	TAB                                 = '\t'
	COMMA                               = ','
	COLON                               = ':'
	NO_DELIMITERS_DETECTED              = "There were no delimiters detected in the file"
	FILE_READ_ERROR                     = "Error reading first 1024 bytes"
	FILE_OPEN_ERROR                     = "Error opening the file"
	FILE_CLOSE_ERROR                    = "Error closing the file"
	CSV_READER_ERROR                    = "Error creating CSV reader"
)

var DelimiterCharacters = map[rune]string{
	0:  UNSUPPORTED,
	9:  TSV,
	44: CSV,
}

type RowValidationResult struct{}

type JSONTransformerResult struct{}
