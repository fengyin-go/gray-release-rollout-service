package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// CreateRelease 创建灰度发布单。
func (s *Service) CreateRelease(input model.Release) (*model.Release, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetVersion(input.VersionID); err != nil {
		return nil, model.NewValidationError("version_id", "版本不存在")
	}
	now := time.Now()
	r := &model.Release{
		ID:          idgen.Hex(),
		Name:        input.Name,
		ServiceID:   input.ServiceID,
		VersionID:   input.VersionID,
		Description: input.Description,
		Status:      model.ReleaseDraft,
		CreatedBy:   input.CreatedBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := s.store.CreateRelease(r); err != nil {
		return nil, err
	}
	s.emitChangeLog(r.ID, model.ChangeCreated, "创建发布单 "+r.Name, input.CreatedBy)
	return r, nil
}

// GetRelease 按 ID 获取发布单。
func (s *Service) GetRelease(id string) (*model.Release, error) {
	return s.store.GetRelease(id)
}

// ListReleases 按筛选条件分页查询发布单。
func (s *Service) ListReleases(filter model.ReleaseFilter, page, size int) ([]*model.Release, int, error) {
	all := s.store.ListReleases()
	matched := make([]*model.Release, 0, len(all))
	for _, r := range all {
		if filter.Match(r) {
			matched = append(matched, r)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Release{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

// UpdateRelease 更新发布单描述等可编辑字段。
func (s *Service) UpdateRelease(id string, input model.Release) (*model.Release, error) {
	existing, err := s.store.GetRelease(id)
	if err != nil {
		return nil, err
	}
	existing.Name = input.Name
	existing.Description = input.Description
	existing.UpdatedAt = time.Now()
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateRelease(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

// ChangeReleaseStatus 按状态机流转发布单状态。
func (s *Service) ChangeReleaseStatus(id, to, operator string) (*model.Release, error) {
	existing, err := s.store.GetRelease(id)
	if err != nil {
		return nil, err
	}
	if !model.CanReleaseTransition(existing.Status, to) {
		return nil, model.NewValidationError("status", "不允许从 "+existing.Status+" 流转到 "+to)
	}
	existing.Status = to
	existing.UpdatedAt = time.Now()
	if err := s.store.UpdateRelease(existing); err != nil {
		return nil, err
	}
	s.emitChangeLog(id, model.ChangeStatusChanged, "状态变更为 "+to, operator)
	return existing, nil
}

// DeleteRelease 删除发布单（仅 draft/cancelled 可删除）。
func (s *Service) DeleteRelease(id string) error {
	existing, err := s.store.GetRelease(id)
	if err != nil {
		return err
	}
	if existing.Status != model.ReleaseDraft && existing.Status != model.ReleaseCancelled {
		return model.NewValidationError("release", "仅草稿或已取消的发布单可删除")
	}
	return s.store.DeleteRelease(id)
}
