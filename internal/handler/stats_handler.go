package handler

import (
	"net/http"

	"grayrelease/pkg/httpx"
)

func (s *Server) registerStatsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/stats/releases", s.statsReleases)
	mux.HandleFunc("GET /api/stats/active-releases", s.activeReleases)
}

func (s *Server) statsReleases(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.StatsReleases())
}

func (s *Server) activeReleases(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ActiveReleases())
}
