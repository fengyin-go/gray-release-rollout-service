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
	if fail {
		return errors.New("operation failed")
	}
	defer f.state.Release(key, token)
	return nil
}
