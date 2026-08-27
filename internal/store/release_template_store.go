package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateReleaseTemplate(t *model.ReleaseTemplate) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.templates {
		if exist.Name == t.Name {
			return ErrConflict
		}
	}
	s.templates[t.ID] = t
	return nil
}

func (s *MemoryStore) GetReleaseTemplate(id string) (*model.ReleaseTemplate, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.templates[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (s *MemoryStore) ListReleaseTemplates() []*model.ReleaseTemplate {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.ReleaseTemplate, 0, len(s.templates))
	for _, t := range s.templates {
		list = append(list, t)
	}
	return list
}

func (s *MemoryStore) UpdateReleaseTemplate(t *model.ReleaseTemplate) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.templates[t.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.templates {
		if exist.ID != t.ID && exist.Name == t.Name {
			return ErrConflict
		}
	}
	s.templates[t.ID] = t
	return nil
}

func (s *MemoryStore) DeleteReleaseTemplate(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.templates[id]; !ok {
		return ErrNotFound
	}
	delete(s.templates, id)
	return nil
}
