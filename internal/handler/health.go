package handler

import (
	"net/http"

	"grayrelease/pkg/httpx"
)

func (s *Server) registerHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /healthz", s.healthz)
	mux.HandleFunc("GET /readyz", s.readyz)
}

func (s *Server) healthz(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, map[string]string{"status": "ok"})
}

func (s *Server) readyz(w http.ResponseWriter, r *http.Request) {
	stats := s.svc.StatsReleases()
	httpx.OK(w, map[string]interface{}{
		"status":         "ready",
		"total_releases": stats.Total,
	})
}
