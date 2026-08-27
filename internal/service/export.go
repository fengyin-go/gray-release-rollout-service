package service

import (
	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// Snapshot 灰度发布全量快照。
type Snapshot struct {
	Versions       []*model.Version       `json:"versions"`
	Releases       []*model.Release       `json:"releases"`
	TrafficRules   []*model.TrafficRule   `json:"traffic_rules"`
	Whitelists     []*model.Whitelist     `json:"whitelists"`
	ReleaseSteps   []*model.ReleaseStep   `json:"release_steps"`
	RolloutRecords []*model.RolloutRecord `json:"rollout_records"`
	ChangeLogs     []*model.ChangeLog     `json:"change_logs"`
}

// ExportSnapshot 导出全量快照。
func (s *Service) ExportSnapshot() Snapshot {
	return Snapshot{
		Versions:       s.store.ListVersions(),
		Releases:       s.store.ListReleases(),
		TrafficRules:   s.store.ListTrafficRules(),
		Whitelists:     s.store.ListWhitelists(),
		ReleaseSteps:   s.store.ListReleaseSteps(),
		RolloutRecords: s.store.ListRolloutRecords(),
		ChangeLogs:     s.store.ListChangeLogs(),
	}
}

// ImportSnapshot 导入快照（跳过重复项），返回各实体导入数量。
func (s *Service) ImportSnapshot(snap Snapshot) (map[string]int, error) {
	imported := map[string]int{
		"versions":        0,
		"releases":        0,
		"traffic_rules":   0,
		"whitelists":      0,
		"release_steps":   0,
		"rollout_records": 0,
		"change_logs":     0,
	}
	for _, v := range snap.Versions {
		if v.ID == "" {
			v.ID = idgen.Hex()
		}
		if err := s.store.CreateVersion(v); err == nil {
			imported["versions"]++
		}
	}
	for _, r := range snap.Releases {
		if r.ID == "" {
			r.ID = idgen.Hex()
		}
		if err := s.store.CreateRelease(r); err == nil {
			imported["releases"]++
		}
	}
	for _, t := range snap.TrafficRules {
		if t.ID == "" {
			t.ID = idgen.Hex()
		}
		if err := s.store.CreateTrafficRule(t); err == nil {
			imported["traffic_rules"]++
		}
	}
	for _, w := range snap.Whitelists {
		if w.ID == "" {
			w.ID = idgen.Hex()
		}
		if err := s.store.CreateWhitelist(w); err == nil {
			imported["whitelists"]++
		}
	}
	for _, st := range snap.ReleaseSteps {
		if st.ID == "" {
			st.ID = idgen.Hex()
		}
		if err := s.store.CreateReleaseStep(st); err == nil {
			imported["release_steps"]++
		}
	}
	for _, r := range snap.RolloutRecords {
		if r.ID == "" {
			r.ID = idgen.Hex()
		}
		if err := s.store.CreateRolloutRecord(r); err == nil {
			imported["rollout_records"]++
		}
	}
	for _, c := range snap.ChangeLogs {
		if c.ID == "" {
			c.ID = idgen.Hex()
		}
		if err := s.store.CreateChangeLog(c); err == nil {
			imported["change_logs"]++
		}
	}
	return imported, nil
}
