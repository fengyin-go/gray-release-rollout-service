package handler

import (
	"net/http"

	"grayrelease/internal/model"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerReleaseRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/releases", s.createRelease)
	mux.HandleFunc("GET /api/releases", s.listReleases)
	mux.HandleFunc("GET /api/releases/{id}", s.getRelease)
	mux.HandleFunc("PUT /api/releases/{id}", s.updateRelease)
	mux.HandleFunc("PATCH /api/releases/{id}/status", s.changeReleaseStatus)
	mux.HandleFunc("DELETE /api/releases/{id}", s.deleteRelease)
}

type createReleaseRequest struct {
	Name        string `json:"name"`
	ServiceID   string `json:"service_id"`
	VersionID   string `json:"version_id"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

func (s *Server) createRelease(w http.ResponseWriter, r *http.Request) {
	var req createReleaseRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	rel, err := s.svc.CreateRelease(model.Release{
		Name:        req.Name,
		ServiceID:   req.ServiceID,
		VersionID:   req.VersionID,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, rel)
}

func (s *Server) listReleases(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.ReleaseFilter{
		ServiceID: r.URL.Query().Get("service_id"),
		Status:    r.URL.Query().Get("status"),
		Keyword:   r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListReleases(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getRelease(w http.ResponseWriter, r *http.Request) {
	rel, err := s.svc.GetRelease(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, rel)
}

func (s *Server) updateRelease(w http.ResponseWriter, r *http.Request) {
	var req createReleaseRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	rel, err := s.svc.UpdateRelease(r.PathValue("id"), model.Release{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, rel)
}

type changeReleaseStatusRequest struct {
	Status   string `json:"status"`
	Operator string `json:"operator"`
}

func (s *Server) changeReleaseStatus(w http.ResponseWriter, r *http.Request) {
	var req changeReleaseStatusRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	rel, err := s.svc.ChangeReleaseStatus(r.PathValue("id"), req.Status, req.Operator)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, rel)
}

func (s *Server) deleteRelease(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteRelease(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
