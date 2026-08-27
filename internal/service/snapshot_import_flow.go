package service

import (
	"errors"

	"grayrelease/internal/store"
)

type SnapshotImportFlow struct {
	state *store.SnapshotImportLeaseState
}

func NewSnapshotImportFlow(state *store.SnapshotImportLeaseState) *SnapshotImportFlow {
	return &SnapshotImportFlow{state: state}
}

func (f *SnapshotImportFlow) Process(key string, fail bool) error {
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
