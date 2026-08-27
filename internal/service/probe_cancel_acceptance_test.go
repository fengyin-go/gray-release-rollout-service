package service

import (
	"context"
	"testing"

	"grayrelease/internal/store"
)

func TestR007ProbeCancel(t *testing.T) {
	state := store.NewProbeCancelState()
	flow := NewProbeCancelFlow(state)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := flow.Dispatch(ctx, "job-a")
	if err == nil || state.Count() != 0 {
		t.Fatalf("健康探测任务 must stop before commit after cancellation: err=%v count=%d", err, state.Count())
	}
}
