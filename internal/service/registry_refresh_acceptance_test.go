package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR022RegistryRefresh(t *testing.T) {
	state := store.NewRegistryRefreshLeaseState()
	flow := NewRegistryRefreshFlow(state)
	firstErr := flow.Process("node-a", true)
	secondErr := flow.Process("node-a", false)
	if firstErr == nil || secondErr != nil || state.Active() != 0 {
		t.Fatalf("注册表刷新租约 must be released after failure so the next run can start: first=%v second=%v active=%d", firstErr, secondErr, state.Active())
	}
}
