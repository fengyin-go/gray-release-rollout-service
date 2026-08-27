package service

import (
	"time"

	"grayrelease/internal/model"
)

// ReleaseETA 发布完成时间预估。
type ReleaseETA struct {
	Release                   *model.Release `json:"release"`
	Progress                  int            `json:"progress"`
	StartedAt                 *time.Time     `json:"started_at,omitempty"`
	LastAdvanceAt             *time.Time     `json:"last_advance_at,omitempty"`
	AvgStepSeconds            int64          `json:"avg_step_seconds"`
	RemainingSteps            int            `json:"remaining_steps"`
	EstimatedRemainingSeconds int64          `json:"estimated_remaining_seconds"`
}

// ReleaseETA 根据放量历史估算发布完成还需的时间。
func (s *Service) ReleaseETA(releaseID string) (*ReleaseETA, error) {
	r, err := s.store.GetRelease(releaseID)
	if err != nil {
		return nil, err
	}
	eta := &ReleaseETA{
		Release:  r,
		Progress: s.RolloutProgress(releaseID),
	}
	records := s.ListRolloutRecords(releaseID)
	if len(records) >= 2 {
		// records 为时间倒序，取最早与最晚放量记录
		first := records[len(records)-1]
		last := records[0]
		eta.StartedAt = &first.CreatedAt
		eta.LastAdvanceAt = &last.CreatedAt
		dur := last.CreatedAt.Sub(first.CreatedAt)
		eta.AvgStepSeconds = int64(dur.Seconds()) / int64(len(records)-1)
	}
	for _, st := range s.ListReleaseSteps(releaseID) {
		if st.Status != model.StepCompleted {
			eta.RemainingSteps++
		}
	}
	eta.EstimatedRemainingSeconds = eta.AvgStepSeconds * int64(eta.RemainingSteps)
	return eta, nil
}
