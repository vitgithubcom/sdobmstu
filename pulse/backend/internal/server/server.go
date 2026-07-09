package server

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "pulse-backend/internal/config"
    "pulse-backend/internal/handlers"
    "pulse-backend/internal/middleware"
    "pulse-backend/internal/repository"
    "pulse-backend/internal/service"

    _ "github.com/lib/pq"
)

type Server struct {
    config *config.Config
}

func NewServer(cfg *config.Config) *Server {
    return &Server{config: cfg}
}

func (s *Server) Run() error {
    // Подключение к БД
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        s.config.DBHost, s.config.DBPort, s.config.DBUser,
        s.config.DBPassword, s.config.DBName)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return err
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        return err
    }
    log.Println("✅ Connected to PostgreSQL")

    // Инициализация репозиториев
    userRepo := repository.NewUserRepository(db)
    kpiRepo := repository.NewKPIRepository(db)
    alertsRepo := repository.NewAlertsRepository(db)
    integrationsRepo := repository.NewIntegrationsRepository(db)
    auditRepo := repository.NewAuditRepository(db)

    // Инициализация сервисов
    authService := service.NewAuthService(userRepo, auditRepo, s.config.JWTSecret)
    kpiService := service.NewKPIService(kpiRepo)
    alertsService := service.NewAlertsService(alertsRepo)
    integrationsService := service.NewIntegrationsService(integrationsRepo)
    userService := service.NewUserService(userRepo)

    // Инициализация хендлеров
    authHandler := handlers.NewAuthHandler(authService)
    kpiHandler := handlers.NewKPIHandler(kpiService)
    alertsHandler := handlers.NewAlertsHandler(alertsService)
    integrationsHandler := handlers.NewIntegrationsHandler(integrationsService)
    userHandler := handlers.NewUserHandler(userService, auditRepo)
    auditHandler := handlers.NewAuditHandler(auditRepo)

    // Настройка роутов
    mux := http.NewServeMux()

    // Публичные
    mux.HandleFunc("POST /api/auth/login", authHandler.Login)

    // Защищённые
    authMiddleware := middleware.AuthMiddleware(authService)

    // KPI
    mux.Handle("GET /api/kpi", authMiddleware(http.HandlerFunc(kpiHandler.GetAll)))
    mux.Handle("GET /api/kpi/{id}", authMiddleware(http.HandlerFunc(kpiHandler.GetByID)))
    mux.Handle("GET /api/chart", authMiddleware(http.HandlerFunc(kpiHandler.GetChartData)))

    // Alerts
    mux.Handle("GET /api/alerts", authMiddleware(http.HandlerFunc(alertsHandler.GetAll)))

    // Integrations
    mux.Handle("GET /api/integrations", authMiddleware(http.HandlerFunc(integrationsHandler.GetAll)))

    // Users
    mux.Handle("GET /api/users", authMiddleware(http.HandlerFunc(userHandler.GetAll)))
    mux.Handle("GET /api/users/profile", authMiddleware(http.HandlerFunc(userHandler.GetProfile)))
    mux.Handle("PUT /api/users/profile", authMiddleware(http.HandlerFunc(userHandler.UpdateProfile)))
    mux.Handle("PUT /api/users/password", authMiddleware(http.HandlerFunc(userHandler.ChangePassword)))
    mux.Handle("POST /api/users", authMiddleware(http.HandlerFunc(userHandler.Create)))
    mux.Handle("PUT /api/users/{id}", authMiddleware(http.HandlerFunc(userHandler.Update)))
    mux.Handle("PATCH /api/users/{id}/toggle", authMiddleware(http.HandlerFunc(userHandler.ToggleActive)))

    // Audit
    mux.Handle("GET /api/audit", authMiddleware(http.HandlerFunc(auditHandler.GetAll)))

    // Обёртка с CORS и логгером
    handler := middleware.CORS(mux)
    handler = middleware.Logger(handler)

    addr := fmt.Sprintf(":%d", s.config.Port)
    log.Printf("🚀 Server starting on %s", addr)
    log.Printf("📊 API endpoints:")
    log.Printf("   POST /api/auth/login")
    log.Printf("   GET  /api/kpi")
    log.Printf("   GET  /api/kpi/{id}")
    log.Printf("   GET  /api/chart")
    log.Printf("   GET  /api/alerts")
    log.Printf("   GET  /api/integrations")
    log.Printf("   GET  /api/users")
    log.Printf("   GET  /api/audit")

    return http.ListenAndServe(addr, handler)
}