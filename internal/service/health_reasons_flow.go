package service

import "grayrelease/internal/store"

type HealthReasonsFlow struct {
	state *store.HealthReasonsState
}

func NewHealthReasonsFlow(state *store.HealthReasonsState) *HealthReasonsFlow {
	return &HealthReasonsFlow{state: state}
}

func (f *HealthReasonsFlow) Publish(key string, values []string) []string {
	f.state.Replace(key, values)
	return f.state.Snapshot(key)
}

func (f *HealthReasonsFlow) Snapshot(key string) []string {
	return f.state.Snapshot(key)
}
