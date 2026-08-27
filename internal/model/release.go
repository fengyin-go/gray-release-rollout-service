package model

import (
	"strings"
	"time"
)

// 发布单状态常量。
const (
	ReleaseDraft      = "draft"
	ReleaseCreating   = "creating"
	ReleaseReleasing  = "releasing"
	ReleaseCompleted  = "completed"
	ReleaseRolledBack = "rolled_back"
	ReleaseCancelled  = "cancelled"
)

// releaseTransitions 定义发布单合法状态流转。
var releaseTransitions = map[string]map[string]bool{
	ReleaseDraft: {
		ReleaseCreating:  true,
		ReleaseCancelled: true,
	},
	ReleaseCreating: {
		ReleaseReleasing: true,
		ReleaseCancelled: true,
	},
	ReleaseReleasing: {
		ReleaseCompleted:  true,
		ReleaseRolledBack: true,
	},
}

// CanReleaseTransition 判断发布单能否从 from 流转到 to。
func CanReleaseTransition(from, to string) bool {
	if m, ok := releaseTransitions[from]; ok {
		return m[to]
	}
	return false
}

// Release 灰度发布单。
type Release struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ServiceID   string    `json:"service_id"`
	VersionID   string    `json:"version_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate 规范化并校验发布单字段。
func (r *Release) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	r.ServiceID = strings.TrimSpace(r.ServiceID)
	r.VersionID = strings.TrimSpace(r.VersionID)
	r.Description = strings.TrimSpace(r.Description)
	r.CreatedBy = strings.TrimSpace(r.CreatedBy)
	if r.Name == "" {
		return NewValidationError("name", "发布单名称不能为空")
	}
	if r.ServiceID == "" {
		return NewValidationError("service_id", "服务 ID 不能为空")
	}
	if r.VersionID == "" {
		return NewValidationError("version_id", "版本 ID 不能为空")
	}
	if r.Status == "" {
		r.Status = ReleaseDraft
	}
	switch r.Status {
	case ReleaseDraft, ReleaseCreating, ReleaseReleasing, ReleaseCompleted, ReleaseRolledBack, ReleaseCancelled:
	default:
		return NewValidationError("status", "发布单状态不合法")
	}
	return nil
}

// ReleaseFilter 发布单筛选条件。
type ReleaseFilter struct {
	ServiceID string
	Status    string
	Keyword   string
}

// Match 判断发布单是否命中筛选条件。
func (f ReleaseFilter) Match(r *Release) bool {
	if f.ServiceID != "" && r.ServiceID != f.ServiceID {
		return false
	}
	if f.Status != "" && r.Status != f.Status {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(r.Name), k) {
			return false
		}
	}
	return true
}
