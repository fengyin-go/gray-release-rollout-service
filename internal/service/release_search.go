package service

import (
	"strings"

	"grayrelease/internal/model"
)

// SearchReleases 按关键字搜索发布单（名称 / 服务 / 创建人）。
func (s *Service) SearchReleases(keyword string) []*model.Release {
	kw := strings.ToLower(strings.TrimSpace(keyword))
	if kw == "" {
		return []*model.Release{}
	}
	result := make([]*model.Release, 0)
	for _, r := range s.store.ListReleases() {
		if strings.Contains(strings.ToLower(r.Name), kw) ||
			strings.Contains(strings.ToLower(r.ServiceID), kw) ||
			strings.Contains(strings.ToLower(r.CreatedBy), kw) {
			result = append(result, r)
		}
	}
	return result
}

// SearchVersions 按关键字搜索版本（版本号 / 制品地址）。
func (s *Service) SearchVersions(keyword string) []*model.Version {
	kw := strings.ToLower(strings.TrimSpace(keyword))
	if kw == "" {
		return []*model.Version{}
	}
	result := make([]*model.Version, 0)
	for _, v := range s.store.ListVersions() {
		if strings.Contains(strings.ToLower(v.Version), kw) ||
			strings.Contains(strings.ToLower(v.ArtifactURL), kw) {
			result = append(result, v)
		}
	}
	return result
}
