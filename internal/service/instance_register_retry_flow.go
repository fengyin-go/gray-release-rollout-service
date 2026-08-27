package service

import "grayrelease/internal/store"

type InstanceRegisterRetryFlow struct {
	state *store.InstanceRegisterRetryRetryState
}

func NewInstanceRegisterRetryFlow(state *store.InstanceRegisterRetryRetryState) *InstanceRegisterRetryFlow {
	return &InstanceRegisterRetryFlow{state: state}
}

// Execute 执行实例注册，遵循以下重试策略：
//   - 成功（nil）立即返回；
//   - 永久拒绝（非临时错误）立即返回，不重试；
//   - 临时错误允许重试，最多 2 次。
func (f *InstanceRegisterRetryFlow) Execute() error {
	var last error
	for attempt := 0; attempt < 2; attempt++ {
		last = f.state.Next()
		if last == nil {
			return nil
		}
		if !store.IsTemporaryRetryFailure(last) {
			return last
		}
	}
	return last
}
