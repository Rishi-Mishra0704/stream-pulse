package data

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Start of Json to Text tests
func TestJSONToText(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests

	// Sample JSON data
	jsonData := []byte(`{"name": "John", "age": 30, "city": "New York"}`)

	// Expected transformed text
	expectedText := "name: John\nage: 30\ncity: New York\n"

	// Call JSONToText function
	actualText, err := JSONToText(jsonData)

	actualLines := sortLines(actualText)
	expectedLines := sortLines(expectedText)

	// Check if there's no error and the transformed text matches the expected text
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, expectedLines, actualLines, "transformed text should match expected text")
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

	actualLines := sortLines(actualText)
	expectedLines := sortLines(expectedText)

	// Check if there's no error and the transformed text matches the expected text
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, expectedLines, actualLines, "transformed text should match expected text")
}
func TestJSONToText_Array(t *testing.T) {
	t.Parallel() // Run this test in parallel with other tests

	// Sample JSON data with an array
	jsonData := []byte(`{"names": ["John", "Alice", "Michael"]}`)

	// Expected transformed text
	expectedText := "names[0]: John\nnames[1]: Alice\nnames[2]: Michael\n"

	// Call JSONToText function
	actualText, err := JSONToText(jsonData)

	actualLines := sortLines(actualText)
	expectedLines := sortLines(expectedText)

	// Check if there's no error and the transformed text matches the expected text
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, expectedLines, actualLines, "transformed text should match expected text")
}

func sortLines(text string) []string {
	lines := strings.Split(text, "\n")
	sort.Strings(lines)
	return lines
}

// End of Json to Text tests

// Start of Json to YAML tests
func TestJSONToYAML(t *testing.T) {
	t.Parallel()
	// Sample JSON data
	jsonData := []byte(`{"name": "John", "age": 30, "city": "New York"}`)

	// Expected YAML data
	expectedYAML := "age: 30\ncity: New York\nname: John\n"

	// Call JSONToYAML function
	actualYAML, err := JSONToYAML(jsonData)

	// Check if there's no error and the converted YAML matches the expected YAML
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, expectedYAML, actualYAML, "converted YAML should match expected YAML")
}

func TestJSONToYAML_InvalidJSON(t *testing.T) {
	t.Parallel()
	// Invalid JSON data
	invalidJSON := []byte(`{"name": "John", "age": 30, "city": New York}`)

	// Call JSONToYAML function with invalid JSON data
	_, err := JSONToYAML(invalidJSON)

	// Check if an error is returned
	assert.Error(t, err, "expected error for invalid JSON data")
}

func TestJSONToYAML_NestedJSON(t *testing.T) {
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

	// Expected YAML data with keys in a specific order and consistent indentation
	expectedYAML := `person:
  name: John
  age: 30
  address:
    city: New York
    zipcode: "10001"
`

	// Call JSONToYAML function with nested JSON data
	actualYAML, err := JSONToYAML(nestedJSON)
	assert.NoError(t, err, "unexpected error")

	// Normalize YAML strings
	expectedLines := normalizeYAML(expectedYAML)
	actualLines := normalizeYAML(actualYAML)

	// Compare normalized YAML strings
	assert.Equal(t, expectedLines, actualLines, "normalized YAML strings should match")
}

// normalizeYAML splits the YAML string into lines and sorts them alphabetically
func normalizeYAML(yamlStr string) []string {
	lines := strings.Split(yamlStr, "\n")
	sort.Strings(lines)
	return lines
}

// End of Json to YAML tests
