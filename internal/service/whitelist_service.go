package service

import (
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// AddWhitelist 添加白名单用户。
func (s *Service) AddWhitelist(releaseID, userID, note string) (*model.Whitelist, error) {
	if _, err := s.store.GetRelease(releaseID); err != nil {
		return nil, model.NewValidationError("release_id", "发布单不存在")
	}
	w := &model.Whitelist{
		ID:        idgen.Hex(),
		ReleaseID: releaseID,
		UserID:    userID,
		Note:      note,
		CreatedAt: time.Now(),
	}
	if err := w.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.CreateWhitelist(w); err != nil {
		return nil, err
	}
	return w, nil
}

// RemoveWhitelist 移除白名单用户。
func (s *Service) RemoveWhitelist(releaseID, userID string) error {
	for _, w := range s.store.ListWhitelists() {
		if w.ReleaseID == releaseID && w.UserID == userID {
			return s.store.DeleteWhitelist(w.ID)
		}
	}
	return model.NewValidationError("whitelist", "白名单记录不存在")
}

// ListWhitelists 返回某发布单的白名单。
func (s *Service) ListWhitelists(releaseID string) []*model.Whitelist {
	list := make([]*model.Whitelist, 0)
	for _, w := range s.store.ListWhitelists() {
		if releaseID == "" || w.ReleaseID == releaseID {
			list = append(list, w)
		}
	}
	return list
}

// IsWhitelisted 判断用户是否在白名单中。
func (s *Service) IsWhitelisted(releaseID, userID string) bool {
	for _, w := range s.store.ListWhitelists() {
		if w.ReleaseID == releaseID && w.UserID == userID {
			return true
		}
	}
	return false
}
