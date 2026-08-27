package service

import (
	"sort"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// AddReleaseSteps 为发布单批量创建放量步骤。
func (s *Service) AddReleaseSteps(releaseID string, steps []model.ReleaseStep) (int, error) {
	if _, err := s.store.GetRelease(releaseID); err != nil {
		return 0, model.NewValidationError("release_id", "发布单不存在")
	}
	created := 0
	for _, st := range steps {
		st.ReleaseID = releaseID
		if err := st.Validate(); err != nil {
			continue
		}
		st.ID = idgen.Hex()
		st.Status = model.StepPending
		if err := s.store.CreateReleaseStep(&st); err == nil {
			created++
		}
	}
	return created, nil
}

// ListReleaseSteps 返回某发布单的步骤（按序号升序）。
func (s *Service) ListReleaseSteps(releaseID string) []*model.ReleaseStep {
	list := make([]*model.ReleaseStep, 0)
	for _, st := range s.store.ListReleaseSteps() {
		if st.ReleaseID == releaseID {
			list = append(list, st)
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].StepNo < list[j].StepNo })
	return list
}

// DeleteReleaseStep 删除步骤。
func (s *Service) DeleteReleaseStep(id string) error {
	return s.store.DeleteReleaseStep(id)
}
