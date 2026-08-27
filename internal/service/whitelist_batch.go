package service

import (
	"grayrelease/internal/model"
)

// BatchAddWhitelist 批量添加白名单用户，返回成功数量。
func (s *Service) BatchAddWhitelist(releaseID string, userIDs []string, note string) (int, error) {
	if _, err := s.store.GetRelease(releaseID); err != nil {
		return 0, model.NewValidationError("release_id", "发布单不存在")
	}
	added := 0
	for _, uid := range userIDs {
		if _, err := s.AddWhitelist(releaseID, uid, note); err == nil {
			added++
		}
	}
	return added, nil
}

// BatchRemoveWhitelist 批量移除白名单用户，返回成功数量。
func (s *Service) BatchRemoveWhitelist(releaseID string, userIDs []string) int {
	removed := 0
	for _, uid := range userIDs {
		if err := s.RemoveWhitelist(releaseID, uid); err == nil {
			removed++
		}
	}
	return removed
}

// WhitelistUsers 返回某发布单白名单的用户 ID 列表。
func (s *Service) WhitelistUsers(releaseID string) []string {
	users := make([]string, 0)
	for _, w := range s.store.ListWhitelists() {
		if w.ReleaseID == releaseID {
			users = append(users, w.UserID)
		}
	}
	return users
}
