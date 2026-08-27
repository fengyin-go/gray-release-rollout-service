package service

import (
	"strconv"

	"grayrelease/internal/model"
)

// PlanRollout 根据目标百分比序列自动生成放量步骤。
func (s *Service) PlanRollout(releaseID string, schedule []int, prefix string) (int, error) {
	if _, err := s.store.GetRelease(releaseID); err != nil {
		return 0, model.NewValidationError("release_id", "发布单不存在")
	}
	if len(schedule) == 0 {
		return 0, model.NewValidationError("schedule", "放量百分比序列不能为空")
	}
	for i := 1; i < len(schedule); i++ {
		if schedule[i] <= schedule[i-1] {
			return 0, model.NewValidationError("schedule", "放量百分比必须严格递增")
		}
	}
	if schedule[len(schedule)-1] > 100 {
		return 0, model.NewValidationError("schedule", "放量百分比不能超过 100")
	}
	steps := make([]model.ReleaseStep, 0, len(schedule))
	for i, pct := range schedule {
		steps = append(steps, model.ReleaseStep{
			StepNo:        i + 1,
			Name:          stepName(prefix, i+1, pct),
			TargetPercent: pct,
		})
	}
	return s.AddReleaseSteps(releaseID, steps)
}

func stepName(prefix string, no, pct int) string {
	if prefix == "" {
		prefix = "步骤"
	}
	return prefix + " " + strconv.Itoa(no) + "（" + strconv.Itoa(pct) + "%）"
}
