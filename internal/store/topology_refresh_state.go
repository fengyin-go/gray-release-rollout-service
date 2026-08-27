package store

import (
	"context"
	"sync"
)

type TopologyRefreshState struct {
	mu        sync.Mutex
	completed []string
}

func NewTopologyRefreshState() *TopologyRefreshState { return &TopologyRefreshState{} }

func (s *TopologyRefreshState) Commit(ctx context.Context, key string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = ctx
	s.completed = append(s.completed, key)
	return nil
}

func (s *TopologyRefreshState) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.completed)
}
