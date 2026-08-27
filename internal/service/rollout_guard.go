package service

import (
	"grayrelease/internal/model"
)

// RolloutGuard 放量守卫结果。
type RolloutGuard struct {
	Allowed      bool     `json:"allowed"`
	Reasons      []string `json:"reasons"`
	CurrentStep  *model.ReleaseStep `json:"current_step,omitempty"`
}

// CheckRolloutGuard 检查发布单是否具备推进条件。
func (s *Service) CheckRolloutGuard(releaseID string) RolloutGuard {
	guard := RolloutGuard{Allowed: true, Reasons: []string{}}
	r, err := s.store.GetRelease(releaseID)
	if err != nil {
		guard.Allowed = false
		guard.Reasons = append(guard.Reasons, "发布单不存在")
		return guard
	}
	if r.Status != model.ReleaseReleasing && r.Status != model.ReleaseCreating {
		guard.Allowed = false
		guard.Reasons = append(guard.Reasons, "发布单状态不允许放量")
		return guard
	}
	if err := s.ValidateStepSequence(releaseID); err != nil {
		guard.Allowed = false
		guard.Reasons = append(guard.Reasons, err.Error())
	}
	next, err := s.NextStep(releaseID)
	if err != nil {
		guard.Allowed = false
		guard.Reasons = append(guard.Reasons, err.Error())
	} else {
		guard.CurrentStep = next
	}
	return guard
}
