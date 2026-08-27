package model

import (
	"strings"
	"time"
)

// RolloutRecord 放量记录。
type RolloutRecord struct {
	ID          string    `json:"id"`
	ReleaseID   string    `json:"release_id"`
	StepID      string    `json:"step_id"`
	FromPercent int       `json:"from_percent"`
	ToPercent   int       `json:"to_percent"`
	Operator    string    `json:"operator"`
	CreatedAt   time.Time `json:"created_at"`
}

// Validate 规范化并校验放量记录字段。
func (r *RolloutRecord) Validate() error {
	r.ReleaseID = strings.TrimSpace(r.ReleaseID)
	r.StepID = strings.TrimSpace(r.StepID)
	r.Operator = strings.TrimSpace(r.Operator)
	if r.ReleaseID == "" {
		return NewValidationError("release_id", "发布单 ID 不能为空")
	}
	if r.FromPercent < 0 || r.FromPercent > 100 {
		return NewValidationError("from_percent", "起始百分比须在 0-100 之间")
	}
	if r.ToPercent < 0 || r.ToPercent > 100 {
		return NewValidationError("to_percent", "目标百分比须在 0-100 之间")
	}
	if r.ToPercent <= r.FromPercent {
		return NewValidationError("to_percent", "目标百分比必须大于起始百分比")
	}
	return nil
}
