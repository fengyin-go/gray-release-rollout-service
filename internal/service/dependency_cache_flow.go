package service

import "grayrelease/internal/store"

type DependencyCacheFlow struct {
	state *store.DependencyCacheState
}

func NewDependencyCacheFlow(state *store.DependencyCacheState) *DependencyCacheFlow {
	return &DependencyCacheFlow{state: state}
}

func (f *DependencyCacheFlow) Publish(key string, values []string) []string {
	f.state.Replace(key, values)
	return f.state.Snapshot(key)
}

func (f *DependencyCacheFlow) Snapshot(key string) []string {
	return f.state.Snapshot(key)
}
