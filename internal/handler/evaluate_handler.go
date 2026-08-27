package handler

import (
	"net/http"

	"grayrelease/internal/service"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerEvaluateRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/releases/{id}/evaluate", s.evaluate)
}

type evaluateRequest struct {
	UserID  string            `json:"user_id"`
	Headers map[string]string `json:"headers"`
	Cookies map[string]string `json:"cookies"`
}

func (s *Server) evaluate(w http.ResponseWriter, r *http.Request) {
	var req evaluateRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	result := s.svc.EvaluateRequest(r.PathValue("id"), service.TrafficRequest{
		UserID:  req.UserID,
		Headers: req.Headers,
		Cookies: req.Cookies,
	})
	httpx.OK(w, result)
}
