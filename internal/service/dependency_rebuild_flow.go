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
	// 依赖重建无论成功还是失败，返回前都必须释放活动租约，
	// 否则同一 key 的下次重建会一直看到资源占用。
	defer f.state.Release(key, token)
	if fail {
		return errors.New("operation failed")
	}
	return nil
}
