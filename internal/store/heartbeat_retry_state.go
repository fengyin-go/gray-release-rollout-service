package store

import "errors"

type RetryFailure struct {
	Temporary bool
	Message   string
}

func (e *RetryFailure) Error() string { return e.Message }

type HeartbeatRetryRetryState struct {
	steps    []error
	attempts int
	commits  int
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
		return errors.New(err.Error())
	}
	s.commits++
	return nil
}

func (s *HeartbeatRetryRetryState) Attempts() int { return s.attempts }
func (s *HeartbeatRetryRetryState) Commits() int  { return s.commits }
