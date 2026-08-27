package service

import (
	"grayrelease/internal/model"
)

// VersionDiff 两个版本的差异。
type VersionDiff struct {
	From            *model.Version `json:"from"`
	To              *model.Version `json:"to"`
	Same            bool           `json:"same"`
	ArtifactChanged bool           `json:"artifact_changed"`
	ChecksumChanged bool           `json:"checksum_changed"`
	SizeDelta       int64          `json:"size_delta"`
}

// DiffVersions 比较两个版本的制品、校验和与大小差异。
func (s *Service) DiffVersions(fromID, toID string) (*VersionDiff, error) {
	from, err := s.store.GetVersion(fromID)
	if err != nil {
		return nil, err
	}
	to, err := s.store.GetVersion(toID)
	if err != nil {
		return nil, err
	}
	diff := &VersionDiff{
		From:            from,
		To:              to,
		ArtifactChanged: from.ArtifactURL != to.ArtifactURL,
		ChecksumChanged: from.Checksum != to.Checksum,
		SizeDelta:       to.SizeBytes - from.SizeBytes,
	}
	diff.Same = !diff.ArtifactChanged && !diff.ChecksumChanged && diff.SizeDelta == 0
	return diff, nil
}
