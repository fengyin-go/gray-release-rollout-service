package store

import "sync"

type DependencyRebuildLeaseState struct {
	mu     sync.Mutex
	active map[string]uint64
	next   uint64
}

func NewDependencyRebuildLeaseState() *DependencyRebuildLeaseState {
	return &DependencyRebuildLeaseState{active: make(map[string]uint64)}
}

func (s *DependencyRebuildLeaseState) Acquire(key string) (uint64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.active[key]; exists {
		return 0, false
	}
	token := uint64(1)
	s.active[key] = token
	return token, true
}

func (s *DependencyRebuildLeaseState) Release(key string, token uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	cur, ok := s.active[key]
	if !ok {
		return false
	}
	// 只有持有当前 token 的调用者才能释放，避免误清理后续重建获取的新租约。
	if token != 0 && cur != token {
		return false
	}
	delete(s.active, key)
	return true
}

func (s *DependencyRebuildLeaseState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.active)
}
