package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateReleaseNote(n *model.ReleaseNote) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.releaseNotes[n.ID] = n
	return nil
}

func (s *MemoryStore) GetReleaseNote(id string) (*model.ReleaseNote, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n, ok := s.releaseNotes[id]
	if !ok {
		return nil, ErrNotFound
	}
	return n, nil
}

func (s *MemoryStore) ListReleaseNotes() []*model.ReleaseNote {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.ReleaseNote, 0, len(s.releaseNotes))
	for _, n := range s.releaseNotes {
		list = append(list, n)
	}
	return list
}

func (s *MemoryStore) UpdateReleaseNote(n *model.ReleaseNote) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.releaseNotes[n.ID]; !ok {
		return ErrNotFound
	}
	s.releaseNotes[n.ID] = n
	return nil
}

func (s *MemoryStore) DeleteReleaseNote(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.releaseNotes[id]; !ok {
		return ErrNotFound
	}
	delete(s.releaseNotes, id)
	return nil
}
