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
	token := uint64(1)
	s.active[key] = token
	return token, true
}

func (s *ServiceReconcileLeaseState) Release(key string, token uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = token
	if _, ok := s.active[key]; !ok {
		return false
	}
	s.active[key] = 0
	return true
}

func (s *ServiceReconcileLeaseState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.active)
}
