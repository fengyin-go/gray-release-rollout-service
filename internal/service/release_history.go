package service

import (
	"sort"

	"grayrelease/internal/model"
	"grayrelease/pkg/semver"
)

// VersionHistory 返回某服务的全部版本历史（semver 降序）。
func (s *Service) VersionHistory(serviceID string) []*model.Version {
	return s.ListVersions(serviceID)
}

// ReleaseHistory 返回某服务的全部发布单（时间倒序）。
func (s *Service) ReleaseHistory(serviceID string) []*model.Release {
	list := make([]*model.Release, 0)
	for _, r := range s.store.ListReleases() {
		if serviceID == "" || r.ServiceID == serviceID {
			list = append(list, r)
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].CreatedAt.After(list[j].CreatedAt) })
	return list
}

// RecentChanges 返回最近的 N 条变更记录。
func (s *Service) RecentChanges(n int) []*model.ChangeLog {
	if n <= 0 {
		n = 20
	}
	logs, _, _ := s.ListChangeLogs(model.ChangeLogFilter{}, 1, n)
	return logs
}

// LatestVersionPerService 返回每个服务的最新版本。
func (s *Service) LatestVersionPerService() map[string]*model.Version {
	result := make(map[string]*model.Version)
	for _, v := range s.store.ListVersions() {
		cur, ok := result[v.ServiceID]
		if !ok || semver.Greater(v.Version, cur.Version) {
			result[v.ServiceID] = v
		}
	}
	return result
}
