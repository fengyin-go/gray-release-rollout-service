package service

import (
	"hash/fnv"

	"grayrelease/internal/model"
)

// TrafficRequest 灰度判定请求上下文。
type TrafficRequest struct {
	UserID  string
	Headers map[string]string
	Cookies map[string]string
}

// EvaluateResult 灰度判定结果。
type EvaluateResult struct {
	InGray   bool   `json:"in_gray"`
	RuleID   string `json:"rule_id,omitempty"`
	RuleName string `json:"rule_name,omitempty"`
	Reason   string `json:"reason"`
}

// EvaluateRequest 判断请求是否命中灰度：白名单优先，其次按优先级匹配流量规则。
func (s *Service) EvaluateRequest(releaseID string, req TrafficRequest) EvaluateResult {
	if req.UserID != "" && s.IsWhitelisted(releaseID, req.UserID) {
		return EvaluateResult{InGray: true, Reason: "whitelist"}
	}
	for _, r := range s.ListTrafficRules(releaseID) {
		if !r.Enabled {
			continue
		}
		if matchRule(r, req) {
			return EvaluateResult{InGray: true, RuleID: r.ID, RuleName: r.Name, Reason: r.Type}
		}
	}
	return EvaluateResult{InGray: false, Reason: "no_rule_matched"}
}

// matchRule 判断单条规则是否命中请求。
func matchRule(r *model.TrafficRule, req TrafficRequest) bool {
	switch r.Type {
	case model.RuleTypeUser:
		return req.UserID != "" && req.UserID == r.MatchValue
	case model.RuleTypeHeader:
		v, ok := req.Headers[r.MatchKey]
		return ok && v == r.MatchValue
	case model.RuleTypeCookie:
		v, ok := req.Cookies[r.MatchKey]
		return ok && v == r.MatchValue
	case model.RuleTypePercentage:
		if req.UserID == "" {
			return false
		}
		return bucket(req.UserID) < r.Percentage
	default:
		return false
	}
}

// bucket 将用户 ID 哈希到 0-99 的稳定桶。
func bucket(userID string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(userID))
	return int(h.Sum32() % 100)
}
