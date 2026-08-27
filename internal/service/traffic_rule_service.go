package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// CreateTrafficRule 创建流量规则。
func (s *Service) CreateTrafficRule(input model.TrafficRule) (*model.TrafficRule, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetRelease(input.ReleaseID); err != nil {
		return nil, model.NewValidationError("release_id", "发布单不存在")
	}
	t := &model.TrafficRule{
		ID:         idgen.Hex(),
		ReleaseID:  input.ReleaseID,
		Name:       input.Name,
		Type:       input.Type,
		MatchKey:   input.MatchKey,
		MatchValue: input.MatchValue,
		Percentage: input.Percentage,
		Priority:   input.Priority,
		Enabled:    true,
		CreatedAt:  time.Now(),
	}
	if err := s.store.CreateTrafficRule(t); err != nil {
		return nil, err
	}
	s.emitChangeLog(input.ReleaseID, model.ChangeRuleChanged, "新增流量规则 "+t.Name, "")
	return t, nil
}

// GetTrafficRule 按 ID 获取流量规则。
func (s *Service) GetTrafficRule(id string) (*model.TrafficRule, error) {
	return s.store.GetTrafficRule(id)
}

// ListTrafficRules 按发布单筛选流量规则（按优先级升序）。
func (s *Service) ListTrafficRules(releaseID string) []*model.TrafficRule {
	all := s.store.ListTrafficRules()
	list := make([]*model.TrafficRule, 0, len(all))
	for _, t := range all {
		if releaseID == "" || t.ReleaseID == releaseID {
			list = append(list, t)
		}
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].Priority != list[j].Priority {
			return list[i].Priority < list[j].Priority
		}
		return list[i].CreatedAt.Before(list[j].CreatedAt)
	})
	return list
}

// UpdateTrafficRule 更新流量规则。
func (s *Service) UpdateTrafficRule(id string, input model.TrafficRule) (*model.TrafficRule, error) {
	existing, err := s.store.GetTrafficRule(id)
	if err != nil {
		return nil, err
	}
	existing.Name = input.Name
	existing.Type = input.Type
	existing.MatchKey = input.MatchKey
	existing.MatchValue = input.MatchValue
	existing.Percentage = input.Percentage
	existing.Priority = input.Priority
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateTrafficRule(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

// ToggleTrafficRule 启用或停用流量规则。
func (s *Service) ToggleTrafficRule(id string, enabled bool) (*model.TrafficRule, error) {
	t, err := s.store.GetTrafficRule(id)
	if err != nil {
		return nil, err
	}
	t.Enabled = enabled
	if err := s.store.UpdateTrafficRule(t); err != nil {
		return nil, err
	}
	return t, nil
}

// DeleteTrafficRule 删除流量规则。
func (s *Service) DeleteTrafficRule(id string) error {
	return s.store.DeleteTrafficRule(id)
}
