package service

// SnapshotValidation 快照校验结果。
type SnapshotValidation struct {
	Valid  bool     `json:"valid"`
	Errors []string `json:"errors"`
}

// ValidateSnapshot 校验快照的引用完整性。
func (s *Service) ValidateSnapshot(snap Snapshot) SnapshotValidation {
	v := SnapshotValidation{Valid: true, Errors: []string{}}
	serviceIDs := make(map[string]bool)
	versionIDs := make(map[string]bool)
	releaseIDs := make(map[string]bool)
	for _, ver := range snap.Versions {
		versionIDs[ver.ID] = true
		serviceIDs[ver.ServiceID] = true
	}
	for _, r := range snap.Releases {
		releaseIDs[r.ID] = true
		if r.VersionID != "" && !versionIDs[r.VersionID] {
			v.Errors = append(v.Errors, "发布单 "+r.Name+" 引用了不存在的版本 "+r.VersionID)
		}
	}
	for _, t := range snap.TrafficRules {
		if !releaseIDs[t.ReleaseID] {
			v.Errors = append(v.Errors, "流量规则 "+t.Name+" 引用了不存在的发布单 "+t.ReleaseID)
		}
	}
	for _, w := range snap.Whitelists {
		if !releaseIDs[w.ReleaseID] {
			v.Errors = append(v.Errors, "白名单引用了不存在的发布单 "+w.ReleaseID)
		}
	}
	for _, st := range snap.ReleaseSteps {
		if !releaseIDs[st.ReleaseID] {
			v.Errors = append(v.Errors, "发布步骤引用了不存在的发布单 "+st.ReleaseID)
		}
	}
	if len(v.Errors) > 0 {
		v.Valid = false
	}
	return v
}
