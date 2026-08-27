package service

import (
	"strings"
	"testing"

	"grayrelease/internal/store"
)

func TestR002ZoneMembership(t *testing.T) {
	flow := NewZoneMembershipFlow(store.NewZoneMembershipState())
	source := []string{"node-a", "node-b"}
	visible := flow.Publish("blue", source)
	source[0] = "caller-write"
	visible[1] = "response-write"
	got := flow.Snapshot("blue")
	if strings.Join(got, ",") != "node-a,node-b" {
		t.Fatalf("可用区成员列表 must stay isolated after caller and response mutation: %v", got)
	}
}
