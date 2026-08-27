package handler

import (
	"net/http"

	"grayrelease/internal/model"
	"grayrelease/pkg/httpx"
)

func (s *Server) registerStepRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/releases/{id}/steps", s.addSteps)
	mux.HandleFunc("GET /api/releases/{id}/steps", s.listSteps)
	mux.HandleFunc("DELETE /api/steps/{id}", s.deleteStep)
}

type addStepsRequest struct {
	Steps []model.ReleaseStep `json:"steps"`
}

func (s *Server) addSteps(w http.ResponseWriter, r *http.Request) {
	var req addStepsRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.AddReleaseSteps(r.PathValue("id"), req.Steps)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, map[string]int{"created": n})
}

func (s *Server) listSteps(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListReleaseSteps(r.PathValue("id")))
}

func (s *Server) deleteStep(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteReleaseStep(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
