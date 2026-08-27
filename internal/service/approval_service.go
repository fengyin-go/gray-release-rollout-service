package service

import (
	"sort"
	"time"

	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// CreateApproval 为发布单发起审批。
func (s *Service) CreateApproval(releaseID, approver string) (*model.ApprovalRecord, error) {
	if _, err := s.store.GetRelease(releaseID); err != nil {
		return nil, model.NewValidationError("release_id", "发布单不存在")
	}
	a := &model.ApprovalRecord{
		ID:        idgen.Hex(),
		ReleaseID: releaseID,
		Approver:  approver,
		Status:    model.ApprovalPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := a.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.CreateApprovalRecord(a); err != nil {
		return nil, err
	}
	return a, nil
}

// Approve 审批通过。
func (s *Service) Approve(id, comment string) (*model.ApprovalRecord, error) {
	return s.setApprovalStatus(id, model.ApprovalApproved, comment)
}

// Reject 审批拒绝。
func (s *Service) Reject(id, comment string) (*model.ApprovalRecord, error) {
	return s.setApprovalStatus(id, model.ApprovalRejected, comment)
}

func (s *Service) setApprovalStatus(id, status, comment string) (*model.ApprovalRecord, error) {
	a, err := s.store.GetApprovalRecord(id)
	if err != nil {
		return nil, err
	}
	if a.Status != model.ApprovalPending {
		return nil, model.NewValidationError("approval", "审批已处理，不能重复操作")
	}
	a.Status = status
	a.Comment = comment
	a.UpdatedAt = time.Now()
	if err := s.store.UpdateApprovalRecord(a); err != nil {
		return nil, err
	}
	s.emitChangeLog(a.ReleaseID, model.ChangeStatusChanged, "审批 "+status+"，审批人 "+a.Approver, "")
	return a, nil
}

// ListApprovals 返回某发布单的审批记录。
func (s *Service) ListApprovals(releaseID string) []*model.ApprovalRecord {
	list := make([]*model.ApprovalRecord, 0)
	for _, a := range s.store.ListApprovalRecords() {
		if releaseID == "" || a.ReleaseID == releaseID {
			list = append(list, a)
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].CreatedAt.Before(list[j].CreatedAt) })
	return list
}

// ApprovalSummary 审批汇总。
type ApprovalSummary struct {
	Total    int `json:"total"`
	Pending  int `json:"pending"`
	Approved int `json:"approved"`
	Rejected int `json:"rejected"`
}

// ApprovalSummary 返回某发布单的审批汇总。
func (s *Service) ApprovalSummary(releaseID string) ApprovalSummary {
	sum := ApprovalSummary{}
	for _, a := range s.store.ListApprovalRecords() {
		if releaseID != "" && a.ReleaseID != releaseID {
			continue
		}
		sum.Total++
		switch a.Status {
		case model.ApprovalPending:
			sum.Pending++
		case model.ApprovalApproved:
			sum.Approved++
		case model.ApprovalRejected:
			sum.Rejected++
		}
	}
	return sum
}
