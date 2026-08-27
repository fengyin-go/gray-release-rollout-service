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

// Acquire reserves the occupation marker for key, returning a lease token
// that the caller must present back to Release. A zero token means busy.
func (s *RegistryRefreshLeaseState) Acquire(key string) (uint64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.active[key]; exists {
		return 0, false
	}
	s.next++
	token := s.next
	s.active[key] = token
	return token, true
}

// Release frees the occupation marker, but only when token matches the one
// handed out by Acquire — a stale caller must not clobber a newer lease.
// The slot is deleted so a same-key refresh can re-acquire immediately.
func (s *RegistryRefreshLeaseState) Release(key string, token uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	cur, ok := s.active[key]
	if !ok || cur != token {
		return false
	}
	delete(s.active, key)
	return true
}

func (s *RegistryRefreshLeaseState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.active)
}
