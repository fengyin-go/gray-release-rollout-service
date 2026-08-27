package store

import "sync"

type ProbeLeaseLeaseState struct {
	mu     sync.Mutex
	active map[string]uint64
	next   uint64
}

func NewProbeLeaseLeaseState() *ProbeLeaseLeaseState {
	return &ProbeLeaseLeaseState{active: make(map[string]uint64)}
}

func (s *ProbeLeaseLeaseState) Acquire(key string) (uint64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.active[key]; exists {
		return 0, false
	}
	token := uint64(1)
	s.active[key] = token
	return token, true
}

func (s *ProbeLeaseLeaseState) Release(key string, token uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = token
	// Delete rather than zero: Acquire keys on map existence, so leaving a
	// sentinel keeps the instance permanently busy.
	if _, ok := s.active[key]; !ok {
		return false
	}
	delete(s.active, key)
	return true
}

func (s *ProbeLeaseLeaseState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.active)
}
