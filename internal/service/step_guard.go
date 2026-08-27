package service

import (
	"grayrelease/internal/model"
)

// ValidateStepSequence 校验发布单步骤百分比是否单调递增。
func (s *Service) ValidateStepSequence(releaseID string) error {
	steps := s.ListReleaseSteps(releaseID)
	for i := 1; i < len(steps); i++ {
		if steps[i].TargetPercent <= steps[i-1].TargetPercent {
			return model.NewValidationError("step", "步骤目标百分比必须严格递增")
		}
	}
	return nil
}

// NextStep 返回发布单下一步可推进的步骤（首个未完成步骤）。
func (s *Service) NextStep(releaseID string) (*model.ReleaseStep, error) {
	steps := s.ListReleaseSteps(releaseID)
	for _, st := range steps {
		if st.Status != model.StepCompleted {
			return st, nil
		}
	}
	return nil, model.NewValidationError("step", "无待推进步骤")
}

// CompleteReleaseCheck 判断发布单是否已放量到 100% 且可标记完成。
func (s *Service) CompleteReleaseCheck(releaseID string) bool {
	return s.RolloutProgress(releaseID) >= 100
}

// HasPendingSteps 判断发布单是否还有未完成步骤。
func (s *Service) HasPendingSteps(releaseID string) bool {
	for _, st := range s.store.ListReleaseSteps() {
		if st.ReleaseID == releaseID && st.Status != model.StepCompleted {
			return true
		}
	}
	return false
}
