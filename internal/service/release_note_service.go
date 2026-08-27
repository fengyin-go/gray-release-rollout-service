package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// CreateReleaseNote 创建发布说明。
func (s *Service) CreateReleaseNote(input model.ReleaseNote) (*model.ReleaseNote, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetRelease(input.ReleaseID); err != nil {
		return nil, model.NewValidationError("release_id", "发布单不存在")
	}
	n := &model.ReleaseNote{
		ID:        idgen.Hex(),
		ReleaseID: input.ReleaseID,
		Content:   input.Content,
		Version:   input.Version,
		CreatedAt: time.Now(),
	}
	if err := s.store.CreateReleaseNote(n); err != nil {
		return nil, err
	}
	return n, nil
}

// GetReleaseNote 按 ID 获取发布说明。
func (s *Service) GetReleaseNote(id string) (*model.ReleaseNote, error) {
	return s.store.GetReleaseNote(id)
}

// ListReleaseNotes 返回某发布单的发布说明（时间倒序）。
func (s *Service) ListReleaseNotes(releaseID string) []*model.ReleaseNote {
	list := make([]*model.ReleaseNote, 0)
	for _, n := range s.store.ListReleaseNotes() {
		if releaseID == "" || n.ReleaseID == releaseID {
			list = append(list, n)
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].CreatedAt.After(list[j].CreatedAt) })
	return list
}

// UpdateReleaseNote 更新发布说明。
func (s *Service) UpdateReleaseNote(id string, input model.ReleaseNote) (*model.ReleaseNote, error) {
	existing, err := s.store.GetReleaseNote(id)
	if err != nil {
		return nil, err
	}
	existing.Content = input.Content
	existing.Version = input.Version
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateReleaseNote(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

// DeleteReleaseNote 删除发布说明。
func (s *Service) DeleteReleaseNote(id string) error {
	return s.store.DeleteReleaseNote(id)
}
