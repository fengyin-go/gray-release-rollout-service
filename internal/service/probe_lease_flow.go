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
	// Release on every post-acquire return path, including the error branch.
	// Otherwise a failed probe leaves the lease held and the same instance can
	// never be probed again.
	defer f.state.Release(key, token)
	if fail {
		return errors.New("operation failed")
	}
	return nil
}
