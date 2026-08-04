cat > go-app/main.go << 'EOF'
package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "time"

    _ "github.com/lib/pq"
)

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    Bio       string    `json:"bio"`
    Avatar    string    `json:"avatar"`
}

var db *sql.DB

func main() {
    connStr := "host=postgres port=5432 user=postgres password=postgres dbname=testdb sslmode=disable"
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Ошибка подключения к БД:", err)
    }

    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    if err := db.Ping(); err != nil {
        log.Fatal("БД недоступна:", err)
    }
    log.Println("Подключение к PostgreSQL установлено")

    http.HandleFunc("/users/", handleUser)
    http.HandleFunc("/health", handleHealth)
    http.HandleFunc("/debug/metrics", handleMetrics)

    log.Println("Сервер запущен на :8084")
    log.Fatal(http.ListenAndServe(":8084", nil))
}

func handleUser(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    idStr := r.URL.Path[len("/users/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil || id < 1 {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // 🔴 ИСПОЛЬЗУЕМ Query() ВМЕСТО QueryRow()
    // 🔴 И НЕ ЗАКРЫВАЕМ rows!
    query := "SELECT id, name, email, created_at, bio, avatar FROM users WHERE id = $1"
    rows, err := db.Query(query, id)
    if err != nil {
        log.Printf("Ошибка запроса: %v", err)
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    // ❌ rows.Close() НЕ ВЫЗВАН!
    // defer rows.Close() // ЗАКОММЕНТИРОВАНО!

    var user User
    var found bool
    for rows.Next() {
        err := rows.Scan(
            &user.ID,
            &user.Name,
            &user.Email,
            &user.CreatedAt,
            &user.Bio,
            &user.Avatar,
        )
        if err != nil {
            log.Printf("Ошибка сканирования: %v", err)
            http.Error(w, "Database error", http.StatusInternalServerError)
            return
        }
        found = true
        break // берем только первого пользователя
    }

    if !found {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // ❌ И здесь rows.Close() не вызывается!
    // rows.Close() // ЗАКОММЕНТИРОВАНО!

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(user); err != nil {
        log.Printf("Ошибка сериализации JSON: %v", err)
        http.Error(w, "JSON encode error", http.StatusInternalServerError)
        return
    }

    log.Printf("Запрос /users/%d выполнен за %v", id, time.Since(start))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"ok"}`))
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
    stats := db.Stats()
    metrics := map[string]interface{}{
        "open_connections": stats.OpenConnections,
        "in_use":           stats.InUse,
        "idle":             stats.Idle,
        "wait_count":       stats.WaitCount,
        "wait_duration":    stats.WaitDuration.String(),
        "max_open":         stats.MaxOpenConnections,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(metrics)
}
EOF