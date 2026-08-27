package store

import "sync"

type SnapshotImportLeaseState struct {
	mu     sync.Mutex
	active map[string]uint64
	next   uint64
}

func NewSnapshotImportLeaseState() *SnapshotImportLeaseState {
	return &SnapshotImportLeaseState{active: make(map[string]uint64)}
}

// Acquire 尝试为 key 获取租约，返回的 token 在释放时需原样回传，
// 用以校验释放者是否仍持有该租约（防止过期持有者误释放）。
func (s *SnapshotImportLeaseState) Acquire(key string) (uint64, bool) {
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

// Release 释放 key 对应的租约；仅当 token 与当前持有者匹配时才真正移除，
// 使该 key 可被下一次 Acquire 占用。
func (s *SnapshotImportLeaseState) Release(key string, token uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	cur, ok := s.active[key]
	if !ok {
		return false
	}
	if cur != token {
		return false
	}
	delete(s.active, key)
	return true
}

func (s *SnapshotImportLeaseState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.active)
}
