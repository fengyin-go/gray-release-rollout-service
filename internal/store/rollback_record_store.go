package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateRollbackRecord(r *model.RollbackRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rollbacks[r.ID] = r
	return nil
}

func (s *MemoryStore) ListRollbackRecords() []*model.RollbackRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.RollbackRecord, 0, len(s.rollbacks))
	for _, r := range s.rollbacks {
		list = append(list, r)
	}
	return list
}
