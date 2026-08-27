package store

import "sync"

type ServiceReconcileLeaseState struct {
	mu     sync.Mutex
	active map[string]uint64
	next   uint64
}

func NewServiceReconcileLeaseState() *ServiceReconcileLeaseState {
	return &ServiceReconcileLeaseState{active: make(map[string]uint64)}
}

func (s *ServiceReconcileLeaseState) Acquire(key string) (uint64, bool) {
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

func (s *ServiceReconcileLeaseState) Release(key string, token uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	cur, ok := s.active[key]
	if !ok || cur != token {
		return false
	}
	delete(s.active, key)
	return true
}

func (s *ServiceReconcileLeaseState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.active)
}
