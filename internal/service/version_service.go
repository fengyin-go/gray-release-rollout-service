package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
	"grayrelease/pkg/semver"
)

// CreateVersion 创建服务版本。
func (s *Service) CreateVersion(input model.Version) (*model.Version, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	v := &model.Version{
		ID:          idgen.Hex(),
		ServiceID:   input.ServiceID,
		Version:     input.Version,
		ArtifactURL: input.ArtifactURL,
		Checksum:    input.Checksum,
		SizeBytes:   input.SizeBytes,
		Description: input.Description,
		CreatedAt:   time.Now(),
	}
	if err := s.store.CreateVersion(v); err != nil {
		return nil, err
	}
	return v, nil
}

// GetVersion 按 ID 获取版本。
func (s *Service) GetVersion(id string) (*model.Version, error) {
	return s.store.GetVersion(id)
}

// ListVersions 按服务 ID 筛选版本（按 semver 降序）。
func (s *Service) ListVersions(serviceID string) []*model.Version {
	all := s.store.ListVersions()
	list := make([]*model.Version, 0, len(all))
	for _, v := range all {
		if serviceID == "" || v.ServiceID == serviceID {
			list = append(list, v)
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return semver.Compare(list[i].Version, list[j].Version) > 0
	})
	return list
}

// LatestVersion 返回某服务的最新版本。
func (s *Service) LatestVersion(serviceID string) (*model.Version, error) {
	versions := s.ListVersions(serviceID)
	if len(versions) == 0 {
		return nil, model.NewValidationError("service_id", "服务无版本记录")
	}
	return versions[0], nil
}

// UpdateVersion 更新版本可编辑字段。
func (s *Service) UpdateVersion(id string, input model.Version) (*model.Version, error) {
	existing, err := s.store.GetVersion(id)
	if err != nil {
		return nil, err
	}
	existing.ArtifactURL = input.ArtifactURL
	existing.Checksum = input.Checksum
	existing.SizeBytes = input.SizeBytes
	existing.Description = input.Description
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateVersion(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

// DeleteVersion 删除版本，若被发布单引用则拒绝。
func (s *Service) DeleteVersion(id string) error {
	if _, err := s.store.GetVersion(id); err != nil {
		return err
	}
	for _, r := range s.store.ListReleases() {
		if r.VersionID == id {
			return model.NewValidationError("version", "版本已被发布单引用，无法删除")
		}
	}
	return s.store.DeleteVersion(id)
}
