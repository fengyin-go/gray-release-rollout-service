package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR013EventSink(t *testing.T) {
	flow := NewEventSinkFlow(store.NewEventSinkResolver(false))
	panicked := false
	value := ""
	func() {
		defer func() { panicked = recover() != nil }()
		value = flow.Resolve("node-a")
	}()
	if panicked || value != "unavailable" {
		t.Fatalf("注册事件接收端 must use the disabled fallback without panic: panic=%v value=%q", panicked, value)
	}
}
