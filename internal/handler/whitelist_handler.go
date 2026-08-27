package handler

import (
	"net/http"

	"grayrelease/pkg/httpx"
)

func (s *Server) registerWhitelistRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/whitelists", s.addWhitelist)
	mux.HandleFunc("POST /api/whitelists/batch", s.batchAddWhitelist)
	mux.HandleFunc("POST /api/whitelists/batch-remove", s.batchRemoveWhitelist)
	mux.HandleFunc("GET /api/whitelists", s.listWhitelists)
	mux.HandleFunc("DELETE /api/whitelists", s.removeWhitelist)
}

type whitelistRequest struct {
	ReleaseID string `json:"release_id"`
	UserID    string `json:"user_id"`
	Note      string `json:"note"`
}

func (s *Server) addWhitelist(w http.ResponseWriter, r *http.Request) {
	var req whitelistRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	wl, err := s.svc.AddWhitelist(req.ReleaseID, req.UserID, req.Note)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, wl)
}

func (s *Server) listWhitelists(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListWhitelists(r.URL.Query().Get("release_id")))
}

func (s *Server) removeWhitelist(w http.ResponseWriter, r *http.Request) {
	releaseID := r.URL.Query().Get("release_id")
	userID := r.URL.Query().Get("user_id")
	if err := s.svc.RemoveWhitelist(releaseID, userID); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type batchWhitelistRequest struct {
	ReleaseID string   `json:"release_id"`
	UserIDs   []string `json:"user_ids"`
	Note      string   `json:"note"`
}

func (s *Server) batchAddWhitelist(w http.ResponseWriter, r *http.Request) {
	var req batchWhitelistRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	n, err := s.svc.BatchAddWhitelist(req.ReleaseID, req.UserIDs, req.Note)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, map[string]int{"added": n})
}

func (s *Server) batchRemoveWhitelist(w http.ResponseWriter, r *http.Request) {
	var req batchWhitelistRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	httpx.OK(w, map[string]int{"removed": s.svc.BatchRemoveWhitelist(req.ReleaseID, req.UserIDs)})
}
