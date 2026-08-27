package service

import (
	"errors"

	"grayrelease/internal/store"
)

type RegistryRefreshFlow struct {
	state *store.RegistryRefreshLeaseState
}

func NewRegistryRefreshFlow(state *store.RegistryRefreshLeaseState) *RegistryRefreshFlow {
	return &RegistryRefreshFlow{state: state}
}

func (f *RegistryRefreshFlow) Process(key string, fail bool) error {
	token, ok := f.state.Acquire(key)
	if !ok {
		return errors.New("lease busy")
	}
	// Register the release before any failure path: a failed refresh must
	// still give up the occupation marker, otherwise a same-key restart is
	// blocked by the stale lease.
	defer f.state.Release(key, token)
	if fail {
		return errors.New("operation failed")
	}
	return nil
}
