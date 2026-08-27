package service

import "grayrelease/internal/store"

type DependencyScanFlow struct{ stream *store.DependencyScanStream }

func NewDependencyScanFlow(stream *store.DependencyScanStream) *DependencyScanFlow {
	return &DependencyScanFlow{stream: stream}
}

func (f *DependencyScanFlow) Collect(items []string, failAt int) ([]string, error) {
	out, errs := f.stream.Start(items, failAt)
	values := make([]string, 0)
	for value := range out {
		values = append(values, value)
	}
	return values, <-errs
}
