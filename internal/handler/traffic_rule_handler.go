package handler

import (
	"net/http"

	"grayrelease/internal/model"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerTrafficRuleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/traffic-rules", s.createTrafficRule)
	mux.HandleFunc("GET /api/traffic-rules", s.listTrafficRules)
	mux.HandleFunc("GET /api/traffic-rules/{id}", s.getTrafficRule)
	mux.HandleFunc("PUT /api/traffic-rules/{id}", s.updateTrafficRule)
	mux.HandleFunc("PATCH /api/traffic-rules/{id}/toggle", s.toggleTrafficRule)
	mux.HandleFunc("DELETE /api/traffic-rules/{id}", s.deleteTrafficRule)
}

type createTrafficRuleRequest struct {
	ReleaseID  string `json:"release_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	MatchKey   string `json:"match_key"`
	MatchValue string `json:"match_value"`
	Percentage int    `json:"percentage"`
	Priority   int    `json:"priority"`
}

func (s *Server) createTrafficRule(w http.ResponseWriter, r *http.Request) {
	var req createTrafficRuleRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	t, err := s.svc.CreateTrafficRule(model.TrafficRule{
		ReleaseID:  req.ReleaseID,
		Name:       req.Name,
		Type:       req.Type,
		MatchKey:   req.MatchKey,
		MatchValue: req.MatchValue,
		Percentage: req.Percentage,
		Priority:   req.Priority,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, t)
}

func (s *Server) listTrafficRules(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListTrafficRules(r.URL.Query().Get("release_id")))
}

func (s *Server) getTrafficRule(w http.ResponseWriter, r *http.Request) {
	t, err := s.svc.GetTrafficRule(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, t)
}

func (s *Server) updateTrafficRule(w http.ResponseWriter, r *http.Request) {
	var req createTrafficRuleRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	t, err := s.svc.UpdateTrafficRule(r.PathValue("id"), model.TrafficRule{
		Name:       req.Name,
		Type:       req.Type,
		MatchKey:   req.MatchKey,
		MatchValue: req.MatchValue,
		Percentage: req.Percentage,
		Priority:   req.Priority,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, t)
}

type toggleTrafficRuleRequest struct {
	Enabled bool `json:"enabled"`
}

func (s *Server) toggleTrafficRule(w http.ResponseWriter, r *http.Request) {
	var req toggleTrafficRuleRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	t, err := s.svc.ToggleTrafficRule(r.PathValue("id"), req.Enabled)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, t)
}

func (s *Server) deleteTrafficRule(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteTrafficRule(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
