package service

import (
	"sort"
	"strconv"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// AdvanceStep 完成一个放量步骤，并记录放量历史。
func (s *Service) AdvanceStep(releaseID, stepID, operator string) (*model.RolloutRecord, error) {
	st, err := s.store.GetReleaseStep(stepID)
	if err != nil {
		return nil, err
	}
	if st.ReleaseID != releaseID {
		return nil, model.NewValidationError("step_id", "步骤不属于该发布单")
	}
	if st.Status == model.StepCompleted {
		return nil, model.NewValidationError("step", "步骤已完成")
	}
	from := s.RolloutProgress(releaseID)
	now := time.Now()
	st.Status = model.StepCompleted
	if st.StartedAt == nil {
		st.StartedAt = &now
	}
	st.FinishedAt = &now
	if err := s.store.UpdateReleaseStep(st); err != nil {
		return nil, err
	}
	rec := &model.RolloutRecord{
		ID:          idgen.Hex(),
		ReleaseID:   releaseID,
		StepID:      stepID,
		FromPercent: from,
		ToPercent:   st.TargetPercent,
		Operator:    operator,
		CreatedAt:   now,
	}
	if err := s.store.CreateRolloutRecord(rec); err != nil {
		return nil, err
	}
	s.emitChangeLog(releaseID, model.ChangeStepAdvanced, "步骤 "+st.Name+" 放量至 "+strconv.Itoa(st.TargetPercent)+"%", operator)
	return rec, nil
}

// RolloutProgress 返回发布单当前放量进度（已完成步骤的最大百分比）。
func (s *Service) RolloutProgress(releaseID string) int {
	progress := 0
	for _, st := range s.store.ListReleaseSteps() {
		if st.ReleaseID == releaseID && st.Status == model.StepCompleted {
			if st.TargetPercent > progress {
				progress = st.TargetPercent
			}
		}
	}
	return progress
}

// ListRolloutRecords 返回某发布单的放量记录（时间倒序）。
func (s *Service) ListRolloutRecords(releaseID string) []*model.RolloutRecord {
	list := make([]*model.RolloutRecord, 0)
	for _, r := range s.store.ListRolloutRecords() {
		if releaseID == "" || r.ReleaseID == releaseID {
			list = append(list, r)
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].CreatedAt.After(list[j].CreatedAt) })
	return list
}
