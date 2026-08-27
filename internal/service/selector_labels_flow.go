package service

import "grayrelease/internal/store"

type SelectorLabelsFlow struct {
	state *store.SelectorLabelsState
}

func NewSelectorLabelsFlow(state *store.SelectorLabelsState) *SelectorLabelsFlow {
	return &SelectorLabelsFlow{state: state}
}

func (f *SelectorLabelsFlow) Publish(key string, values []string) []string {
	f.state.Replace(key, values)
	return f.state.Snapshot(key)
}

func (f *SelectorLabelsFlow) Snapshot(key string) []string {
	return f.state.Snapshot(key)
}
