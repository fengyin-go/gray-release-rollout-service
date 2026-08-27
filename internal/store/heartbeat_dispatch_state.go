package store

import (
	"context"
	"sync"
)

type HeartbeatDispatchState struct {
	mu        sync.Mutex
	completed []string
}

func NewHeartbeatDispatchState() *HeartbeatDispatchState { return &HeartbeatDispatchState{} }

func (s *HeartbeatDispatchState) Commit(ctx context.Context, key string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = ctx
	s.completed = append(s.completed, key)
	return nil
}

func (s *HeartbeatDispatchState) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.completed)
}
