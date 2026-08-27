package service

import "grayrelease/internal/store"

type ConfigUpdateRetryFlow struct {
	state *store.ConfigUpdateRetryRetryState
}

func NewConfigUpdateRetryFlow(state *store.ConfigUpdateRetryRetryState) *ConfigUpdateRetryFlow {
	return &ConfigUpdateRetryFlow{state: state}
}

func (f *ConfigUpdateRetryFlow) Execute() error {
	var last error
	for attempt := 0; attempt < 2; attempt++ {
		last = f.state.Next()
		if last == nil {
			return nil
		}
	}
	return last
}
