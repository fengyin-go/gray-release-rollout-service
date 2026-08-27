package service

import (
	"testing"
	"time"

	"grayrelease/internal/store"
)

func TestR020HealthAggregation(t *testing.T) {
	flow := NewHealthAggregationFlow(store.NewHealthAggregationStream())
	items := []string{"node-a", "node-b", "node-c"}
	type outcome struct {
		values []string
		err    error
	}
	done := make(chan outcome, 1)
	go func() {
		values, err := flow.Collect(items, 1)
		done <- outcome{values: values, err: err}
	}()
	got := outcome{}
	timedOut := false
	select {
	case got = <-done:
	case <-time.After(500 * time.Millisecond):
		timedOut = true
	}
	if timedOut || got.err == nil || len(got.values) != 1 {
		t.Fatalf("健康状态归并器 must close promptly and preserve the partial result on an error: timeout=%v err=%v values=%v", timedOut, got.err, got.values)
	}
}
