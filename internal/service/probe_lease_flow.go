package service

import (
	"errors"

	"grayrelease/internal/store"
)

type ProbeLeaseFlow struct{ state *store.ProbeLeaseLeaseState }

func NewProbeLeaseFlow(state *store.ProbeLeaseLeaseState) *ProbeLeaseFlow {
	return &ProbeLeaseFlow{state: state}
}

func (f *ProbeLeaseFlow) Process(key string, fail bool) error {
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
