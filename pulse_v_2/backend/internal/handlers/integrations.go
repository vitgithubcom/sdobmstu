package handlers

import (
    "encoding/json"
    "net/http"
    "pulse-backend/internal/service"
)

type IntegrationsHandler struct {
    service *service.IntegrationsService
}

func NewIntegrationsHandler(service *service.IntegrationsService) *IntegrationsHandler {
    return &IntegrationsHandler{service: service}
}

func (h *IntegrationsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    integrations, err := h.service.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(integrations)
}