package store

import (
	"context"
	"sync"
)

type SnapshotCancelState struct {
	mu        sync.Mutex
	completed []string
}

func NewSnapshotCancelState() *SnapshotCancelState { return &SnapshotCancelState{} }

func (s *SnapshotCancelState) Commit(ctx context.Context, key string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = ctx
	s.completed = append(s.completed, key)
	return nil
}

func (s *SnapshotCancelState) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.completed)
}
