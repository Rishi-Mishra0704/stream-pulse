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

	var text string

	for key, value := range data {
		text += fmt.Sprintf("%s: %v\n", key, value)

	}
	return text, nil
}
