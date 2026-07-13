package handlers

import (
    "encoding/json"
    "net/http"
    "pulse-backend/internal/service"
)

type AlertsHandler struct {
    service *service.AlertsService
}

func NewAlertsHandler(service *service.AlertsService) *AlertsHandler {
    return &AlertsHandler{service: service}
}

func (h *AlertsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    alerts, err := h.service.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(alerts)
}