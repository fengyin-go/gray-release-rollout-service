package service

import "grayrelease/internal/store"

type InstanceRegisterRetryFlow struct {
	state *store.InstanceRegisterRetryRetryState
}

func NewInstanceRegisterRetryFlow(state *store.InstanceRegisterRetryRetryState) *InstanceRegisterRetryFlow {
	return &InstanceRegisterRetryFlow{state: state}
}

func (f *InstanceRegisterRetryFlow) Execute() error {
	var last error
	for attempt := 0; attempt < 2; attempt++ {
		last = f.state.Next()
		if last == nil {
			return nil
		}
	}
	return last
}
