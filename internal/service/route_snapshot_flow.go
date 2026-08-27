package service

import "grayrelease/internal/store"

type RouteSnapshotFlow struct {
	state *store.RouteSnapshotState
}

func NewRouteSnapshotFlow(state *store.RouteSnapshotState) *RouteSnapshotFlow {
	return &RouteSnapshotFlow{state: state}
}

func (f *RouteSnapshotFlow) Publish(key string, values []string) []string {
	f.state.Replace(key, values)
	return f.state.Snapshot(key)
}

func (f *RouteSnapshotFlow) Snapshot(key string) []string {
	return f.state.Snapshot(key)
}
