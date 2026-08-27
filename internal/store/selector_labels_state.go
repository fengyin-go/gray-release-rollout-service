package store

import "sync"

type SelectorLabelsState struct {
	mu     sync.RWMutex
	values map[string][]string
}

func NewSelectorLabelsState() *SelectorLabelsState {
	return &SelectorLabelsState{values: make(map[string][]string)}
}

func (s *SelectorLabelsState) Replace(key string, values []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = values
}

func (s *SelectorLabelsState) Snapshot(key string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.values[key]
}
