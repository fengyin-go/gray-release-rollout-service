// Package store 定义数据访问接口与内存实现。
package store

import (
	"errors"

	"grayrelease/internal/model"
)

var (
	ErrNotFound = errors.New("记录不存在")
	ErrConflict = errors.New("记录已存在或状态冲突")
)

// Store 聚合全部实体的数据访问方法，便于测试时替换实现。
type Store interface {
	// 版本
	CreateVersion(v *model.Version) error
	GetVersion(id string) (*model.Version, error)
	ListVersions() []*model.Version
	UpdateVersion(v *model.Version) error
	DeleteVersion(id string) error

	// 发布单
	CreateRelease(r *model.Release) error
	GetRelease(id string) (*model.Release, error)
	ListReleases() []*model.Release
	UpdateRelease(r *model.Release) error
	DeleteRelease(id string) error

	// 流量规则
	CreateTrafficRule(t *model.TrafficRule) error
	GetTrafficRule(id string) (*model.TrafficRule, error)
	ListTrafficRules() []*model.TrafficRule
	UpdateTrafficRule(t *model.TrafficRule) error
	DeleteTrafficRule(id string) error

	// 白名单
	CreateWhitelist(w *model.Whitelist) error
	GetWhitelist(id string) (*model.Whitelist, error)
	ListWhitelists() []*model.Whitelist
	DeleteWhitelist(id string) error

	// 发布步骤
	CreateReleaseStep(s *model.ReleaseStep) error
	GetReleaseStep(id string) (*model.ReleaseStep, error)
	ListReleaseSteps() []*model.ReleaseStep
	UpdateReleaseStep(s *model.ReleaseStep) error
	DeleteReleaseStep(id string) error

	// 放量记录
	CreateRolloutRecord(r *model.RolloutRecord) error
	ListRolloutRecords() []*model.RolloutRecord

	// 变更记录
	CreateChangeLog(c *model.ChangeLog) error
	ListChangeLogs() []*model.ChangeLog

	// 回滚记录
	CreateRollbackRecord(r *model.RollbackRecord) error
	ListRollbackRecords() []*model.RollbackRecord

	// 审批记录
	CreateApprovalRecord(a *model.ApprovalRecord) error
	GetApprovalRecord(id string) (*model.ApprovalRecord, error)
	ListApprovalRecords() []*model.ApprovalRecord
	UpdateApprovalRecord(a *model.ApprovalRecord) error

	// 放量模板
	CreateReleaseTemplate(t *model.ReleaseTemplate) error
	GetReleaseTemplate(id string) (*model.ReleaseTemplate, error)
	ListReleaseTemplates() []*model.ReleaseTemplate
	UpdateReleaseTemplate(t *model.ReleaseTemplate) error
	DeleteReleaseTemplate(id string) error

	// 发布说明
	CreateReleaseNote(n *model.ReleaseNote) error
	GetReleaseNote(id string) (*model.ReleaseNote, error)
	ListReleaseNotes() []*model.ReleaseNote
	UpdateReleaseNote(n *model.ReleaseNote) error
	DeleteReleaseNote(id string) error
}
