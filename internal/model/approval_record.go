package model

import (
	"strings"
	"time"
)

// 审批状态常量。
const (
	ApprovalPending  = "pending"
	ApprovalApproved = "approved"
	ApprovalRejected = "rejected"
)

// ApprovalRecord 发布审批记录。
type ApprovalRecord struct {
	ID        string    `json:"id"`
	ReleaseID string    `json:"release_id"`
	Approver  string    `json:"approver"`
	Status    string    `json:"status"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate 规范化并校验审批记录字段。
func (a *ApprovalRecord) Validate() error {
	a.ReleaseID = strings.TrimSpace(a.ReleaseID)
	a.Approver = strings.TrimSpace(a.Approver)
	a.Comment = strings.TrimSpace(a.Comment)
	if a.ReleaseID == "" {
		return NewValidationError("release_id", "发布单 ID 不能为空")
	}
	if a.Approver == "" {
		return NewValidationError("approver", "审批人不能为空")
	}
	if a.Status == "" {
		a.Status = ApprovalPending
	}
	if a.Status != ApprovalPending && a.Status != ApprovalApproved && a.Status != ApprovalRejected {
		return NewValidationError("status", "审批状态不合法")
	}
	return nil
}
