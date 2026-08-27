package service

import "grayrelease/internal/store"

type HeartbeatRetryFlow struct {
	state *store.HeartbeatRetryRetryState
}

func NewHeartbeatRetryFlow(state *store.HeartbeatRetryRetryState) *HeartbeatRetryFlow {
	return &HeartbeatRetryFlow{state: state}
}

// Execute 推进心跳写入流程，仅对临时繁忙重试，遇到永久拒绝立即停止且不提交。
func (f *HeartbeatRetryFlow) Execute() error {
	// 最多重试一次临时繁忙：初次写入失败后，对临时错误再尝试一次。
	for attempt := 0; attempt < 2; attempt++ {
		err := f.state.Next()
		if err == nil {
			return nil
		}
		// 永久拒绝不可重试、不可提交，立即返回，不再继续。
		if !store.IsTemporary(err) {
			return err
		}
		// 临时繁忙：仍有重试配额时继续，否则跳出后返回该临时错误。
	}
	return f.state.Last()
}
