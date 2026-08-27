package store

import "errors"

type ProbeBatchStream struct{}

func NewProbeBatchStream() *ProbeBatchStream { return &ProbeBatchStream{} }

func (s *ProbeBatchStream) Start(items []string, failAt int) (<-chan string, <-chan error) {
	out := make(chan string)
	errs := make(chan error, 1)
	go func() {
		// close(out) so consumers draining the result stream (Collect's range)
		// can terminate on both the normal-completion and error paths;
		// otherwise the range blocks forever and the error is never returned.
		defer close(out)
		defer close(errs)
		for index, item := range items {
			if index == failAt {
				errs <- errors.New("stream interrupted")
				return
			}
			out <- item
		}
	}()
	return out, errs
}
