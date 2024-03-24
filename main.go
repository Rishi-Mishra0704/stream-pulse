package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Rishi-Mishra0704/stream-pulse/internal/pkg/data"
)

func main() {
	// Read JSON data from file
	var err error
	jsonData, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Measure runtime of JSONToText function
	startTimeText := time.Now()
	_, err = data.JSONToText(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	endTimeText := time.Now()
	fmt.Printf("JSONToText runtime: %s\n", formatDuration(endTimeText.Sub(startTimeText)))

	// Measure runtime of JSONToYAML function
	startTimeYAML := time.Now()
	_, err = data.JSONToYAML(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	endTimeYAML := time.Now()
	fmt.Printf("JSONToYAML runtime: %s\n", formatDuration(endTimeYAML.Sub(startTimeYAML)))

	// // Print the output
	// fmt.Println("JSONToText output:")
	// fmt.Println(text)
	// fmt.Println("JSONToYAML output:")
	// fmt.Println(yaml)
}

// formatDuration formats the given duration for readability
func formatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return fmt.Sprintf("%dµs", d.Microseconds())
	}
	return d.String()
}
