// Package data provides functionality for processing real-time data.
package data

import (
	"strings"
	"sync"
)

// DataProcessor defines an interface for processing data and retrieving processing results.
type DataProcessor interface {
	// Process processes the incoming data and updates the processing results.
	Process(data []byte) error
	// Results returns the processing results as a map of words to their respective counts.
	Results() map[string]int
}

// RealtimeDataProcessor is a real-time data processor that implements the DataProcessor interface.
type RealtimeDataProcessor struct {
	results map[string]int // results stores the processing results.
	mutex   sync.Mutex     // mutex provides synchronization for concurrent access to results.
}

// NewRealTimeDataProcessor creates a new RealtimeDataProcessor instance with an initialized results map.
func NewRealTimeDataProcessor() *RealtimeDataProcessor {
	return &RealtimeDataProcessor{
		results: make(map[string]int),
	}
}

// Process processes the incoming data by splitting it into words and updating the processing results.
func (p *RealtimeDataProcessor) Process(data string) {
	words := strings.Fields(data)

	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, word := range words {
		p.results[word]++
	}
}

// Results returns the processing results as a map of words to their respective counts.
// It locks the mutex to ensure exclusive access to the results map and creates a copy of the map to avoid data race.
func (p *RealtimeDataProcessor) Results() map[string]int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	resultCopy := make(map[string]int)
	for key, value := range p.results {
		resultCopy[key] = value
	}
	return resultCopy
}
