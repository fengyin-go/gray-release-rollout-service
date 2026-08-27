package service

import (
	"context"
	"testing"

	"grayrelease/internal/store"
)

func TestR006HeartbeatDispatch(t *testing.T) {
	state := store.NewHeartbeatDispatchState()
	flow := NewHeartbeatDispatchFlow(state)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := flow.Dispatch(ctx, "job-a")
	if err == nil || state.Count() != 0 {
		t.Fatalf("心跳派发任务 must stop before commit after cancellation: err=%v count=%d", err, state.Count())
	}
}
