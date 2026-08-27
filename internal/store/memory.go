package store

import (
	"sync"

	"grayrelease/internal/model"
)

// MemoryStore 基于内存的 Store 实现，线程安全。
type MemoryStore struct {
	mu             sync.RWMutex
	versions       map[string]*model.Version
	releases       map[string]*model.Release
	trafficRules   map[string]*model.TrafficRule
	whitelists     map[string]*model.Whitelist
	releaseSteps   map[string]*model.ReleaseStep
	rolloutRecords map[string]*model.RolloutRecord
	changeLogs     map[string]*model.ChangeLog
	rollbacks      map[string]*model.RollbackRecord
	approvals      map[string]*model.ApprovalRecord
	templates      map[string]*model.ReleaseTemplate
	releaseNotes   map[string]*model.ReleaseNote
}

// NewMemoryStore 构造空的内存存储。
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		versions:       make(map[string]*model.Version),
		releases:       make(map[string]*model.Release),
		trafficRules:   make(map[string]*model.TrafficRule),
		whitelists:     make(map[string]*model.Whitelist),
		releaseSteps:   make(map[string]*model.ReleaseStep),
		rolloutRecords: make(map[string]*model.RolloutRecord),
		changeLogs:     make(map[string]*model.ChangeLog),
		rollbacks:      make(map[string]*model.RollbackRecord),
		approvals:      make(map[string]*model.ApprovalRecord),
		templates:      make(map[string]*model.ReleaseTemplate),
		releaseNotes:   make(map[string]*model.ReleaseNote),
	}
}

var _ Store = (*MemoryStore)(nil)
