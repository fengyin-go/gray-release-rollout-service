package service

import "grayrelease/internal/store"

type EventSinkFlow struct{ resolver store.EventSinkResolver }

func NewEventSinkFlow(resolver store.EventSinkResolver) *EventSinkFlow {
	return &EventSinkFlow{resolver: resolver}
}

func (f *EventSinkFlow) Resolve(key string) string {
	if f.resolver == nil {
		return "unavailable"
	}
	return f.resolver.Resolve(key)
}
