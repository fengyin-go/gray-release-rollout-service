package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateTrafficRule(t *model.TrafficRule) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.trafficRules {
		if exist.ReleaseID == t.ReleaseID && exist.Name == t.Name {
			return ErrConflict
		}
	}
	s.trafficRules[t.ID] = t
	return nil
}

func (s *MemoryStore) GetTrafficRule(id string) (*model.TrafficRule, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.trafficRules[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (s *MemoryStore) ListTrafficRules() []*model.TrafficRule {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.TrafficRule, 0, len(s.trafficRules))
	for _, t := range s.trafficRules {
		list = append(list, t)
	}
	return list
}

func (s *MemoryStore) UpdateTrafficRule(t *model.TrafficRule) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.trafficRules[t.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.trafficRules {
		if exist.ID != t.ID && exist.ReleaseID == t.ReleaseID && exist.Name == t.Name {
			return ErrConflict
		}
	}
	s.trafficRules[t.ID] = t
	return nil
}

func (s *MemoryStore) DeleteTrafficRule(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.trafficRules[id]; !ok {
		return ErrNotFound
	}
	delete(s.trafficRules, id)
	return nil
}
