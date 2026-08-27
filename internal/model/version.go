package model

import (
	"strings"
	"time"
)

// Version 服务版本实体。
type Version struct {
	ID          string    `json:"id"`
	ServiceID   string    `json:"service_id"`
	Version     string    `json:"version"`
	ArtifactURL string    `json:"artifact_url"`
	Checksum    string    `json:"checksum"`
	SizeBytes   int64     `json:"size_bytes"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// Validate 规范化并校验版本字段。
func (v *Version) Validate() error {
	v.ServiceID = strings.TrimSpace(v.ServiceID)
	v.Version = strings.TrimSpace(v.Version)
	v.ArtifactURL = strings.TrimSpace(v.ArtifactURL)
	v.Checksum = strings.TrimSpace(v.Checksum)
	v.Description = strings.TrimSpace(v.Description)
	if v.ServiceID == "" {
		return NewValidationError("service_id", "服务 ID 不能为空")
	}
	if v.Version == "" {
		return NewValidationError("version", "版本号不能为空")
	}
	if v.ArtifactURL == "" {
		return NewValidationError("artifact_url", "制品地址不能为空")
	}
	if v.SizeBytes < 0 {
		return NewValidationError("size_bytes", "制品大小不能为负数")
	}
	return nil
}
