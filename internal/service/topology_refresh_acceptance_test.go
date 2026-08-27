package service

import (
	"context"
	"testing"

	"grayrelease/internal/store"
)

func TestR008TopologyRefresh(t *testing.T) {
	state := store.NewTopologyRefreshState()
	flow := NewTopologyRefreshFlow(state)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := flow.Dispatch(ctx, "job-a")
	if err == nil || state.Count() != 0 {
		t.Fatalf("拓扑刷新任务 must stop before commit after cancellation: err=%v count=%d", err, state.Count())
	}
}
