package constants

import (
	"time"
)

type EncodingType string

const (
	MAX_READ_THRESHOLD                               = 1024
	MAX_EXECUTION_TIME                               = 500 * time.Millisecond
	BOM_LENGTH                                       = 3
	UTF8                                EncodingType = "UTF-8"
	UTF8_BOM                            EncodingType = "UTF-8 WITH BOM"
	USASCII                             EncodingType = "US-ASCII"
	ISO8859_1                           EncodingType = "ISO-8859-1"
	WINDOWS1252                         EncodingType = "windows-1252"
	UNDEF                               EncodingType = "UNDEFINED"
	CSV                                              = "CSV"
	TSV                                              = "TSV"
	UNSUPPORTED                         rune         = 0
	TAB                                              = '\t'
	COMMA                                            = ','
	STATUS_SUCCESS                                   = "success"
	STATUS_FAILED                                    = "failed"
	NO_DELIMITERS_DETECTED                           = "No delimiters were detected in the file. Please ensure the file has the correct format."
	FILE_READ_ERROR                                  = "Error reading the file. Check if the file is accessible and not corrupted."
	FILE_OPEN_ERROR                                  = "Error opening the file. Verify the file path and permissions."
	FILE_WRITE_ERROR                                 = "Error writing to the file. Ensure you have the necessary write permissions and the file is not locked."
	FILE_CLOSE_ERROR                                 = "Error closing the file. This may indicate an issue with file system resources."
	FILE_CREATE_ERROR                                = "Error creating the temp file."
	CSV_READER_ERROR                                 = "Error creating CSV reader. Please check the CSV format and ensure it is correctly formatted."
	DIRECTORY_CREATE_ERROR                           = "Failed to create temporary directory."
	DIRECTORY_REMOVE_ERROR                           = "Failed to remove the test directory. Verify that the directory exists and you have the necessary permissions."
	ERROR_CONVERTING_STRUCT_TO_JSON                  = "Error converting the struct to JSON. Check the struct definition for compatibility with JSON marshalling."
	ERROR_UNMARSHALING_JSON                          = "Invalid JSON. Please check JSON format."
	JSON_EXTENSION                                   = ".json"
	CSV_FILENAME_WITH_BOM                            = "HasBOM.csv"
	CSV_FILENAME_WITHOUT_BOM                         = "NoBOM.csv"
	UNSUPPORTED_DELIMITER_ERROR                      = "Unsupported delimiter found in the file. Please use a supported delimiter and try again."
	UNSUPPORTED_ENCODING_ERROR                       = "Unsupported encoding detected. Ensure the file is encoded in a supported format."
	BOM_NOT_DETECTED_ERROR                           = "Byte Order Mark was not detected."
	INTERFACE_TO_SLICE_CONVERSION_ERROR              = "Error occurred while converting interface{} to slice."
	FILE_MISSING_ERROR                               = "received_filename is a required metadata field."
	INVALID_CONFIG_FILE                              = "Missing config_identifiers in config file."
)

type Severity string

const (
	Warning Severity = "warning"
	Failure Severity = "failure"
)

const (
	EMPTY_INPUT          = "Empty Input Test"
	VALID_US_ASCII       = "US-ASCII Valid Input"
	INVALID_US_ASCII     = "US-ASCII Invalid Input"
	VALID_WINDOWS_1252   = "Windows 1252 Valid Input"
	INVALID_WINDOWS_1252 = "Windows 1252 Invalid Input"
	VALID_ISO_8859_1     = "ISO 8859-1 Valid Input"
	INVALID_ISO_8859_1   = "ISO 8859-1 Invalid Input"
	VALID_UTF_8          = "UTF-8 Valid Input"
	INVALID_UTF_8        = "UTF-8 Invalid Input"
)

const (
	DATA_STREAM_ID     = "data_stream_id"
	SENDER_ID          = "sender_id"
	RECEIVED_FILENAME  = "received_filename"
	DATA_PRODUCER_ID   = "data_producer_id"
	DATA_STREAM_ROUTE  = "data_stream_route"
	JURISDICTION       = "jurisdiction"
	VERSION            = "version"
	HEADER             = "header"
	CONFIG             = "config"
	CONFIG_IDENTIFIERS = "config_identifiers"

	CSV_DATA_STREAM_ID    = "dex-csv"
	CSV_DATA_PRODUCER_ID  = "dex-csv"
	CSV_DATA_STREAM_ROUTE = "dex-csv"
	CSV_SENDER_ID         = "nrss-csv"
	CSV_JURISDICTION      = "NJ"
)

const (
	DEAD_LETTER_QUEUE       = "storage/dead-letter-queue"
	FILE_REPORTS            = "storage/filereports"
	ROW_REPORTS             = "storage/rowreports"
	TRANSFORMED_ROW_REPORTS = "storage/transformedrows"
	CONFIG_FILE             = "internal/config/config.json"
)

const (
	MSBMask                 rune = 0x80 // most significant bit, binary:10000000 decimal: 128
	Windows1252RunThreshold rune = 0x9F // as decimal 159
	SingleByteSequenceEnd   rune = 0xFF // as decimal 255

)

const (
	ERR_MISMATCHED_FIELD_COUNTS              = "Mismatched field count. Please ensure each row contains the correct number of fields."
	ERR_UNESCAPED_QUOTES                     = "Unescaped quotes found in field. Please ensure quotes within a quoted field are escaped by preceding them with another double quote."
	ERR_BARE_QUOTE                           = "Bare quote character found in unquoted field. Please ensure the field is correctly quoted."
	ERR_HEADER_LAST_FIELD_TRAILING_DELIMITER = "Trailing delimiter found in the last field of the header. Please ensure the last field is not followed by a delimiter (comma or tab)."
	ERR_HEADER_VALIDATION                    = "Expected header and actual header do not match."
)

var DelimiterCharacters = map[rune]rune{
	0:  UNSUPPORTED,
	9:  TAB,
	44: COMMA,
}

var (
	UTF8Bom   = []byte{0xEF, 0xBB, 0xBF}
	UTF8NoBom = []byte("Name, Role, Age")
)
var Windows1252Map = map[rune]byte{
	0x20AC: 0x80, // € (Euro sign)
	0x201A: 0x82, // ‚ (Single low-9 quotation mark)
	0x0192: 0x83, // ƒ (Latin small letter f with hook)
	0x201E: 0x84, // „ (Double low-9 quotation mark)
	0x2026: 0x85, // … (Horizontal ellipsis)
	0x2020: 0x86, // † (Dagger)
	0x2021: 0x87, // ‡ (Double dagger)
	0x02C6: 0x88, // ˆ (Modifier letter circumflex accent)
	0x2030: 0x89, // ‰ (Per mille sign)
	0x0160: 0x8A, // Š (Latin capital letter S with caron)
	0x2039: 0x8B, // ‹ (Single left-pointing angle quotation mark)
	0x0152: 0x8C, // Œ (Latin capital ligature OE)
	0x017D: 0x8E, // Ž (Latin capital letter Z with caron)
	0x2018: 0x91, // ‘ (Left single quotation mark)
	0x2019: 0x92, // ’ (Right single quotation mark)
	0x201C: 0x93, // “ (Left double quotation mark)
	0x201D: 0x94, // ” (Right double quotation mark)
	0x2022: 0x95, // • (Bullet)
	0x2013: 0x96, // – (En dash)
	0x2014: 0x97, // — (Em dash)
	0x02DC: 0x98, // ˜ (Small tilde)
	0x2122: 0x99, // ™ (Trade mark sign)
	0x0161: 0x9A, // š (Latin small letter s with caron)
	0x203A: 0x9B, // › (Single right-pointing angle quotation mark)
	0x0153: 0x9C, // œ (Latin small ligature oe)
	0x017E: 0x9E, // ž (Latin small letter z with caron)
	0x0178: 0x9F, // Ÿ (Latin capital letter Y with diaeresis)
}
var ExtendedASCIIMap = map[rune]byte{
	0x00A0: 0xA0, //   (No-break space)
	0x00A1: 0xA1, // ¡ (Inverted exclamation mark)
	0x00A2: 0xA2, // ¢ (Cent sign)
	0x00A3: 0xA3, // £ (Pound sign)
	0x00A4: 0xA4, // ¤ (Currency sign)
	0x00A5: 0xA5, // ¥ (Yen sign)
	0x00A6: 0xA6, // ¦ (Broken bar)
	0x00A7: 0xA7, // § (Section sign)
	0x00A8: 0xA8, // ¨ (Diaeresis)
	0x00A9: 0xA9, // © (Copyright sign)
	0x00AA: 0xAA, // ª (Feminine ordinal indicator)
	0x00AB: 0xAB, // « (Left-pointing double angle quotation mark)
	0x00AC: 0xAC, // ¬ (Not sign)
	0x00AD: 0xAD, // ­ (Soft hyphen)
	0x00AE: 0xAE, // ® (Registered sign)
	0x00AF: 0xAF, // ¯ (Macron)
	0x00B0: 0xB0, // ° (Degree sign)
	0x00B1: 0xB1, // ± (Plus-minus sign)
	0x00B2: 0xB2, // ² (Superscript two)
	0x00B3: 0xB3, // ³ (Superscript three)
	0x00B4: 0xB4, // ´ (Acute accent)
	0x00B5: 0xB5, // µ (Micro sign)
	0x00B6: 0xB6, // ¶ (Pilcrow sign)
	0x00B7: 0xB7, // · (Middle dot)
	0x00B8: 0xB8, // ¸ (Cedilla)
	0x00B9: 0xB9, // ¹ (Superscript one)
	0x00BA: 0xBA, // º (Masculine ordinal indicator)
	0x00BB: 0xBB, // » (Right-pointing double angle quotation mark)
	0x00BC: 0xBC, // ¼ (Vulgar fraction one quarter)
	0x00BD: 0xBD, // ½ (Vulgar fraction one half)
	0x00BE: 0xBE, // ¾ (Vulgar fraction three quarters)
	0x00BF: 0xBF, // ¿ (Inverted question mark)
	0x00C0: 0xC0, // À (Latin capital letter A with grave)
	0x00C1: 0xC1, // Á (Latin capital letter A with acute)
	0x00C2: 0xC2, // Â (Latin capital letter A with circumflex)
	0x00C3: 0xC3, // Ã (Latin capital letter A with tilde)
	0x00C4: 0xC4, // Ä (Latin capital letter A with diaeresis)
	0x00C5: 0xC5, // Å (Latin capital letter A with ring above)
	0x00C6: 0xC6, // Æ (Latin capital letter AE)
	0x00C7: 0xC7, // Ç (Latin capital letter C with cedilla)
	0x00C8: 0xC8, // È (Latin capital letter E with grave)
	0x00C9: 0xC9, // É (Latin capital letter E with acute)
	0x00CA: 0xCA, // Ê (Latin capital letter E with circumflex)
	0x00CB: 0xCB, // Ë (Latin capital letter E with diaeresis)
	0x00CC: 0xCC, // Ì (Latin capital letter I with grave)
	0x00CD: 0xCD, // Í (Latin capital letter I with acute)
	0x00CE: 0xCE, // Î (Latin capital letter I with circumflex)
	0x00CF: 0xCF, // Ï (Latin capital letter I with diaeresis)
	0x00D0: 0xD0, // Ð (Latin capital letter Eth)
	0x00D1: 0xD1, // Ñ (Latin capital letter N with tilde)
	0x00D2: 0xD2, // Ò (Latin capital letter O with grave)
	0x00D3: 0xD3, // Ó (Latin capital letter O with acute)
	0x00D4: 0xD4, // Ô (Latin capital letter O with circumflex)
	0x00D5: 0xD5, // Õ (Latin capital letter O with tilde)
	0x00D6: 0xD6, // Ö (Latin capital letter O with diaeresis)
	0x00D7: 0xD7, // × (Multiplication sign)
	0x00D8: 0xD8, // Ø (Latin capital letter O with stroke)
	0x00D9: 0xD9, // Ù (Latin capital letter U with grave)
	0x00DA: 0xDA, // Ú (Latin capital letter U with acute)
	0x00DB: 0xDB, // Û (Latin capital letter U with circumflex)
	0x00DC: 0xDC, // Ü (Latin capital letter U with diaeresis)
	0x00DD: 0xDD, // Ý (Latin capital letter Y with acute)
	0x00DE: 0xDE, // Þ (Latin capital letter Thorn)
	0x00DF: 0xDF, // ß (Latin small letter sharp S)
	0x00E0: 0xE0, // à (Latin small letter a with grave)
	0x00E1: 0xE1, // á (Latin small letter a with acute)
	0x00E2: 0xE2, // â (Latin small letter a with circumflex)
	0x00E3: 0xE3, // ã (Latin small letter a with tilde)
	0x00E4: 0xE4, // ä (Latin small letter a with diaeresis)
	0x00E5: 0xE5, // å (Latin small letter a with ring above)
	0x00E6: 0xE6, // æ (Latin small letter ae)
	0x00E7: 0xE7, // ç (Latin small letter c with cedilla)
	0x00E8: 0xE8, // è (Latin small letter e with grave)
	0x00E9: 0xE9, // é (Latin small letter e with acute)
	0x00EA: 0xEA, // ê (Latin small letter e with circumflex)
	0x00EB: 0xEB, // ë (Latin small letter e with diaeresis)
	0x00EC: 0xEC, // ì (Latin small letter i with grave)
	0x00ED: 0xED, // í (Latin small letter i with acute)
	0x00EE: 0xEE, // î (Latin small letter i with circumflex)
	0x00EF: 0xEF, // ï (Latin small letter i with diaeresis)
	0x00F0: 0xF0, // ð (Latin small letter eth)
	0x00F1: 0xF1, // ñ (Latin small letter n with tilde)
	0x00F2: 0xF2, // ò (Latin small letter o with grave)
	0x00F3: 0xF3, // ó (Latin small letter o with acute)
	0x00F4: 0xF4, // ô (Latin small letter o with circumflex)
	0x00F5: 0xF5, // õ (Latin small letter o with tilde)
	0x00F6: 0xF6, // ö (Latin small letter o with diaeresis)
	0x00F7: 0xF7, // ÷ (Division sign)
	0x00F8: 0xF8, // ø (Latin small letter o with stroke)
	0x00F9: 0xF9, // ù (Latin small letter u with grave)
	0x00FA: 0xFA, // ú (Latin small letter u with acute)
	0x00FB: 0xFB, // û (Latin small letter u with circumflex)
	0x00FC: 0xFC, // ü (Latin small letter u with diaeresis)
	0x00FD: 0xFD, // ý (Latin small letter y with acute)
	0x00FE: 0xFE, // þ (Latin small letter thorn)
	0x00FF: 0xFF, // ÿ (Latin small letter y with diaeresis)
}
