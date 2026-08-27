package store

import "sync"

type HealthReasonsState struct {
	mu     sync.RWMutex
	values map[string][]string
}

func NewHealthReasonsState() *HealthReasonsState {
	return &HealthReasonsState{values: make(map[string][]string)}
}

func (s *HealthReasonsState) Replace(key string, values []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = values
}

func (s *HealthReasonsState) Snapshot(key string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.values[key]
}
