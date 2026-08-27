package service

import "grayrelease/internal/store"

type ConfigUpdateRetryFlow struct {
	state *store.ConfigUpdateRetryRetryState
}

func NewConfigUpdateRetryFlow(state *store.ConfigUpdateRetryRetryState) *ConfigUpdateRetryFlow {
	return &ConfigUpdateRetryFlow{state: state}
}

// Execute 推进配置更新。仅当遇到临时占用时重试，
// 永久拒绝（不可恢复）须原样返回，不重试、不提交。
func (f *ConfigUpdateRetryFlow) Execute() error {
	var last error
	for attempt := 0; attempt < 2; attempt++ {
		last = f.state.Next()
		if last == nil {
			return nil
		}
		// 临时占用可重试；永久拒绝原样返回，不再尝试。
		if !store.IsRetryable(last) {
			return last
		}
	}
	return last
}
