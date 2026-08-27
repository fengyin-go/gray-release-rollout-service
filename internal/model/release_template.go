package model

import (
	"strings"
	"time"
)

// ReleaseTemplate 放量模板。
type ReleaseTemplate struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Schedule    []int     `json:"schedule"`
	CreatedAt   time.Time `json:"created_at"`
}

// Validate 规范化并校验模板字段。
func (t *ReleaseTemplate) Validate() error {
	t.Name = strings.TrimSpace(t.Name)
	t.Description = strings.TrimSpace(t.Description)
	if t.Name == "" {
		return NewValidationError("name", "模板名称不能为空")
	}
	if len(t.Schedule) == 0 {
		return NewValidationError("schedule", "放量百分比序列不能为空")
	}
	for i := 1; i < len(t.Schedule); i++ {
		if t.Schedule[i] <= t.Schedule[i-1] {
			return NewValidationError("schedule", "放量百分比必须严格递增")
		}
	}
	if t.Schedule[len(t.Schedule)-1] > 100 {
		return NewValidationError("schedule", "放量百分比不能超过 100")
	}
	return nil
}
