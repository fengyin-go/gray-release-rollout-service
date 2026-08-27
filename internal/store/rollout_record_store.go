package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateRolloutRecord(r *model.RolloutRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rolloutRecords[r.ID] = r
	return nil
}

func (s *MemoryStore) ListRolloutRecords() []*model.RolloutRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.RolloutRecord, 0, len(s.rolloutRecords))
	for _, r := range s.rolloutRecords {
		list = append(list, r)
	}
	return list
}
