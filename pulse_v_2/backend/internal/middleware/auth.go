package middleware

import (
    "context"
    "net/http"
    "pulse-backend/internal/service"
    "strings"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Authorization header required", http.StatusUnauthorized)
                return
            }

            parts := strings.Split(authHeader, " ")
            if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
                http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
                return
            }

            user, err := authService.ValidateToken(parts[1])
            if err != nil {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            ctx := context.WithValue(r.Context(), UserContextKey, user)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func GetUserFromContext(r *http.Request) interface{} {
    return r.Context().Value(UserContextKey)
}