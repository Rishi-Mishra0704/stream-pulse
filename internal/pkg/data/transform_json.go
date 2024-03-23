package data

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
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

func JSONToYAML(jsonData []byte) (string, error) {
	var data interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", err
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}
