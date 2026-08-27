package store

import "grayrelease/internal/model"

func (s *MemoryStore) CreateApprovalRecord(a *model.ApprovalRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.approvals {
		if exist.ReleaseID == a.ReleaseID && exist.Approver == a.Approver {
			return ErrConflict
		}
	}
	s.approvals[a.ID] = a
	return nil
}

func (s *MemoryStore) GetApprovalRecord(id string) (*model.ApprovalRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.approvals[id]
	if !ok {
		return nil, ErrNotFound
	}
	return a, nil
}

func (s *MemoryStore) ListApprovalRecords() []*model.ApprovalRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.ApprovalRecord, 0, len(s.approvals))
	for _, a := range s.approvals {
		list = append(list, a)
	}
	return list
}

func (s *MemoryStore) UpdateApprovalRecord(a *model.ApprovalRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.approvals[a.ID]; !ok {
		return ErrNotFound
	}
	s.approvals[a.ID] = a
	return nil
}
