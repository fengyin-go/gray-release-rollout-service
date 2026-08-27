package service

import (
	"grayrelease/internal/model"
)

// ImpactAnalysis 灰度影响分析。
type ImpactAnalysis struct {
	Release        *model.Release `json:"release"`
	RuleCount      int            `json:"rule_count"`
	WhitelistCount int            `json:"whitelist_count"`
	TargetPercent  int            `json:"target_percent"`
	EstimatedUsers int            `json:"estimated_users"`
	RiskLevel      string         `json:"risk_level"`
}

// AnalyzeImpact 分析灰度发布的影响面与风险等级。
func (s *Service) AnalyzeImpact(releaseID string) (*ImpactAnalysis, error) {
	r, err := s.store.GetRelease(releaseID)
	if err != nil {
		return nil, err
	}
	analysis := &ImpactAnalysis{
		Release:        r,
		RuleCount:      len(s.ListTrafficRules(releaseID)),
		WhitelistCount: len(s.ListWhitelists(releaseID)),
		TargetPercent:  s.RolloutProgress(releaseID),
	}
	maxPct := 0
	for _, t := range s.ListTrafficRules(releaseID) {
		if t.Type == model.RuleTypePercentage && t.Percentage > maxPct {
			maxPct = t.Percentage
		}
	}
	// 假设服务总用户规模为 10000，按百分比估算受影响用户。
	analysis.EstimatedUsers = maxPct * 100
	analysis.RiskLevel = riskLevel(maxPct)
	return analysis, nil
}

// riskLevel 根据放量百分比划分风险等级。
func riskLevel(pct int) string {
	switch {
	case pct >= 50:
		return "high"
	case pct >= 20:
		return "medium"
	default:
		return "low"
	}
}
