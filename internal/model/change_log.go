package model

import (
	"strings"
	"time"
)

// 变更动作常量。
const (
	ChangeCreated        = "created"
	ChangeStatusChanged  = "status_changed"
	ChangeStepAdvanced   = "step_advanced"
	ChangeRolledBack     = "rolled_back"
	ChangeCancelled      = "cancelled"
	ChangeRuleChanged    = "rule_changed"
)

// ChangeLog 发布单变更记录。
type ChangeLog struct {
	ID        string    `json:"id"`
	ReleaseID string    `json:"release_id"`
	Action    string    `json:"action"`
	Detail    string    `json:"detail"`
	Operator  string    `json:"operator"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate 规范化并校验变更记录字段。
func (c *ChangeLog) Validate() error {
	c.ReleaseID = strings.TrimSpace(c.ReleaseID)
	c.Action = strings.TrimSpace(c.Action)
	c.Detail = strings.TrimSpace(c.Detail)
	c.Operator = strings.TrimSpace(c.Operator)
	if c.Action == "" {
		return NewValidationError("action", "变更动作不能为空")
	}
	return nil
}

// ChangeLogFilter 变更记录筛选条件。
type ChangeLogFilter struct {
	ReleaseID string
	Action    string
	Operator  string
}

// Match 判断变更记录是否命中筛选条件。
func (f ChangeLogFilter) Match(c *ChangeLog) bool {
	if f.ReleaseID != "" && c.ReleaseID != f.ReleaseID {
		return false
	}
	if f.Action != "" && c.Action != f.Action {
		return false
	}
	if f.Operator != "" && c.Operator != f.Operator {
		return false
	}
	return true
}
