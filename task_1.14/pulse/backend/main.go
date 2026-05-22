package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type KPI struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Value  string  `json:"value"`
	Unit   string  `json:"unit"`
	Plan   int     `json:"plan"`
	Delta  float64 `json:"delta"`
	Trend  string  `json:"trend"`
	Status string  `json:"status"`
}

type ChartData struct {
	Name string `json:"name"`
	Fact int    `json:"факт"`
	Plan int    `json:"план"`
}

type Alert struct {
	ID       int    `json:"id"`
	System   string `json:"system"`
	Msg      string `json:"msg"`
	Severity string `json:"severity"`
	Time     string `json:"time"`
}

type Integration struct {
	Name     string `json:"name"`
	LastSync string `json:"lastSync"`
	Status   string `json:"status"`
	Lag      string `json:"lag"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type User struct {
	Login string `json:"login"`
	Role  string `json:"role"`
}


var mockKPI = []KPI{
	{ID: 1, Title: "Выручка", Value: "14 280", Unit: "тыс. руб", Plan: 15200, Delta: -6.1, Trend: "down", Status: "critical"},
	{ID: 2, Title: "OEE (общая эффективность)", Value: "73.5", Unit: "%", Plan: 78, Delta: -4.5, Trend: "down", Status: "warning"},
	{ID: 3, Title: "Простои", Value: "124", Unit: "часов", Plan: 80, Delta: 55, Trend: "down", Status: "critical"},
	{ID: 4, Title: "Заказы в срок", Value: "87.2", Unit: "%", Plan: 92, Delta: -4.8, Trend: "down", Status: "warning"},
	{ID: 5, Title: "Склад (оборачив.)", Value: "8.3", Unit: "дней", Plan: 7.2, Delta: 15, Trend: "down", Status: "critical"},
}

var mockChartData = []ChartData{
	{Name: "Пн", Fact: 1820, Plan: 2100},
	{Name: "Вт", Fact: 2050, Plan: 2100},
	{Name: "Ср", Fact: 1980, Plan: 2100},
	{Name: "Чт", Fact: 2140, Plan: 2100},
	{Name: "Пт", Fact: 1780, Plan: 2100},
	{Name: "Сб", Fact: 1620, Plan: 1800},
	{Name: "Вс", Fact: 1450, Plan: 1600},
}

var mockAlerts = []Alert{
	{ID: 1, System: "Mock-ERP", Msg: "План продаж под угрозой (выполнение 68%)", Severity: "critical", Time: "12:34"},
	{ID: 2, System: "Mock-MES", Msg: "Станок #1042: превышение времени цикла", Severity: "warning", Time: "12:28"},
	{ID: 3, System: "Mock-CRM", Msg: "Не загружены сделки за последний час", Severity: "critical", Time: "12:15"},
	{ID: 4, System: "Mock-Warehouse", Msg: "Остатки по группе А ниже нормы", Severity: "info", Time: "11:58"},
}

var mockIntegrations = []Integration{
	{Name: "Mock-ERP", LastSync: "12:44:23", Status: "ok", Lag: "12 сек"},
	{Name: "Mock-MES", LastSync: "12:44:01", Status: "ok", Lag: "34 сек"},
	{Name: "Mock-CRM", LastSync: "12:30:05", Status: "warning", Lag: "14 мин"},
	{Name: "Mock-Warehouse", LastSync: "12:42:10", Status: "ok", Lag: "2 мин"},
}


func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func getKPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockKPI)
}

func getChartData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockChartData)
}

func getAlerts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockAlerts)
}

func getIntegrations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockIntegrations)
}

func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Login == "admin" && req.Password == "admin123" {
		resp := LoginResponse{
			Token: "mock-jwt-token-123",
			User:  User{Login: "admin", Role: "admin"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
}

func verifyToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "Bearer mock-jwt-token-123" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"valid": "true"})
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"valid": "false"})
}


func main() {
	http.HandleFunc("GET /api/kpi", corsMiddleware(getKPI))
	http.HandleFunc("GET /api/chart", corsMiddleware(getChartData))
	http.HandleFunc("GET /api/alerts", corsMiddleware(getAlerts))
	http.HandleFunc("GET /api/integrations", corsMiddleware(getIntegrations))
	http.HandleFunc("POST /api/auth/login", corsMiddleware(login))
	http.HandleFunc("GET /api/auth/verify", corsMiddleware(verifyToken))

	log.Println("🚀 Go-сервер запущен на http://localhost:8080")
	log.Println("📊 API доступны:")
	log.Println("   GET  /api/kpi")
	log.Println("   GET  /api/chart")
	log.Println("   GET  /api/alerts")
	log.Println("   GET  /api/integrations")
	log.Println("   POST /api/auth/login")
	log.Println("   GET  /api/auth/verify")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}