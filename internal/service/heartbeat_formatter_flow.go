package service

import "grayrelease/internal/store"

type HeartbeatFormatterFlow struct {
	resolver store.HeartbeatFormatterResolver
}

func NewHeartbeatFormatterFlow(resolver store.HeartbeatFormatterResolver) *HeartbeatFormatterFlow {
	return &HeartbeatFormatterFlow{resolver: resolver}
}

func (f *HeartbeatFormatterFlow) Resolve(key string) string {
	if f.resolver == nil {
		return "unavailable"
	}
	return f.resolver.Resolve(key)
}
