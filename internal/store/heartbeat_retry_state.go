package store

import "errors"

// RetryFailure 表示一次心跳写入的失败结果。
// Temporary 为 true 表示临时性繁忙（可重试），false 表示永久拒绝（不可提交）。
type RetryFailure struct {
	Temporary bool
	Message   string
}

func (e *RetryFailure) Error() string { return e.Message }

// IsTemporary 判断错误是否为可重试的临时繁忙。
// 非临时性（永久拒绝）与无法识别的错误一律视为不可重试，避免把永久拒绝误当临时繁忙重提。
func IsTemporary(err error) bool {
	if err == nil {
		return false
	}
	var f *RetryFailure
	if errors.As(err, &f) {
		return f.Temporary
	}
	return false
}

type HeartbeatRetryRetryState struct {
	steps    []error
	attempts int
	commits  int
	last     error
}

func NewHeartbeatRetryRetryState(steps ...error) *HeartbeatRetryRetryState {
	return &HeartbeatRetryRetryState{steps: append([]error(nil), steps...)}
}

func (s *HeartbeatRetryRetryState) Next() error {
	s.attempts++
	var err error
	if len(s.steps) > 0 {
		err = s.steps[0]
		s.steps = s.steps[1:]
	}
	if err != nil {
		// 保留原始错误类型，使调用方能够区分临时繁忙与永久拒绝。
		s.last = err
		return err
	}
	s.commits++
	return nil
}

// Last 返回最近一次 Next 遇到的错误，供流程在重试配额用尽后回带临时繁忙。
func (s *HeartbeatRetryRetryState) Last() error { return s.last }

func (s *HeartbeatRetryRetryState) Attempts() int { return s.attempts }
func (s *HeartbeatRetryRetryState) Commits() int  { return s.commits }
