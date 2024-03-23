package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONToText(t *testing.T) {
	// Sample JSON data
	jsonData := []byte(`{"name": "John", "age": 30, "city": "New York"}`)

	// Expected transformed text
	expectedText := "name: John\nage: 30\ncity: New York\n"

	// Call JSONToText function
	actualText, err := JSONToText(jsonData)

	// Check if there's no error and the transformed text matches the expected text
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, expectedText, actualText, "transformed text should match expected text")
}

func TestJSONToText_InvalidJSON(t *testing.T) {
	// Invalid JSON data
	invalidJSON := []byte(`{"name": "John", "age": 30, "city": New York}`)

	// Call JSONToText function with invalid JSON data
	_, err := JSONToText(invalidJSON)

	// Check if an error is returned
	assert.Error(t, err, "expected error for invalid JSON data")
}
