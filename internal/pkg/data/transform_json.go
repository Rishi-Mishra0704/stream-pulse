package data

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// JSONToText converts JSON data to text format.
//
// It takes JSON data as input and returns the data formatted as text, where each key-value pair is represented as a separate line.
//
// Example:
//
//	jsonData := []byte(`{"name": "John", "age": 30, "city": "New York"}`)
//	text, err := JSONToText(jsonData)
//	// Output:
//	// name: John
//	// age: 30
//	// city: New York
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

// JSONToYAML converts JSON data to YAML format.
//
// It takes JSON data as input and returns the data formatted as YAML.
//
// Example:
//
//	jsonData := []byte(`{
//	  "name": "John",
//	  "age": 30,
//	  "city": "New York",
//	  "contact": {
//	    "email": "john@example.com",
//	    "phone": "123-456-7890"
//	  }
//	}`)
//	yamlData, err := JSONToYAML(jsonData)
//	// Output:
//	// age: 30
//	// city: New York
//	// contact:
//	//   email: john@example.com
//	//   phone: "123-456-7890"
//	// name: John
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
