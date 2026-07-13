package handlers

import (
    "encoding/json"
    "net/http"
    "pulse-backend/internal/domain"
    "pulse-backend/internal/middleware"
    "pulse-backend/internal/service"
)

type UserHandler struct {
    userService *service.UserService
    auditRepo   interface {
        Create(userID int, action, details, ip string) error
    }
}

func NewUserHandler(userService *service.UserService, auditRepo interface {
    Create(userID int, action, details, ip string) error
}) *UserHandler {
    return &UserHandler{userService: userService, auditRepo: auditRepo}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    users, err := h.userService.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
    user := middleware.GetUserFromContext(r).(*domain.User)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
    currentUser := middleware.GetUserFromContext(r).(*domain.User)

    var req struct {
        FullName string `json:"full_name"`
        Email    string `json:"email"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    currentUser.FullName = req.FullName
    currentUser.Email = req.Email

    if err := h.userService.Update(currentUser); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    ip := r.Header.Get("X-Real-IP")
    h.auditRepo.Create(currentUser.ID, "update_profile", "Обновление профиля", ip)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(currentUser)
}

func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
    currentUser := middleware.GetUserFromContext(r).(*domain.User)

    var req struct {
        OldPassword string `json:"old_password"`
        NewPassword string `json:"new_password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    if err := h.userService.UpdatePassword(currentUser.ID, req.OldPassword, req.NewPassword); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    ip := r.Header.Get("X-Real-IP")
    h.auditRepo.Create(currentUser.ID, "change_password", "Смена пароля", ip)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Пароль успешно изменён"})
}

// ===== Create — POST /api/users/create =====
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        FullName string `json:"full_name"`
        Password string `json:"password"`
        Role     string `json:"role"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    user := &domain.User{
        Username: req.Username,
        Email:    req.Email,
        FullName: req.FullName,
        Role:     req.Role,
        IsActive: true,
    }

    if err := h.userService.Create(user, req.Password); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    currentUser := middleware.GetUserFromContext(r).(*domain.User)
    ip := r.Header.Get("X-Real-IP")
    h.auditRepo.Create(currentUser.ID, "create_user", "Создание пользователя "+req.Username, ip)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// ===== Update — PUT /api/users/update =====
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
    var req struct {
        ID       int    `json:"id"`
        Email    string `json:"email"`
        FullName string `json:"full_name"`
        Role     string `json:"role"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    user, err := h.userService.GetByID(req.ID)
    if err != nil || user == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    user.Email = req.Email
    user.FullName = req.FullName
    user.Role = req.Role

    if err := h.userService.Update(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    currentUser := middleware.GetUserFromContext(r).(*domain.User)
    ip := r.Header.Get("X-Real-IP")
    h.auditRepo.Create(currentUser.ID, "update_user", "Обновление пользователя "+user.Username, ip)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// ===== ToggleActive — PATCH /api/users/toggle =====
func (h *UserHandler) ToggleActive(w http.ResponseWriter, r *http.Request) {
    var req struct {
        ID       int  `json:"id"`
        IsActive bool `json:"is_active"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    if err := h.userService.ToggleActive(req.ID, req.IsActive); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    currentUser := middleware.GetUserFromContext(r).(*domain.User)
    ip := r.Header.Get("X-Real-IP")
    h.auditRepo.Create(currentUser.ID, "toggle_user", "Изменение статуса пользователя", ip)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Статус обновлён"})
}