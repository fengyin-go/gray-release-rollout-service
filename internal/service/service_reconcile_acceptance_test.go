package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR024ServiceReconcile(t *testing.T) {
	state := store.NewServiceReconcileLeaseState()
	flow := NewServiceReconcileFlow(state)
	firstErr := flow.Process("node-a", true)
	secondErr := flow.Process("node-a", false)
	if firstErr == nil || secondErr != nil || state.Active() != 0 {
		t.Fatalf("服务校准租约 must be released after failure so the next run can start: first=%v second=%v active=%d", firstErr, secondErr, state.Active())
	}
}
