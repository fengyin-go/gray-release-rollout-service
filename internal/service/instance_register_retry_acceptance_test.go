package service

import (
	"testing"

	"grayrelease/internal/store"
)

func TestR028InstanceRegisterRetry(t *testing.T) {
	permanent := store.NewInstanceRegisterRetryRetryState(&store.RetryFailure{Temporary: false, Message: "rejected"}, nil)
	permanentErr := NewInstanceRegisterRetryFlow(permanent).Execute()
	temporary := store.NewInstanceRegisterRetryRetryState(&store.RetryFailure{Temporary: true, Message: "busy"}, nil)
	temporaryErr := NewInstanceRegisterRetryFlow(temporary).Execute()
	if permanentErr == nil || permanent.Attempts() != 1 || permanent.Commits() != 0 || temporaryErr != nil || temporary.Attempts() != 2 || temporary.Commits() != 1 {
		t.Fatalf("实例注册重试 must stop permanent failures and retry temporary failures exactly once: permanent=(%v,%d,%d) temporary=(%v,%d,%d)", permanentErr, permanent.Attempts(), permanent.Commits(), temporaryErr, temporary.Attempts(), temporary.Commits())
	}
}
