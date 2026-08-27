package service

import "grayrelease/internal/store"

type HealthAggregationFlow struct {
	stream *store.HealthAggregationStream
}

func NewHealthAggregationFlow(stream *store.HealthAggregationStream) *HealthAggregationFlow {
	return &HealthAggregationFlow{stream: stream}
}

func (f *HealthAggregationFlow) Collect(items []string, failAt int) ([]string, error) {
	out, errs := f.stream.Start(items, failAt)
	values := make([]string, 0)
	for value := range out {
		values = append(values, value)
	}
	return values, <-errs
}
