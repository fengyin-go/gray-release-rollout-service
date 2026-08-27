package store

import "errors"

type RetryFailure struct {
	Temporary bool
	Message   string
}

func (e *RetryFailure) Error() string { return e.Message }

// IsTemporaryRetryFailure 报告 err 是否为可重试的临时错误。
// 非 *RetryFailure 错误一律视为永久失败，不可重试。
func IsTemporaryRetryFailure(err error) bool {
	var rf *RetryFailure
	if errors.As(err, &rf) {
		return rf.Temporary
	}
	return false
}

type InstanceRegisterRetryRetryState struct {
	steps    []error
	attempts int
	commits  int
}

func NewInstanceRegisterRetryRetryState(steps ...error) *InstanceRegisterRetryRetryState {
	return &InstanceRegisterRetryRetryState{steps: append([]error(nil), steps...)}
}

func (s *InstanceRegisterRetryRetryState) Next() error {
	s.attempts++
	var err error
	if len(s.steps) > 0 {
		err = s.steps[0]
		s.steps = s.steps[1:]
	}
	if err != nil {
		// 保留原始错误类型，使调用方能够通过 Temporary 区分
		// 永久拒绝（不可重试）与临时错误（允许重试）。
		return err
	}
	s.commits++
	return nil
}

func (s *InstanceRegisterRetryRetryState) Attempts() int { return s.attempts }
func (s *InstanceRegisterRetryRetryState) Commits() int  { return s.commits }
