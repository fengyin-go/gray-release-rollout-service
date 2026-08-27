package service

import "grayrelease/internal/store"

type EventExportFlow struct{ stream *store.EventExportStream }

func NewEventExportFlow(stream *store.EventExportStream) *EventExportFlow {
	return &EventExportFlow{stream: stream}
}

func (f *EventExportFlow) Collect(items []string, failAt int) ([]string, error) {
	out, errs := f.stream.Start(items, failAt)
	values := make([]string, 0)
	for value := range out {
		values = append(values, value)
	}
	return values, <-errs
}
