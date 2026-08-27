package handler

import (
	"net/http"

	"grayrelease/pkg/httpx"
)

func (s *Server) registerRolloutRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/releases/{id}/advance", s.advanceStep)
	mux.HandleFunc("GET /api/releases/{id}/progress", s.rolloutProgress)
	mux.HandleFunc("GET /api/releases/{id}/records", s.rolloutRecords)
	mux.HandleFunc("GET /api/releases/{id}/overview", s.rolloutOverview)
	mux.HandleFunc("GET /api/releases/{id}/guard", s.rolloutGuard)
}

type advanceStepRequest struct {
	StepID   string `json:"step_id"`
	Operator string `json:"operator"`
}

func (s *Server) advanceStep(w http.ResponseWriter, r *http.Request) {
	var req advanceStepRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	rec, err := s.svc.AdvanceStep(r.PathValue("id"), req.StepID, req.Operator)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, rec)
}

func (s *Server) rolloutProgress(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, map[string]int{"progress": s.svc.RolloutProgress(r.PathValue("id"))})
}

func (s *Server) rolloutRecords(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListRolloutRecords(r.PathValue("id")))
}

func (s *Server) rolloutOverview(w http.ResponseWriter, r *http.Request) {
	ov, err := s.svc.RolloutOverview(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, ov)
}

func (s *Server) rolloutGuard(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.CheckRolloutGuard(r.PathValue("id")))
}
