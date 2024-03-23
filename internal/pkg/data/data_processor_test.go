package data

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealtimeDataProcessor_Process(t *testing.T) {
	processor := NewRealTimeDataProcessor()

	// Process some data
	processor.Process("hello world")
	processor.Process("hello")
	processor.Process("world world")

	// Check if the results are as expected
	expected := map[string]int{
		"hello": 2,
		"world": 3,
	}
	assert.Equal(t, expected, processor.Results(), "processing results should match expected values")
}

func TestRealtimeDataProcessor_Process_EmptyData(t *testing.T) {
	processor := NewRealTimeDataProcessor()

	// Process empty data
	processor.Process("")

	// Check if the results are empty
	expected := map[string]int{}
	assert.Equal(t, expected, processor.Results(), "processing results should be empty for empty data")
}

func TestRealtimeDataProcessor_Results_Concurrency(t *testing.T) {
	processor := NewRealTimeDataProcessor()

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Perform concurrent data processing
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			processor.Process("hello world")
		}()
	}

	// Wait for all processing to complete
	wg.Wait()

	// Check if the results are as expected
	expected := map[string]int{
		"hello": 1000,
		"world": 1000,
	}
	assert.Equal(t, expected, processor.Results(), "processing results should match expected values")
}
