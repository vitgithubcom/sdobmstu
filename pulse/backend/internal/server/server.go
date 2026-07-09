package server

import (
    "database/sql"
    "fmt"
    "net/http"
    "pulse-backend/internal/config"
    "pulse-backend/internal/handlers"
    "pulse-backend/internal/middleware"
    "pulse-backend/internal/repository"
    "pulse-backend/internal/service"
    "pulse-backend/pkg/logger"

    _ "github.com/lib/pq"
)

type Server struct {
    config *config.Config
    log    *logger.Logger
}

func NewServer(cfg *config.Config) *Server {
    log := logger.Default()
    log.SetPrefix("PULSE")

    if cfg.Env == "development" {
        log.SetUseColor(true)
        log.SetLevel(logger.DEBUG)
    } else {
        log.SetUseColor(false)
        log.SetLevel(logger.INFO)
    }

    return &Server{
        config: cfg,
        log:    log,
    }
}

func (s *Server) Run() error {
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        s.config.DBHost, s.config.DBPort, s.config.DBUser,
        s.config.DBPassword, s.config.DBName)

    s.log.Debug("Connecting to PostgreSQL: %s", connStr)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        s.log.Fatal("Failed to open database: %v", err)
        return err
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        s.log.Fatal("Failed to ping database: %v", err)
        return err
    }
    s.log.Info("✅ Connected to PostgreSQL on %s:%d", s.config.DBHost, s.config.DBPort)

    // ===== ИНИЦИАЛИЗАЦИЯ РЕПОЗИТОРИЕВ =====
    userRepo := repository.NewUserRepository(db)
    kpiRepo := repository.NewKPIRepository(db)
    alertsRepo := repository.NewAlertsRepository(db)
    integrationsRepo := repository.NewIntegrationsRepository(db)
    auditRepo := repository.NewAuditRepository(db)

    // ===== ИНИЦИАЛИЗАЦИЯ СЕРВИСОВ =====
    authService := service.NewAuthService(userRepo, auditRepo, s.config.JWTSecret)
    kpiService := service.NewKPIService(kpiRepo)
    alertsService := service.NewAlertsService(alertsRepo)
    integrationsService := service.NewIntegrationsService(integrationsRepo)
    userService := service.NewUserService(userRepo)

    // ===== ИНИЦИАЛИЗАЦИЯ ХЕНДЛЕРОВ =====
    authHandler := handlers.NewAuthHandler(authService)
    kpiHandler := handlers.NewKPIHandler(kpiService)
    alertsHandler := handlers.NewAlertsHandler(alertsService)
    integrationsHandler := handlers.NewIntegrationsHandler(integrationsService)
    userHandler := handlers.NewUserHandler(userService, auditRepo)
    auditHandler := handlers.NewAuditHandler(auditRepo)

    // ===== НАСТРОЙКА РОУТОВ =====
    mux := http.NewServeMux()

    // ---- Публичные (без авторизации) ----
    mux.HandleFunc("/api/auth/login", authHandler.Login)

    // ---- Защищённые (с авторизацией) ----
    authMiddleware := middleware.AuthMiddleware(authService)

    // KPI
    mux.Handle("/api/kpi", authMiddleware(http.HandlerFunc(kpiHandler.GetAll)))
    mux.Handle("/api/kpi/{id}", authMiddleware(http.HandlerFunc(kpiHandler.GetByID)))
    mux.Handle("/api/chart", authMiddleware(http.HandlerFunc(kpiHandler.GetChartData)))

    // Alerts
    mux.Handle("/api/alerts", authMiddleware(http.HandlerFunc(alertsHandler.GetAll)))

    // Integrations
    mux.Handle("/api/integrations", authMiddleware(http.HandlerFunc(integrationsHandler.GetAll)))

    // Users
    mux.Handle("/api/users", authMiddleware(http.HandlerFunc(userHandler.GetAll)))
    mux.Handle("/api/users/profile", authMiddleware(http.HandlerFunc(userHandler.GetProfile)))
    mux.Handle("/api/users/profile", authMiddleware(http.HandlerFunc(userHandler.UpdateProfile)))
    mux.Handle("/api/users/password", authMiddleware(http.HandlerFunc(userHandler.ChangePassword)))
    mux.Handle("/api/users", authMiddleware(http.HandlerFunc(userHandler.Create)))
    mux.Handle("/api/users/{id}", authMiddleware(http.HandlerFunc(userHandler.Update)))
    mux.Handle("/api/users/{id}/toggle", authMiddleware(http.HandlerFunc(userHandler.ToggleActive)))

    // Audit
    mux.Handle("/api/audit", authMiddleware(http.HandlerFunc(auditHandler.GetAll)))

    // ---- Обёртка с CORS и логгером ----
    handler := middleware.CORS(mux)
    handler = middleware.Logger(handler)

    addr := fmt.Sprintf(":%d", s.config.Port)

    s.log.Info("🚀 Server starting on http://localhost%s", addr)
    s.log.Info("📊 API endpoints:")
    s.log.Info("   POST /api/auth/login")
    s.log.Info("   GET  /api/kpi")
    s.log.Info("   GET  /api/kpi/{id}")
    s.log.Info("   GET  /api/chart")
    s.log.Info("   GET  /api/alerts")
    s.log.Info("   GET  /api/integrations")
    s.log.Info("   GET  /api/users")
    s.log.Info("   GET  /api/audit")
    s.log.Info("🌍 Environment: %s", s.config.Env)

    return http.ListenAndServe(addr, handler)
}