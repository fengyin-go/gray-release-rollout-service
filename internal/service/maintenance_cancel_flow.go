package service

import (
	"context"

	"grayrelease/internal/store"
)

type MaintenanceCancelFlow struct{ state *store.MaintenanceCancelState }

func NewMaintenanceCancelFlow(state *store.MaintenanceCancelState) *MaintenanceCancelFlow {
	return &MaintenanceCancelFlow{state: state}
}

func (f *MaintenanceCancelFlow) Dispatch(ctx context.Context, key string) error {
	return f.state.Commit(context.Background(), key)
}
