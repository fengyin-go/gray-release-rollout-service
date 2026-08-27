package model

import (
	"strings"
	"time"
)

// ReleaseNote 发布说明。
type ReleaseNote struct {
	ID        string    `json:"id"`
	ReleaseID string    `json:"release_id"`
	Content   string    `json:"content"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate 规范化并校验发布说明字段。
func (n *ReleaseNote) Validate() error {
	n.ReleaseID = strings.TrimSpace(n.ReleaseID)
	n.Content = strings.TrimSpace(n.Content)
	n.Version = strings.TrimSpace(n.Version)
	if n.ReleaseID == "" {
		return NewValidationError("release_id", "发布单 ID 不能为空")
	}
	if n.Content == "" {
		return NewValidationError("content", "发布说明内容不能为空")
	}
	return nil
}
