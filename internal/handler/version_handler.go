package handler

import (
	"net/http"

	"grayrelease/internal/model"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerVersionRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/versions", s.createVersion)
	mux.HandleFunc("GET /api/versions", s.listVersions)
	mux.HandleFunc("GET /api/versions/{id}", s.getVersion)
	mux.HandleFunc("PUT /api/versions/{id}", s.updateVersion)
	mux.HandleFunc("DELETE /api/versions/{id}", s.deleteVersion)
}

type createVersionRequest struct {
	ServiceID   string `json:"service_id"`
	Version     string `json:"version"`
	ArtifactURL string `json:"artifact_url"`
	Checksum    string `json:"checksum"`
	SizeBytes   int64  `json:"size_bytes"`
	Description string `json:"description"`
}

func (s *Server) createVersion(w http.ResponseWriter, r *http.Request) {
	var req createVersionRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	v, err := s.svc.CreateVersion(model.Version{
		ServiceID:   req.ServiceID,
		Version:     req.Version,
		ArtifactURL: req.ArtifactURL,
		Checksum:    req.Checksum,
		SizeBytes:   req.SizeBytes,
		Description: req.Description,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, v)
}

func (s *Server) listVersions(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListVersions(r.URL.Query().Get("service_id")))
}

func (s *Server) getVersion(w http.ResponseWriter, r *http.Request) {
	v, err := s.svc.GetVersion(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, v)
}

func (s *Server) updateVersion(w http.ResponseWriter, r *http.Request) {
	var req createVersionRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	v, err := s.svc.UpdateVersion(r.PathValue("id"), model.Version{
		ArtifactURL: req.ArtifactURL,
		Checksum:    req.Checksum,
		SizeBytes:   req.SizeBytes,
		Description: req.Description,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, v)
}

func (s *Server) deleteVersion(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteVersion(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
