package store

import (
	"context"
	"sync"
)

type MaintenanceCancelState struct {
	mu        sync.Mutex
	completed []string
}

func NewMaintenanceCancelState() *MaintenanceCancelState { return &MaintenanceCancelState{} }

func (s *MaintenanceCancelState) Commit(ctx context.Context, key string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = ctx
	s.completed = append(s.completed, key)
	return nil
}

func (s *MaintenanceCancelState) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.completed)
}
