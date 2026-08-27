package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR015HeartbeatFormatter(t *testing.T) {
	flow := NewHeartbeatFormatterFlow(store.NewHeartbeatFormatterResolver(false))
	panicked := false
	value := ""
	func() {
		defer func() { panicked = recover() != nil }()
		value = flow.Resolve("node-a")
	}()
	if panicked || value != "unavailable" {
		t.Fatalf("心跳文本格式组件 must use the disabled fallback without panic: panic=%v value=%q", panicked, value)
	}
}
