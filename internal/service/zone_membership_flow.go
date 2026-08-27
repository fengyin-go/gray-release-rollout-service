package service

import "grayrelease/internal/store"

type ZoneMembershipFlow struct {
	state *store.ZoneMembershipState
}

func NewZoneMembershipFlow(state *store.ZoneMembershipState) *ZoneMembershipFlow {
	return &ZoneMembershipFlow{state: state}
}

func (f *ZoneMembershipFlow) Publish(key string, values []string) []string {
	f.state.Replace(key, values)
	return f.state.Snapshot(key)
}

func (f *ZoneMembershipFlow) Snapshot(key string) []string {
	return f.state.Snapshot(key)
}
