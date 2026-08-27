package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateReleaseStep(st *model.ReleaseStep) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.releaseSteps {
		if exist.ReleaseID == st.ReleaseID && exist.StepNo == st.StepNo {
			return ErrConflict
		}
	}
	s.releaseSteps[st.ID] = st
	return nil
}

func (s *MemoryStore) GetReleaseStep(id string) (*model.ReleaseStep, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	st, ok := s.releaseSteps[id]
	if !ok {
		return nil, ErrNotFound
	}
	return st, nil
}

func (s *MemoryStore) ListReleaseSteps() []*model.ReleaseStep {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.ReleaseStep, 0, len(s.releaseSteps))
	for _, st := range s.releaseSteps {
		list = append(list, st)
	}
	return list
}

func (s *MemoryStore) UpdateReleaseStep(st *model.ReleaseStep) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.releaseSteps[st.ID]; !ok {
		return ErrNotFound
	}
	s.releaseSteps[st.ID] = st
	return nil
}

func (s *MemoryStore) DeleteReleaseStep(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.releaseSteps[id]; !ok {
		return ErrNotFound
	}
	delete(s.releaseSteps, id)
	return nil
}
