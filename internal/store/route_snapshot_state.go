package store

import "sync"

type RouteSnapshotState struct {
	mu     sync.RWMutex
	values map[string][]string
}

func NewRouteSnapshotState() *RouteSnapshotState {
	return &RouteSnapshotState{values: make(map[string][]string)}
}

func (s *RouteSnapshotState) Replace(key string, values []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = values
}

func (s *RouteSnapshotState) Snapshot(key string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.values[key]
}
