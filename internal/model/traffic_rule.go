package model

import (
	"strings"
	"time"
)

// 流量规则类型常量。
const (
	RuleTypePercentage = "percentage"
	RuleTypeHeader     = "header"
	RuleTypeCookie     = "cookie"
	RuleTypeUser       = "user"
)

// TrafficRule 灰度流量规则。
type TrafficRule struct {
	ID         string    `json:"id"`
	ReleaseID  string    `json:"release_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	MatchKey   string    `json:"match_key"`
	MatchValue string    `json:"match_value"`
	Percentage int       `json:"percentage"`
	Priority   int       `json:"priority"`
	Enabled    bool      `json:"enabled"`
	CreatedAt  time.Time `json:"created_at"`
}

// Validate 规范化并校验流量规则字段。
func (t *TrafficRule) Validate() error {
	t.ReleaseID = strings.TrimSpace(t.ReleaseID)
	t.Name = strings.TrimSpace(t.Name)
	t.MatchKey = strings.TrimSpace(t.MatchKey)
	t.MatchValue = strings.TrimSpace(t.MatchValue)
	if t.ReleaseID == "" {
		return NewValidationError("release_id", "发布单 ID 不能为空")
	}
	if t.Name == "" {
		return NewValidationError("name", "规则名称不能为空")
	}
	if t.Type == "" {
		t.Type = RuleTypePercentage
	}
	switch t.Type {
	case RuleTypePercentage, RuleTypeHeader, RuleTypeCookie, RuleTypeUser:
	default:
		return NewValidationError("type", "流量规则类型不合法")
	}
	if t.Type != RuleTypePercentage && t.MatchKey == "" {
		return NewValidationError("match_key", "匹配键不能为空")
	}
	if t.Percentage < 0 || t.Percentage > 100 {
		return NewValidationError("percentage", "百分比须在 0-100 之间")
	}
	if t.Priority < 0 {
		return NewValidationError("priority", "优先级不能为负数")
	}
	return nil
}
