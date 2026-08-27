package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR011BalancerResolver(t *testing.T) {
	flow := NewBalancerResolverFlow(store.NewBalancerResolverResolver(false))
	panicked := false
	value := ""
	func() {
		defer func() { panicked = recover() != nil }()
		value = flow.Resolve("node-a")
	}()
	if panicked || value != "unavailable" {
		t.Fatalf("负载算法解析组件 must use the disabled fallback without panic: panic=%v value=%q", panicked, value)
	}
}
