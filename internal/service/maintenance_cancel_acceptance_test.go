package service

import (
	"context"
	"testing"

	"grayrelease/internal/store"
)

func TestR010MaintenanceCancel(t *testing.T) {
	state := store.NewMaintenanceCancelState()
	flow := NewMaintenanceCancelFlow(state)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := flow.Dispatch(ctx, "job-a")
	if err == nil || state.Count() != 0 {
		t.Fatalf("维护扫描任务 must stop before commit after cancellation: err=%v count=%d", err, state.Count())
	}
}
