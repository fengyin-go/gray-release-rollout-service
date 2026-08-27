package service

import (
	"context"

	"grayrelease/internal/store"
)

type ProbeCancelFlow struct{ state *store.ProbeCancelState }

func NewProbeCancelFlow(state *store.ProbeCancelState) *ProbeCancelFlow {
	return &ProbeCancelFlow{state: state}
}

func (f *ProbeCancelFlow) Dispatch(ctx context.Context, key string) error {
	return f.state.Commit(context.Background(), key)
}
