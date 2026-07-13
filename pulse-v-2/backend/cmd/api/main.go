package main

import (
    "log"
    "pulse-backend/internal/config"
    "pulse-backend/internal/server"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    srv := server.NewServer(cfg)
    if err := srv.Run(); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}