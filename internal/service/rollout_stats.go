package service

import (
	"sort"

	"grayrelease/internal/model"
)

// RolloutStats 跨发布单的放量统计。
type RolloutStats struct {
	TotalReleases   int            `json:"total_releases"`
	ActiveReleases  int            `json:"active_releases"`
	Completed       int            `json:"completed"`
	RolledBack      int            `json:"rolled_back"`
	RollbackRate    float64        `json:"rollback_rate"`
	AvgProgress     float64        `json:"avg_progress"`
	ByProgress      map[string]int `json:"by_progress"`
	TotalRolloutRecords int        `json:"total_rollout_records"`
}

// RolloutStats 计算跨发布单的放量统计。
func (s *Service) RolloutStats() RolloutStats {
	stats := RolloutStats{ByProgress: make(map[string]int)}
	releases := s.store.ListReleases()
	stats.TotalReleases = len(releases)
	sum := 0
	for _, r := range releases {
		switch r.Status {
		case model.ReleaseCompleted:
			stats.Completed++
		case model.ReleaseRolledBack:
			stats.RolledBack++
		case model.ReleaseCreating, model.ReleaseReleasing:
			stats.ActiveReleases++
		}
		p := s.RolloutProgress(r.ID)
		sum += p
		stats.ByProgress[progressBucket(p)]++
	}
	if stats.TotalReleases > 0 {
		stats.AvgProgress = float64(sum) / float64(stats.TotalReleases)
		stats.RollbackRate = float64(stats.RolledBack) / float64(stats.TotalReleases)
	}
	stats.TotalRolloutRecords = len(s.store.ListRolloutRecords())
	return stats
}

func progressBucket(pct int) string {
	switch {
	case pct <= 0:
		return "0%"
	case pct <= 25:
		return "1-25%"
	case pct <= 50:
		return "26-50%"
	case pct < 100:
		return "51-99%"
	default:
		return "100%"
	}
}

// MostActiveReleases 返回放量记录最多的 N 个发布单。
type ReleaseActivity struct {
	Release *model.Release `json:"release"`
	Records int            `json:"records"`
}

// MostActiveReleases 按放量记录数对发布单排序。
func (s *Service) MostActiveReleases(n int) []ReleaseActivity {
	if n <= 0 {
		n = 10
	}
	counts := make(map[string]int)
	for _, r := range s.store.ListRolloutRecords() {
		counts[r.ReleaseID]++
	}
	list := make([]ReleaseActivity, 0)
	for _, r := range s.store.ListReleases() {
		if counts[r.ID] > 0 {
			list = append(list, ReleaseActivity{Release: r, Records: counts[r.ID]})
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Records > list[j].Records })
	if len(list) > n {
		list = list[:n]
	}
	return list
}
