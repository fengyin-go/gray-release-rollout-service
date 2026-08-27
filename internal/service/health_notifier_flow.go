package service

import "grayrelease/internal/store"

type HealthNotifierFlow struct{ resolver store.HealthNotifierResolver }

func NewHealthNotifierFlow(resolver store.HealthNotifierResolver) *HealthNotifierFlow {
	return &HealthNotifierFlow{resolver: resolver}
}

func (f *HealthNotifierFlow) Resolve(key string) string {
	if f.resolver == nil {
		return "unavailable"
	}
	return f.resolver.Resolve(key)
}
