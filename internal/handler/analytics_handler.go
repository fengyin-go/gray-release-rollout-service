package handler

import (
	"net/http"
	"strconv"

	"grayrelease/pkg/httpx"
)

func (s *Server) registerAnalyticsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/releases/{id}/plan", s.planRollout)
	mux.HandleFunc("POST /api/releases/{id}/plan-strategy", s.planByStrategy)
	mux.HandleFunc("POST /api/releases/{id}/batch-evaluate", s.batchEvaluate)
	mux.HandleFunc("GET /api/releases/{id}/timeline", s.rolloutTimeline)
	mux.HandleFunc("GET /api/releases/{id}/metrics", s.releaseMetrics)
	mux.HandleFunc("GET /api/releases/{id}/next-step", s.nextStep)
	mux.HandleFunc("GET /api/releases/{id}/impact", s.impactAnalysis)
	mux.HandleFunc("GET /api/releases/{id}/progress-report", s.progressReport)
	mux.HandleFunc("GET /api/releases/{id}/eta", s.releaseETA)
	mux.HandleFunc("GET /api/strategies", s.listStrategies)
	mux.HandleFunc("GET /api/reports/rule-types", s.ruleTypeDistribution)
	mux.HandleFunc("GET /api/reports/service-counts", s.serviceReleaseCounts)
	mux.HandleFunc("GET /api/versions/diff", s.diffVersions)
	mux.HandleFunc("GET /api/versions/history", s.versionHistory)
	mux.HandleFunc("GET /api/versions/latest-per-service", s.latestVersionPerService)
	mux.HandleFunc("GET /api/releases/history", s.releaseHistory)
	mux.HandleFunc("GET /api/releases/compare", s.compareReleases)
	mux.HandleFunc("GET /api/stats/rollout", s.rolloutStats)
	mux.HandleFunc("GET /api/stats/most-active", s.mostActiveReleases)
	mux.HandleFunc("POST /api/templates/seed", s.seedPresets)
	mux.HandleFunc("GET /api/search/releases", s.searchReleases)
	mux.HandleFunc("GET /api/search/versions", s.searchVersions)
}

type planRolloutRequest struct {
	Schedule []int  `json:"schedule"`
	Prefix   string `json:"prefix"`
}

func (s *Server) planRollout(w http.ResponseWriter, r *http.Request) {
	var req planRolloutRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.PlanRollout(r.PathValue("id"), req.Schedule, req.Prefix)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, map[string]int{"created": n})
}

type planStrategyRequest struct {
	Strategy string `json:"strategy"`
}

func (s *Server) planByStrategy(w http.ResponseWriter, r *http.Request) {
	var req planStrategyRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.PlanByStrategy(r.PathValue("id"), req.Strategy)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, map[string]int{"created": n})
}

type batchEvaluateRequest struct {
	UserIDs []string `json:"user_ids"`
}

func (s *Server) batchEvaluate(w http.ResponseWriter, r *http.Request) {
	var req batchEvaluateRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	httpx.OK(w, s.svc.BatchEvaluate(r.PathValue("id"), req.UserIDs))
}

func (s *Server) rolloutTimeline(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.RolloutTimeline(r.PathValue("id")))
}

func (s *Server) releaseMetrics(w http.ResponseWriter, r *http.Request) {
	m, err := s.svc.ReleaseMetrics(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, m)
}

func (s *Server) nextStep(w http.ResponseWriter, r *http.Request) {
	st, err := s.svc.NextStep(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, st)
}

func (s *Server) listStrategies(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListStrategies())
}

func (s *Server) ruleTypeDistribution(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.RuleTypeDistribution(r.URL.Query().Get("release_id")))
}

func (s *Server) serviceReleaseCounts(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ServiceReleaseCounts())
}

func (s *Server) diffVersions(w http.ResponseWriter, r *http.Request) {
	diff, err := s.svc.DiffVersions(r.URL.Query().Get("from"), r.URL.Query().Get("to"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, diff)
}

func (s *Server) searchReleases(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.SearchReleases(r.URL.Query().Get("q")))
}

func (s *Server) searchVersions(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.SearchVersions(r.URL.Query().Get("q")))
}

func (s *Server) impactAnalysis(w http.ResponseWriter, r *http.Request) {
	analysis, err := s.svc.AnalyzeImpact(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, analysis)
}

func (s *Server) progressReport(w http.ResponseWriter, r *http.Request) {
	report, err := s.svc.ProgressReport(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, report)
}

func (s *Server) releaseETA(w http.ResponseWriter, r *http.Request) {
	eta, err := s.svc.ReleaseETA(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, eta)
}

func (s *Server) versionHistory(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.VersionHistory(r.URL.Query().Get("service_id")))
}

func (s *Server) latestVersionPerService(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.LatestVersionPerService())
}

func (s *Server) releaseHistory(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ReleaseHistory(r.URL.Query().Get("service_id")))
}

func (s *Server) compareReleases(w http.ResponseWriter, r *http.Request) {
	c, err := s.svc.CompareReleases(r.URL.Query().Get("left"), r.URL.Query().Get("right"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, c)
}

func (s *Server) rolloutStats(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.RolloutStats())
}

func (s *Server) mostActiveReleases(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	httpx.OK(w, s.svc.MostActiveReleases(n))
}

func (s *Server) seedPresets(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, map[string]int{"created": s.svc.SeedPresetTemplates()})
}
