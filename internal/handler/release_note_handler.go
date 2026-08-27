package handler

import (
	"net/http"

	"grayrelease/internal/model"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerReleaseNoteRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/release-notes", s.createReleaseNote)
	mux.HandleFunc("GET /api/release-notes", s.listReleaseNotes)
	mux.HandleFunc("GET /api/release-notes/{id}", s.getReleaseNote)
	mux.HandleFunc("PUT /api/release-notes/{id}", s.updateReleaseNote)
	mux.HandleFunc("DELETE /api/release-notes/{id}", s.deleteReleaseNote)
}

type createReleaseNoteRequest struct {
	ReleaseID string `json:"release_id"`
	Content   string `json:"content"`
	Version   string `json:"version"`
}

func (s *Server) createReleaseNote(w http.ResponseWriter, r *http.Request) {
	var req createReleaseNoteRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.CreateReleaseNote(model.ReleaseNote{
		ReleaseID: req.ReleaseID,
		Content:   req.Content,
		Version:   req.Version,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, n)
}

func (s *Server) listReleaseNotes(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListReleaseNotes(r.URL.Query().Get("release_id")))
}

func (s *Server) getReleaseNote(w http.ResponseWriter, r *http.Request) {
	n, err := s.svc.GetReleaseNote(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, n)
}

func (s *Server) updateReleaseNote(w http.ResponseWriter, r *http.Request) {
	var req createReleaseNoteRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.UpdateReleaseNote(r.PathValue("id"), model.ReleaseNote{
		Content: req.Content,
		Version: req.Version,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, n)
}

func (s *Server) deleteReleaseNote(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteReleaseNote(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
