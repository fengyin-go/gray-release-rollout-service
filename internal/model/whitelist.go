package model

import (
	"strings"
	"time"
)

// Whitelist 灰度白名单。
type Whitelist struct {
	ID        string    `json:"id"`
	ReleaseID string    `json:"release_id"`
	UserID    string    `json:"user_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate 规范化并校验白名单字段。
func (w *Whitelist) Validate() error {
	w.ReleaseID = strings.TrimSpace(w.ReleaseID)
	w.UserID = strings.TrimSpace(w.UserID)
	w.Note = strings.TrimSpace(w.Note)
	if w.ReleaseID == "" {
		return NewValidationError("release_id", "发布单 ID 不能为空")
	}
	if w.UserID == "" {
		return NewValidationError("user_id", "用户 ID 不能为空")
	}
	return nil
}
