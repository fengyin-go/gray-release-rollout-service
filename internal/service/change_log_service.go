package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// emitChangeLog 写入一条发布单变更记录。
func (s *Service) emitChangeLog(releaseID, action, detail, operator string) {
	c := &model.ChangeLog{
		ID:        idgen.Hex(),
		ReleaseID: releaseID,
		Action:    action,
		Detail:    detail,
		Operator:  operator,
		CreatedAt: time.Now(),
	}
	_ = s.store.CreateChangeLog(c)
}

// ListChangeLogs 按筛选条件分页查询变更记录（时间倒序）。
func (s *Service) ListChangeLogs(filter model.ChangeLogFilter, page, size int) ([]*model.ChangeLog, int, error) {
	all := s.store.ListChangeLogs()
	matched := make([]*model.ChangeLog, 0, len(all))
	for _, c := range all {
		if filter.Match(c) {
			matched = append(matched, c)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.ChangeLog{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}
