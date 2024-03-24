package data

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// JSONToText converts JSON data into a text format.
// It takes JSON data as input and returns a string representing the data in a text format.
// If there is an error during unmarshalling or processing, it returns an error.
/*
Example:
	{
		"person": {
			"name": "John",
			"age": 30,
			"address": {
				"city": "New York",
				"zipcode": "10001"
			}
		}
	}
Output:
	person.name: John
	person.age: 30
	person.address.city: New York
	person.address.zipcode: 10001
*/
func JSONToText(jsonData []byte) (string, error) {
	var data interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", err
	}

	text := processJSONForText(data, "")

	return text, nil
}

// processJSON recursively processes nested JSON data and converts it into text.
// It takes a JSON object and a prefix string as input and returns the converted text.
func processJSONForText(data interface{}, prefix string) string {
	var text strings.Builder

	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if _, ok := value.(map[string]interface{}); ok {
				text.WriteString(processJSONForText(value, prefix+key+"."))
			} else {
				text.WriteString(processJSONForText(value, prefix+key+""))
			}
		}
	case []interface{}:
		for i, item := range v {
			text.WriteString(processJSONForText(item, fmt.Sprintf("%s[%d]", prefix, i)))
		}
	default:
		text.WriteString(fmt.Sprintf("%s: %v\n", prefix, v))
	}

	return text.String()
}

// JSONToYAML converts JSON data into YAML format.
// It takes JSON data as input and returns a string representing the data in YAML format.
// If there is an error during unmarshalling or processing, it returns an error.
func JSONToYAML(jsonData []byte) (string, error) {
	var data interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", err
	}

	yamlData, err := toJSONYAML(data)
	if err != nil {
		return "", err
	}

	return yamlData, nil
}

// toJSONYAML converts the given data structure into YAML format.
// It takes any data structure as input and returns a string representing the YAML format.
// If there is an error during encoding, it returns an empty string and the error.
func toJSONYAML(data interface{}) (string, error) {
	// Create a string builder to store the YAML output
	var buf strings.Builder

	// Create a new YAML encoder with an indentation level of 2 spaces
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)

	// Encode the data into YAML format and write it to the string builder
	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}

	// Return the YAML string and no error
	return buf.String(), nil
}
