package store

import "sync"

type ZoneMembershipState struct {
	mu     sync.RWMutex
	values map[string][]string
}

func NewZoneMembershipState() *ZoneMembershipState {
	return &ZoneMembershipState{values: make(map[string][]string)}
}

func (s *ZoneMembershipState) Replace(key string, values []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = values
}

func (s *ZoneMembershipState) Snapshot(key string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.values[key]
}
