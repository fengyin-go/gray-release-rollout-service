package service

import "grayrelease/internal/store"

type BalancerResolverFlow struct {
	resolver store.BalancerResolverResolver
}

func NewBalancerResolverFlow(resolver store.BalancerResolverResolver) *BalancerResolverFlow {
	return &BalancerResolverFlow{resolver: resolver}
}

func (f *BalancerResolverFlow) Resolve(key string) string {
	if f.resolver == nil {
		return "unavailable"
	}
	return f.resolver.Resolve(key)
}
