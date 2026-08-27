package service

import (
	"strings"
	"testing"

	"grayrelease/internal/store"
)

func TestR005SelectorLabels(t *testing.T) {
	flow := NewSelectorLabelsFlow(store.NewSelectorLabelsState())
	source := []string{"node-a", "node-b"}
	visible := flow.Publish("blue", source)
	source[0] = "caller-write"
	visible[1] = "response-write"
	got := flow.Snapshot("blue")
	if strings.Join(got, ",") != "node-a,node-b" {
		t.Fatalf("实例筛选标签 must stay isolated after caller and response mutation: %v", got)
	}
}
