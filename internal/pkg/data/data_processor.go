package data

import "sync"

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
