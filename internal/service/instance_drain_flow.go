package service

import "grayrelease/internal/store"

type InstanceDrainFlow struct{ stream *store.InstanceDrainStream }

func NewInstanceDrainFlow(stream *store.InstanceDrainStream) *InstanceDrainFlow {
	return &InstanceDrainFlow{stream: stream}
}

func (f *InstanceDrainFlow) Collect(items []string, failAt int) ([]string, error) {
	out, errs := f.stream.Start(items, failAt)
	values := make([]string, 0)
	for value := range out {
		values = append(values, value)
	}
	return values, <-errs
}
