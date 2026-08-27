package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR012HealthNotifier(t *testing.T) {
	flow := NewHealthNotifierFlow(store.NewHealthNotifierResolver(false))
	panicked := false
	value := ""
	func() {
		defer func() { panicked = recover() != nil }()
		value = flow.Resolve("node-a")
	}()
	if panicked || value != "unavailable" {
		t.Fatalf("健康变更通知出口 must use the disabled fallback without panic: panic=%v value=%q", panicked, value)
	}
}
