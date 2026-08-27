package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR026HeartbeatRetry(t *testing.T) {
	permanent := store.NewHeartbeatRetryRetryState(&store.RetryFailure{Temporary: false, Message: "rejected"}, nil)
	permanentErr := NewHeartbeatRetryFlow(permanent).Execute()
	temporary := store.NewHeartbeatRetryRetryState(&store.RetryFailure{Temporary: true, Message: "busy"}, nil)
	temporaryErr := NewHeartbeatRetryFlow(temporary).Execute()
	if permanentErr == nil || permanent.Attempts() != 1 || permanent.Commits() != 0 || temporaryErr != nil || temporary.Attempts() != 2 || temporary.Commits() != 1 {
		t.Fatalf("心跳写入重试 must stop permanent failures and retry temporary failures exactly once: permanent=(%v,%d,%d) temporary=(%v,%d,%d)", permanentErr, permanent.Attempts(), permanent.Commits(), temporaryErr, temporary.Attempts(), temporary.Commits())
	}
}
