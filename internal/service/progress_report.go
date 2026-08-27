package service

import (
	"grayrelease/internal/model"
)

// ProgressReport 发布进度综合报告。
type ProgressReport struct {
	Release        *model.Release        `json:"release"`
	Progress       int                   `json:"progress"`
	Steps          []*model.ReleaseStep  `json:"steps"`
	NextStep       *model.ReleaseStep    `json:"next_step,omitempty"`
	Timeline       []TimelinePoint       `json:"timeline"`
	Rollbacks      int                   `json:"rollbacks"`
	RolloutRecords []*model.RolloutRecord `json:"rollout_records"`
	ReadyToComplete bool                 `json:"ready_to_complete"`
}

// ProgressReport 生成发布单进度综合报告。
func (s *Service) ProgressReport(releaseID string) (*ProgressReport, error) {
	r, err := s.store.GetRelease(releaseID)
	if err != nil {
		return nil, err
	}
	report := &ProgressReport{
		Release:         r,
		Progress:        s.RolloutProgress(releaseID),
		Steps:           s.ListReleaseSteps(releaseID),
		Timeline:        s.RolloutTimeline(releaseID),
		Rollbacks:       len(s.ListRollbackRecords(releaseID)),
		RolloutRecords:  s.ListRolloutRecords(releaseID),
		ReadyToComplete: s.CompleteReleaseCheck(releaseID),
	}
	if next, err := s.NextStep(releaseID); err == nil {
		report.NextStep = next
	}
	return report, nil
}
