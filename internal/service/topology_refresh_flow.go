package service

import (
	"context"

	"grayrelease/internal/store"
)

type TopologyRefreshFlow struct{ state *store.TopologyRefreshState }

func NewTopologyRefreshFlow(state *store.TopologyRefreshState) *TopologyRefreshFlow {
	return &TopologyRefreshFlow{state: state}
}

func (f *TopologyRefreshFlow) Dispatch(ctx context.Context, key string) error {
	return f.state.Commit(context.Background(), key)
}
