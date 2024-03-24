package data

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

var encoderPool = sync.Pool{
	New: func() interface{} {
		// Initialize a new YAML encoder with default settings
		return yaml.NewEncoder(new(strings.Builder))
	},
}

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

	// Get a YAML encoder from the pool
	encoder := encoderPool.Get().(*yaml.Encoder)
	defer encoderPool.Put(encoder)

	// Create a new strings.Builder for each conversion
	var buf strings.Builder
	buf.Grow(len(jsonData) * 2) // Preallocate buffer size
	// Encode the data into YAML format and write it to the buffer
	encoder.SetIndent(2)
	err = encoder.Encode(data)
	if err != nil {
		return "", err
	}

	// Return the YAML string
	return buf.String(), nil
}
