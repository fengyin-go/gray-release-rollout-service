package store

import "sync"

type DependencyCacheState struct {
	mu     sync.RWMutex
	values map[string][]string
}

func NewDependencyCacheState() *DependencyCacheState {
	return &DependencyCacheState{values: make(map[string][]string)}
}

func (s *DependencyCacheState) Replace(key string, values []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = values
}

func (s *DependencyCacheState) Snapshot(key string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.values[key]
}
