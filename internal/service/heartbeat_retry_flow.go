package service

import "grayrelease/internal/store"

type HeartbeatRetryFlow struct {
	state *store.HeartbeatRetryRetryState
}

func NewHeartbeatRetryFlow(state *store.HeartbeatRetryRetryState) *HeartbeatRetryFlow {
	return &HeartbeatRetryFlow{state: state}
}

func (f *HeartbeatRetryFlow) Execute() error {
	var last error
	for attempt := 0; attempt < 2; attempt++ {
		last = f.state.Next()
		if last == nil {
			return nil
		}
	}
	return last
}
