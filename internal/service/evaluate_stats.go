package service

// BatchEvaluateResult 批量灰度判定统计。
type BatchEvaluateResult struct {
	Total   int     `json:"total"`
	InGray  int     `json:"in_gray"`
	HitRate float64 `json:"hit_rate"`
}

// BatchEvaluate 对一批用户做灰度判定并统计命中率。
func (s *Service) BatchEvaluate(releaseID string, userIDs []string) BatchEvaluateResult {
	res := BatchEvaluateResult{Total: len(userIDs)}
	for _, uid := range userIDs {
		if s.EvaluateRequest(releaseID, TrafficRequest{UserID: uid}).InGray {
			res.InGray++
		}
	}
	if res.Total > 0 {
		res.HitRate = float64(res.InGray) / float64(res.Total)
	}
	return res
}
