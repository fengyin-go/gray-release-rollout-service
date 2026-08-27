package handler

import (
	"net/http"

	"grayrelease/internal/model"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerChangeLogRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/change-logs", s.listChangeLogs)
}

func (s *Server) listChangeLogs(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.ChangeLogFilter{
		ReleaseID: r.URL.Query().Get("release_id"),
		Action:    r.URL.Query().Get("action"),
		Operator:  r.URL.Query().Get("operator"),
	}
	items, total, err := s.svc.ListChangeLogs(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}
