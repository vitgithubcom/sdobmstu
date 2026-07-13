package handlers

import (
    "encoding/json"
    "net/http"
    "pulse-backend/internal/service"
    "strconv"
    "strings"
)

type KPIHandler struct {
    service *service.KPIService
}

func NewKPIHandler(service *service.KPIService) *KPIHandler {
    return &KPIHandler{service: service}
}

func (h *KPIHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    kpis, err := h.service.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(kpis)
}

func (h *KPIHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/api/kpi/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    kpi, err := h.service.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if kpi == nil {
        http.Error(w, "KPI not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(kpi)
}

func (h *KPIHandler) GetChartData(w http.ResponseWriter, r *http.Request) {
    data, err := h.service.GetChartData()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}