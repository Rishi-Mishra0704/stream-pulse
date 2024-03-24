package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONToText(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests

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
	t.Parallel() // Run this test in parallel with other tests

	// Invalid JSON data
	invalidJSON := []byte(`{"name": "John", "age": 30, "city": New York}`)

	// Call JSONToText function with invalid JSON data
	_, err := JSONToText(invalidJSON)

	// Check if an error is returned
	assert.Error(t, err, "expected error for invalid JSON data")
}

func TestJSONToText_NestedJSON(t *testing.T) {
	// This test relies on shared resources and cannot run concurrently
	// So, it should not have t.Parallel()

	// Nested JSON data
	nestedJSON := []byte(`{
		"person": {
			"name": "John",
			"age": 30,
			"address": {
				"city": "New York",
				"zipcode": "10001"
			}
		}
	}`)

	// Expected transformed text
	expectedText := "person.name: John\nperson.age: 30\nperson.address.city: New York\nperson.address.zipcode: 10001\n"

	// Call JSONToText function with nested JSON data
	actualText, err := JSONToText(nestedJSON)

	// Check if there's no error and the transformed text matches the expected text
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, expectedText, actualText, "transformed text should match expected text")
}

// Reset the test state after each test execution
