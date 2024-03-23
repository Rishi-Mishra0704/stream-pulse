package data

import (
	"strings"
	"sync"
)

type DataProcessor interface {
	Process(data []byte) error
	Results() map[string]int
}

type RealtimeDataProcessor struct {
	results map[string]int
	mutex   sync.Mutex
}

func NewRealTimeDataProcessor() *RealtimeDataProcessor {
	return &RealtimeDataProcessor{
		results: make(map[string]int),
	}
}

func (p *RealtimeDataProcessor) Process(data string) {
	words := strings.Fields(data)

	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, word := range words {
		p.results[word]++
	}
}

func (p *RealtimeDataProcessor) Result() map[string]int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	resultCopy := make(map[string]int)

	for key, value := range p.results {
		resultCopy[key] = value
	}
	return resultCopy
}
