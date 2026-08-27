package service

import "grayrelease/internal/store"

type HealthReportRetryFlow struct {
	state *store.HealthReportRetryRetryState
}

func NewHealthReportRetryFlow(state *store.HealthReportRetryRetryState) *HealthReportRetryFlow {
	return &HealthReportRetryFlow{state: state}
}

func (f *HealthReportRetryFlow) Execute() error {
	last := f.state.Next()
	if last == nil {
		return nil
	}
	// 只让临时错误再试一次；不可重试的错误立即返回，并保留原始错误类别。
	if rf, ok := last.(*store.RetryFailure); ok && rf.Temporary {
		last = f.state.Next()
		if last == nil {
			return nil
		}
	}
	return last
}
