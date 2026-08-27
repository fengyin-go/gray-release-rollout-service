package service

import (
	"errors"

	"grayrelease/internal/store"
)

type DependencyRebuildFlow struct {
	state *store.DependencyRebuildLeaseState
}

func NewDependencyRebuildFlow(state *store.DependencyRebuildLeaseState) *DependencyRebuildFlow {
	return &DependencyRebuildFlow{state: state}
}

func (f *DependencyRebuildFlow) Process(key string, fail bool) error {
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
