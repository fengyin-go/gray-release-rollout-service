package service

import (
	"strings"
	"testing"

	"grayrelease/internal/store"
)

func TestR004DependencyCache(t *testing.T) {
	flow := NewDependencyCacheFlow(store.NewDependencyCacheState())
	source := []string{"node-a", "node-b"}
	visible := flow.Publish("blue", source)
	source[0] = "caller-write"
	visible[1] = "response-write"
	got := flow.Snapshot("blue")
	if strings.Join(got, ",") != "node-a,node-b" {
		t.Fatalf("依赖路径缓存 must stay isolated after caller and response mutation: %v", got)
	}
}
