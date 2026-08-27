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
		return errors.New(err.Error())
	}
	s.commits++
	return nil
}

func (s *ConfigUpdateRetryRetryState) Attempts() int { return s.attempts }
func (s *ConfigUpdateRetryRetryState) Commits() int  { return s.commits }
