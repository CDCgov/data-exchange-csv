package transform

func transformRowToJson(row string, rowCount int) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	data["col1"] = "value1"

	//convert to json
	//jsonData, err := json.Marshal(data)

	//note to self-> check type []byte
	return data, nil
}
