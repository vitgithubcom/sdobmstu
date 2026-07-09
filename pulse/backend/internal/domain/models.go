package domain

import "time"

// ========== USER ==========
type User struct {
    ID           int        `json:"id"`
    Username     string     `json:"username"`
    Email        string     `json:"email"`
    PasswordHash string     `json:"-"`
    FullName     string     `json:"full_name"`
    Role         string     `json:"role"`
    IsActive     bool       `json:"is_active"`
    CreatedAt    time.Time  `json:"created_at"`
    UpdatedAt    time.Time  `json:"updated_at"`
    LastLogin    *time.Time `json:"last_login,omitempty"` // ← указатель, чтобы принимать NULL
}

type LoginRequest struct {
    Login    string `json:"login"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}

// ========== KPI ==========
type KPI struct {
    ID         int       `json:"id"`
    Code       string    `json:"code"`
    Name       string    `json:"name"`
    Unit       string    `json:"unit"`
    Value      float64   `json:"value"`
    Plan       float64   `json:"plan"`
    Delta      float64   `json:"delta"`
    Direction  string    `json:"direction"`
    Status     string    `json:"status"`
    Source     string    `json:"source_system"`
    Completion float64   `json:"completion"`
    History    []History `json:"history,omitempty"`
}

type History struct {
    Period string  `json:"period"`
    Fact   float64 `json:"факт"`
    Plan   float64 `json:"план"`
}

type ChartData struct {
    Name string `json:"name"`
    Fact int    `json:"факт"`
    Plan int    `json:"план"`
}

// ========== ALERTS ==========
type Alert struct {
    ID         int       `json:"id"`
    System     string    `json:"system"`
    Message    string    `json:"message"`
    Severity   string    `json:"severity"`
    IsActive   bool      `json:"is_active"`
    CreatedAt  time.Time `json:"created_at"`
    ResolvedAt *time.Time `json:"resolved_at,omitempty"` // ← тоже указатель для NULL
}

// ========== INTEGRATIONS ==========
type Integration struct {
    ID           int       `json:"id"`
    Name         string    `json:"name"`
    Status       string    `json:"status"`
    LastSync     time.Time `json:"last_sync"`
    LagSeconds   int       `json:"lag_seconds"`
    ErrorMessage string    `json:"error_message"`
}

// ========== AUDIT ==========
type AuditLog struct {
    ID         int       `json:"id"`
    UserID     int       `json:"user_id"`
    Username   string    `json:"username"`
    FullName   string    `json:"full_name"`
    Action     string    `json:"action"`
    Details    string    `json:"details"`
    IPAddress  string    `json:"ip_address"`
    CreatedAt  time.Time `json:"created_at"`
}