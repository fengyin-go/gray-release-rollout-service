package handler

import (
	"net/http"

	"grayrelease/pkg/httpx"
)

func (s *Server) registerApprovalRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/approvals", s.createApproval)
	mux.HandleFunc("GET /api/approvals", s.listApprovals)
	mux.HandleFunc("POST /api/approvals/{id}/approve", s.approve)
	mux.HandleFunc("POST /api/approvals/{id}/reject", s.reject)
	mux.HandleFunc("GET /api/releases/{id}/approval-summary", s.approvalSummary)
}

type createApprovalRequest struct {
	ReleaseID string `json:"release_id"`
	Approver  string `json:"approver"`
}

func (s *Server) createApproval(w http.ResponseWriter, r *http.Request) {
	var req createApprovalRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	a, err := s.svc.CreateApproval(req.ReleaseID, req.Approver)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, a)
}

func (s *Server) listApprovals(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ListApprovals(r.URL.Query().Get("release_id")))
}

type approvalActionRequest struct {
	Comment string `json:"comment"`
}

func (s *Server) approve(w http.ResponseWriter, r *http.Request) {
	var req approvalActionRequest
	_ = httpx.Decode(r, &req)
	a, err := s.svc.Approve(r.PathValue("id"), req.Comment)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, a)
}

func (s *Server) reject(w http.ResponseWriter, r *http.Request) {
	var req approvalActionRequest
	_ = httpx.Decode(r, &req)
	a, err := s.svc.Reject(r.PathValue("id"), req.Comment)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, a)
}

func (s *Server) approvalSummary(w http.ResponseWriter, r *http.Request) {
	httpx.OK(w, s.svc.ApprovalSummary(r.PathValue("id")))
}
