package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// RollbackRelease 回滚发布单，记录回滚历史。
func (s *Service) RollbackRelease(releaseID, reason, operator string) (*model.RollbackRecord, error) {
	r, err := s.store.GetRelease(releaseID)
	if err != nil {
		return nil, err
	}
	if r.Status != model.ReleaseReleasing && r.Status != model.ReleaseCreating {
		return nil, model.NewValidationError("release", "仅进行中的发布单可回滚")
	}
	from := r.Status
	r.Status = model.ReleaseRolledBack
	r.UpdatedAt = time.Now()
	if err := s.store.UpdateRelease(r); err != nil {
		return nil, err
	}
	rec := &model.RollbackRecord{
		ID:         idgen.Hex(),
		ReleaseID:  releaseID,
		FromStatus: from,
		Reason:     reason,
		Operator:   operator,
		CreatedAt:  time.Now(),
	}
	if err := rec.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.CreateRollbackRecord(rec); err != nil {
		return nil, err
	}
	s.emitChangeLog(releaseID, model.ChangeRolledBack, "回滚发布单，原因: "+reason, operator)
	return rec, nil
}

// ListRollbackRecords 返回回滚记录（可按发布单筛选，时间倒序）。
func (s *Service) ListRollbackRecords(releaseID string) []*model.RollbackRecord {
	list := make([]*model.RollbackRecord, 0)
	for _, r := range s.store.ListRollbackRecords() {
		if releaseID == "" || r.ReleaseID == releaseID {
			list = append(list, r)
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].CreatedAt.After(list[j].CreatedAt) })
	return list
}
