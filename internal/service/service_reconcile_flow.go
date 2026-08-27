package service

import (
	"errors"

	"grayrelease/internal/store"
)

type ServiceReconcileFlow struct {
	state *store.ServiceReconcileLeaseState
}

func NewServiceReconcileFlow(state *store.ServiceReconcileLeaseState) *ServiceReconcileFlow {
	return &ServiceReconcileFlow{state: state}
}

func (f *ServiceReconcileFlow) Process(key string, fail bool) error {
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
