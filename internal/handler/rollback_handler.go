package handler

import (
	"net/http"

	"grayrelease/pkg/httpx"
)

func (s *Server) registerRollbackRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/releases/{id}/rollback", s.rollback)
	mux.HandleFunc("GET /api/releases/{id}/rollbacks", s.listRollbacks)
}

type rollbackRequest struct {
	Reason   string `json:"reason"`
	Operator string `json:"operator"`
}

func (s *Server) rollback(w http.ResponseWriter, r *http.Request) {
	var req rollbackRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	rec, err := s.svc.RollbackRelease(r.PathValue("id"), req.Reason, req.Operator)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, rec)
}

func (s *Server) listRollbacks(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListRollbackRecords(r.PathValue("id")))
}
