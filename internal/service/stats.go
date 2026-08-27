package service

import (
	"sort"

	"grayrelease/internal/model"
)

// ReleaseStats 发布单维度统计。
type ReleaseStats struct {
	Total     int            `json:"total"`
	ByStatus  map[string]int `json:"by_status"`
	TotalVersions int        `json:"total_versions"`
	TotalRules    int        `json:"total_rules"`
}

// StatsReleases 返回发布单与版本、规则的汇总统计。
func (s *Service) StatsReleases() ReleaseStats {
	stats := ReleaseStats{ByStatus: make(map[string]int)}
	for _, r := range s.store.ListReleases() {
		stats.Total++
		stats.ByStatus[r.Status]++
	}
	stats.TotalVersions = len(s.store.ListVersions())
	stats.TotalRules = len(s.store.ListTrafficRules())
	return stats
}

// RolloutOverview 发布单放量概览。
type RolloutOverview struct {
	Release    *model.Release `json:"release"`
	Progress   int            `json:"progress"`
	TotalSteps int            `json:"total_steps"`
	DoneSteps  int            `json:"done_steps"`
	Whitelist  int            `json:"whitelist_count"`
	Rules      int            `json:"rule_count"`
}

// RolloutOverview 返回单个发布单的放量概览。
func (s *Service) RolloutOverview(releaseID string) (*RolloutOverview, error) {
	r, err := s.store.GetRelease(releaseID)
	if err != nil {
		return nil, err
	}
	ov := &RolloutOverview{
		Release:   r,
		Progress:  s.RolloutProgress(releaseID),
		Whitelist: len(s.ListWhitelists(releaseID)),
		Rules:     len(s.ListTrafficRules(releaseID)),
	}
	for _, st := range s.store.ListReleaseSteps() {
		if st.ReleaseID != releaseID {
			continue
		}
		ov.TotalSteps++
		if st.Status == model.StepCompleted {
			ov.DoneSteps++
		}
	}
	return ov, nil
}

// ActiveReleases 返回处于进行中状态的发布单列表。
func (s *Service) ActiveReleases() []*model.Release {
	active := make([]*model.Release, 0)
	for _, r := range s.store.ListReleases() {
		if r.Status == model.ReleaseCreating || r.Status == model.ReleaseReleasing {
			active = append(active, r)
		}
	}
	sort.Slice(active, func(i, j int) bool { return active[i].UpdatedAt.After(active[j].UpdatedAt) })
	return active
}
