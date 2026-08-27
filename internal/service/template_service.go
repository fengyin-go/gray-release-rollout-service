package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// CreateTemplate 创建放量模板。
func (s *Service) CreateTemplate(input model.ReleaseTemplate) (*model.ReleaseTemplate, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	t := &model.ReleaseTemplate{
		ID:          idgen.Hex(),
		Name:        input.Name,
		Description: input.Description,
		Schedule:    append([]int{}, input.Schedule...),
		CreatedAt:   time.Now(),
	}
	if err := s.store.CreateReleaseTemplate(t); err != nil {
		return nil, err
	}
	return t, nil
}

// ListTemplates 返回全部模板（按名称排序）。
func (s *Service) ListTemplates() []*model.ReleaseTemplate {
	list := s.store.ListReleaseTemplates()
	sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })
	return list
}

// GetTemplate 按 ID 获取模板。
func (s *Service) GetTemplate(id string) (*model.ReleaseTemplate, error) {
	return s.store.GetReleaseTemplate(id)
}

// UpdateTemplate 更新模板。
func (s *Service) UpdateTemplate(id string, input model.ReleaseTemplate) (*model.ReleaseTemplate, error) {
	existing, err := s.store.GetReleaseTemplate(id)
	if err != nil {
		return nil, err
	}
	existing.Name = input.Name
	existing.Description = input.Description
	existing.Schedule = append([]int{}, input.Schedule...)
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateReleaseTemplate(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

// DeleteTemplate 删除模板。
func (s *Service) DeleteTemplate(id string) error {
	return s.store.DeleteReleaseTemplate(id)
}

// ApplyTemplate 将模板应用到发布单，生成放量步骤。
func (s *Service) ApplyTemplate(templateID, releaseID string) (int, error) {
	t, err := s.store.GetReleaseTemplate(templateID)
	if err != nil {
		return 0, err
	}
	return s.PlanRollout(releaseID, t.Schedule, "")
}
