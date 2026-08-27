package model

import (
	"strings"
	"time"
)

// RollbackRecord 回滚记录。
type RollbackRecord struct {
	ID         string    `json:"id"`
	ReleaseID  string    `json:"release_id"`
	FromStatus string    `json:"from_status"`
	Reason     string    `json:"reason"`
	Operator   string    `json:"operator"`
	CreatedAt  time.Time `json:"created_at"`
}

// Validate 规范化并校验回滚记录字段。
func (r *RollbackRecord) Validate() error {
	r.ReleaseID = strings.TrimSpace(r.ReleaseID)
	r.FromStatus = strings.TrimSpace(r.FromStatus)
	r.Reason = strings.TrimSpace(r.Reason)
	r.Operator = strings.TrimSpace(r.Operator)
	if r.ReleaseID == "" {
		return NewValidationError("release_id", "发布单 ID 不能为空")
	}
	if r.Reason == "" {
		return NewValidationError("reason", "回滚原因不能为空")
	}
	return nil
}
