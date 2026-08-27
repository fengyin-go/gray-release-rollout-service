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
	// 在任何可能 return 的分支之前注册释放，确保失败路径也能释放租约，
	// 否则下一次对相同 key 的恢复会被 "lease busy" 挡住。
	defer f.state.Release(key, token)
	if fail {
		return errors.New("operation failed")
	}
	return nil
}
