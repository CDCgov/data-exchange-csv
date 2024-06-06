package file

import (
	"encoding/json"
	"io"
	"os"

	"github.com/CDCgov/data-exchange-csv/cmd/internal/constants"
)

func (vr *ValidationResult) processMetadataFields(configFile string) {

	file, err := os.Open(configFile)
	if err != nil {
		vr.Error = &Error{Message: constants.FILE_OPEN_ERROR, Code: 13}
		copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
	}
	defer file.Close()

	fields, err := io.ReadAll(file)
	if err != nil {
		vr.Error = &Error{Message: err.Error(), Code: 13}
		copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
		return
	}

	var metadataMap map[string]string
	err = json.Unmarshal(fields, &metadataMap)
	if err != nil {
		vr.Error = &Error{Message: err.Error(), Code: 13}
		copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
		return
	}

	vr.Jurisdiction = metadataMap[constants.JURISDICTION]
	if filename, ok := metadataMap[constants.RECEIVED_FILENAME]; ok {
		vr.ReceivedFile = filename
	} else {
		vr.Error = &Error{Message: constants.RECEIVED_FILENAME, Code: 13}
		copyToDestination(vr, constants.DEAD_LETTER_QUEUE)
		return
	}

	vr.DataStreamID = metadataMap[constants.DATA_STREAM_ID]
	vr.DataProducerID = metadataMap[constants.DATA_PRODUCER_ID]
	vr.Version = metadataMap[constants.VERSION]
}
