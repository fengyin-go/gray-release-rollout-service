package service

import (
	"strings"
	"testing"

	"grayrelease/internal/store"
)

func TestR001RouteSnapshot(t *testing.T) {
	flow := NewRouteSnapshotFlow(store.NewRouteSnapshotState())
	source := []string{"node-a", "node-b"}
	visible := flow.Publish("blue", source)
	source[0] = "caller-write"
	visible[1] = "response-write"
	got := flow.Snapshot("blue")
	if strings.Join(got, ",") != "node-a,node-b" {
		t.Fatalf("路由端点快照 must stay isolated after caller and response mutation: %v", got)
	}
}
