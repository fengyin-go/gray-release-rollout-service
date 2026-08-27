package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
)

// RuleTypeDistribution 返回某发布单流量规则的类型分布。
func (s *Service) RuleTypeDistribution(releaseID string) map[string]int {
	dist := make(map[string]int)
	for _, t := range s.store.ListTrafficRules() {
		if releaseID == "" || t.ReleaseID == releaseID {
			dist[t.Type]++
		}
	}
	return dist
}

// TimelinePoint 放量时间线节点。
type TimelinePoint struct {
	Percent   int       `json:"percent"`
	ReachedAt time.Time `json:"reached_at"`
}

// RolloutTimeline 返回发布单放量时间线（按时间升序）。
func (s *Service) RolloutTimeline(releaseID string) []TimelinePoint {
	records := s.ListRolloutRecords(releaseID)
	points := make([]TimelinePoint, 0, len(records))
	for _, r := range records {
		points = append(points, TimelinePoint{Percent: r.ToPercent, ReachedAt: r.CreatedAt})
	}
	sort.Slice(points, func(i, j int) bool { return points[i].ReachedAt.Before(points[j].ReachedAt) })
	return points
}

// ServiceReleaseCounts 返回每个服务的发布单数量。
func (s *Service) ServiceReleaseCounts() map[string]int {
	counts := make(map[string]int)
	for _, r := range s.store.ListReleases() {
		counts[r.ServiceID]++
	}
	return counts
}

// StatusTimeline 返回某发布单状态变更时间线（从变更记录）。
type StatusPoint struct {
	Status   string    `json:"status"`
	ChangedAt time.Time `json:"changed_at"`
}

// StatusTimeline 返回某发布单的状态变更时间线。
func (s *Service) StatusTimeline(releaseID string) []StatusPoint {
	logs, _, _ := s.ListChangeLogs(model.ChangeLogFilter{ReleaseID: releaseID, Action: model.ChangeStatusChanged}, 1, 1000)
	points := make([]StatusPoint, 0, len(logs))
	for _, c := range logs {
		points = append(points, StatusPoint{Status: c.Detail, ChangedAt: c.CreatedAt})
	}
	sort.Slice(points, func(i, j int) bool { return points[i].ChangedAt.Before(points[j].ChangedAt) })
	return points
}
