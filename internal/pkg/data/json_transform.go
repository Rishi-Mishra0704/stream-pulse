package data

import (
	"encoding/json"
	"fmt"
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
	var data map[string]interface{}

	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		return "", err
	}

	text := processJSON(data, "")

	return text, nil
}

// processJSON recursively processes nested JSON data and converts it into text.
// It takes a JSON object and a prefix string as input and returns the converted text.
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
