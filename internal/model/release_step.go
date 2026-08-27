package model

import (
	"strings"
	"time"
)

// 发布步骤状态常量。
const (
	StepPending    = "pending"
	StepInProgress = "in_progress"
	StepCompleted  = "completed"
)

// ReleaseStep 灰度发布步骤（放量阶段）。
type ReleaseStep struct {
	ID            string     `json:"id"`
	ReleaseID     string     `json:"release_id"`
	StepNo        int        `json:"step_no"`
	Name          string     `json:"name"`
	TargetPercent int        `json:"target_percent"`
	Status        string     `json:"status"`
	StartedAt     *time.Time `json:"started_at,omitempty"`
	FinishedAt    *time.Time `json:"finished_at,omitempty"`
}

// Validate 规范化并校验发布步骤字段。
func (s *ReleaseStep) Validate() error {
	s.ReleaseID = strings.TrimSpace(s.ReleaseID)
	s.Name = strings.TrimSpace(s.Name)
	if s.ReleaseID == "" {
		return NewValidationError("release_id", "发布单 ID 不能为空")
	}
	if s.Name == "" {
		return NewValidationError("name", "步骤名称不能为空")
	}
	if s.StepNo <= 0 {
		return NewValidationError("step_no", "步骤序号必须大于 0")
	}
	if s.TargetPercent <= 0 || s.TargetPercent > 100 {
		return NewValidationError("target_percent", "目标百分比须在 1-100 之间")
	}
	if s.Status == "" {
		s.Status = StepPending
	}
	if s.Status != StepPending && s.Status != StepInProgress && s.Status != StepCompleted {
		return NewValidationError("status", "步骤状态不合法")
	}
	return nil
}
