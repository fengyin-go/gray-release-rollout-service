package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR014TopologyFilter(t *testing.T) {
	flow := NewTopologyFilterFlow(store.NewTopologyFilterResolver(false))
	panicked := false
	value := ""
	func() {
		defer func() { panicked = recover() != nil }()
		value = flow.Resolve("node-a")
	}()
	if panicked || value != "unavailable" {
		t.Fatalf("拓扑筛选组件 must use the disabled fallback without panic: panic=%v value=%q", panicked, value)
	}
}
