package store

import (
	"context"
	"sync"
)

type ProbeCancelState struct {
	mu        sync.Mutex
	completed []string
}

func NewProbeCancelState() *ProbeCancelState { return &ProbeCancelState{} }

func (s *ProbeCancelState) Commit(ctx context.Context, key string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = ctx
	s.completed = append(s.completed, key)
	return nil
}

func (s *ProbeCancelState) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.completed)
}
