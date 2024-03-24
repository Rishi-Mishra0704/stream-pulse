package data

import (
	"encoding/json"
	"fmt"
)

func JSONToText(jsonData []byte) (string, error) {
	var data map[string]interface{}

	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		return "", err
	}

	text := processJSON(data, "")

	return text, nil
}

func processJSON(data map[string]interface{}, prefix string) string {
	var text string

	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			text += processJSON(v, prefix+key+".")
		default:
			text += fmt.Sprintf("%s%s: %v\n", prefix, key, value)
		}
	}

	return text
}
