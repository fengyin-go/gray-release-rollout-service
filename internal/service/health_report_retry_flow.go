package service

import "grayrelease/internal/store"

type HealthReportRetryFlow struct {
	state *store.HealthReportRetryRetryState
}

func NewHealthReportRetryFlow(state *store.HealthReportRetryRetryState) *HealthReportRetryFlow {
	return &HealthReportRetryFlow{state: state}
}

func (f *HealthReportRetryFlow) Execute() error {
	var last error
	for attempt := 0; attempt < 2; attempt++ {
		last = f.state.Next()
		if last == nil {
			return nil
		}
	}
	return last
}
