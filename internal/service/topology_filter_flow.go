package service

import "grayrelease/internal/store"

type TopologyFilterFlow struct{ resolver store.TopologyFilterResolver }

func NewTopologyFilterFlow(resolver store.TopologyFilterResolver) *TopologyFilterFlow {
	return &TopologyFilterFlow{resolver: resolver}
}

func (f *TopologyFilterFlow) Resolve(key string) string {
	if f.resolver == nil {
		return "unavailable"
	}
	return f.resolver.Resolve(key)
}
