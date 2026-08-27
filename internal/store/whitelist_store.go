package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateWhitelist(w *model.Whitelist) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.whitelists {
		if exist.ReleaseID == w.ReleaseID && exist.UserID == w.UserID {
			return ErrConflict
		}
	}
	s.whitelists[w.ID] = w
	return nil
}

func (s *MemoryStore) GetWhitelist(id string) (*model.Whitelist, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	w, ok := s.whitelists[id]
	if !ok {
		return nil, ErrNotFound
	}
	return w, nil
}

func (s *MemoryStore) ListWhitelists() []*model.Whitelist {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Whitelist, 0, len(s.whitelists))
	for _, w := range s.whitelists {
		list = append(list, w)
	}
	return list
}

func (s *MemoryStore) DeleteWhitelist(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.whitelists[id]; !ok {
		return ErrNotFound
	}
	delete(s.whitelists, id)
	return nil
}
