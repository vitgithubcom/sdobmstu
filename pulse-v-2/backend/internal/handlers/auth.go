package handlers

import (
    "encoding/json"
    "net/http"
    "pulse-backend/internal/domain"
    "pulse-backend/internal/service"
)

type AuthHandler struct {
    authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req domain.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    ip := r.Header.Get("X-Real-IP")
    if ip == "" {
        ip = r.RemoteAddr
    }

    resp, err := h.authService.Login(req, ip)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}