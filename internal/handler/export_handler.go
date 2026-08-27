package handler

import (
	"net/http"

	"grayrelease/internal/service"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerExportRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/export", s.exportSnapshot)
	mux.HandleFunc("POST /api/import", s.importSnapshot)
}

func (s *Server) exportSnapshot(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ExportSnapshot())
}

func (s *Server) importSnapshot(w http.ResponseWriter, r *http.Request) {
	var snap service.Snapshot
	if err := httpx.Decode(r, &snap); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	imported, err := s.svc.ImportSnapshot(snap)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, imported)
}
