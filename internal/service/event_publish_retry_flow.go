package service

import (
	"errors"

	"grayrelease/internal/store"
)

type EventPublishRetryFlow struct {
	state *store.EventPublishRetryRetryState
}

func NewEventPublishRetryFlow(state *store.EventPublishRetryRetryState) *EventPublishRetryFlow {
	return &EventPublishRetryFlow{state: state}
}

func (f *EventPublishRetryFlow) Execute() error {
	var last error
	for attempt := 0; attempt < 2; attempt++ {
		last = f.state.Next()
		if last == nil {
			return nil
		}
		var failure *store.RetryFailure
		if errors.As(last, &failure) && !failure.Temporary {
			return last
		}
	}
	return last
}
