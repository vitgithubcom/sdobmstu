package handlers

import (
    "encoding/json"
    "net/http"
    "pulse-backend/internal/repository"
)

type AuditHandler struct {
    repo *repository.AuditRepository
}

func NewAuditHandler(repo *repository.AuditRepository) *AuditHandler {
    return &AuditHandler{repo: repo}
}

func (h *AuditHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    logs, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(logs)
}