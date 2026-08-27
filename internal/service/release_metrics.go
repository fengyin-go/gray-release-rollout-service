package service

import (
	"time"

	"grayrelease/internal/model"
)

// ReleaseMetrics 发布单运行指标。
type ReleaseMetrics struct {
	Release        *model.Release `json:"release"`
	AgeSeconds     int64          `json:"age_seconds"`
	TotalSteps     int            `json:"total_steps"`
	CompletedSteps int            `json:"completed_steps"`
	Progress       int            `json:"progress"`
	RollbackCount  int            `json:"rollback_count"`
	RuleCount      int            `json:"rule_count"`
	WhitelistCount int            `json:"whitelist_count"`
	ChangeLogCount int            `json:"change_log_count"`
}

// ReleaseMetrics 计算发布单运行指标。
func (s *Service) ReleaseMetrics(releaseID string) (*ReleaseMetrics, error) {
	r, err := s.store.GetRelease(releaseID)
	if err != nil {
		return nil, err
	}
	m := &ReleaseMetrics{
		Release:        r,
		AgeSeconds:     int64(time.Since(r.CreatedAt).Seconds()),
		Progress:       s.RolloutProgress(releaseID),
		RuleCount:      len(s.ListTrafficRules(releaseID)),
		WhitelistCount: len(s.ListWhitelists(releaseID)),
		RollbackCount:  len(s.ListRollbackRecords(releaseID)),
	}
	for _, st := range s.store.ListReleaseSteps() {
		if st.ReleaseID != releaseID {
			continue
		}
		m.TotalSteps++
		if st.Status == model.StepCompleted {
			m.CompletedSteps++
		}
	}
	for _, c := range s.store.ListChangeLogs() {
		if c.ReleaseID == releaseID {
			m.ChangeLogCount++
		}
	}
	return m, nil
}
