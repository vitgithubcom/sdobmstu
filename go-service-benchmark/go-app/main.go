package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	// Подключение к PostgreSQL (новый порт 5434)
	connStr := "host=postgres port=5432 user=postgres password=postgres dbname=testdb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Проверка подключения
	if err := db.Ping(); err != nil {
		log.Fatal("БД недоступна:", err)
	}
	log.Println("Подключение к PostgreSQL установлено")

	// Эндпоинты
	http.HandleFunc("/users/", handleUser)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/debug/metrics", handleMetrics)

	// Запуск сервера на порту 8084
	log.Println("Сервер запущен на :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}

// Обработчик получения пользователя
func handleUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Извлекаем ID из URL
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Имитация нагрузки - можно включать/отключать для тестов
	// time.Sleep(50 * time.Millisecond) // Раскомментировать для симуляции тормозов

	// Запрос к БД
	var user User
	query := "SELECT id, name, email, created_at, bio, avatar FROM users WHERE id = $1"
	err = db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.Bio,
		&user.Avatar,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Ошибка запроса к БД для id=%d: %v", id, err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Сериализация в JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Ошибка сериализации JSON: %v", err)
		http.Error(w, "JSON encode error", http.StatusInternalServerError)
		return
	}

	log.Printf("Запрос /users/%d выполнен за %v", id, time.Since(start))
}

// Health check
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

// Метрики для диагностики
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
