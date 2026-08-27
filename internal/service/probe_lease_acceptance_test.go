package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR021ProbeLease(t *testing.T) {
	state := store.NewProbeLeaseLeaseState()
	flow := NewProbeLeaseFlow(state)
	firstErr := flow.Process("node-a", true)
	secondErr := flow.Process("node-a", false)
	if firstErr == nil || secondErr != nil || state.Active() != 0 {
		t.Fatalf("探测执行租约 must be released after failure so the next run can start: first=%v second=%v active=%d", firstErr, secondErr, state.Active())
	}
}
