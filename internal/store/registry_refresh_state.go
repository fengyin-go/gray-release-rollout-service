package store

import "sync"

type RegistryRefreshLeaseState struct {
	mu     sync.Mutex
	active map[string]uint64
	next   uint64
}

func NewRegistryRefreshLeaseState() *RegistryRefreshLeaseState {
	return &RegistryRefreshLeaseState{active: make(map[string]uint64)}
}

func (s *RegistryRefreshLeaseState) Acquire(key string) (uint64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.active[key]; exists {
		return 0, false
	}
	token := uint64(1)
	s.active[key] = token
	return token, true
}

func (s *RegistryRefreshLeaseState) Release(key string, token uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = token
	if _, ok := s.active[key]; !ok {
		return false
	}
	s.active[key] = 0
	return true
}

func (s *RegistryRefreshLeaseState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.active)
}
