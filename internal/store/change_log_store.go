package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateChangeLog(c *model.ChangeLog) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.changeLogs[c.ID] = c
	return nil
}

func (s *MemoryStore) ListChangeLogs() []*model.ChangeLog {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.ChangeLog, 0, len(s.changeLogs))
	for _, c := range s.changeLogs {
		list = append(list, c)
	}
	return list
}
