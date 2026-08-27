package service

import "grayrelease/internal/store"

type ProbeBatchFlow struct{ stream *store.ProbeBatchStream }

func NewProbeBatchFlow(stream *store.ProbeBatchStream) *ProbeBatchFlow {
	return &ProbeBatchFlow{stream: stream}
}

func (f *ProbeBatchFlow) Collect(items []string, failAt int) ([]string, error) {
	out, errs := f.stream.Start(items, failAt)
	values := make([]string, 0)
	for value := range out {
		values = append(values, value)
	}
	return values, <-errs
}
