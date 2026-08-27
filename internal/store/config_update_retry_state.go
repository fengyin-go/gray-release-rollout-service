package store

import "errors"

type RetryFailure struct {
	Temporary bool
	Message   string
}

func (e *RetryFailure) Error() string { return e.Message }

type ConfigUpdateRetryRetryState struct {
	steps    []error
	attempts int
	commits  int
}

func NewConfigUpdateRetryRetryState(steps ...error) *ConfigUpdateRetryRetryState {
	return &ConfigUpdateRetryRetryState{steps: append([]error(nil), steps...)}
}

func (s *ConfigUpdateRetryRetryState) Next() error {
	s.attempts++
	var err error
	if len(s.steps) > 0 {
		err = s.steps[0]
		s.steps = s.steps[1:]
	}
	if err != nil {
		// 返回原始错误，保留 *RetryFailure 类型，
		// 以便上层根据 Temporary 区分可重试与不可恢复。
		return err
	}
	s.commits++
	return nil
}

// IsRetryable 报告错误是否为可重试的临时占用。
// 永久拒绝（Temporary=false 或非 *RetryFailure）不可重试，须原样返回。
func IsRetryable(err error) bool {
	if err == nil {
		return false
	}
	var f *RetryFailure
	if !errors.As(err, &f) {
		return false
	}
	return f.Temporary
}

func (s *ConfigUpdateRetryRetryState) Attempts() int { return s.attempts }
func (s *ConfigUpdateRetryRetryState) Commits() int  { return s.commits }
