package service

import (
	"grayrelease/internal/model"
)

// ReleaseCompare 两个发布单的对比结果。
type ReleaseCompare struct {
	Left            *model.Release `json:"left"`
	Right           *model.Release `json:"right"`
	LeftRules       int            `json:"left_rules"`
	RightRules      int            `json:"right_rules"`
	LeftWhitelist   int            `json:"left_whitelist"`
	RightWhitelist  int            `json:"right_whitelist"`
	LeftProgress    int            `json:"left_progress"`
	RightProgress   int            `json:"right_progress"`
	ProgressDelta   int            `json:"progress_delta"`
}

// CompareReleases 对比两个发布单的规则、白名单与进度。
func (s *Service) CompareReleases(leftID, rightID string) (*ReleaseCompare, error) {
	left, err := s.store.GetRelease(leftID)
	if err != nil {
		return nil, err
	}
	right, err := s.store.GetRelease(rightID)
	if err != nil {
		return nil, err
	}
	c := &ReleaseCompare{
		Left:           left,
		Right:          right,
		LeftRules:      len(s.ListTrafficRules(leftID)),
		RightRules:     len(s.ListTrafficRules(rightID)),
		LeftWhitelist:  len(s.ListWhitelists(leftID)),
		RightWhitelist: len(s.ListWhitelists(rightID)),
		LeftProgress:   s.RolloutProgress(leftID),
		RightProgress:  s.RolloutProgress(rightID),
	}
	c.ProgressDelta = c.RightProgress - c.LeftProgress
	return c, nil
}
