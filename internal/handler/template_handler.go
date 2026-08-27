package handler

import (
	"net/http"

	"grayrelease/internal/model"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerTemplateRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/templates", s.createTemplate)
	mux.HandleFunc("GET /api/templates", s.listTemplates)
	mux.HandleFunc("GET /api/templates/{id}", s.getTemplate)
	mux.HandleFunc("PUT /api/templates/{id}", s.updateTemplate)
	mux.HandleFunc("DELETE /api/templates/{id}", s.deleteTemplate)
	mux.HandleFunc("POST /api/templates/{id}/apply", s.applyTemplate)
}

type createTemplateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schedule    []int  `json:"schedule"`
}

func (s *Server) createTemplate(w http.ResponseWriter, r *http.Request) {
	var req createTemplateRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	t, err := s.svc.CreateTemplate(model.ReleaseTemplate{
		Name:        req.Name,
		Description: req.Description,
		Schedule:    req.Schedule,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, t)
}

func (s *Server) listTemplates(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListTemplates())
}

func (s *Server) getTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := s.svc.GetTemplate(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, t)
}

func (s *Server) updateTemplate(w http.ResponseWriter, r *http.Request) {
	var req createTemplateRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	t, err := s.svc.UpdateTemplate(r.PathValue("id"), model.ReleaseTemplate{
		Name:        req.Name,
		Description: req.Description,
		Schedule:    req.Schedule,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, t)
}

func (s *Server) deleteTemplate(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteTemplate(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type applyTemplateRequest struct {
	ReleaseID string `json:"release_id"`
}

func (s *Server) applyTemplate(w http.ResponseWriter, r *http.Request) {
	var req applyTemplateRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.ApplyTemplate(r.PathValue("id"), req.ReleaseID)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, map[string]int{"created": n})
}
